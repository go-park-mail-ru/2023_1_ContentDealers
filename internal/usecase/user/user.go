package user

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/dlclark/regexp2"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

var (
	passwordRegexp = regexp2.MustCompile(`^(?=.*[a-zA-Z])(?=.*[0-9])[a-zA-Z0-9]{8,}$`, 0)
	emailRegexp    = regexp2.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, 0)
)

type User struct {
	repo   Repository
	logger logging.Logger
}

func NewUser(repo Repository, logger logging.Logger) *User {
	return &User{repo: repo, logger: logger}
}

func (uc *User) Register(ctx context.Context, user domain.User) (domain.User, error) {
	log.Println("register")
	err := validateCredentials(user)
	if err != nil {
		log.Println("1")
		uc.logger.Trace(err)
		return domain.User{}, err
	}
	log.Println("3")
	passwordHash, err := uc.hashPassword(user.PasswordHash)
	if err != nil {
		return domain.User{}, err
	}
	user.PasswordHash = passwordHash
	return uc.repo.Add(ctx, user)
}

func (uc *User) Auth(ctx context.Context, user domain.User) (domain.User, error) {
	realUser, err := uc.repo.GetByEmail(ctx, user.Email)
	if err != nil {
		// может быть domain.ErrUserNotFound
		return domain.User{}, err
	}
	isVaild, err := uc.verifyPassword(user.PasswordHash, realUser.PasswordHash)
	if err != nil {
		return domain.User{}, domain.ErrWrongCredentials
	}
	if !isVaild {
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

func (uc *User) Update(ctx context.Context, user domain.User) error {
	if user.PasswordHash == "" {
		userTmp, err := uc.repo.GetByID(ctx, user.ID)
		if err != nil {
			return err
		}
		// оставляем тот же пароль
		user.PasswordHash = userTmp.PasswordHash
	} else {
		passwordHashTmp, err := uc.hashPassword(user.PasswordHash)
		if err != nil {
			return err
		}
		user.PasswordHash = passwordHashTmp
	}
	return uc.repo.Update(ctx, user)
}

func validateCredentials(credentials domain.User) error {
	validPass, err := passwordRegexp.MatchString(credentials.PasswordHash)
	fmt.Println(validPass)
	if err != nil {
		return err
	}
	if !validPass {
		return domain.ErrNotValidPassword
	}
	validEmail, err := emailRegexp.MatchString(credentials.Email)
	fmt.Println(validEmail)
	if err != nil {
		return err
	}
	if !validEmail {
		return domain.ErrNotValidEmail
	}
	return nil
}
