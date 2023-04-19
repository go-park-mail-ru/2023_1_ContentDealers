package genre

import (
	"context"
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestRepository_GetByContentIDs(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	contentIDs := []uint64{0, 1, 2}
	rows := sqlmock.
		NewRows([]string{"content_id", "id", "content_url"})
	expect := map[uint64][]domain.Genre{
		0: {
			{0, "Drama"}, {1, "Criminal"},
		},
		1: {
			{2, "Science fiction"}, {0, "Drama"},
		},
		2: {
			{0, "Drama"}, {1, "Criminal"}, {3, "Detective"},
		},
	}
	for contentID, genres := range expect {
		for _, genre := range genres {
			rows = rows.AddRow(contentID, genre.ID, genre.Name)
		}
	}

	mock.
		ExpectQuery(`select c.id, g.id, g.name from genres g join content_genres cg on g.id = cg.genre_id
								join content c on cg.content_id = c.id where`).
		WithArgs(pq.Array(contentIDs)).
		WillReturnRows(rows)

	logger, err := logging.NewLogger(logging.LoggingConfig{})
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	logger.Out = ioutil.Discard
	repo := NewRepository(db, logger)
	content, err := repo.GetByContentIDs(context.Background(), contentIDs)
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
		ExpectQuery(`select c.id, g.id, g.name from genres g join content_genres cg on g.id = cg.genre_id
								join content c on cg.content_id = c.id where`).
		WithArgs(pq.Array(contentIDs)).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetByContentIDs(context.Background(), contentIDs)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	require.NotNil(t, err)
}

func TestRepository_GetByContentID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	contentID := uint64(0)
	rows := sqlmock.
		NewRows([]string{"content_id", "id", "content_url"})
	expect := []domain.Genre{
		{0, "Drama"}, {1, "Criminal"},
	}
	for _, genre := range expect {
		rows = rows.AddRow(contentID, genre.ID, genre.Name)
	}

	mock.
		ExpectQuery(`select c.id, g.id, g.name from genres g join content_genres cg on g.id = cg.genre_id
								join content c on cg.content_id = c.id where`).
		WithArgs(pq.Array([]uint64{contentID})).
		WillReturnRows(rows)

	logger, err := logging.NewLogger(logging.LoggingConfig{})
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	logger.Out = ioutil.Discard
	repo := NewRepository(db, logger)
	content, err := repo.GetByContentID(context.Background(), contentID)
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
		ExpectQuery(`select c.id, g.id, g.name from genres g join content_genres cg on g.id = cg.genre_id
								join content c on cg.content_id = c.id where`).
		WithArgs(pq.Array([]uint64{contentID})).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetByContentID(context.Background(), contentID)
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

	personID := uint64(0)
	rows := sqlmock.
		NewRows([]string{"id", "content_url"})
	expect := []domain.Genre{
		{0, "Drama"}, {1, "Criminal"},
	}
	for _, genre := range expect {
		rows = rows.AddRow(genre.ID, genre.Name)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`select distinct(g.id), g.name from genres g
			join content_genres cg on g.id = cg.genre_id
			join content_roles_persons crp on cg.content_id = crp.content_id
			where crp.person_id = $1
			order by g.id`)).
		WithArgs(personID).
		WillReturnRows(rows)

	logger, err := logging.NewLogger(logging.LoggingConfig{})
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	logger.Out = ioutil.Discard
	repo := NewRepository(db, logger)
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
		ExpectQuery(regexp.QuoteMeta(`select distinct(g.id), g.name from genres g
			join content_genres cg on g.id = cg.genre_id
			join content_roles_persons crp on cg.content_id = crp.content_id
			where crp.person_id = $1
			order by g.id`)).
		WithArgs(personID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetByPersonID(context.Background(), personID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	require.NotNil(t, err)
}
