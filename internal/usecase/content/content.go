package content

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type Content struct {
	repo        Repository
	personRoles PersonRolesRepository
	genre       GenreRepository
	selection   SelectionRepository
	country     CountryRepository
}

type Options struct {
	ContentRepo     Repository
	PersonRolesRepo PersonRolesRepository
	GenreRepo       GenreRepository
	SelectionRepo   SelectionRepository
	CountryRepo     CountryRepository
}

func NewContent(options Options) *Content {
	return &Content{
		repo:        options.ContentRepo,
		personRoles: options.PersonRolesRepo,
		genre:       options.GenreRepo,
		selection:   options.SelectionRepo,
		country:     options.CountryRepo,
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
