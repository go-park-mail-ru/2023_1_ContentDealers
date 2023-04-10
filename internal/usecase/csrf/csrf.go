package csrf

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type CSRF struct {
	Secret []byte
}

func NewCSRF(secret string) (CSRF, error) {
	tmp := []byte(secret)
	// проверка секрета на валидность (длина секрета 16, 24, 32 байта)
	_, err := aes.NewCipher(tmp)
	if err != nil {
		return CSRF{}, fmt.Errorf("cypher problem %v", err)
	}
	return CSRF{Secret: tmp}, nil
}

type tokenData struct {
	SessionID string
	UserID    uint64
	Exp       int64
}

func (c *CSRF) Create(s domain.Session, tokenExpTime int64) (string, error) {
	block, err := aes.NewCipher(c.Secret)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	td := &tokenData{SessionID: s.ID.String(), UserID: s.UserID, Exp: tokenExpTime}
	data, _ := json.Marshal(td)
	ciphertext := aesgcm.Seal(nil, nonce, data, nil)

	res := append([]byte(nil), nonce...)
	res = append(res, ciphertext...)

	token := base64.StdEncoding.EncodeToString(res)
	return token, nil
}

func (c *CSRF) Check(s domain.Session, inputToken string) (bool, error) {
	// объект блочного шифра AES
	block, err := aes.NewCipher(c.Secret)
	if err != nil {
		return false, err
	}
	// GCM - режим аутентифицированного шифрования
	// на основе tk.Secret генерируется ключ аутентификации
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return false, err
	}
	ciphertext, err := base64.StdEncoding.DecodeString(inputToken)
	if err != nil {
		return false, err
	}
	// nonce - случайное число для защиты от повторных атак
	// длина aes nonce равна 12 байтам
	nonceSize := aesgcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return false, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return false, fmt.Errorf("decrypt fail: %v", err)
	}

	td := tokenData{}
	err = json.Unmarshal(plaintext, &td)
	if err != nil {
		return false, fmt.Errorf("bad json: %v", err)
	}

	if td.Exp < time.Now().Unix() {
		return false, fmt.Errorf("token expired")
	}

	expected := tokenData{SessionID: s.ID.String(), UserID: s.UserID}
	td.Exp = 0
	return td == expected, nil
}
