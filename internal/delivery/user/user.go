package user

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	domainSession "github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/domain"
	domainUser "github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/domain"
	"golang.org/x/crypto/argon2"
)

// TODO: может, имеет смысл для констант ввести префикс ("kNameFormFile", "cNameFormFile")
const buffSize = 512

type Handler struct {
	userGateway    UserGateway
	userUsecase    UserUsecase
	sessionGateway SessionGateway
	logger         logging.Logger
	avatarCfg      AvatarConfig
}

func NewHandler(user UserGateway, userUsecase UserUsecase, session SessionGateway, logger logging.Logger, avatarCfg AvatarConfig) *Handler {
	return &Handler{
		userGateway:    user,
		userUsecase:    userUsecase,
		sessionGateway: session,
		logger:         logger,
		avatarCfg:      avatarCfg,
	}
}

func (h *Handler) DeleteAvatar(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()
	sessionRaw := ctx.Value("session")
	session, ok := sessionRaw.(domainSession.Session)
	if !ok {
		h.logger.WithRequestID(ctx).Trace(domainSession.ErrSessionInvalid)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := h.userGateway.GetByID(ctx, session.UserID)
	if err != nil {
		// domain.ErrUserNotFound
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.userGateway.DeleteAvatar(ctx, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateAvatar(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()
	sessionRaw := ctx.Value("session")
	session, ok := sessionRaw.(domainSession.Session)
	if !ok {
		h.logger.WithRequestID(ctx).Trace(domainSession.ErrSessionInvalid)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := h.userGateway.GetByID(ctx, session.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// nolint:gomnd
	maxSizeBody := int64(h.avatarCfg.MaxSizeBody) << 20
	r.Body = http.MaxBytesReader(w, r.Body, maxSizeBody)
	file, _, err := r.FormFile(h.avatarCfg.NameFormFile)
	if err != nil {
		if errors.As(err, new(*http.MaxBytesError)) {
			h.logger.WithRequestID(ctx).Tracef("the size exceeded the maximum size equal to %d mb: %v", maxSizeBody, err)

			w.WriteHeader(http.StatusRequestEntityTooLarge)
			io.WriteString(w, `{"status": 5, "message":"the size exceeded the maximum size equal to %d mb"}`)
			// для совместимости с nginx
			return
		}
		h.logger.WithRequestID(ctx).Tracef("failed to parse avatar file from the body: %v", err)
		io.WriteString(w, `{"message":"failed to parse avatar file from the body"}`)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	buff := make([]byte, buffSize)
	_, err = file.Read(buff)
	if err != nil {
		h.logger.WithRequestID(ctx).Tracef("avatar file can't be read: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"avatar file can't be read"}`)
	}
	filetype := http.DetectContentType(buff)

	// TODO: можно еще проверить расширение header.filename
	if filetype != "image/jpeg" {
		h.logger.WithRequestID(ctx).Trace("avatar does not have type: image/jpeg")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"status": 6, "message":"avatar does not have type: image/jpeg"}`)
		return
	}

	file.Seek(0, io.SeekStart)

	_, err = h.userGateway.UpdateAvatar(ctx, user, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Info(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()
	sessionRaw := ctx.Value("session")
	session, ok := sessionRaw.(domainSession.Session)
	if !ok {
		h.logger.WithRequestID(ctx).Trace(domainSession.ErrSessionInvalid)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := h.userGateway.GetByID(ctx, session.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.Email = html.EscapeString(user.Email)

	hasSub := user.SubscriptionExpiryDate.After(time.Now())

	response, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"user": map[string]interface{}{
				"email":          user.Email,
				"avatar_url":     user.AvatarURL,
				"hasSub":         hasSub,
				"sub_expiration": user.SubscriptionExpiryDate.Format("2006-01-02"),
			},
		},
	})

	if err != nil {
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()

	sessionRaw := ctx.Value("session")
	session, ok := sessionRaw.(domainSession.Session)
	if !ok {
		h.logger.WithRequestID(ctx).Trace(domainSession.ErrSessionInvalid)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := h.userGateway.GetByID(ctx, session.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(r.Body)
	userUpdate := userDTO{}
	err = decoder.Decode(&userUpdate)
	if err != nil {
		h.logger.WithRequestID(ctx).Tracef("failed to parse json string from the body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"failed to parse json string from the body"}`)
		return
	}

	user.Email = userUpdate.Email
	user.PasswordHash = userUpdate.Password

	err = h.userGateway.Update(ctx, user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if errors.Is(err, domainUser.ErrUserAlreadyExists) {
			io.WriteString(w, `{"status": 7, "message":"user with this email already exists"}`)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) PasswordValidate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()

	sessionRaw := ctx.Value("session")
	session, ok := sessionRaw.(domainSession.Session)
	if !ok {
		h.logger.WithRequestID(ctx).Trace(domainSession.ErrSessionInvalid)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := h.userGateway.GetByID(ctx, session.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(r.Body)
	userPass := userPasswordDTO{}
	err = decoder.Decode(&userPass)
	if err != nil {
		h.logger.WithRequestID(ctx).Tracef("failed to parse json string from the body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"failed to parse json string from the body"}`)
		return
	}

	isValid, err := verifyPassword(ctx, userPass.Password, user.PasswordHash)
	if err != nil {
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !isValid {
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"password mismatch"}`)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// TODO: Для ускорения написания ручки

const (
	saltSize    = 16
	hashSize    = 32
	iterations  = 1
	memory      = 64 * 1024
	parallelism = 4
)

func verifyPassword(ctx context.Context, password, encodedPassword string) (bool, error) {
	if len(encodedPassword) <= saltSize {
		err := fmt.Errorf("Invalid encoded password format")
		return false, err
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
