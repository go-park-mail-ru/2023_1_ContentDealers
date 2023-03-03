package delivery

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase"
)

type MovieSelectionHandler struct {
	useCase *usecase.MovieSelectionUseCase
}

func NewMovieSelectionHandler(useCase *usecase.MovieSelectionUseCase) MovieSelectionHandler {
	return MovieSelectionHandler{useCase: useCase}
}

func (h *MovieSelectionHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	selections, err := h.useCase.GetAll()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status":500}`)
		return
	}

	response, err := json.Marshal(map[string]interface{}{
		"status": 200,
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

func (h *MovieSelectionHandler) GetById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	idRaw := mux.Vars(r)["id"]
	var id uint64

	_, err := fmt.Sscanf(idRaw, "%d", &id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"status":400}`)
		return
	}

	movieSelection, err := h.useCase.GetById(id)
	if err != nil {
		switch err {
		case repository.ErrMovieSelectionNotFound:
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
		"status": 200,
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
