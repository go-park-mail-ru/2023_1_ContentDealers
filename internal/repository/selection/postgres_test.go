package selection

import (
	"context"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/stretchr/testify/require"
)

func TestRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	rows := sqlmock.
		NewRows([]string{"id", "title"})
	expect := []domain.Selection{
		{ID: 0, Title: "Best films"}, {ID: 1, Title: "Filmium recommend"}, {ID: 2, Title: "Choice of viewers"},
	}
	limit := uint(15)
	offset := uint(0)
	for _, selection := range expect {
		rows = rows.AddRow(selection.ID, selection.Title)
	}

	mock.
		ExpectQuery(`select s.id, s.title from selections s order by id desc`).
		WithArgs(limit, offset).
		WillReturnRows(rows)

	repo := NewRepository(db)
	content, err := repo.GetAll(context.Background(), limit, offset)
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
		ExpectQuery(`select s.id, s.title from selections s order by id desc`).
		WithArgs(limit, offset).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetAll(context.Background(), limit, offset)
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
		NewRows([]string{"id", "title"})
	expect := []domain.Selection{
		{ID: 0, Title: "Best films"}, {ID: 1, Title: "Filmium recommend"}, {ID: 2, Title: "Choice of viewers"},
	}
	for _, selection := range expect {
		rows = rows.AddRow(selection.ID, selection.Title)
	}
	mock.
		ExpectQuery(`select s.id, s.title from selections s 
    				join content_selections cs on cs.selection_id = s.id
					join content c on c.id = cs.content_id where`).
		WithArgs(contentID).
		WillReturnRows(rows)

	repo := NewRepository(db)
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
		ExpectQuery(`select s.id, s.title from selections s 
    				join content_selections cs on cs.selection_id = s.id
					join content c on c.id = cs.content_id where`).
		WithArgs(contentID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetByContentID(context.Background(), contentID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	require.NotNil(t, err)
}

func TestRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	ID := uint64(0)
	rows := sqlmock.
		NewRows([]string{"id", "title"})
	expect := domain.Selection{ID: ID, Title: "Best films"}

	rows = rows.AddRow(expect.ID, expect.Title)

	mock.
		ExpectQuery(`select s.id, s.title from selections s where id`).
		WithArgs(ID).
		WillReturnRows(rows)

	repo := NewRepository(db)
	content, err := repo.GetByID(context.Background(), ID)
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
		ExpectQuery(`select s.id, s.title from selections s where id`).
		WithArgs(ID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetByID(context.Background(), ID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	require.NotNil(t, err)
}
