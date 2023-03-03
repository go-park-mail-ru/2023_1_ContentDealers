package setup

import (
	"io"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/middleware"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, `{"status":404}`)
}

func Routes(userHandler delivery.UserHandler,
	movieSelectionHandler delivery.MovieSelectionHandler,
	sessionUseCase *usecase.SessionUseCase) *mux.Router {
	corsMiddleware := cors.New(cors.Options{
		// TODO: поменять настройки CORS, когда будет известен домен
		AllowedOrigins:   []string{"example.com"},
		AllowCredentials: true,
		Debug:            true,
	})
	authMiddleware := middleware.NewAuthMiddleware(sessionUseCase)

	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(NotFound)

	authRouter := router.Methods("GET", "POST").Subrouter()
	unAuthRouter := router.Methods("GET", "POST").Subrouter()

	router.Use(corsMiddleware.Handler)
	router.Use(middleware.SetContentTypeJSON)
	authRouter.Use(authMiddleware.Authorized)
	unAuthRouter.Use(authMiddleware.UnAuthorized)

	router.HandleFunc("/selections", movieSelectionHandler.GetAll)
	router.HandleFunc("/selections/{id:[0-9]+}", movieSelectionHandler.GetById)

	unAuthRouter.HandleFunc("/signin", userHandler.SignIn).Methods("POST")
	unAuthRouter.HandleFunc("/signup", userHandler.SignUp).Methods("POST")

	authRouter.HandleFunc("/logout", userHandler.Logout).Methods("POST")
	authRouter.HandleFunc("/profile", userHandler.GetUserInfo).Methods("GET")

	return router
}
