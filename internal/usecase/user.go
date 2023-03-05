package usecase

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

const salt = "hjqrhjqw124617ajfhajs"

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

var ErrWrongCredentials = errors.New("wrong credentials")

type UserRepository interface {
	Add(user domain.UserCredentials) (domain.User, error)
	GetAll() ([]domain.User, error)
	GetByEmail(email string) (domain.User, error)
	GetById(id uint64) (domain.User, error)
}

type UserUseCase struct {
	repo UserRepository
}

const maxLenPassword = 30
const minLenPassword = 3

func ValidateCredentials(data domain.UserCredentials) error {
	if data.Email == "" || data.Password == "" {
		return fmt.Errorf("password or email is empty")
	}
	if len([]rune(data.Password)) < minLenPassword || len([]rune(data.Password)) > maxLenPassword {
		return fmt.Errorf("password length is incorrect")
	}
	if !emailRegex.MatchString(data.Email) {
		return fmt.Errorf("mail not validated")
	}
	return nil
}

func NewUser(repo UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) RegisterUser(credentials domain.UserCredentials) (domain.User, error) {
	if err := ValidateCredentials(credentials); err != nil {
		log.Printf("data has not been validated: %s", err)
		return domain.User{}, err
	}
	credentials.Password = generatePasswordHash(credentials.Password)
	return uc.repo.Add(credentials)
}

func (uc *UserUseCase) AuthUser(credentials domain.UserCredentials) (domain.User, error) {
	realUser, err := uc.repo.GetByEmail(credentials.Email)
	if err != nil {
		return domain.User{}, err
	}
	credentials.Password = generatePasswordHash(credentials.Password)
	if realUser.Password != credentials.Password {
		return domain.User{}, ErrWrongCredentials
	}
	return realUser, nil
}

func (uc *UserUseCase) GetById(id uint64) (domain.User, error) {
	return uc.repo.GetById(id)
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
