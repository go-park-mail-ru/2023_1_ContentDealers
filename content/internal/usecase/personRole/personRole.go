package personRole

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type PersonRole struct {
	personRepo PersonRepository
	roleRepo   RoleRepository
	logger     logging.Logger
}

func NewPersonRole(person PersonRepository, role RoleRepository, logger logging.Logger) *PersonRole {
	return &PersonRole{personRepo: person, roleRepo: role, logger: logger}
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
	personIDToIdx := make(map[uint64]int, len(persons))
	for idx, person := range persons {
		personIDToIdx[person.ID] = idx
	}
	result := make([]domain.PersonRoles, 0, len(persons))
	for personID, role := range roles {
		idx := personIDToIdx[personID]
		result = append(result, domain.PersonRoles{
			Person: persons[idx],
			Role:   role,
		})
	}
	return result, nil
}
