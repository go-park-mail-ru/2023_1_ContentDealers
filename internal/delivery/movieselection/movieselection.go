package movieselection

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/gorilla/mux"
)

type Handler struct {
	useCase MovieSelectionUseCase
	logger  logging.Logger
}

func NewHandler(useCase MovieSelectionUseCase, logger logging.Logger) Handler {
	return Handler{useCase: useCase, logger: logger}
}

// @Summary Get All
// @Tags movieselection
// @Description Получить все подборки
// @Description Через параметры можно указать лимит фильмов / сериалов в одной подборке
// @Produce  json
// @Param limit query integer false "Ограничение на количество"
// @Param offset query integer false "Смешение от первого фильма"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /selections [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	selections, err := h.useCase.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"movie_selections": selections,
		},
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// @Summary Get By Id
// @Tags movieselection
// @Description Получить конкретную подборку
// @Description Через параметры можно указать лимит фильмов / сериалов в одной подборке
// @Produce  json
// @Param limit query integer false "Ограничение на количество"
// @Param offset query integer false "Смешение от первого фильма"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /selections/id [get]
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	idRaw := mux.Vars(r)["id"]
	var id uint64

	_, err := fmt.Sscanf(idRaw, "%d", &id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"movie selection id is not numeric"}`)
		return
	}

	movieSelection, err := h.useCase.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrMovieSelectionNotFound):
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, `{"message":"movie selection not found"}`)
		default:
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	response, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"selection": movieSelection,
		},
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
