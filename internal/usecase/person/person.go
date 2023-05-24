package person

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type UseCase struct {
	contentGateway ContentGateway
	logger         logging.Logger
}

func NewUseCase(contentGateway ContentGateway, logger logging.Logger) *UseCase {
	return &UseCase{
		contentGateway: contentGateway,
		logger:         logger,
	}
}

func (uc *UseCase) GetByID(ctx context.Context, id uint64) (domain.Person, error) {
	return uc.contentGateway.GetPersonByID(ctx, id)
}
