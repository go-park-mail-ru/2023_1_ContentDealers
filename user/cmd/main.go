package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os/signal"
	"syscall"

	config "github.com/go-park-mail-ru/2023_1_ContentDealers/config"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/client/postgresql"
	interceptorServer "github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/grpc/interceptor/server"
	pingDelivery "github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/grpc/ping"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/proto/ping"
	delivery "github.com/go-park-mail-ru/2023_1_ContentDealers/user/internal/delivery/user"
	repository "github.com/go-park-mail-ru/2023_1_ContentDealers/user/internal/repository/user"
	usecase "github.com/go-park-mail-ru/2023_1_ContentDealers/user/internal/usecase/user"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/proto/user"
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

	cfgGeneral, err := config.GetCfg(*configPtr)
	if err != nil {
		return fmt.Errorf("Fail to parse config yml file: %w", err)
	}

	cfg := cfgGeneral.User

	logger, err := logging.NewLogger(cfg.Logging, "user service")
	if err != nil {
		return fmt.Errorf("Fail to initialization logger: %w", err)
	}

	db, err := postgresql.NewClientPostgres(cfg.Postgres)
	if err != nil {
		logger.Error(err)
		return err
	}

	userRepository := repository.NewRepository(db, logger)
	userUseCase := usecase.NewUser(&userRepository, logger)
	userService := delivery.NewGrpc(userUseCase, logger)
	pingService := pingDelivery.NewGrpc()

	interceptor := interceptorServer.NewInterceptorServer("user", logger)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.LogAndMetrics),
	)

	user.RegisterUserServiceServer(server, userService)
	ping.RegisterPingServiceServer(server, pingService)

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
