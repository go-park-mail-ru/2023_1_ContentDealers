package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/setup"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

type TestEnv struct {
	userRepository repository.UserInMemoryRepository

	sessionRepository        repository.SessionInMemoryRepository
	movieRepository          repository.MovieInMemoryRepository
	movieSelectionRepository repository.MovieSelectionInMemoryRepository

	userUseCase           *usecase.UserUseCase
	sessionUseCase        *usecase.SessionUseCase
	movieSelectionUseCase *usecase.MovieSelectionUseCase

	userHandler           delivery.UserHandler
	movieSelectionHandler delivery.MovieSelectionHandler

	router *mux.Router
}

func NewTestEnv() *TestEnv {
	userRepository := repository.NewUserInMemoryRepository()
	sessionRepository := repository.NewSessionInMemoryRepository()
	movieRepository := repository.NewMovieInMemoryRepository()
	movieSelectionRepository := repository.NewMovieSelectionInMemoryRepository()

	setup.Content(&movieRepository, &movieSelectionRepository)

	userUseCase := usecase.NewUser(&userRepository)
	sessionUseCase := usecase.NewSession(&sessionRepository)
	movieSelectionUseCase := usecase.NewMovieSelection(&movieSelectionRepository)

	userHandler := delivery.NewUserHandler(userUseCase, sessionUseCase)
	movieSelectionHandler := delivery.NewMovieSelectionHandler(movieSelectionUseCase)
	router := setup.Routes(userHandler, movieSelectionHandler, sessionUseCase)

	return &TestEnv{
		userRepository:           userRepository,
		sessionRepository:        sessionRepository,
		movieRepository:          movieRepository,
		movieSelectionRepository: movieSelectionRepository,
		userUseCase:              userUseCase,
		sessionUseCase:           sessionUseCase,
		movieSelectionUseCase:    movieSelectionUseCase,
		userHandler:              userHandler,
		movieSelectionHandler:    movieSelectionHandler,
		router:                   router,
	}

}

type TestCase struct {
	Path         string
	Method       string
	RequestBody  map[string]interface{}
	ResponseBody map[string]interface{}
	StatusCode   int
}

var testCasesMovies = []TestCase{
	{
		Path:   "/selections",
		Method: "GET",
		ResponseBody: map[string]interface{}{
			"status": 200,
			"body": map[string]interface{}{
				"movie_selections": setup.MovieSelections,
			},
		},
		StatusCode: 200,
	},
	{
		Path:   "/selections/1",
		Method: "GET",
		ResponseBody: map[string]interface{}{
			"status": 200,
			"body": map[string]interface{}{
				"selection": setup.MovieSelections[1],
			},
		},
		StatusCode: 200,
	},
	{
		Path:         "/selections/33",
		Method:       "GET",
		ResponseBody: map[string]interface{}{"status": 404},
		StatusCode:   404,
	},
	{
		Path:         "/selections/hello",
		Method:       "GET",
		ResponseBody: map[string]interface{}{"status": 404},
		StatusCode:   404,
	},
	{
		Path:         "/selections/hello",
		Method:       "GET",
		ResponseBody: map[string]interface{}{"status": 404},
		StatusCode:   404,
	},
}

func TestApiMovies(t *testing.T) {
	testEnv := NewTestEnv()

	for numCase, testCase := range testCasesMovies {
		req := httptest.NewRequest(testCase.Method, testCase.Path, nil)
		w := httptest.NewRecorder()
		testEnv.router.ServeHTTP(w, req)
		require.Equal(t, testCase.StatusCode, w.Code, fmt.Sprintf("TestApiMovies %d, test case %v, wrong status", numCase, testCase))

		expectedResBody, err := json.Marshal(testCase.ResponseBody)
		if err != nil {
			t.Errorf("internal error: error while unmarshalling JSON: %s", err)
		}
		resBody, err := ioutil.ReadAll(w.Body)
		if err != nil {
			t.Errorf("error while reading response body: %s", err)
		}
		// сравниваем []byte
		require.Equal(t, resBody, expectedResBody, fmt.Sprintf("TestApiMovies %d, test case %v, wrong body", numCase, testCase))
	}
}

var testCasesWithoutCookie = []TestCase{
	{
		// регистрация: успех
		Path:   "/signup",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    "roma@mail.ru",
			"password": "password",
		},
		StatusCode: 201,
	},
	{
		// регистрация: пользователь уже существует
		Path:   "/signup",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    "roma@mail.ru",
			"password": "password",
		},
		StatusCode: 409,
	},
	{
		// вход: несуществующий пользователь
		Path:   "/signin",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    "NotExist@mail.ru",
			"password": "password",
		},
		StatusCode: 404,
	},
	{
		// вход: неверный пароль
		Path:   "/signin",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    "roma@mail.ru",
			"password": "wrongPassword",
		},
		StatusCode: 404,
	},
	{
		// вход: успех
		Path:   "/signin",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    "roma@mail.ru",
			"password": "password",
		},
		StatusCode: 200,
	},
	{
		// вход: повторно можно авторизироваться, если в запросе не было кук
		Path:   "/signin",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    "roma@mail.ru",
			"password": "password",
		},
		StatusCode: 200,
	},
	{
		// регистрация: невалидная почта
		Path:   "/signup",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    "vova.mail.ru",
			"password": "password",
		},
		StatusCode: 400,
	},
	{
		// регистрация: короткий пароль
		Path:   "/signup",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    "vova@mail.ru",
			"password": "1",
		},
		StatusCode: 400,
	},
}

func TestApiWithoutCookie(t *testing.T) {
	testEnv := NewTestEnv()

	for numCase, testCase := range testCasesWithoutCookie {
		reqBody, err := json.Marshal(testCase.RequestBody)
		if err != nil {
			t.Errorf("internal error: error while unmarshalling JSON: %s", err)
		}
		reqBodyReader := bytes.NewReader(reqBody)

		req := httptest.NewRequest(testCase.Method, testCase.Path, reqBodyReader)
		w := httptest.NewRecorder()
		testEnv.router.ServeHTTP(w, req)
		require.Equal(t, testCase.StatusCode, w.Code, fmt.Sprintf("TestApiMovies %d, test case %v, wrong status", numCase, testCase))
	}
}
