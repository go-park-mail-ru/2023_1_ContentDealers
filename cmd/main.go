package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/movieselection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user"
	movieSelectionRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/movieselection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/session"
	userRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/user"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/setup/logger"
	"github.com/joho/godotenv"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user/csrf"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/setup"
	movieSelectionUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/movieselection"
	sessionUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/session"
	userUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/user"
)

const ReadHeaderTimeout = 5 * time.Second

// @title Filmium Backend API
// @version 1.0
// @description API Server for Filmium Application

// @host localhost:8080
// @BasePath /
func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	logger, err := logger.NewLogger()

	config, err := setup.GetConfig()
	if err != nil {
		logger.Error(err)
		return err
	}

	db, err := setup.NewClientPostgres(config.Storage)
	if err != nil {
		logger.Error(err)
		return err
	}

	redisClient, err := setup.NewClientRedis()
	if err != nil {
		logger.Error(err)
		return err
	}

	err = godotenv.Load()
	if err != nil {
		logger.Error(err)
		return err
	}
	cryptToken, err := csrf.NewCryptToken(os.Getenv("CSRF_TOKEN"))
	if err != nil {
		logger.Error(err)
		return err
	}

	userRepository := userRepo.NewRepository(db)
	sessionRepository := session.NewRepository(redisClient)
	movieSelectionRepository := movieSelectionRepo.NewRepository(db)

	userUseCase := userUseCase.NewUser(&userRepository)
	sessionUseCase := sessionUseCase.NewSession(&sessionRepository)
	movieSelectionUseCase := movieSelectionUseCase.NewMovieSelection(&movieSelectionRepository)

	userHandler := user.NewHandler(userUseCase, sessionUseCase, cryptToken, logger)
	movieSelectionHandler := movieselection.NewHandler(movieSelectionUseCase, logger)

	router := setup.Routes(&setup.SettingsRouter{
		UserHandler:           userHandler,
		MovieSelectionHandler: movieSelectionHandler,
		SessionUseCase:        sessionUseCase,
		AllowedOrigins:        []string{config.CORS.AllowedOrigins},
		CryptToken:            cryptToken,
		Logger:                logger,
	})

	addr := fmt.Sprintf("%s:%s", config.Listen.BindIP, config.Listen.Port)

	server := http.Server{
		Addr:              addr,
		Handler:           router,
		ReadHeaderTimeout: ReadHeaderTimeout,
	}

	logger.Infoln("start listening on", addr)

	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
