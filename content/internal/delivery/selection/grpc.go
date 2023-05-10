package selection

import (
	"context"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/selection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type Grpc struct {
	selection.UnimplementedSelectionServiceServer

	useCase UseCase
	logger  logging.Logger
}

func NewGrpc(useCase UseCase, logger logging.Logger) *Grpc {
	return &Grpc{useCase: useCase, logger: logger}
}

const (
	defaultLimit  = 15
	defaultOffset = 0
)

func (service *Grpc) GetAll(ctx context.Context, cfg *selection.GetAllCfg) (*selection.Selections, error) {
	limit := uint(cfg.GetLimit())
	offset := uint(cfg.GetOffset())

	all, err := service.useCase.GetAll(ctx, limit, offset)
	if err != nil {
		service.logger.WithRequestID(ctx).Error(err)
		return nil, err
	}

	var response selection.Selections
	err = dto.Map(&response.Selections, all)
	if err != nil {
		service.logger.WithRequestID(ctx).Error(err)
		return nil, err
	}

	return &response, nil
}

func (service *Grpc) GetByID(ctx context.Context, selectionID *selection.ID) (*selection.Selection, error) {
	id := selectionID.GetID()

	foundSelection, err := service.useCase.GetByID(ctx, id)
	if err != nil {
		service.logger.WithRequestID(ctx).Error(err)
		return nil, err
	}

	var response selection.Selection
	err = dto.Map(&response, foundSelection)
	if err != nil {
		service.logger.WithRequestID(ctx).Error(err)
		return nil, err
	}

	return &response, nil
}
