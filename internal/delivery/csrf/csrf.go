package csrf

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/domain"
)

type Handler struct {
	csrfUseCase UseCase
	logger      logging.Logger
	expiresAt   time.Duration
}

func NewHandler(csrfUseCase UseCase, logger logging.Logger, cfg CSRFConfig) *Handler {
	return &Handler{
		csrfUseCase: csrfUseCase,
		logger:      logger,
		expiresAt:   time.Second * time.Duration(cfg.ExpiresAt),
	}
}

func (h *Handler) GetCSRF(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sessionRaw := ctx.Value("session")
	session, ok := sessionRaw.(domain.Session)
	if !ok {
		h.logger.WithRequestID(ctx).Trace(domain.ErrSessionInvalid)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	token, err := h.csrfUseCase.Create(r.Context(), session, time.Now().Add(h.expiresAt).Unix())
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
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
