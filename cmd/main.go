package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/movieselection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user"
	movieSelectionRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/movieselection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/session"
	userRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/user"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/setup"
	movieSelectionUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/movieselection"
	sessionUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/session"
	userUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/user"
)

const addr = ":8080"
const ReadHeaderTimeout = 5 * time.Second

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {

	db, err := setup.NewClientPostgres()
	if err != nil {
		return err
	}

	redisClient, err := setup.NewClientRedis()
	if err != nil {
		return err
	}

	userRepository := userRepo.NewRepository(db)
	sessionRepository := session.NewRepository(redisClient)
	movieSelectionRepository := movieSelectionRepo.NewRepository(db)

	userUseCase := userUseCase.NewUser(&userRepository)
	sessionUseCase := sessionUseCase.NewSession(&sessionRepository)
	movieSelectionUseCase := movieSelectionUseCase.NewMovieSelection(&movieSelectionRepository)

	userHandler := user.NewHandler(userUseCase, sessionUseCase)
	movieSelectionHandler := movieselection.NewHandler(movieSelectionUseCase)

	router := setup.Routes(&setup.SettingsRouter{
		UserHandler:           userHandler,
		MovieSelectionHandler: movieSelectionHandler,
		SessionUseCase:        sessionUseCase,
		AllowedOrigins:        []string{"89.208.199.170"},
	})

	server := http.Server{
		Addr:              addr,
		Handler:           router,
		ReadHeaderTimeout: ReadHeaderTimeout,
	}

	log.Println("start listening on", addr)

	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
