package content

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var contentID uint64 = 0
	rows := sqlmock.
		NewRows([]string{"id", "title", "description", "rating", "year", "is_free", "age_limit", "trailer_url",
			"preview_url", "type"})
	expect := []domain.Content{
		{
			ID:           contentID,
			Title:        "Title",
			Description:  "Description",
			Rating:       10,
			Year:         2000,
			IsFree:       true,
			AgeLimit:     18,
			TrailerURL:   "0.mp4",
			PreviewURL:   "0.jpg",
			Type:         "film",
			PersonsRoles: nil,
			Genres:       nil,
			Selections:   nil,
			Countries:    nil,
		},
	}
	for _, content := range expect {
		rows = rows.AddRow(content.ID, content.Title, content.Description, content.Rating, content.Year, content.IsFree,
			content.AgeLimit, content.TrailerURL, content.PreviewURL, content.Type)
	}

	mock.
		ExpectQuery(`select c.id, c.title, c.description, c.rating, c.year, c.is_free, c.age_limit,
       							c.trailer_url, c.preview_url, c.type from content c where`).
		WithArgs(pq.Array([]uint64{contentID})).
		WillReturnRows(rows)

	repo := NewRepository(db)

	content, err := repo.GetByID(context.Background(), contentID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	require.Equal(t, content, expect[0])

	mock.
		ExpectQuery(`select c.id, c.title, c.description, c.rating, c.year, c.is_free, c.age_limit,
       							c.trailer_url, c.preview_url, c.type from content c where`).
		WithArgs(pq.Array([]uint64{contentID})).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetByID(context.Background(), contentID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	require.NotNil(t, err)
}

func TestRepository_GetByPersonID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var personID uint64 = 0
	rows := sqlmock.
		NewRows([]string{"id", "title", "description", "rating", "year", "is_free", "age_limit", "trailer_url",
			"preview_url", "type"})
	expect := []domain.Content{
		{
			ID:           0,
			Title:        "Title",
			Description:  "Description",
			Rating:       10,
			Year:         2000,
			IsFree:       true,
			AgeLimit:     18,
			TrailerURL:   "0.mp4",
			PreviewURL:   "0.jpg",
			Type:         "film",
			PersonsRoles: nil,
			Genres:       nil,
			Selections:   nil,
			Countries:    nil,
		},
		{
			ID:           1,
			Title:        "Title1",
			Description:  "Description1",
			Rating:       0,
			Year:         2003,
			IsFree:       false,
			AgeLimit:     18,
			TrailerURL:   "1.mp4",
			PreviewURL:   "1.jpg",
			Type:         "series",
			PersonsRoles: nil,
			Genres:       nil,
			Selections:   nil,
			Countries:    nil,
		},
	}
	for _, content := range expect {
		rows = rows.AddRow(content.ID, content.Title, content.Description, content.Rating, content.Year, content.IsFree,
			content.AgeLimit, content.TrailerURL, content.PreviewURL, content.Type)
	}

	mock.
		ExpectQuery(`select c.id, c.title, c.description, c.rating, c.year, c.is_free, c.age_limit,
       							c.trailer_url, c.preview_url, c.type from content c
       							join content_roles_persons crp on c.id = crp.content_id`).
		WithArgs(personID).
		WillReturnRows(rows)

	repo := NewRepository(db)
	content, err := repo.GetByPersonID(context.Background(), personID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	require.Equal(t, content, expect)

	mock.
		ExpectQuery(`select c.id, c.title, c.description, c.rating, c.year, c.is_free, c.age_limit,
       							c.trailer_url, c.preview_url, c.type from content c
       							join content_roles_persons crp on c.id = crp.content_id`).
		WithArgs(personID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetByPersonID(context.Background(), personID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	require.NotNil(t, err)
}

func TestRepository_GetBySelectionIDs(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	selectionsIDs := []uint64{0, 1, 2}
	rows := sqlmock.
		NewRows([]string{"content_id", "id", "title", "description", "rating", "year", "is_free", "age_limit",
			"trailer_url", "preview_url", "type"})
	expect := map[uint64][]domain.Content{
		0: {{
			ID:           0,
			Title:        "Title",
			Description:  "Description",
			Rating:       10,
			Year:         2000,
			IsFree:       true,
			AgeLimit:     18,
			TrailerURL:   "0.mp4",
			PreviewURL:   "0.jpg",
			Type:         "film",
			PersonsRoles: nil,
			Genres:       nil,
			Selections:   nil,
			Countries:    nil,
		}},
		1: {
			{
				ID:           1,
				Title:        "Title1",
				Description:  "Description1",
				Rating:       0,
				Year:         2003,
				IsFree:       false,
				AgeLimit:     18,
				TrailerURL:   "1.mp4",
				PreviewURL:   "1.jpg",
				Type:         "series",
				PersonsRoles: nil,
				Genres:       nil,
				Selections:   nil,
				Countries:    nil,
			},
		},
		2: {
			{
				ID:           2,
				Title:        "Title2",
				Description:  "Description2",
				Rating:       5,
				Year:         2003,
				IsFree:       false,
				AgeLimit:     18,
				TrailerURL:   "2.mp4",
				PreviewURL:   "2.jpg",
				Type:         "series",
				PersonsRoles: nil,
				Genres:       nil,
				Selections:   nil,
				Countries:    nil,
			},
			{
				ID:           3,
				Title:        "Title3",
				Description:  "Description3",
				Rating:       5,
				Year:         2007,
				IsFree:       false,
				AgeLimit:     18,
				TrailerURL:   "3.mp4",
				PreviewURL:   "3.jpg",
				Type:         "film",
				PersonsRoles: nil,
				Genres:       nil,
				Selections:   nil,
				Countries:    nil,
			},
		},
	}
	for selectionID, contentSlice := range expect {
		for _, content := range contentSlice {
			rows = rows.AddRow(selectionID,
				content.ID, content.Title, content.Description, content.Rating, content.Year, content.IsFree,
				content.AgeLimit, content.TrailerURL, content.PreviewURL, content.Type)
		}
	}

	mock.
		ExpectQuery(`select cs.selection_id, c.id, c.title, c.description, c.rating, c.year,
       							c.is_free, c.age_limit, c.trailer_url, c.preview_url, c.type from content c 
       		   					join content_selections cs on c.id = cs.content_id
       		   					where cs.selection_id`).
		WithArgs(pq.Array(selectionsIDs)).
		WillReturnRows(rows)

	repo := NewRepository(db)
	content, err := repo.GetBySelectionIDs(context.Background(), selectionsIDs)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	require.Equal(t, content, expect)

	mock.
		ExpectQuery(`select cs.selection_id, c.id, c.title, c.description, c.rating, c.year,
       							c.is_free, c.age_limit, c.trailer_url, c.preview_url, c.type from content c 
       		   					join content_selections cs on c.id = cs.content_id
       		   					where cs.selection_id`).
		WithArgs(pq.Array(selectionsIDs)).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetBySelectionIDs(context.Background(), selectionsIDs)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	require.NotNil(t, err)
}

func TestRepository_GetFilmByContentID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var contentID uint64 = 0
	rows := sqlmock.
		NewRows([]string{"id", "content_url"})
	expect := domain.Film{ID: 0, ContentURL: "film/0"}
	rows = rows.AddRow(expect.ID, expect.ContentURL)

	mock.
		ExpectQuery(`select id, content_url from films where`).
		WithArgs(contentID).
		WillReturnRows(rows)

	repo := NewRepository(db)
	content, err := repo.GetFilmByContentID(context.Background(), contentID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	require.Equal(t, content, expect)

	mock.
		ExpectQuery(`select id, content_url from films where`).
		WithArgs(contentID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetFilmByContentID(context.Background(), contentID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	require.NotNil(t, err)
}
