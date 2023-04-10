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
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/selection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user"
	contentRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/content"
	countryRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/country"
	filmRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/film"
	genreRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/genre"
	personRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/person"
	roleRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/role"
	selectionRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/selection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/session"
	userRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/user"
	filmUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/film"
	personUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/person"
	personRoleUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/personRole"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/setup"
	contentUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/content"
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
	logger, err := logging.NewLogger()

	// go run cmd/main.go --config [-c] config.yml
	configPtr := flag.String("config", "", "Config file")
	flag.StringVar(configPtr, "c", "", "Config file (short)")

	flag.Parse()

	if *configPtr == "" {
		return fmt.Errorf("Needed to pass config file")
	}

	cfg, err := config.GetCfg(*configPtr)
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
	selectionRepository := selectionRepo.NewRepository(db)
	contentRepository := contentRepo.NewRepository(db)
	filmRepository := filmRepo.NewRepository(db)
	genreRepository := genreRepo.NewRepository(db)
	roleRepository := roleRepo.NewRepository(db)
	countryRepository := countryRepo.NewRepository(db)
	personRepository := personRepo.NewRepository(db)

	userUseCase := userUseCase.NewUser(&userRepository)
	sessionUseCase := sessionUseCase.NewSession(&sessionRepository)
	selectionUseCase := selectionUseCase.NewSelection(&selectionRepository, &contentRepository)
	personRolesUseCase := personRoleUseCase.NewPersonRole(&personRepository, &roleRepository)
	contentUseCase := contentUseCase.NewContent(contentUseCase.Options{
		ContentRepo:        &contentRepository,
		GenreRepo:          &genreRepository,
		SelectionRepo:      &selectionRepository,
		CountryRepo:        &countryRepository,
		PersonRolesUseCase: personRolesUseCase,
	})
	filmUseCase := filmUseCase.NewFilm(&filmRepository, contentUseCase)
	personUseCase := personUseCase.NewPerson(personUseCase.Options{
		Repo:    &personRepository,
		Content: &contentRepository,
		Role:    &roleRepository,
		Genre:   &genreRepository,
	})

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

	selectionHandler := selection.NewHandler(selectionUseCase)
	filmHandler := film.NewHandler(filmUseCase)
	personHandler := person.NewHandler(personUseCase)
	userHandler := user.NewHandler(userUseCase, sessionUseCase, logger)
	csrfHandler := csrf.NewHandler(csrfUseCase)

	router := setup.Routes(&setup.SettingsRouter{
		UserHandler:      userHandler,
		CSRFHandler:      csrfHandler,
		SelectionHandler: selectionHandler,
		FilmHandler:      filmHandler,
		PersonHandler:    personHandler,
		SessionUseCase:   sessionUseCase,
		AllowedOrigins:   []string{cfg.CORS.AllowedOrigins},
		CSRFUseCase:      csrfUseCase,
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
			log.Fatalf("listen and server: %v", err)
		}
	}()
	logger.Infoln("start listening on", addr)

	<-ctx.Done()

	logger.Infoln("server shutdown")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	err = server.Shutdown(shutdownCtx)
	if err != nil {
		return fmt.Errorf("shutdown : %w", err)
	}
	return nil
}
