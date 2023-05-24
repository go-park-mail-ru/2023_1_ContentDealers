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
	paymentDelivery "github.com/go-park-mail-ru/2023_1_ContentDealers/payment/internal/delivery/payment"
	user "github.com/go-park-mail-ru/2023_1_ContentDealers/payment/internal/gateway/user"
	paymentUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/payment/internal/usecase/payment"
	payment "github.com/go-park-mail-ru/2023_1_ContentDealers/payment/pkg/proto/payment"
	interceptorServer "github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/grpc/interceptor/server"
	pingDelivery "github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/grpc/ping"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/proto/ping"
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
		return fmt.Errorf("needed to pass config file")
	}

	cfgGeneral, err := config.GetCfg(*configPtr)
	if err != nil {
		return fmt.Errorf("fail to parse config yml file: %w", err)
	}
	cfg := cfgGeneral.Payment

	logger, err := logging.NewLogger(cfg.Logging, "payment serivce")
	if err != nil {
		return fmt.Errorf("fail to initialization logger: %w", err)
	}

	userGateway, err := user.NewGateway(user.ServiceUserConfig(cfg.ServiceUser), logger)
	if err != nil {
		return fmt.Errorf("failed to create new user gateway: %w", err)
	}

	useCase := paymentUseCase.NewUseCase(userGateway, paymentUseCase.Config{
		Secret:            cfg.Secret,
		Secret2:           cfg.Secret2,
		MerchantID:        cfg.MerchantID,
		Currency:          cfg.Currency,
		SubscriptionPrice: cfg.SubscriptionPrice,
	})

	paymentService := paymentDelivery.NewGrpc(useCase, logger)
	pingService := pingDelivery.NewGrpc()

	interceptor := interceptorServer.NewInterceptorServer("payment", logger)
	server := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.LogAndMetrics),
	)
	payment.RegisterPaymentServiceServer(server, paymentService)
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
