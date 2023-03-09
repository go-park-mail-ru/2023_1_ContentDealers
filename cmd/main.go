package main

import (
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delievery"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/middleware"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"io"
	"log"
	"net/http"
)

func main() {
	userRepository := repository.NewUserInMemoryRepository()
	sessionRepository := repository.NewSessionInMemoryRepository()

	userUseCase := usecase.NewUserUseCase(&userRepository)
	sessionUseCase := usecase.NewSessionUseCase(&sessionRepository)

	userHandler := delievery.NewUserHandler(userUseCase, sessionUseCase)

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"status": 200}`)
	}).Methods("GET")

	authMiddleware := middleware.NewAuthMiddleware(sessionUseCase)

	authRouter := router.Methods("GET", "POST").Subrouter()
	authRouter.Use(authMiddleware.Authorized)

	unAuthRouter := router.Methods("GET", "POST").Subrouter()
	unAuthRouter.Use(authMiddleware.UnAuthorized)

	unAuthRouter.HandleFunc("/signin", userHandler.SignIn).Methods("POST")
	unAuthRouter.HandleFunc("/signup", userHandler.SignUp).Methods("POST")
	authRouter.HandleFunc("/logout", userHandler.Logout).Methods("POST")
	authRouter.HandleFunc("/profile", userHandler.GetUserInfo).Methods("GET")

	corsHandler := cors.New(cors.Options{
		// TODO: поменять настройки CORS, когда будет известен домен
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	})

	router.Use(corsHandler.Handler)

	addr := ":8080"

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Println("start listening on", addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
