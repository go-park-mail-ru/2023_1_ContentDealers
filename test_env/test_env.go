package test_env

import (
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/setup"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase"
	"github.com/gorilla/mux"
)

const TestOrigin = "example.com"

type TestCase struct {
	Path         string
	Method       string
	WithCookie   bool
	RequestBody  map[string]interface{}
	ResponseBody map[string]interface{}
	StatusCode   int
}

var TestUser = domain.UserCredentials{
	Email:    "vova@mail.ru",
	Password: "password",
}

type TestEnv struct {
	UserRepository repository.UserInMemoryRepository

	SessionRepository        repository.SessionInMemoryRepository
	MovieRepository          repository.MovieInMemoryRepository
	MovieSelectionRepository repository.MovieSelectionInMemoryRepository

	UserUseCase           *usecase.UserUseCase
	SessionUseCase        *usecase.SessionUseCase
	MovieSelectionUseCase *usecase.MovieSelectionUseCase

	UserHandler           delivery.UserHandler
	MovieSelectionHandler delivery.MovieSelectionHandler

	Router *mux.Router
}

func NewTestEnv() *TestEnv {
	UserRepository := repository.NewUserInMemoryRepository()
	SessionRepository := repository.NewSessionInMemoryRepository()
	MovieRepository := repository.NewMovieInMemoryRepository()
	MovieSelectionRepository := repository.NewMovieSelectionInMemoryRepository()

	setup.Content(&MovieRepository, &MovieSelectionRepository)

	UserUseCase := usecase.NewUser(&UserRepository)
	SessionUseCase := usecase.NewSession(&SessionRepository)
	MovieSelectionUseCase := usecase.NewMovieSelection(&MovieSelectionRepository)

	UserHandler := delivery.NewUserHandler(UserUseCase, SessionUseCase)
	MovieSelectionHandler := delivery.NewMovieSelectionHandler(MovieSelectionUseCase)

	Router := setup.Routes(&setup.SettingsRouter{
		UserHandler:           UserHandler,
		MovieSelectionHandler: MovieSelectionHandler,
		SessionUseCase:        SessionUseCase,
		AllowedOrigins:        []string{TestOrigin},
	})

	return &TestEnv{
		UserRepository:           UserRepository,
		SessionRepository:        SessionRepository,
		MovieRepository:          MovieRepository,
		MovieSelectionRepository: MovieSelectionRepository,
		UserUseCase:              UserUseCase,
		SessionUseCase:           SessionUseCase,
		MovieSelectionUseCase:    MovieSelectionUseCase,
		UserHandler:              UserHandler,
		MovieSelectionHandler:    MovieSelectionHandler,
		Router:                   Router,
	}

}
