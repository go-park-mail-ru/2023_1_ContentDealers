package person

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/repository/domain"
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

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	idRaw := mux.Vars(r)["id"]
	var id uint64

	_, err := fmt.Sscanf(idRaw, "%d", &id)
	if err != nil {
		h.logger.Trace("person id is not numeric: %w", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"person id is not numeric"}`)
		return
	}

	person, err := h.useCase.GetByID(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrRepoNotFound):
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, `{"message":"person not found"}`)
		default:
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	personResponse := personDTO{}
	err = dto.Map(&personResponse, person)
	if err != nil {
		h.logger.Trace(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	response, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"person": personResponse,
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
