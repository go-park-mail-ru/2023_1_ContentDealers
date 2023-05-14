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
	delivery "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/internal/delivery/favcontent"
	deliveryViews "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/internal/delivery/history_views"
	deliveryRating "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/internal/delivery/rating"
	repository "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/internal/repository/favcontent"
	repositoryViews "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/internal/repository/history_views"
	repositoryRating "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/internal/repository/rating"
	usecase "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/internal/usecase/favcontent"
	usecaseViews "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/internal/usecase/history_views"
	usecaseRating "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/internal/usecase/rating"
	favContentProto "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/proto/favcontent"
	viewsProto "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/proto/history_views"
	rateProto "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/proto/rating"
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

	rateRepository := repositoryRating.NewRepository(db, logger)
	rateUseCase := usecaseRating.NewUseCase(&rateRepository, logger)
	rateService := deliveryRating.NewGrpc(rateUseCase, logger)

	viewsRepository := repositoryViews.NewRepository(db, logger)
	viewsUseCase := usecaseViews.NewUseCase(&viewsRepository, logger)
	viewsService := deliveryViews.NewGrpc(viewsUseCase, logger)

	pingService := pingDelivery.NewGrpc()

	interceptor := interceptorServer.NewInterceptorServer("favorites", logger)

	server := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.LogAndMetrics),
	)

	favContentProto.RegisterFavoritesContentServiceServer(server, favContentService)
	rateProto.RegisterRatingServiceServer(server, rateService)
	viewsProto.RegisterHistoryViewsServiceServer(server, viewsService)
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
