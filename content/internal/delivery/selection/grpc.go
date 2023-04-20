package selection

import (
	"context"
	"fmt"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/selection"
	"google.golang.org/grpc/metadata"
)

type Grpc struct {
	selection.UnimplementedSelectionServiceServer

	useCase UseCase
}

func NewGrpc(useCase UseCase) *Grpc {
	return &Grpc{useCase: useCase}
}

const (
	defaultLimit  = 15
	defaultOffset = 0
)

func (service *Grpc) GetAll(ctx context.Context, _ *selection.Nothing) (*selection.Selections, error) {
	var limit uint = defaultLimit
	var offset uint = defaultOffset

	meta, ok := metadata.FromIncomingContext(ctx)
	if ok == true {
		limitArg := meta.Get("limit")
		if len(limitArg) >= 1 {
			limitRaw := limitArg[0]
			_, _ = fmt.Sscanf(limitRaw, "%d", &limit)
		}
		offsetArg := meta.Get("offset")
		if len(offsetArg) >= 1 {
			offsetRaw := offsetArg[0]
			_, _ = fmt.Sscanf(offsetRaw, "%d", &offset)
		}
	}

	all, err := service.useCase.GetAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	var response selection.Selections
	err = dto.Map(&response.Selections, all)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (service *Grpc) GetByID(ctx context.Context, selectionID *selection.ID) (*selection.Selection, error) {
	id := selectionID.GetID()

	foundSelection, err := service.useCase.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	var response selection.Selection
	err = dto.Map(&response, foundSelection)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
