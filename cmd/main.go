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
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/setup"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/joho/godotenv"

	config "github.com/go-park-mail-ru/2023_1_ContentDealers/config"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/csrf"
	csrfUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/csrf"
	movieSelectionUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/movieselection"
	sessionUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/session"
	userUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/user"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/client/postgresql"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/client/redis"
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
	logger, err := logging.NewLogger()

	cfg, err := config.GetConfig()
	if err != nil {
		logger.Error(err)
		return err
	}

	db, err := postgresql.NewClientPostgres(cfg.Storage)
	if err != nil {
		logger.Error(err)
		return err
	}

	redisClient, err := redis.NewClientRedis(cfg.Redis)
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

	err = godotenv.Load()
	if err != nil {
		logger.Error(err)
		return err
	}
	csrfUseCase, err := csrfUseCase.NewCSRF(os.Getenv("CSRF_TOKEN"))
	if err != nil {
		logger.Error(err)
		return err
	}

	userHandler := user.NewHandler(userUseCase, sessionUseCase, logger)
	csrfHandler := csrf.NewHandler(csrfUseCase)
	movieSelectionHandler := movieselection.NewHandler(movieSelectionUseCase, logger)

	router := setup.Routes(&setup.SettingsRouter{
		UserHandler:           userHandler,
		CSRFHandler:           csrfHandler,
		MovieSelectionHandler: movieSelectionHandler,
		SessionUseCase:        sessionUseCase,
		AllowedOrigins:        []string{cfg.CORS.AllowedOrigins},
		CSRFUseCase:           csrfUseCase,
		Logger:                logger,
	})

	addr := fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port)

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
