package testenv

import (
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/selection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/film"
	movieSelectionRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/selection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/session"
	userRepo "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository/user"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/setup"
	movieSelectionUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/selection"
	sessionUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/session"
	userUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/user"
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
	Password: "passW_0rd",
}

type TestEnv struct {
	UserRepository           UserRepository
	SessionRepository        SessionRepository
	MovieRepository          MovieRepository
	MovieSelectionRepository MovieSelectionRepository

	UserUseCase           UserUseCase
	SessionUseCase        SessionUseCase
	MovieSelectionUseCase MovieSelectionUseCase

	UserHandler           user.Handler
	MovieSelectionHandler selection.Handler

	Router *mux.Router
}

func NewTestEnv() *TestEnv {
	userRepository := userRepo.NewInMemoryRepository()
	sessionRepository := session.NewInMemoryRepository()
	movieRepository := film.NewInMemoryRepository()
	movieSelectionRepository := movieSelectionRepo.NewInMemoryRepository()

	setup.Content(&movieRepository, &movieSelectionRepository)

	userUseCase := userUseCase.NewUser(&userRepository)
	sessionUseCase := sessionUseCase.NewSession(&sessionRepository)
	movieSelectionUseCase := movieSelectionUseCase.NewMovieSelection(&movieSelectionRepository)

	userHandler := user.NewHandler(userUseCase, sessionUseCase)
	movieSelectionHandler := selection.NewHandler(movieSelectionUseCase)

	router := setup.Routes(&setup.SettingsRouter{
		UserHandler:      userHandler,
		SelectionHandler: movieSelectionHandler,
		SessionUseCase:   sessionUseCase,
		AllowedOrigins:   []string{TestOrigin},
	})

	return &TestEnv{
		UserRepository:           &userRepository,
		SessionRepository:        &sessionRepository,
		MovieRepository:          &movieRepository,
		MovieSelectionRepository: &movieSelectionRepository,
		UserUseCase:              userUseCase,
		SessionUseCase:           sessionUseCase,
		MovieSelectionUseCase:    movieSelectionUseCase,
		UserHandler:              userHandler,
		MovieSelectionHandler:    movieSelectionHandler,
		Router:                   router,
	}

}
