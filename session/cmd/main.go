package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os/signal"
	"syscall"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/client/redis"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	config "github.com/go-park-mail-ru/2023_1_ContentDealers/session/config"
	delivery "github.com/go-park-mail-ru/2023_1_ContentDealers/session/internal/delivery/session"
	repository "github.com/go-park-mail-ru/2023_1_ContentDealers/session/internal/repository/session"
	usecase "github.com/go-park-mail-ru/2023_1_ContentDealers/session/internal/usecase/session"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/proto/session"
	"google.golang.org/grpc"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	configPtr := flag.String("config", "", "Config file")
	flag.StringVar(configPtr, "c", "", "Config file (short)")

	flag.Parse()

	if *configPtr == "" {
		return fmt.Errorf("Needed to pass config file")
	}

	cfg, err := config.GetCfg(*configPtr)
	if err != nil {
		return fmt.Errorf("Fail to parse config yml file: %w", err)
	}

	logger, err := logging.NewLogger(cfg.Logging, "session service")
	if err != nil {
		return fmt.Errorf("Fail to initialization logger: %w", err)
	}

	redisClient, err := redis.NewClientRedis(cfg.Redis)
	if err != nil {
		logger.Error(err)
		return err
	}

	sessionRepository := repository.NewRepository(redisClient, logger)
	sessionUseCase := usecase.NewSession(&sessionRepository, logger, cfg.Session.ExpiresAt)
	sessionService := delivery.NewGrpc(sessionUseCase, logger)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(sessionService.LogInterceptor),
	)

	session.RegisterSessionServiceServer(server, sessionService)

	addr := fmt.Sprintf("%s:%s", cfg.Server.BindIP, cfg.Server.Port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Fatalln("cant listen port", err)
	}
	logger.Infoln("start listening on", addr)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	go func() {
		err = server.Serve(lis)
		if err != nil {
			logger.Error(err)
		}
	}()

	<-ctx.Done()
	server.GracefulStop()
	return nil
}
