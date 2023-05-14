package history_views

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/dranikpg/dto-mapper"
	contentDomain "github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/domain"
	"github.com/gorilla/mux"
)

const (
	defaultLimit  = 15
	defaultOffset = 0
)

type Handler struct {
	useCase ViewsUseCase
	logger  logging.Logger
}

func NewHandler(useCase ViewsUseCase, logger logging.Logger) *Handler {
	return &Handler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h *Handler) UpdateProgressView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	viewDTO := viewDTO{}
	err := decoder.Decode(&viewDTO)
	if err != nil {
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	stopView, err := time.ParseDuration(viewDTO.StopView)
	if err != nil {
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	duration, err := time.ParseDuration(viewDTO.Duration)
	if err != nil {
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	view := domain.View{
		ContentID: viewDTO.ContentID,
		StopView:  stopView,
		Duration:  duration,
	}

	err = h.useCase.UpdateProgressView(ctx, view)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) HasView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer r.Body.Close()

	idRaw := mux.Vars(r)["id"]
	var id uint64

	_, err := fmt.Sscanf(idRaw, "%d", &id)
	if err != nil {
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"content id is not numeric"}`)
		return
	}

	view := domain.View{
		ContentID: id,
	}

	hasView, err := h.useCase.HasView(ctx, view)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	viewDTO := viewDTO{
		UserID:    hasView.View.UserID,
		ContentID: hasView.View.ContentID,
		StopView:  hasView.View.StopView.String(),
		Duration:  hasView.View.Duration.String(),
	}

	response, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"has":  hasView.HasView,
			"view": viewDTO,
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

// чекает query string (all, partially)
func (h *Handler) GetViewsByUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer r.Body.Close()

	// query string

	var limit int32 = defaultLimit
	var offset int32 = defaultOffset
	var order string
	var typeReq string

	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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
	// пустая строка будет обработана
	order = query.Get("order")
	typeReq = query.Get("type")

	options := domain.ViewsOptions{
		SortDate: order,
		Limit:    uint32(limit),
		Offset:   uint32(offset),
	}

	var content []contentDomain.Content
	var isLast bool
	if typeReq == "" || typeReq == "all" {
		content, isLast, err = h.useCase.GetAllViewsByUser(ctx, options)
	} else if typeReq == "part" {
		content, isLast, err = h.useCase.GetPartiallyViewsByUser(ctx, options)
	}
	if err != nil {
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contentResponse := []contentDTO{}
	err = dto.Map(&contentResponse, content)
	if err != nil {
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"content": contentResponse,
			"is_last": isLast,
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
