package film

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/stretchr/testify/require"
)

func TestRepository_GetByContentID(t *testing.T) {
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
		ExpectQuery(`select id, content_url from films where`).
		WithArgs(contentID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetByContentID(context.Background(), contentID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	require.NotNil(t, err)
}
