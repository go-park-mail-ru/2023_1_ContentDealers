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

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/favorites"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/film"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/person"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/selection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user"
	contentRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/gateway/content"
	countryRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/gateway/country"
	favGate "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/gateway/favorites"
	filmRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/gateway/film"
	genreRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/gateway/genre"
	personRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/gateway/person"
	roleRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/gateway/role"
	selectionRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/gateway/selection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/gateway/session"
	userGate "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/gateway/user"
	favUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/favorites"
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
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/client/postgresql"
)

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

	logger, err := logging.NewLogger(cfg.Logging, "api-gateway")
	if err != nil {
		return fmt.Errorf("Fail to initialization logger: %w", err)
	}

	db, err := postgresql.NewClientPostgres(cfg.Postgres)
	if err != nil {
		logger.Error(err)
		return err
	}

	userGateway, err := userGate.NewGateway(logger, cfg.ServiceUser)
	if err != nil {
		return err
	}

	sessionGateway, err := session.NewGateway(logger, cfg.ServiceSession)
	if err != nil {
		return err
	}

	favGateway, err := favGate.NewGateway(logger, cfg.ServiceFavorites)
	if err != nil {
		return err
	}

	selectionRepository := selectionRepo.NewRepository(db, logger)
	contentRepository := contentRepo.NewRepository(db, logger)
	filmRepository := filmRepo.NewRepository(db, logger)
	genreRepository := genreRepo.NewRepository(db, logger)
	roleRepository := roleRepo.NewRepository(db, logger)
	countryRepository := countryRepo.NewRepository(db, logger)
	personRepository := personRepo.NewRepository(db, logger)

	selectionUseCase := selectionUseCase.NewSelection(&selectionRepository, &contentRepository, logger)
	personRolesUseCase := personRoleUseCase.NewPersonRole(&personRepository, &roleRepository, logger)

	contentUseCase := contentUseCase.NewContent(contentUseCase.Options{
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

	favUseCase := favUseCase.NewUseCase(favGateway, sessionGateway, contentUseCase, logger)

	err = godotenv.Load()
	if err != nil {
		logger.Error(err)
		return err
	}
	csrfUseCase, err := csrfUseCase.NewUseCase(os.Getenv("CSRF_TOKEN"), logger)
	if err != nil {
		logger.Error(err)
		return err
	}

	selectionHandler := selection.NewHandler(selectionUseCase, logger)

	filmHandler := film.NewHandler(filmUseCase, logger)

	personHandler := person.NewHandler(personUseCase, logger)
	userHandler := user.NewHandler(userGateway, sessionGateway, logger, cfg.Avatar)
	csrfHandler := csrf.NewHandler(csrfUseCase, logger, cfg.CSRF)

	favHandler := favorites.NewHandler(favUseCase, logger)

	router := setup.Routes(&setup.SettingsRouter{
		UserHandler:      *userHandler,
		FavHandler:       *favHandler,
		CSRFHandler:      *csrfHandler,
		SelectionHandler: selectionHandler,
		FilmHandler:      filmHandler,
		PersonHandler:    personHandler,
		SessionGateway:   sessionGateway,
		AllowedOrigins:   []string{cfg.CORS.AllowedOrigins},
		CSRFUseCase:      *csrfUseCase,
		Logger:           logger,
		CSRFConfig:       cfg.CSRF,
	})

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	addr := fmt.Sprintf("%s:%s", cfg.Server.BindIP, cfg.Server.Port)

	server := http.Server{
		Addr:              addr,
		Handler:           router,
		ReadHeaderTimeout: time.Second * time.Duration(cfg.Server.ReadHeaderTimeout),
		WriteTimeout:      time.Second * time.Duration(cfg.Server.WriteTimeout),
		ReadTimeout:       time.Second * time.Duration(cfg.Server.ReadTimeout),
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

	shutdownCtx, cancel := context.WithTimeout(context.Background(),
		time.Second*time.Duration(cfg.Server.ShutdownTimeout))
	defer cancel()

	err = server.Shutdown(shutdownCtx)
	if err != nil {
		logger.Error(err)
		return fmt.Errorf("shutdown: %w", err)
	}
	return nil
}
