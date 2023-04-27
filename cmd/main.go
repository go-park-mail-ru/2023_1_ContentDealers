package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/film"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/person"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/search"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/selection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/gateway/content"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/session"
	userRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/user"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/setup"
	filmUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/film"
	personUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/person"
	searchUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/search"
	selectionUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/selection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/joho/godotenv"

	config "github.com/go-park-mail-ru/2023_1_ContentDealers/config"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/csrf"
	csrfUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/csrf"
	sessionUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/session"
	userUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/user"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/client/postgresql"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/client/redis"
)

const (
	ReadHeaderTimeout = 5 * time.Second
	shutdownTimeout   = 5 * time.Second
)

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
	// go run cmd/main.go --config [-c] config.yml
	configPtr := flag.String("config", "", "Config file")
	flag.StringVar(configPtr, "c", "", "Config file (short)")

	flag.Parse()

	if *configPtr == "" {
		return fmt.Errorf("Needed to pass config file")
	}

	cfg, err := config.GetCfg(*configPtr)
	if err != nil {
		return fmt.Errorf("Fail to parse config yml file: %w", err)
	}

	logger, err := logging.NewLogger(cfg.Logging)
	if err != nil {
		return fmt.Errorf("Fail to initialization logger: %w", err)
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

	contentAddr := fmt.Sprintf("%s:%s", cfg.Content.Host, cfg.Content.Port)
	userRepository := userRepo.NewRepository(db, logger)
	sessionRepository := session.NewRepository(redisClient, logger)
	contentGateway, err := content.NewGrpc(contentAddr, logger)
	if err != nil {
		logger.Error(err)
		return err
	}

	userUsecase := userUseCase.NewUser(&userRepository, logger)
	sessionUsecase := sessionUseCase.NewSession(&sessionRepository, logger)
	selectionUsecase := selectionUseCase.NewUseCase(contentGateway, logger)
	filmUsecase := filmUseCase.NewUseCase(contentGateway, logger)
	personUsecase := personUseCase.NewUseCase(contentGateway, logger)
	searchUsecase := searchUseCase.NewUseCase(contentGateway, logger)

	err = godotenv.Load()
	if err != nil {
		logger.Error(err)
		return err
	}
	csrfUsecase, err := csrfUseCase.NewCSRF(os.Getenv("CSRF_TOKEN"), logger)
	if err != nil {
		logger.Error(err)
		return err
	}

	selectionHandler := selection.NewHandler(selectionUsecase, logger)
	filmHandler := film.NewHandler(filmUsecase, logger)
	personHandler := person.NewHandler(personUsecase, logger)
	userHandler := user.NewHandler(userUsecase, sessionUsecase, logger)
	csrfHandler := csrf.NewHandler(csrfUsecase, logger)
	searchHandler := search.NewHandler(searchUsecase, logger)

	router := setup.Routes(&setup.SettingsRouter{
		UserHandler:      userHandler,
		CSRFHandler:      csrfHandler,
		SelectionHandler: selectionHandler,
		FilmHandler:      filmHandler,
		PersonHandler:    personHandler,
		SearchHandler:    searchHandler,
		SessionUseCase:   sessionUsecase,
		AllowedOrigins:   []string{cfg.CORS.AllowedOrigins},
		CSRFUseCase:      *csrfUsecase,
		Logger:           logger,
	})

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	addr := fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port)

	server := http.Server{
		Addr:              addr,
		Handler:           router,
		ReadHeaderTimeout: ReadHeaderTimeout,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Error("listen and server", err)
		}
	}()
	logger.Infoln("start listening on", addr)

	<-ctx.Done()

	logger.Infoln("server shutdown")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	err = server.Shutdown(shutdownCtx)
	if err != nil {
		logger.Error(err)
		return fmt.Errorf("shutdown: %w", err)
	}
	return nil
}
