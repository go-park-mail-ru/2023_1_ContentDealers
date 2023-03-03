package main

import (
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/setup"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase"
)

const addr = ":8080"

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	userRepository := repository.NewUserInMemoryRepository()
	sessionRepository := repository.NewSessionInMemoryRepository()
	movieRepository := repository.NewMovieInMemoryRepository()
	movieSelectionRepository := repository.NewMovieSelectionInMemoryRepository()

	setup.Content(&movieRepository, &movieSelectionRepository)

	userUseCase := usecase.NewUser(&userRepository)
	sessionUseCase := usecase.NewSession(&sessionRepository)
	movieSelectionUseCase := usecase.NewMovieSelection(&movieSelectionRepository)

	userHandler := delivery.NewUserHandler(userUseCase, sessionUseCase)
	movieSelectionHandler := delivery.NewMovieSelectionHandler(movieSelectionUseCase)

	router := setup.Routes(userHandler, movieSelectionHandler, sessionUseCase)

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Println("start listening on", addr)

	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
