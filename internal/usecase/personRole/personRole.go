package personRole

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type PersonRole struct {
	personRepo PersonRepository
	roleRepo   RoleRepository
}

func NewPersonRole(person PersonRepository, role RoleRepository) *PersonRole {
	return &PersonRole{personRepo: person, roleRepo: role}
}

func (uc *PersonRole) GetByContentID(ctx context.Context, ContentID uint64) ([]domain.PersonRoles, error) {
	persons, err := uc.personRepo.GetByContentID(ctx, ContentID)
	if err != nil {
		return nil, err
	}
	roles, err := uc.roleRepo.GetByContentID(ctx, ContentID)
	if err != nil {
		return nil, err
	}
	personIdToIdx := make(map[uint64]int, len(persons))
	for idx, person := range persons {
		personIdToIdx[person.ID] = idx
	}
	result := make([]domain.PersonRoles, 0, len(persons))
	for personID, role := range roles {
		idx := personIdToIdx[personID]
		result = append(result, domain.PersonRoles{
			Person: persons[idx],
			Role:   role,
		})
	}
	return result, nil
}
