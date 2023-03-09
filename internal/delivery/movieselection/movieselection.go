package movieselection

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/contract"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/gorilla/mux"
)

type Handler struct {
	useCase contract.MovieSelectionUseCase
}

func NewHandler(useCase contract.MovieSelectionUseCase) Handler {
	return Handler{useCase: useCase}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	selections, err := h.useCase.GetAll()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status":500}`)
		return
	}

	response, err := json.Marshal(map[string]interface{}{
		"status": http.StatusOK,
		"body": map[string]interface{}{
			"movie_selections": selections,
		},
	})

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status":500}`)
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
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"status":400}`)
		return
	}

	movieSelection, err := h.useCase.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrMovieSelectionNotFound):
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, `{"status":404}`)
		default:
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, `{"status":500}`)
		}
		return
	}

	response, err := json.Marshal(map[string]interface{}{
		"status": http.StatusOK,
		"body": map[string]interface{}{
			"selection": movieSelection,
		},
	})

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status":500}`)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
