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
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

type CSRF struct {
	Secret []byte
	logger logging.Logger
}

func NewCSRF(secret string, logger logging.Logger) (*CSRF, error) {
	tmp := []byte(secret)
	// проверка секрета на валидность (длина секрета 16, 24, 32 байта)
	_, err := aes.NewCipher(tmp)
	if err != nil {
		return &CSRF{}, fmt.Errorf("cypher problem %v", err)
	}
	return &CSRF{Secret: tmp, logger: logger}, nil
}

type tokenData struct {
	SessionID string
	UserID    uint64
	Exp       int64
}

func (c *CSRF) Create(s domain.Session, tokenExpTime int64) (string, error) {
	block, err := aes.NewCipher(c.Secret)
	if err != nil {
		c.logger.Trace(err)
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		c.logger.Trace(err)
		return "", err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		c.logger.Trace(err)
		return "", err
	}

	td := &tokenData{SessionID: s.ID.String(), UserID: s.UserID, Exp: tokenExpTime}
	data, err := json.Marshal(td)
	if err != nil {
		c.logger.Trace(err)
		return "", err
	}
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
		c.logger.Trace(err)
		return false, err
	}
	// GCM - режим аутентифицированного шифрования
	// на основе tk.Secret генерируется ключ аутентификации
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		c.logger.Trace(err)
		return false, err
	}
	ciphertext, err := base64.StdEncoding.DecodeString(inputToken)
	if err != nil {
		c.logger.Trace(err)
		return false, err
	}
	// nonce - случайное число для защиты от повторных атак
	// длина aes nonce равна 12 байтам
	nonceSize := aesgcm.NonceSize()
	if len(ciphertext) < nonceSize {
		err := fmt.Errorf("ciphertext too short")
		c.logger.Trace(err)
		return false, err
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		c.logger.Trace(err)
		return false, err
	}

	td := tokenData{}
	err = json.Unmarshal(plaintext, &td)
	if err != nil {
		c.logger.Trace(err)
		return false, err
	}

	if td.Exp < time.Now().Unix() {
		err := fmt.Errorf("token expired")
		c.logger.Trace(err)
		return false, err
	}

	expected := tokenData{SessionID: s.ID.String(), UserID: s.UserID}
	td.Exp = 0
	return td == expected, nil
}
