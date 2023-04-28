package content

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type UseCase struct {
	repo        Repository
	personRoles PersonRolesUseCase
	genre       GenreRepository
	selection   SelectionRepository
	country     CountryRepository
	logger      logging.Logger
}

type Options struct {
	ContentRepo        Repository
	GenreRepo          GenreRepository
	SelectionRepo      SelectionRepository
	CountryRepo        CountryRepository
	PersonRolesUseCase PersonRolesUseCase
	Logger             logging.Logger
}

func NewUseCase(options Options) *UseCase {
	return &UseCase{
		repo:        options.ContentRepo,
		personRoles: options.PersonRolesUseCase,
		genre:       options.GenreRepo,
		selection:   options.SelectionRepo,
		country:     options.CountryRepo,
		logger:      options.Logger,
	}
}

func (uc *UseCase) GetByID(ctx context.Context, id uint64) (domain.Content, error) {
	content, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return domain.Content{}, err
	}
	content.PersonsRoles, err = uc.personRoles.GetByContentID(ctx, id)
	if err != nil {
		return domain.Content{}, err
	}
	content.Genres, err = uc.genre.GetByContentID(ctx, id)
	if err != nil {
		return domain.Content{}, err
	}
	content.Selections, err = uc.selection.GetByContentID(ctx, id)
	if err != nil {
		return domain.Content{}, err
	}
	content.Countries, err = uc.country.GetByContentID(ctx, id)
	if err != nil {
		return domain.Content{}, err
	}
	return content, nil
}
