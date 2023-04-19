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

	filmRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/repository/film"
	personRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/repository/person"
	selectionRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/repository/selection"
	filmUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/usecase/film"
	personUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/usecase/person"
	selectionUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/usecase/selection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/film"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/person"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/selection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/session"
	userRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/user"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/setup"
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

	userRepository := userRepo.NewRepository(db, logger)
	sessionRepository := session.NewRepository(redisClient, logger)
	selectionRepository := selectionRepo.NewRepository(db, logger)
	contentRepository := content.NewRepository(db, logger)
	filmRepository := filmRepo.NewRepository(db, logger)
	genreRepository := genre.NewRepository(db, logger)
	roleRepository := role.NewRepository(db, logger)
	countryRepository := country.NewRepository(db, logger)
	personRepository := personRepo.NewRepository(db, logger)

	userUseCase := userUseCase.NewUser(&userRepository, logger)
	sessionUseCase := sessionUseCase.NewSession(&sessionRepository, logger)
	selectionUseCase := selectionUseCase.NewSelection(&selectionRepository, &contentRepository, logger)
	personRolesUseCase := personRole.NewPersonRole(&personRepository, &roleRepository, logger)

	contentUseCase := content.NewContent(content.Options{
		ContentRepo:        &contentRepository,
		GenreRepo:          &genreRepository,
		SelectionRepo:      &selectionRepository,
		CountryRepo:        &countryRepository,
		PersonRolesUseCase: personRolesUseCase,
	}, logger)
	filmUseCase := filmUseCase.NewFilm(&filmRepository, contentUseCase, logger)
	personUseCase := personUseCase.NewPerson(personUseCase.Options{
		Repo:    &personRepository,
		Content: &contentRepository,
		Role:    &roleRepository,
		Genre:   &genreRepository,
	}, logger)

	err = godotenv.Load()
	if err != nil {
		logger.Error(err)
		return err
	}
	csrfUseCase, err := csrfUseCase.NewCSRF(os.Getenv("CSRF_TOKEN"), logger)
	if err != nil {
		logger.Error(err)
		return err
	}

	selectionHandler := selection.NewHandler(selectionUseCase, logger)
	filmHandler := film.NewHandler(filmUseCase, logger)
	personHandler := person.NewHandler(personUseCase, logger)
	userHandler := user.NewHandler(userUseCase, sessionUseCase, logger)
	csrfHandler := csrf.NewHandler(csrfUseCase, logger)

	router := setup.Routes(&setup.SettingsRouter{
		UserHandler:      userHandler,
		CSRFHandler:      csrfHandler,
		SelectionHandler: selectionHandler,
		FilmHandler:      filmHandler,
		PersonHandler:    personHandler,
		SessionUseCase:   sessionUseCase,
		AllowedOrigins:   []string{cfg.CORS.AllowedOrigins},
		CSRFUseCase:      *csrfUseCase,
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
