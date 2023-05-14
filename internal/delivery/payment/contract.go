package payment

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/payment/pkg/domain"
)

type Gateway interface {
	Accept(ctx context.Context, payment domain.Payment) error
	GetPaymentLink(ctx context.Context, userID uint64) (string, error)
}
