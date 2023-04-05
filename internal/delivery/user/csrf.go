package user

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

const ExpirationTimeCSRF = 2 * time.Hour

func (h *Handler) GetCSRF(w http.ResponseWriter, r *http.Request) {
	sessionRaw := r.Context().Value("session")
	session, ok := sessionRaw.(domain.Session)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	token, err := h.cryptToken.Create(session, time.Now().Add(ExpirationTimeCSRF).Unix())
	if err != nil {
		// log "csrf token creation error"
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"csrf-token": token,
		},
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}