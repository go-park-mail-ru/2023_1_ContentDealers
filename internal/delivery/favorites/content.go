package favorites

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/favorites/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
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

func (h *Handler) GetFavContent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer r.Body.Close()

	var order string

	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// пустая строка будет обработана
	order = query.Get("order")

	options := domain.FavoritesOptions{
		SortDate: order,
	}

	content, err := h.useCase.Get(ctx, options)
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
		},
	})
	if err != nil {
		h.logger.Trace(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)

	// favsContent, err := h.useCase.Get(ctx, options)
	// if err != nil {
	// 	h.logger.WithRequestID(ctx).Trace(err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	// var favContentResponse []FavoriteContentDTO
	// err = dto.Map(&favContentResponse, favsContent)
	// if err != nil {
	// 	h.logger.WithRequestID(ctx).Trace(err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// }

	// response, err := json.Marshal(map[string]interface{}{
	// 	"body": map[string]interface{}{
	// 		"favcontent": favContentResponse,
	// 	},
	// })
	// if err != nil {
	// 	h.logger.WithRequestID(ctx).Trace(err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	// w.WriteHeader(http.StatusOK)
	// w.Write(response)
}
