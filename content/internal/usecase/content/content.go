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

func (uc *UseCase) GetFilmByContentID(ctx context.Context, ContentID uint64) (domain.Film, error) {
	film, err := uc.repo.GetFilmByContentID(ctx, ContentID)
	if err != nil {
		return domain.Film{}, err
	}
	film.Content, err = uc.GetByID(ctx, ContentID)
	return film, err
}

func (uc *UseCase) GetSeriesByContentID(ctx context.Context, ContentID uint64) (domain.Series, error) {
	series, err := uc.repo.GetSeriesByContentID(ctx, ContentID)
	if err != nil {
		return domain.Series{}, err
	}
	series.Content, err = uc.GetByID(ctx, ContentID)
	return series, err
}

func (uc *UseCase) GetContentByContentIDs(ctx context.Context, ContentIDs []uint64) ([]domain.Content, error) {
	return uc.repo.GetByIDs(ctx, ContentIDs)
}
