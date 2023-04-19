package content

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type Content struct {
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
}

func NewContent(options Options, logger logging.Logger) *Content {
	return &Content{
		repo:        options.ContentRepo,
		personRoles: options.PersonRolesUseCase,
		genre:       options.GenreRepo,
		selection:   options.SelectionRepo,
		country:     options.CountryRepo,
		logger:      logger,
	}
}

func (uc *Content) GetByID(ctx context.Context, id uint64) (domain.Content, error) {
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
