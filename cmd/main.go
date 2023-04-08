package main

import (
	"fmt"
	"log"
	"net/http"
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
	sessionUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/session"
	userUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/user"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/setup"
	contentUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/content"
	selectionUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/selection"
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
	personUseCase := personUseCase.NewPerson(&personRepository)

	userHandler := user.NewHandler(userUseCase, sessionUseCase)
	selectionHandler := selection.NewHandler(selectionUseCase)
	filmHandler := film.NewHandler(filmUseCase)
	personHandler := person.NewHandler(personUseCase)

	router := setup.Routes(&setup.SettingsRouter{
		UserHandler:      userHandler,
		SelectionHandler: selectionHandler,
		SessionUseCase:   sessionUseCase,
		FilmHandler:      filmHandler,
		PersonHandler:    personHandler,
		AllowedOrigins:   []string{config.CORS.AllowedOrigins},
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
