package user

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

const (
	shortFormDate = "2006-Jan-02"
)

// SignUp TODO: SignUp принимает только json данные для регистрации
// аватар устанавливается или обновляется другой ручкой
// если аватар нужно устанавливать при регистрации, фроненд вызовет отдельную ручку

// @Summary SignUp
// @Tags auth
// @Description Создать аккаунт
// @Accept  json
// @Produce  json
// @Param input body userCreateDTO true "Информация об аккаунте"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /user/signup [post]
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	userCreate := userCreateDTO{}
	err := decoder.Decode(&userCreate)
	if err != nil {
		h.logger.Tracef("failed to parse json string from the body: %w", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"failed to parse json string from the body"}`)
		return
	}

	birthdayTime, err := time.Parse(shortFormDate, userCreate.DateBirth)
	if err != nil {
		h.logger.Tracef("failed to parse birthday from string to birthdayTime: %w", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"failed to parse birthday from string to birthdayTime"}`)
		return
	}

	if userCreate.AvatarURL == "" {
		userCreate.AvatarURL = "media/avatars/default_avatar.jpg"
	}
	user := domain.User{
		Email:        userCreate.Email,
		PasswordHash: userCreate.Password,
		AvatarURL:    userCreate.AvatarURL,
		DateBirth:    birthdayTime,
	}

	_, err = h.userUseCase.Register(r.Context(), user)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrUserAlreadyExists):
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"status": 1, "message":"user already exists"}`)
		case errors.Is(err, domain.ErrNotValidEmail):
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"status": 2, "message":"email not validated"}`)
		case errors.Is(err, domain.ErrNotValidPassword):
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"status": 3, "message":"password not validated"}`)

		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary SignIn
// @Tags auth
// @Description Войти в аккаунт
// @Accept  json
// @Produce  json
// @Param input body userCredentialsDTO true "Данные для входа в аккаунт"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /user/signin [post]
func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	credentials := userCredentialsDTO{}
	err := decoder.Decode(&credentials)
	if err != nil {
		h.logger.Tracef("failed to parse json string from the body: %w", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"failed to parse json string from the body"}`)
		return
	}

	user := domain.User{
		Email:        credentials.Email,
		PasswordHash: credentials.Password, // no hash
	}

	// TODO: перезаписывание user, стоит ли так делать?
	user, err = h.userUseCase.Auth(r.Context(), user)
	if err != nil {
		// не важно, не валиден пароль или не найден пользователь с почтой
		// ответ: пользователь не найден
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"status": 4, "message":"auth wrong credentials"}`)
		return
	}

	session, err := h.sessionUseCase.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sessionCookie := http.Cookie{
		Name:     "session_id",
		Value:    session.ID.String(),
		Expires:  session.ExpiresAt,
		HttpOnly: true,
		Path:     "/",
		// FIXME: Secure: true,
	}

	http.SetCookie(w, &sessionCookie)
	w.WriteHeader(http.StatusOK)
}

// @Summary Logout
// @Tags auth
// @Description Выйти из аккаунта
// @Description Необходимы куки
// @Description Необходим csrf токен
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 500
// @Router /user/logout [post]
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()
	sessionRaw := ctx.Value("session")
	session, ok := sessionRaw.(domain.Session)
	if !ok {
		h.logger.Trace(domain.ErrSessionInvalid)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := h.sessionUseCase.Delete(session.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sessionCookie := http.Cookie{
		Name:     "session_id",
		Value:    "",
		Expires:  time.Now().AddDate(0, 0, -1),
		HttpOnly: true,
		Path:     "/",
		// FIXME: Secure: true,
	}

	http.SetCookie(w, &sessionCookie)
	w.WriteHeader(http.StatusOK)
}
