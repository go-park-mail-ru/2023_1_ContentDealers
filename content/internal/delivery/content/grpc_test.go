package content

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/dranikpg/dto-mapper"
	contentUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/usecase/content"
	mock_content "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/usecase/content/mock"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/content"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGrpc_GetFilmByContentID(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	contentRepo := mock_content.NewMockRepository(ctrl)
	genreRepo := mock_content.NewMockGenreRepository(ctrl)
	selectionRepo := mock_content.NewMockSelectionRepository(ctrl)
	countryRepo := mock_content.NewMockCountryRepository(ctrl)
	personRolesUseCase := mock_content.NewMockPersonRolesUseCase(ctrl)
	usecase := contentUseCase.NewUseCase(contentUseCase.Options{
		ContentRepo:        contentRepo,
		GenreRepo:          genreRepo,
		SelectionRepo:      selectionRepo,
		CountryRepo:        countryRepo,
		PersonRolesUseCase: personRolesUseCase,
	})
	logger, err := logging.NewLogger(logging.LoggingConfig{}, "")
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
		return
	}
	logger.Out = ioutil.Discard
	service := NewGrpc(usecase, logger)

	personRoles := []domain.PersonRoles{
		{
			Person: domain.Person{
				ID:   1,
				Name: "Name1",
			},
			Role: domain.Role{
				ID:    1,
				Title: "Actor",
			},
		},
		{
			Person: domain.Person{
				ID:   1,
				Name: "Name2",
			},
			Role: domain.Role{
				ID:    2,
				Title: "Producer",
			},
		},
	}

	genres := []domain.Genre{
		{1, "Drama"},
		{2, "Horror"},
	}

	selections := []domain.Selection{
		{ID: 1, Title: "Best films"},
	}

	countries := []domain.Country{
		{1, "USA"},
	}
	expectedContent := domain.Content{
		ID:           1,
		Title:        "Title",
		Description:  "Description",
		PersonsRoles: personRoles,
		Genres:       genres,
		Selections:   selections,
		Countries:    countries,
	}

	expectedFilm := domain.Film{
		ID:         1,
		ContentURL: "",
		Content:    expectedContent,
	}

	contentRepo.EXPECT().GetFilmByContentID(context.Background(), uint64(1)).Times(1).Return(domain.Film{ID: 1}, nil)
	contentRepo.EXPECT().GetByID(context.Background(), uint64(1)).Times(1).Return(domain.Content{
		ID:          1,
		Title:       "Title",
		Description: "Description",
	}, nil)

	genreRepo.EXPECT().GetByContentID(context.Background(), uint64(1)).Times(1).Return(genres, nil)
	personRolesUseCase.EXPECT().GetByContentID(context.Background(), uint64(1)).Times(1).Return(personRoles, nil)
	selectionRepo.EXPECT().GetByContentID(context.Background(), uint64(1)).Times(1).Return(selections, nil)
	countryRepo.EXPECT().GetByContentID(context.Background(), uint64(1)).Times(1).Return(countries, nil)

	film, err := service.GetFilmByContentID(context.Background(), &content.ContentID{ID: 1})
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
		return
	}

	var expected content.Film
	err = dto.Map(&expected, expectedFilm)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
		return
	}

	require.Equal(t, &expected, film)
}

func TestGrpc_GetContentByContentIDs(t *testing.T) {

}

func TestGrpc_GetSeriesByContentID(t *testing.T) {

}
