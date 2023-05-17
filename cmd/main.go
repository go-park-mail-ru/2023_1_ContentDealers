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

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/content"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/favorites"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/genre"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/history_views"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/payment"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/person"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/rating"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/search"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/selection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user"
	contentGate "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/gateway/content"
	paymentGate "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/gateway/payment"
	sessGate "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/gateway/session"
	userGate "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/gateway/user"
	userActionGate "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/gateway/user_action"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/setup"
	filmUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/content"
	favUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/favorites"
	genreUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/genre"
	viewsUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/history_views"
	personUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/person"
	rateUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/rating"
	searchUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/search"
	selectionUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/selection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/joho/godotenv"

	config "github.com/go-park-mail-ru/2023_1_ContentDealers/config"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/csrf"
	csrfUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/csrf"
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

	cfgGeneral, err := config.GetCfg(*configPtr)
	if err != nil {
		return fmt.Errorf("Fail to parse config yml file: %w", err)
	}

	cfg := cfgGeneral.ApiGateway

	logger, err := logging.NewLogger(cfg.Logging, "api-gateway")
	if err != nil {
		return fmt.Errorf("Fail to initialization logger: %w", err)
	}

	userGateway, err := userGate.NewGateway(logger, cfg.ServiceUser)
	if err != nil {
		return err
	}

	sessionGateway, err := sessGate.NewGateway(logger, cfg.ServiceSession)
	if err != nil {
		return err
	}

	userActionGateway, err := userActionGate.NewGateway(logger, cfg.ServiceUserAction)
	if err != nil {
		return err
	}

	contentGateway, err := contentGate.NewGateway(cfg.ServiceContent, logger)
	if err != nil {
		logger.Error(err)
		return err
	}

	paymentGateway, err := paymentGate.NewGateway(logger, cfg.ServicePayment)
	if err != nil {
		logger.Error(err)
		return err
	}

	selectionUsecase := selectionUseCase.NewUseCase(contentGateway, logger)
	filmUsecase := filmUseCase.NewUseCase(contentGateway, logger)
	personUsecase := personUseCase.NewUseCase(contentGateway, logger)
	searchUsecase := searchUseCase.NewUseCase(contentGateway, logger)
	genreUsecase := genreUseCase.NewUseCase(contentGateway, logger)

	favUseCase := favUseCase.NewUseCase(userActionGateway, sessionGateway, contentGateway, logger)
	ratingUseCase := rateUseCase.NewUseCase(userActionGateway, sessionGateway, contentGateway, logger)
	viewsUseCase := viewsUseCase.NewUseCase(userActionGateway, sessionGateway, contentGateway, logger)

	err = godotenv.Load()
	if err != nil {
		logger.Error(err)
	}

	csrfUseCase, err := csrfUseCase.NewUseCase(os.Getenv("CSRF_TOKEN"), logger)
	if err != nil {
		logger.Error(err)
		return err
	}

	userHandler := user.NewHandler(userGateway, sessionGateway, logger, cfg.Avatar)
	favHandler := favorites.NewHandler(favUseCase, logger)
	viewsHandler := history_views.NewHandler(viewsUseCase, logger)
	rateHandler := rating.NewHandler(ratingUseCase, logger)
	selectionHandler := selection.NewHandler(selectionUsecase, logger)
	contentHandler := content.NewHandler(filmUsecase, logger)
	personHandler := person.NewHandler(personUsecase, logger)
	csrfHandler := csrf.NewHandler(csrfUseCase, logger, cfg.CSRF)
	searchHandler := search.NewHandler(searchUsecase, logger)
	genreHandler := genre.NewHandler(genreUsecase, logger)
	paymentHandler := payment.NewHandler(paymentGateway, logger)

	router := setup.Routes(&setup.SettingsRouter{
		PaymentHandler:   paymentHandler,
		UserHandler:      *userHandler,
		FavHandler:       *favHandler,
		RateHandler:      *rateHandler,
		ViewsHandler:     *viewsHandler,
		CSRFHandler:      *csrfHandler,
		SelectionHandler: selectionHandler,
		ContentHandler:   contentHandler,
		PersonHandler:    personHandler,
		SessionGateway:   sessionGateway,
		SearchHandler:    searchHandler,
		GenreHandler:     genreHandler,
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
