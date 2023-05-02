package person

import (
	"context"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/person"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type Grpc struct {
	person.UnimplementedPersonServiceServer

	useCase UseCase
	logger  logging.Logger
}

func NewGrpc(useCase UseCase, logger logging.Logger) *Grpc {
	return &Grpc{useCase: useCase, logger: logger}
}

func (service *Grpc) GetByID(ctx context.Context, personID *person.ID) (*person.Person, error) {
	id := personID.GetID()

	foundPerson, err := service.useCase.GetByID(ctx, id)
	if err != nil {
		service.logger.WithRequestID(ctx).Error(err)
		return nil, err
	}

	var response person.Person

	err = dto.Map(&response, foundPerson)
	if err != nil {
		service.logger.WithRequestID(ctx).Error(err)
		return nil, err
	}
	return &response, nil
}
