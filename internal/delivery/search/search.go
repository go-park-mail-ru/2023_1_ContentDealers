package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

const (
	wordDelimiter = '+'
	defaultLimit  = 6
	defaultOffset = 0
)

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
	var limit int32 = defaultLimit
	var offset int32 = defaultOffset

	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		h.logger.WithRequestID(r.Context()).Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	searchQuery = query.Get("query")

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

	slug := query.Get("slug")

	searchQuery = strings.Join(strings.FieldsFunc(searchQuery, func(r rune) bool {
		return r == wordDelimiter
	}), " ")

	search, err := h.useCase.Search(r.Context(), domain.SearchQuery{
		Query:      searchQuery,
		TargetSlug: slug,
		Limit:      uint32(limit),
		Offset:     uint32(offset),
	})

	if err != nil {
		h.logger.Trace(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

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
