package user

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	saltSize    = 16
	hashSize    = 32
	iterations  = 1
	memory      = 64 * 1024
	parallelism = 4
)

func hashPassword(password string) (string, error) {
	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, hashSize)
	encodedSalt := base64.StdEncoding.EncodeToString(salt)
	encodedHash := base64.StdEncoding.EncodeToString(hash)

	// разделитель точка
	return fmt.Sprintf("%s.%s", encodedSalt, encodedHash), nil
}

func verifyPassword(password, encodedPassword string) (bool, error) {
	if len(encodedPassword) <= saltSize {
		return false, fmt.Errorf("Invalid encoded password format")
	}
	tmp := strings.Split(encodedPassword, ".")
	saltString, hashString := tmp[0], tmp[1]
	salt, err := base64.StdEncoding.DecodeString(saltString)
	if err != nil {
		return false, err
	}
	hash, err := base64.StdEncoding.DecodeString(hashString)
	if err != nil {
		return false, err
	}

	expectedHash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, hashSize)
	return bytes.Equal(hash, expectedHash), nil
}
