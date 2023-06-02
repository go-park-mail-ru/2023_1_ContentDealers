package personrole

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
)

type UseCase struct {
	personRepo PersonRepository
	roleRepo   RoleRepository
}

func NewUseCase(person PersonRepository, role RoleRepository) *UseCase {
	return &UseCase{personRepo: person, roleRepo: role}
}

func (uc *UseCase) GetByContentID(ctx context.Context, ContentID uint64) ([]domain.PersonRoles, error) {
	persons, err := uc.personRepo.GetByContentID(ctx, ContentID)
	if err != nil {
		return nil, err
	}
	roles, err := uc.roleRepo.GetByContentID(ctx, ContentID)
	if err != nil {
		return nil, err
	}
	result := make([]domain.PersonRoles, 0, len(persons))
	for _, person := range persons {
		result = append(result, domain.PersonRoles{
			Person: person,
			Role:   roles[person.ID],
		})
	}
	return result, nil
}
