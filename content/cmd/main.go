package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os/signal"
	"syscall"

	config "github.com/go-park-mail-ru/2023_1_ContentDealers/config"
	filmDelivery "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/delivery/film"
	personDelivery "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/delivery/person"
	searchDelivery "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/delivery/search"
	selectionDelivery "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/delivery/selection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/repository/content"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/repository/country"
	filmRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/repository/film"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/repository/genre"
	personRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/repository/person"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/repository/role"
	selectionRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/repository/selection"
	contentUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/usecase/content"
	filmUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/usecase/film"
	personUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/usecase/person"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/usecase/personRole"
	searchUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/usecase/search"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/usecase/search/extender"
	selectionUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/usecase/selection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/film"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/person"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/search"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/selection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/client/postgresql"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"google.golang.org/grpc"
)

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	configPtr := flag.String("config", "", "Config file")
	flag.StringVar(configPtr, "c", "", "Config file (short)")

	flag.Parse()

	if *configPtr == "" {
		return fmt.Errorf("needed to pass config file")
	}

	cfg, err := config.GetCfg(*configPtr)
	if err != nil {
		return fmt.Errorf("fail to parse config yml file: %w", err)
	}

	logger, err := logging.NewLogger(cfg.Logging)
	if err != nil {
		return fmt.Errorf("fail to initialization logger: %w", err)
	}

	db, err := postgresql.NewClientPostgres(cfg.Storage)
	if err != nil {
		logger.Error(err)
		return err
	}

	filmRepository := filmRepo.NewRepository(db, logger)
	contentRepository := content.NewRepository(db, logger)
	genreRepository := genre.NewRepository(db, logger)
	selectionRepository := selectionRepo.NewRepository(db, logger)
	countryRepository := country.NewRepository(db, logger)
	personRepository := personRepo.NewRepository(db, logger)
	roleRepository := role.NewRepository(db, logger)

	personRolesUseCase := personRole.NewPersonRole(&personRepository, &roleRepository, logger)
	contentUsecase := contentUseCase.NewContent(contentUseCase.Options{
		ContentRepo:        &contentRepository,
		GenreRepo:          &genreRepository,
		SelectionRepo:      &selectionRepository,
		CountryRepo:        &countryRepository,
		PersonRolesUseCase: personRolesUseCase,
		Logger:             logger,
	})
	filmUsecase := filmUseCase.NewFilm(&filmRepository, contentUsecase, logger)
	personUsecase := personUseCase.NewPerson(personUseCase.Options{
		Repo:    &personRepository,
		Content: &contentRepository,
		Role:    &roleRepository,
		Genre:   &genreRepository,
		Logger:  logger,
	})
	selectionUsecase := selectionUseCase.NewSelection(&selectionRepository, &contentRepository, logger)

	searchExtenders := []searchUseCase.Extender{
		extender.NewContentExtender(&contentRepository, logger),
		extender.NewPersonExtender(&personRepository, logger),
	}
	searchUsecase := searchUseCase.NewSearch(searchExtenders, logger)

	filmService := filmDelivery.NewGrpc(filmUsecase)
	personService := personDelivery.NewGrpc(personUsecase)
	selectionService := selectionDelivery.NewGrpc(selectionUsecase)
	searchService := searchDelivery.NewGrpc(searchUsecase)

	server := grpc.NewServer()
	film.RegisterFilmServiceServer(server, filmService)
	person.RegisterPersonServiceServer(server, personService)
	selection.RegisterSelectionServiceServer(server, selectionService)
	search.RegisterSearchServiceServer(server, searchService)

	addr := cfg.ContentAddr

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Fatalln("cant listen port", err)
	}
	logger.Infoln("start listening on", addr)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	go func() {
		err = server.Serve(lis)
		if err != nil {
			logger.Error(err)
		}
	}()

	<-ctx.Done()
	server.GracefulStop()
	return nil
}
