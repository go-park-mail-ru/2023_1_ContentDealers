package search

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

const wordDelimiter = '+'

type Handler struct {
	useCase UseCase
	logger  logging.Logger
}

func NewHandler(useCase UseCase, logger logging.Logger) Handler {
	return Handler{useCase: useCase, logger: logger}
}

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var searchQuery string
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err == nil {
		searchQuery = query.Get("query")
	}

	searchQuery = strings.Join(strings.FieldsFunc(searchQuery, func(r rune) bool {
		return r == wordDelimiter
	}), " ")

	search, err := h.useCase.Search(r.Context(), searchQuery)

	searchResponse := searchDTO{}
	err = dto.Map(&searchResponse, search)
	if err != nil {
		h.logger.Trace(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	response, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"search": searchResponse,
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
