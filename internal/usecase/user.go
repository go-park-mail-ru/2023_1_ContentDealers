package usecase

import (
	"crypto/sha256"
	"fmt"
	"regexp"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/contract"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

const salt = "hjqrhjqw124617ajfhajs"

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

var incorrentPasswordRegex = regexp.MustCompile(`(^(.{0,7}|[^0-9]*|[^A-Z]*|[^a-z]*|[a-zA-Z0-9]*)$)`)

var _ contract.UserUseCase = (*User)(nil)

type User struct {
	repo contract.UserRepository
}

func NewUser(repo contract.UserRepository) *User {
	return &User{repo: repo}
}

func (uc *User) Register(credentials domain.UserCredentials) (domain.User, error) {
	if err := validateCredentials(credentials); err != nil {
		return domain.User{}, err
	}
	credentials.Password = generatePasswordHash(credentials.Password)
	return uc.repo.Add(credentials)
}

func (uc *User) Auth(credentials domain.UserCredentials) (domain.User, error) {
	realUser, err := uc.repo.GetByEmail(credentials.Email)
	if err != nil {
		return domain.User{}, err
	}
	credentials.Password = generatePasswordHash(credentials.Password)
	if realUser.Password != credentials.Password {
		return domain.User{}, domain.ErrWrongCredentials
	}
	return realUser, nil
}

func (uc *User) GetByID(id uint64) (domain.User, error) {
	return uc.repo.GetByID(id)
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func validateCredentials(credentials domain.UserCredentials) error {
	if incorrentPasswordRegex.MatchString(credentials.Password) {
		return domain.ErrNotValidPassword
	}
	if !emailRegex.MatchString(credentials.Email) {
		return domain.ErrNotValidEmail
	}
	return nil
}
