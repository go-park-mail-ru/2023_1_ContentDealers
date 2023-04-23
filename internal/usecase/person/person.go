package person

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type Person struct {
	contentGateway ContentGateway
	logger         logging.Logger
}

func NewPerson(contentGateway ContentGateway, logger logging.Logger) *Person {
	return &Person{
		contentGateway: contentGateway,
		logger:         logger,
	}
}

func (uc *Person) GetByID(ctx context.Context, id uint64) (domain.Person, error) {
	return uc.contentGateway.GetPersonByID(ctx, id)
}
