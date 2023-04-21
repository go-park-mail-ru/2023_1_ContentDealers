package role

import (
	"context"
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/stretchr/testify/require"
)

func TestRepository_GetByContentID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	contentID := uint64(0)
	rows := sqlmock.
		NewRows([]string{"person_id", "id", "title"})
	expect := map[uint64]domain.Role{
		0: {0, "Actor"},
		1: {1, "Author"},
		2: {2, "Producer"},
	}
	for personID, role := range expect {
		rows = rows.AddRow(personID, role.ID, role.Title)
	}

	mock.
		ExpectQuery(`select crp.person_id, r.id, r.title from roles r 
    		  					join content_roles_persons crp on r.id = crp.role_id
    		  					where`).
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
		ExpectQuery(`select crp.person_id, r.id, r.title from roles r 
    		  					join content_roles_persons crp on r.id = crp.role_id
    		  					where`).
		WithArgs(contentID).
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
		NewRows([]string{"id", "title"})
	expect := []domain.Role{
		{0, "Actor"},
		{1, "Author"},
		{2, "Producer"},
	}
	for _, role := range expect {
		rows = rows.AddRow(role.ID, role.Title)
	}

	mock.
		ExpectQuery(regexp.QuoteMeta(`select distinct(r.id), r.title from roles r 
    		  join content_roles_persons crp on r.id = crp.role_id
    		  where crp.person_id = $1
			  order by r.id`)).
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
		ExpectQuery(regexp.QuoteMeta(`select distinct(r.id), r.title from roles r 
    		  join content_roles_persons crp on r.id = crp.role_id
    		  where crp.person_id = $1
			  order by r.id`)).
		WithArgs(personID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetByPersonID(context.Background(), personID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	require.NotNil(t, err)
}
