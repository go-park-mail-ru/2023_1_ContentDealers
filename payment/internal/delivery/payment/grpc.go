package payment

import (
	"context"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/payment/pkg/domain"
	proto "github.com/go-park-mail-ru/2023_1_ContentDealers/payment/pkg/proto/payment"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type Grpc struct {
	useCase UseCase
	logger  logging.Logger

	proto.UnimplementedPaymentServiceServer
}

func NewGrpc(useCase UseCase, logger logging.Logger) *Grpc {
	return &Grpc{
		useCase: useCase,
		logger:  logger,
	}
}

func (service *Grpc) Accept(ctx context.Context, protoPayment *proto.Payment) (*proto.Nothing, error) {
	payment := domain.Payment{}
	err := dto.Map(&payment, protoPayment)
	if err != nil {
		service.logger.Trace(err)
		return nil, err
	}

	return &proto.Nothing{}, service.useCase.Accept(ctx, payment)
}

func (service *Grpc) GetPaymentLink(ctx context.Context, protoUserID *proto.UserID) (*proto.PaymentLink, error) {
	userID := protoUserID.GetID()
	return &proto.PaymentLink{Link: service.useCase.GetPaymentLink(ctx, userID)}, nil
}
