package content

import (
	"context"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/film"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/person"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/selection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"google.golang.org/grpc"
)

type Grpc struct {
	filmService      film.FilmServiceClient
	personService    person.PersonServiceClient
	selectionService selection.SelectionServiceClient

	logger logging.Logger
}

func NewGrpc(addr string, logger logging.Logger) (Grpc, error) {
	grpcConn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return Grpc{}, err
	}

	result := Grpc{}
	result.filmService = film.NewFilmServiceClient(grpcConn)
	result.personService = person.NewPersonServiceClient(grpcConn)
	result.selectionService = selection.NewSelectionServiceClient(grpcConn)

	return result, nil
}

func (gateway *Grpc) GetAllSelections(ctx context.Context, limit, offset uint32) ([]domain.Selection, error) {
	selections, err := gateway.selectionService.GetAll(ctx, &selection.GetAllCfg{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}
	var result []domain.Selection

	err = dto.Map(&result, selections.GetSelections())
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (gateway *Grpc) GetSelectionByID(ctx context.Context, id uint64) (domain.Selection, error) {
	selectionDTO, err := gateway.selectionService.GetByID(ctx, &selection.ID{ID: id})
	if err != nil {
		return domain.Selection{}, err
	}

	var result domain.Selection
	err = dto.Map(&result, selectionDTO)
	if err != nil {
		return domain.Selection{}, err
	}

	return result, nil
}
