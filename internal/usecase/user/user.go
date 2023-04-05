package user

import (
	"io"
	"regexp"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

// TODO: нужно изменить регулярки
var (
	emailRegex             = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	incorrentPasswordRegex = regexp.MustCompile(`(^(.{0,7}|[^0-9]*|[^A-Z]*|[^a-z]*|[a-zA-Z0-9]*)$)`)
)

type User struct {
	repo UserRepository
}

func NewUser(repo UserRepository) *User {
	return &User{repo: repo}
}

func (uc *User) Register(user domain.User) (domain.User, error) {
	if err := validateCredentials(user); err != nil {
		return domain.User{}, err
	}
	passwordHash, err := hashPassword(user.PasswordHash)
	if err != nil {
		return domain.User{}, err
	}
	user.PasswordHash = passwordHash
	return uc.repo.Add(user)
}

func (uc *User) Auth(user domain.User) (domain.User, error) {
	realUser, err := uc.repo.GetByEmail(user.Email)
	if err != nil {
		return domain.User{}, err
	}
	isVaild, err := verifyPassword(user.PasswordHash, realUser.PasswordHash)
	if err != nil {
		return domain.User{}, err
	}
	if !isVaild {
		return domain.User{}, domain.ErrWrongCredentials
	}
	return realUser, nil
}

func (uc *User) GetByID(id uint64) (domain.User, error) {
	return uc.repo.GetByID(id)
}

func (uc *User) UpdateAvatar(user domain.User, file io.Reader) (domain.User, error) {
	return uc.repo.UpdateAvatar(user, file)
}

func validateCredentials(credentials domain.User) error {
	// TODO: регулярки нужно изменить (снизить строгость пароля)
	return nil
	// if incorrentPasswordRegex.MatchString(credentials.Password) {
	// 	return domain.ErrNotValidPassword
	// }
	// if !emailRegex.MatchString(credentials.Email) {
	// 	return domain.ErrNotValidEmail
	// }
	// return nil
}
