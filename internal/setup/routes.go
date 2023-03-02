package setup

import (
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delievery"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/middleware"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Routes() *mux.Router {
	userRepository := repository.NewUserInMemoryRepository()
	sessionRepository := repository.NewSessionInMemoryRepository()
	movieRepository := repository.NewMovieInMemoryRepository()
	movieSelectionRepository := repository.NewMovieSelectionInMemoryRepository()

	Content(&movieRepository, &movieSelectionRepository)

	userUseCase := usecase.NewUserUseCase(&userRepository)
	sessionUseCase := usecase.NewSessionUseCase(&sessionRepository)
	movieSelectionUseCase := usecase.NewMovieSelectionUseCase(&movieSelectionRepository)

	userHandler := delievery.NewUserHandler(userUseCase, sessionUseCase)
	movieSelectionHandler := delievery.NewMovieSelectionHandler(&movieSelectionUseCase)

	corsMiddleware := cors.New(cors.Options{
		// TODO: поменять настройки CORS, когда будет известен домен
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	})
	authMiddleware := middleware.NewAuthMiddleware(sessionUseCase)

	router := mux.NewRouter()
	authRouter := router.Methods("GET", "POST").Subrouter()
	unAuthRouter := router.Methods("GET", "POST").Subrouter()

	router.Use(corsMiddleware.Handler)
	authRouter.Use(authMiddleware.Authorized)
	unAuthRouter.Use(authMiddleware.UnAuthorized)

	router.HandleFunc("/selections", movieSelectionHandler.GetAll)
	router.HandleFunc("/selection/{id}", movieSelectionHandler.GetById)

	unAuthRouter.HandleFunc("/signin", userHandler.SignIn).Methods("POST")
	unAuthRouter.HandleFunc("/signup", userHandler.SignUp).Methods("POST")

	authRouter.HandleFunc("/logout", userHandler.Logout).Methods("POST")
	authRouter.HandleFunc("/profile", userHandler.GetUserInfo).Methods("GET")

	return router
}
