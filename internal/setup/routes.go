package setup

import (
	"io"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/movieselection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/middleware"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, `{"status":404}`)
}

type SettingsRouter struct {
	AllowedOrigins        []string
	UserHandler           user.Handler
	MovieSelectionHandler movieselection.Handler
	SessionUseCase        *usecase.Session
}

func Routes(s *SettingsRouter) *mux.Router {
	corsMiddleware := cors.New(cors.Options{
		// TODO: поменять настройки CORS, когда будет известен домен
		AllowedOrigins:   s.AllowedOrigins,
		AllowedMethods:   []string{"POST", "GET", "OPTIONS"},
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
	authRouter.Use(authMiddleware.Authorized)
	unAuthRouter.Use(authMiddleware.UnAuthorized)

	router.HandleFunc("/selections", s.MovieSelectionHandler.GetAll)
	router.HandleFunc("/selections/{id:[0-9]+}", s.MovieSelectionHandler.GetByID)

	unAuthRouter.HandleFunc("/signin", s.UserHandler.SignIn).Methods("POST")
	unAuthRouter.HandleFunc("/signup", s.UserHandler.SignUp).Methods("POST")

	authRouter.HandleFunc("/logout", s.UserHandler.Logout).Methods("POST")
	authRouter.HandleFunc("/profile", s.UserHandler.Info).Methods("GET")

	return router
}
