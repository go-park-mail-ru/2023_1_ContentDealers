package user

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"regexp"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

const salt = "hjqrhjqw124617ajfhajs"

// TODO: нужно изменить регулярки
var (
	emailRegex             = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	incorrentPasswordRegex = regexp.MustCompile(`(^(.{0,7}|[^0-9]*|[^A-Z]*|[^a-z]*|[a-zA-Z0-9]*)$)`)
)

type User struct {
	repo Repository
}

func NewUser(repo Repository) *User {
	return &User{repo: repo}
}

func (uc *User) Register(ctx context.Context, user domain.User) (domain.User, error) {
	if err := validateCredentials(user); err != nil {
		return domain.User{}, err
	}
	user.PasswordHash = generatePasswordHash(user.PasswordHash)
	return uc.repo.Add(ctx, user)
}

func (uc *User) Auth(ctx context.Context, user domain.User) (domain.User, error) {
	realUser, err := uc.repo.GetByEmail(ctx, user.Email)
	if err != nil {
		return domain.User{}, err
	}
	// кажется, можно сделать красивей
	user.PasswordHash = generatePasswordHash(user.PasswordHash)
	if realUser.PasswordHash != user.PasswordHash {
		return domain.User{}, domain.ErrWrongCredentials
	}
	return realUser, nil
}

func (uc *User) GetByID(ctx context.Context, id uint64) (domain.User, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *User) UpdateAvatar(ctx context.Context, user domain.User, file io.Reader) (domain.User, error) {
	return uc.repo.UpdateAvatar(ctx, user, file)
}

func (uc *User) DeleteAvatar(ctx context.Context, user domain.User) error {
	return uc.repo.DeleteAvatar(ctx, user)
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
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
