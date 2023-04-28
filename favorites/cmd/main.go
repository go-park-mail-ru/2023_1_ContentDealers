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
	delivery "github.com/go-park-mail-ru/2023_1_ContentDealers/favorites/internal/delivery/content"
	repository "github.com/go-park-mail-ru/2023_1_ContentDealers/favorites/internal/repository/content"
	usecase "github.com/go-park-mail-ru/2023_1_ContentDealers/favorites/internal/usecase/content"
	favContentProto "github.com/go-park-mail-ru/2023_1_ContentDealers/favorites/pkg/proto/content"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/client/postgresql"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
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

	cfg := cfgGeneral.Favorites

	logger, err := logging.NewLogger(cfg.Logging, "favorites service")
	if err != nil {
		return fmt.Errorf("Fail to initialization logger: %w", err)
	}

	db, err := postgresql.NewClientPostgres(cfg.Postgres)
	if err != nil {
		logger.Error(err)
		return err
	}

	favContentRepository := repository.NewRepository(db, logger)
	favContentUseCase := usecase.NewUseCase(&favContentRepository, logger)
	favContentService := delivery.NewGrpc(favContentUseCase, logger)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(favContentService.LogInterceptor),
	)

	favContentProto.RegisterFavoritesContentServiceServer(server, favContentService)

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
