package person

import (
	"context"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/stretchr/testify/require"
)

func TestRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	personID := uint64(0)
	rows := sqlmock.
		NewRows([]string{"id", "name", "gender", "growth", "birthplace", "avatar_url", "age"})
	expect := domain.Person{
		ID:         0,
		Name:       "Ivan Ivanov",
		Gender:     "M",
		Growth:     nil,
		Birthplace: nil,
		AvatarURL:  "media/avatar",
		Age:        18,
	}
	rows.AddRow(expect.ID, expect.Name, expect.Gender, expect.Growth, expect.Birthplace, expect.AvatarURL, expect.Age)

	mock.
		ExpectQuery(`select p.id, p.name, p.gender, p.growth,
       							p.birthplace, p.avatar_url, p.age from persons p where`).
		WithArgs(personID).
		WillReturnRows(rows)

	repo := NewRepository(db)
	content, err := repo.GetByID(context.Background(), personID)
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
		ExpectQuery(`select p.id, p.name, p.gender, p.growth,
       							p.birthplace, p.avatar_url, p.age from persons p where`).
		WithArgs(personID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.GetByID(context.Background(), personID)
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
		NewRows([]string{"id", "name", "gender", "growth", "birthplace", "avatar_url", "age"})
	expect := []domain.Person{
		{
			ID:         0,
			Name:       "Ivan Ivanov",
			Gender:     "M",
			Growth:     nil,
			Birthplace: nil,
			AvatarURL:  "media/avatar0",
			Age:        18,
		},
		{
			ID:         0,
			Name:       "Elena Khromova",
			Gender:     "F",
			Growth:     nil,
			Birthplace: nil,
			AvatarURL:  "media/avatar1",
			Age:        31,
		},
	}
	for _, person := range expect {
		rows.AddRow(person.ID, person.Name, person.Gender, person.Growth, person.Birthplace,
			person.AvatarURL, person.Age)
	}

	mock.
		ExpectQuery(`select p.id, p.name, p.gender, p.growth,
       							p.birthplace, p.avatar_url, p.age from persons p
       							join content_roles_persons crp on crp.person_id = p.id
       							where`).
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
		ExpectQuery(`select p.id, p.name, p.gender, p.growth,
       							p.birthplace, p.avatar_url, p.age from persons p
       							join content_roles_persons crp on crp.person_id = p.id
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
