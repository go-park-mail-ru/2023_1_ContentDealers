package search

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

func (uc *UseCase) Search(ctx context.Context, query string) (domain.Search, error) {
	return uc.contentGateway.Search(ctx, query)
}
