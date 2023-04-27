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
	domainContent "github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
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

func (h *Handler) GetContentByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var limit int32 = defaultLimit
	var offset int32 = defaultOffset

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

	idRaw := mux.Vars(r)["id"]
	var id uint64

	_, err = fmt.Sscanf(idRaw, "%d", &id)
	if err != nil {
		h.logger.Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"film selection id is not numeric"}`)
		return
	}

	genre, err := h.useCase.GetContentByID(r.Context(), domainContent.ContentFilter{
		ID:     id,
		Limit:  0,
		Offset: 0,
	})
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

	genreResponse := genreDTO{}
	err = dto.Map(&genreResponse, genre)
	if err != nil {
		h.logger.Trace(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	response, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"genre": genreResponse,
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
