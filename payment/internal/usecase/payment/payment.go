package payment

import (
	"context"
	// nolint:gosec
	"crypto/md5"
	"errors"
	"fmt"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/payment/pkg/domain"
)

var ErrIncorrectPaymentAmount = errors.New("incorrect payment amount")
var ErrIncorrectPaymentSign = errors.New("incorrect payment sign")

type Config struct {
	Secret            string
	Secret2           string
	MerchantID        string
	Currency          string
	SubscriptionPrice uint32
}

type UseCase struct {
	gateway Gateway
	cfg     Config
}

func NewUseCase(gateway Gateway, config Config) *UseCase {
	return &UseCase{gateway: gateway, cfg: config}
}

func (uc *UseCase) Accept(ctx context.Context, payment domain.Payment) error {
	if payment.Amount < uc.cfg.SubscriptionPrice {
		return ErrIncorrectPaymentAmount
	}

	// nolint:gosec
	sign := md5.Sum([]byte(fmt.Sprintf("%s:%d:%s:%s", uc.cfg.MerchantID, payment.Amount, uc.cfg.Secret2,
		payment.OrderID)))
	if fmt.Sprintf("%x", sign) != payment.Sign {
		return ErrIncorrectPaymentSign
	}

	var userID uint64
	if _, err := fmt.Sscanf(payment.OrderID, "%d", &userID); err != nil {
		return err
	}

	return uc.gateway.Subscribe(ctx, userID)
}

func (uc *UseCase) GetPaymentLink(ctx context.Context, userID uint64) string {
	// nolint:gosec
	sign := md5.Sum([]byte(fmt.Sprintf("%s:%d:%s:%s:%d", uc.cfg.MerchantID, uc.cfg.SubscriptionPrice,
		uc.cfg.Secret, uc.cfg.Currency, userID)))
	return fmt.Sprintf("https://pay.freekassa.ru/?m=%s&oa=%d&currency=%s&o=%d&s=%x", uc.cfg.MerchantID,
		uc.cfg.SubscriptionPrice, uc.cfg.Currency, userID, sign)
}
