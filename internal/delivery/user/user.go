package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user/csrf"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/setup/logger"
)

// TODO: может, имеет смысл для констант ввести префикс ("kNameFormFile", "cNameFormFile")
const (
	nameFormFile = "avatar"
	// 10Mb (32Mb по умолчанию)
	maxSizeBody = 10 << 20
)

type Handler struct {
	userUseCase    UserUseCase
	sessionUseCase SessionUseCase
	cryptToken     csrf.CryptToken
	logger         logger.Logger
}

func NewHandler(user UserUseCase, session SessionUseCase, cryptToken csrf.CryptToken, logger logger.Logger) Handler {
	return Handler{
		userUseCase:    user,
		sessionUseCase: session,
		cryptToken:     cryptToken,
	}
}

func (h *Handler) UpdateAvatar(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()
	sessionRaw := ctx.Value("session")
	session, ok := sessionRaw.(domain.Session)
	if !ok {
		h.logger.Error("can't update avatar without session")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := h.userUseCase.GetByID(session.UserID)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxSizeBody)
	file, header, err := r.FormFile(nameFormFile)
	if err != nil {
		h.logger.Error(err)
		if errors.As(err, new(*http.MaxBytesError)) {
			io.WriteString(w, fmt.Sprintf(`{"message":"the size exceeded the maximum size equal to %d mb"}`, maxSizeBody))
		} else {
			io.WriteString(w, `{"message":"failed to parse avatar file from the body"}`)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()
	// TODO: здесь ведь нужно закрывать? это не ответственность репозитория? его зона ответственности - просто сохранить?

	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"avatar file can't be read"}`)
	}
	filetype := http.DetectContentType(buff)

	// TODO: можно еще проверить расширение header.filename
	if header.Header["Content-Type"][0] != "image/jpeg" || filetype != "image/jpeg" {
		h.logger.Error("avatar does not have type: image/jpeg")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"avatar does not have type: image/jpeg"}`)
	}

	file.Seek(0, io.SeekStart)

	_, err = h.userUseCase.UpdateAvatar(user, file)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Info(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()
	sessionRaw := ctx.Value("session")
	session, ok := sessionRaw.(domain.Session)
	if !ok {
		h.logger.Error("can't get info without session")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := h.userUseCase.GetByID(session.UserID)
	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.Email = html.EscapeString(user.Email)

	response, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"user": map[string]string{
				"email": user.Email,
			},
		},
	})

	if err != nil {
		h.logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
