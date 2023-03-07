package testenv

import (
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/contract"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/movieselection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/movie"
	movieSelectionRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/movieselection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/session"
	userRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/user"
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
	UserRepository           contract.UserRepository
	SessionRepository        contract.SessionRepository
	MovieRepository          contract.MovieRepository
	MovieSelectionRepository contract.MovieSelectionRepository

	UserUseCase           contract.UserUseCase
	SessionUseCase        contract.SessionUseCase
	MovieSelectionUseCase contract.MovieSelectionUseCase

	UserHandler           user.Handler
	MovieSelectionHandler movieselection.Handler

	Router *mux.Router
}

func NewTestEnv() *TestEnv {
	UserRepository := userRepo.NewInMemoryRepository()
	SessionRepository := session.NewInMemoryRepository()
	MovieRepository := movie.NewInMemoryRepository()
	MovieSelectionRepository := movieSelectionRepo.NewInMemoryRepository()

	setup.Content(&MovieRepository, &MovieSelectionRepository)

	UserUseCase := usecase.NewUser(&UserRepository)
	SessionUseCase := usecase.NewSession(&SessionRepository)
	MovieSelectionUseCase := usecase.NewMovieSelection(&MovieSelectionRepository)

	UserHandler := user.NewHandler(UserUseCase, SessionUseCase)
	MovieSelectionHandler := movieselection.NewHandler(MovieSelectionUseCase)

	Router := setup.Routes(&setup.SettingsRouter{
		UserHandler:           UserHandler,
		MovieSelectionHandler: MovieSelectionHandler,
		SessionUseCase:        SessionUseCase,
		AllowedOrigins:        []string{TestOrigin},
	})

	return &TestEnv{
		UserRepository:           &UserRepository,
		SessionRepository:        &SessionRepository,
		MovieRepository:          &MovieRepository,
		MovieSelectionRepository: &MovieSelectionRepository,
		UserUseCase:              UserUseCase,
		SessionUseCase:           SessionUseCase,
		MovieSelectionUseCase:    MovieSelectionUseCase,
		UserHandler:              UserHandler,
		MovieSelectionHandler:    MovieSelectionHandler,
		Router:                   Router,
	}

}
