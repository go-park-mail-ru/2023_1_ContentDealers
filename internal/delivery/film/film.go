package film

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"

	"github.com/gorilla/mux"
)

type Handler struct {
	useCase UseCase
	logger  logging.Logger
}

func NewHandler(useCase UseCase, logger logging.Logger) Handler {
	return Handler{useCase: useCase, logger: logger}
}

func (h *Handler) GetByContentID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer r.Body.Close()

	idRaw := mux.Vars(r)["content_id"]
	var id uint64

	_, err := fmt.Sscanf(idRaw, "%d", &id)
	if err != nil {
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"film id is not numeric"}`)
		return
	}

	film, err := h.useCase.GetByContentID(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrRepoNotFound):
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, `{"message":"film not found"}`)
		default:
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	filmResponse := filmDTO{}
	err = dto.Map(&filmResponse, film)
	if err != nil {
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	response, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"film": filmResponse,
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
