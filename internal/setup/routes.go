package setup

import (
	"io"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/film"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/person"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/selection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user/middleware"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"status":404}`)
}

type SettingsRouter struct {
	AllowedOrigins   []string
	UserHandler      user.Handler
	SelectionHandler selection.Handler
	FilmHandler      film.Handler
	PersonHandler    person.Handler
	SessionUseCase   SessionUseCase
}

func Routes(s *SettingsRouter) *mux.Router {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   s.AllowedOrigins,
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PATCH"},
		AllowCredentials: true,
		Debug:            true,
	})
	authMiddleware := middleware.NewAuth(s.SessionUseCase)

	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(NotFound)

	authRouter := router.Methods("GET", "POST").Subrouter()
	unAuthRouter := router.Methods("GET", "POST").Subrouter()

	router.Use(corsMiddleware.Handler)
	router.Use(middleware.SetContentTypeJSON)
	authRouter.Use(authMiddleware.RequireAuth)
	unAuthRouter.Use(authMiddleware.RequireUnAuth)

	router.HandleFunc("/selections", s.SelectionHandler.GetAll)
	router.HandleFunc("/selections/{id:[0-9]+}", s.SelectionHandler.GetByID)
	router.HandleFunc("/person/{id:[0-9]+}", s.PersonHandler.GetByID)
	router.HandleFunc("/film/{content_id:[0-9]+}", s.FilmHandler.GetByContentID)

	unAuthRouter.HandleFunc("/user/signin", s.UserHandler.SignIn).Methods("POST")
	unAuthRouter.HandleFunc("/user/signup", s.UserHandler.SignUp).Methods("POST")

	authRouter.HandleFunc("/user/logout", s.UserHandler.Logout).Methods("POST")
	authRouter.HandleFunc("/user/profile", s.UserHandler.Info).Methods("GET")

	// TODO: PATCH в постмане выдавал 405 Method not allowed
	authRouter.HandleFunc("/user/avatar/update", s.UserHandler.UploadAvatar).Methods("POST")
	authRouter.HandleFunc("/user/avatar/delete", s.UserHandler.DeleteAvatar).Methods("POST")
	// authRouter.HandleFunc("/user/update/profile", s.UserHandler.UploadAvatar).Methods("GET")

	return router
}
