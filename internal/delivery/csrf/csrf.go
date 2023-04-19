package csrf

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/repository/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

const ExpirationTimeCSRF = 2 * time.Hour

type Handler struct {
	csrfUseCase CSRFUseCase
	logger      logging.Logger
}

func NewHandler(csrfUseCase CSRFUseCase, logger logging.Logger) Handler {
	return Handler{
		csrfUseCase: csrfUseCase,
		logger:      logger,
	}
}

// @Summary CSRF
// @Tags user
// @Description Получить CSRF токен
// @Description Необходимы куки
// @Produce  json
// @Success 200 {object} tokenDTO
// @Failure 400
// @Failure 500
// @Router /user/csrf [get]
func (h *Handler) GetCSRF(w http.ResponseWriter, r *http.Request) {
	sessionRaw := r.Context().Value("session")
	session, ok := sessionRaw.(domain.Session)
	if !ok {
		h.logger.Trace(domain.ErrSessionInvalid)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	token, err := h.csrfUseCase.Create(session, time.Now().Add(ExpirationTimeCSRF).Unix())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"csrf-token": token,
		},
	})
	if err != nil {
		h.logger.Trace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
