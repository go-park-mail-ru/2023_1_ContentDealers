package rating

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
	useCase RatingUseCase
	logger  logging.Logger
}

func NewHandler(useCase RatingUseCase, logger logging.Logger) *Handler {
	return &Handler{
		useCase: useCase,
		logger:  logger,
	}
}

func (h *Handler) DeleteRating(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	ratingDTO := RatingDTO{}
	err := decoder.Decode(&ratingDTO)
	if err != nil {
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rating := domain.Rating{
		ContentID: ratingDTO.ContentID,
	}

	err = h.useCase.Delete(ctx, rating)
	if err != nil {
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) AddRating(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	ratingDTO := RatingDTO{}
	err := decoder.Decode(&ratingDTO)
	if err != nil {
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if ratingDTO.Rating == 0 {
		msg := "rating not found to add"
		h.logger.WithRequestID(ctx).Trace(msg)
		io.WriteString(w, fmt.Sprintf(`{"message":"%s"}`, msg))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rating := domain.Rating{
		ContentID: ratingDTO.ContentID,
		Rating:    ratingDTO.Rating,
	}

	newRating, countRatings, err := h.useCase.Add(ctx, rating)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"new_rating":    newRating,
			"count_ratings": countRatings,
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

func (h *Handler) HasRating(w http.ResponseWriter, r *http.Request) {
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

	rating := domain.Rating{
		ContentID: id,
	}

	hasRating, err := h.useCase.Has(ctx, rating)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"has":    hasRating.HasRating,
			"rating": hasRating.Rating.Rating,
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

func (h *Handler) GetRatingByUser(w http.ResponseWriter, r *http.Request) {
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

	options := domain.RatingsOptions{
		SortDate: order,
		Limit:    uint32(limit),
		Offset:   uint32(offset),
	}

	content, ratings, isLast, err := h.useCase.GetByUser(ctx, options)
	if err != nil {
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	contentResponse := []contentRatingDTO{}

	for idx, rate := range ratings {
		// nolint:govet
		var contentDTO contentDTO
		err = dto.Map(&contentDTO, content[idx])
		if err != nil {
			h.logger.Trace(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		contentResponse = append(contentResponse, contentRatingDTO{
			Content: contentDTO,
			Rating:  rate.Rating,
		})
	}

	response, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"content_with_ratings": contentResponse,
			"is_last":              isLast,
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
