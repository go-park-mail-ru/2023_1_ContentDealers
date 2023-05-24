package payment

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/payment/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/payment/pkg/proto/payment"
	interceptorClient "github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/grpc/interceptor/client"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/grpc/ping"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"google.golang.org/grpc"
)

type Gateway struct {
	payment payment.PaymentServiceClient
	logger  logging.Logger
}

func NewGateway(logger logging.Logger, cfg ServicePaymentConfig) (*Gateway, error) {
	interceptor := interceptorClient.NewInterceptorClient("payment", logger)

	grpcConn, err := grpc.Dial(
		cfg.Addr,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor.AccessLog),
	)
	if err != nil {
		logger.Error("cant connect to grpc payment service")
		return nil, err
	}

	gateway := Gateway{
		payment: payment.NewPaymentServiceClient(grpcConn),
		logger:  logger,
	}

	err = ping.Ping(grpcConn)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &gateway, nil
}

func (gate *Gateway) Accept(ctx context.Context, p domain.Payment) error {
	_, err := gate.payment.Accept(ctx, &payment.Payment{
		Amount:  p.Amount,
		OrderID: p.OrderID,
		Sign:    p.Sign,
	})
	return err
}

func (gate *Gateway) GetPaymentLink(ctx context.Context, userID uint64) (string, error) {
	link, err := gate.payment.GetPaymentLink(ctx, &payment.UserID{ID: userID})
	return link.GetLink(), err
}
