package main

import (
	"fmt"

	"github.com/dlclark/regexp2"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

var passwordRegexp = regexp2.MustCompile(`^(?=.*[a-zA-Z])(?=.*[0-9])[a-zA-Z0-9]{8,}$`, 0)
var emailRegexp = regexp2.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, 0)

func validateCredentials(credentials domain.User) error {
	validPass, err := passwordRegexp.MatchString(credentials.PasswordHash)
	if err != nil {
		return err
	}
	fmt.Println("pass", validPass)
	validEmail, err := emailRegexp.MatchString(credentials.Email)
	if err != nil {
		return err
	}
	fmt.Println("email", validEmail)
	return nil
}

func main() {
	pass := "3sdffsf4"
	email := "1@m.ru"
	validateCredentials(domain.User{
		PasswordHash: pass,
		Email:        email,
	})
}
