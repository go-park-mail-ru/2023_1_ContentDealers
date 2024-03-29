package favorites

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/domain"
	"github.com/gorilla/mux"
)

const (
	defaultLimit  = 15
	defaultOffset = 0
)

type Handler struct {
	useCase FavContentUseCase
	logger  logging.Logger
}

func NewHandler(useCase FavContentUseCase, logger logging.Logger) *Handler {
	return &Handler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h *Handler) DeleteFavContent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	favContentDTO := FavoriteContentDTO{}
	err := decoder.Decode(&favContentDTO)
	if err != nil {
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	favContent := domain.FavoriteContent{
		ContentID: favContentDTO.ContentID,
	}

	err = h.useCase.Delete(ctx, favContent)
	if err != nil {
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) AddFavContent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	favContentDTO := FavoriteContentDTO{}
	err := decoder.Decode(&favContentDTO)
	if err != nil {
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	favContent := domain.FavoriteContent{
		ContentID: favContentDTO.ContentID,
	}

	err = h.useCase.Add(ctx, favContent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) HasFavContent(w http.ResponseWriter, r *http.Request) {
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

	favContent := domain.FavoriteContent{
		ContentID: id,
	}

	isFav, err := h.useCase.HasFav(ctx, favContent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"has": isFav,
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

func (h *Handler) GetFavContent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer r.Body.Close()

	// query string

	var limit int32 = defaultLimit
	var offset int32 = defaultOffset
	var order string

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

	options := domain.FavoritesOptions{
		SortDate: order,
		Limit:    uint32(limit),
		Offset:   uint32(offset),
	}

	content, isLast, err := h.useCase.Get(ctx, options)
	if err != nil {
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var contentResponse []contentDTO
	err = dto.Map(&contentResponse, content)
	if err != nil {
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusInternalServerError)
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
