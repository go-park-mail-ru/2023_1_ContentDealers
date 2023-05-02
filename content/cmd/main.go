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
	contentDelivery "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/delivery/content"
	genreDelivery "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/delivery/genre"
	personDelivery "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/delivery/person"
	searchDelivery "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/delivery/search"
	selectionDelivery "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/delivery/selection"
	contentRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/repository/content"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/repository/country"
	genreRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/repository/genre"
	personRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/repository/person"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/repository/role"
	selectionRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/repository/selection"
	contentUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/usecase/content"
	genreUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/usecase/genre"
	personUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/usecase/person"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/usecase/personrole"
	searchUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/usecase/search"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/usecase/search/extender"
	selectionUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/usecase/selection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/content"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/genre"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/person"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/search"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/selection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/client/postgresql"
	interceptorServer "github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/grpc/interceptor/server"
	pingDelivery "github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/grpc/ping"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/proto/ping"
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

	cfgGeneral, err := config.GetCfg(*configPtr)
	if err != nil {
		return fmt.Errorf("fail to parse config yml file: %w", err)
	}
	cfg := cfgGeneral.Content

	logger, err := logging.NewLogger(cfg.Logging, "content serivce")
	if err != nil {
		return fmt.Errorf("fail to initialization logger: %w", err)
	}

	db, err := postgresql.NewClientPostgres(cfg.Postgres)
	if err != nil {
		logger.Error(err)
		return err
	}

	contentRepository := contentRepo.NewRepository(db)
	genreRepository := genreRepo.NewRepository(db)
	selectionRepository := selectionRepo.NewRepository(db)
	countryRepository := country.NewRepository(db)
	personRepository := personRepo.NewRepository(db)
	roleRepository := role.NewRepository(db)

	personRolesUseCase := personrole.NewUseCase(&personRepository, &roleRepository)
	contentUsecase := contentUseCase.NewUseCase(contentUseCase.Options{
		ContentRepo:        &contentRepository,
		GenreRepo:          &genreRepository,
		SelectionRepo:      &selectionRepository,
		CountryRepo:        &countryRepository,
		PersonRolesUseCase: personRolesUseCase,
	})
	personUsecase := personUseCase.NewUseCase(personUseCase.Options{
		Repo:    &personRepository,
		Content: &contentRepository,
		Role:    &roleRepository,
		Genre:   &genreRepository,
	})
	selectionUsecase := selectionUseCase.NewUseCase(&selectionRepository, &contentRepository)
	genreUsecase := genreUseCase.NewUseCase(&genreRepository, &contentRepository)

	searchExtenders := []searchUseCase.Extender{
		extender.NewContentExtender(&contentRepository),
		extender.NewPersonExtender(&personRepository),
	}
	searchUsecase := searchUseCase.NewUseCase(searchExtenders)

	contentService := contentDelivery.NewGrpc(contentUsecase, logger)
	personService := personDelivery.NewGrpc(personUsecase, logger)
	selectionService := selectionDelivery.NewGrpc(selectionUsecase, logger)
	searchService := searchDelivery.NewGrpc(searchUsecase, logger)
	genreService := genreDelivery.NewGrpc(genreUsecase, logger)
	pingService := pingDelivery.NewGrpc()

	interceptor := interceptorServer.NewInterceptorServer("content", logger)
	server := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.AccessLog),
	)
	content.RegisterContentServiceServer(server, contentService)
	person.RegisterPersonServiceServer(server, personService)
	selection.RegisterSelectionServiceServer(server, selectionService)
	search.RegisterSearchServiceServer(server, searchService)
	genre.RegisterGenreServiceServer(server, genreService)
	ping.RegisterPingServiceServer(server, pingService)

	addr := fmt.Sprintf("%s:%s", cfg.Server.BindIP, cfg.Server.Port)

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
