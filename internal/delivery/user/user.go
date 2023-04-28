package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	domainSession "github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/domain"
	domainUser "github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/domain"
)

// TODO: может, имеет смысл для констант ввести префикс ("kNameFormFile", "cNameFormFile")
const buffSize = 512

type Handler struct {
	userGateway    UserGateway
	sessionUseCase SessionUseCase
	logger         logging.Logger
	avatarCfg      AvatarConfig
}

func NewHandler(user UserGateway, session SessionUseCase, logger logging.Logger, avatarCfg AvatarConfig) Handler {
	return Handler{
		userGateway:    user,
		sessionUseCase: session,
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

	maxSizeBody := int64(h.avatarCfg.MaxSizeBody) << 20
	r.Body = http.MaxBytesReader(w, r.Body, maxSizeBody)
	file, header, err := r.FormFile(h.avatarCfg.NameFormFile)
	if err != nil {
		if errors.As(err, new(*http.MaxBytesError)) {
			h.logger.WithRequestID(ctx).Tracef("the size exceeded the maximum size equal to %d mb: %w", maxSizeBody, err)

			io.WriteString(w, fmt.Sprintf(`{"status": 5, "message":"the size exceeded the maximum size equal to %d mb"}`, maxSizeBody))
			// для совместимости с nginx
			w.WriteHeader(http.StatusRequestEntityTooLarge)
			return
		} else {
			h.logger.WithRequestID(ctx).Tracef("failed to parse avatar file from the body: %w", err)
			io.WriteString(w, `{"message":"failed to parse avatar file from the body"}`)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	buff := make([]byte, buffSize)
	_, err = file.Read(buff)
	if err != nil {
		h.logger.WithRequestID(ctx).Tracef("avatar file can't be read: %w", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"avatar file can't be read"}`)
	}
	filetype := http.DetectContentType(buff)

	// TODO: можно еще проверить расширение header.filename
	if header.Header["Content-Type"][0] != "image/jpeg" || filetype != "image/jpeg" {
		h.logger.WithRequestID(ctx).Trace("avatar does not have type: image/jpeg")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"status": 6, "message":"avatar does not have type: image/jpeg"}`)
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

	response, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"user": map[string]string{
				"email":      user.Email,
				"avatar_url": user.AvatarURL,
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
	userUpdate := userUpdateDTO{}
	err = decoder.Decode(&userUpdate)
	if err != nil {
		h.logger.WithRequestID(ctx).Tracef("failed to parse json string from the body: %w", err)
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
