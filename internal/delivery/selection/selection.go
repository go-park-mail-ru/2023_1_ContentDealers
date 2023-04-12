package selection

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

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

const (
	defaultLimit  = 15
	defaultOffset = 0
)

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var limit int = defaultLimit
	var offset int = defaultOffset

	query, err := url.ParseQuery(r.URL.RawQuery)
	if err == nil {
		limitRaw := query.Get("limit")
		_, err = fmt.Sscanf(limitRaw, "%d", &limit)
		if err != nil || limit <= 0 {
			limit = defaultLimit
		}

		offsetRaw := query.Get("offset")
		_, err = fmt.Sscanf(offsetRaw, "%d", &offset)
		if err != nil || offset < 0 {
			offset = defaultOffset
		}
	}

	selections, err := h.useCase.GetAll(r.Context(), uint(limit), uint(offset))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var selectionsResponse []selectionDTO
	err = dto.Map(&selectionsResponse, selections)
	if err != nil {
		h.logger.Trace(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	response, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"selections": selectionsResponse,
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

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	idRaw := mux.Vars(r)["id"]
	var id uint64

	_, err := fmt.Sscanf(idRaw, "%d", &id)
	if err != nil {
		h.logger.Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"film selection id is not numeric"}`)
		return
	}

	selection, err := h.useCase.GetByID(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrRepoNotFound):
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, `{"message":"film selection not found"}`)
		default:
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	selectionResponse := selectionDTO{}
	err = dto.Map(&selectionResponse, selection)
	if err != nil {
		h.logger.Trace(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	response, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"selection": selectionResponse,
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
