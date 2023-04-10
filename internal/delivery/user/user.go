package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

// TODO: может, имеет смысл для констант ввести префикс ("kNameFormFile", "cNameFormFile")
const (
	nameFormFile = "avatar"
	// 10Mb (32Mb по умолчанию)
	maxSizeBody = 10 << 20
	buffSize    = 512
)

type Handler struct {
	userUseCase    UserUseCase
	sessionUseCase SessionUseCase
}

func NewHandler(user UserUseCase, session SessionUseCase) Handler {
	return Handler{
		userUseCase:    user,
		sessionUseCase: session,
	}
}

func (h *Handler) DeleteAvatar(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()
	sessionRaw := ctx.Value("session")
	session, ok := sessionRaw.(domain.Session)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := h.userUseCase.GetByID(ctx, session.UserID)
	if err != nil {
		// domain.ErrUserNotFound
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.userUseCase.DeleteAvatar(ctx, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()
	sessionRaw := ctx.Value("session")
	session, ok := sessionRaw.(domain.Session)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := h.userUseCase.GetByID(ctx, session.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxSizeBody)
	file, header, err := r.FormFile(nameFormFile)
	if err != nil {
		log.Println(err)
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

	buff := make([]byte, buffSize)
	_, err = file.Read(buff)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"avatar file can't be read"}`)
	}
	filetype := http.DetectContentType(buff)

	// TODO: можно еще проверить расширение header.filename
	if header.Header["Content-Type"][0] != "image/jpeg" || filetype != "image/jpeg" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"avatar does not have type: image/jpeg"}`)
	}

	file.Seek(0, io.SeekStart)

	_, err = h.userUseCase.UpdateAvatar(ctx, user, file)
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
	session, ok := sessionRaw.(domain.Session)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := h.userUseCase.GetByID(ctx, session.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"user": map[string]string{
				"email":      user.Email,
				"date_birth": user.DateBirth.Format("2006-Jan-02"),
				"avatar_url": user.AvatarURL,
			},
		},
	})

	if err != nil {
		log.Println(err)
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
	session, ok := sessionRaw.(domain.Session)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := h.userUseCase.GetByID(ctx, session.UserID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(r.Body)
	userUpdate := userUpdateDTO{}
	err = decoder.Decode(&userUpdate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"failed to parse json string from the body"}`)
		return
	}

	birthdayTime, err := time.Parse(shortFormDate, userUpdate.Birthday)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"failed to parse birthday from string to birthdayTime"}`)
		return
	}

	user.Email = userUpdate.Email
	user.DateBirth = birthdayTime
	user.PasswordHash = userUpdate.Password

	err = h.userUseCase.Update(ctx, user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
