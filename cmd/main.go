package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/selection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user"
	movieSelectionRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/movieselection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/session"
	userRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/user"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/setup"
	movieSelectionUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/selection"
	sessionUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/session"
	userUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/user"
)

const ReadHeaderTimeout = 5 * time.Second

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {

	config, err := setup.GetConfig()
	if err != nil {
		return err
	}

	fmt.Printf("%v", config)

	db, err := setup.NewClientPostgres(config.Storage)
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
	movieSelectionHandler := selection.NewHandler(movieSelectionUseCase)

	router := setup.Routes(&setup.SettingsRouter{
		UserHandler:           userHandler,
		MovieSelectionHandler: movieSelectionHandler,
		SessionUseCase:        sessionUseCase,
		AllowedOrigins:        []string{config.CORS.AllowedOrigins},
	})

	addr := fmt.Sprintf("%s:%s", config.Listen.BindIP, config.Listen.Port)

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
