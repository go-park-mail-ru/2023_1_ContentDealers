package person

import (
	"context"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/person"
)

type Grpc struct {
	person.UnimplementedPersonServiceServer

	useCase UseCase
}

func NewGrpc(useCase UseCase) *Grpc {
	return &Grpc{useCase: useCase}
}

func (service *Grpc) GetByID(ctx context.Context, personID *person.ID) (*person.Person, error) {
	id := personID.GetID()

	foundPerson, err := service.useCase.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	var response person.Person

	err = dto.Map(&response, foundPerson)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
