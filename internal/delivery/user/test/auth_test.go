package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/testenv"
	"github.com/stretchr/testify/require"
)

var testCasesSignInUp = []testenv.TestCase{
	{
		// регистрация: успех
		Path:   "/user/signup",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    testenv.TestUser.Email,
			"password": testenv.TestUser.Password,
		},
		StatusCode: 201,
	},
	{
		// регистрация: пользователь уже существует
		Path:   "/user/signup",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    testenv.TestUser.Email,
			"password": testenv.TestUser.Password,
		},
		StatusCode: 409,
	},
	{
		// вход: несуществующий пользователь
		Path:   "/user/signin",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    "NotExist@mail.ru",
			"password": testenv.TestUser.Password,
		},
		StatusCode: 404,
	},
	{
		// вход: неверный пароль
		Path:   "/user/signin",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    testenv.TestUser.Email,
			"password": "wrongPassword",
		},
		StatusCode: 404,
	},
	{
		// вход: успех
		Path:   "/user/signin",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    testenv.TestUser.Email,
			"password": testenv.TestUser.Password,
		},
		StatusCode: 200,
	},
	{
		// вход: повторно можно авторизироваться, если в запросе не было кук
		Path:   "/user/signin",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    testenv.TestUser.Email,
			"password": testenv.TestUser.Password,
		},
		StatusCode: 200,
	},
	{
		// регистрация: невалидная почта
		Path:   "/user/signup",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    "roma.mail.ru",
			"password": testenv.TestUser.Password,
		},
		StatusCode: 400,
	},
	{
		// регистрация: короткий пароль
		Path:   "/user/signup",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    testenv.TestUser.Email,
			"password": "1",
		},
		StatusCode: 400,
	},
	{
		// регистрация: отсутствие пароля
		Path:   "/user/signup",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email": testenv.TestUser.Email,
		},
		StatusCode: 400,
	},
}

func TestApiSignInUp(t *testing.T) {
	testEnv := testenv.NewTestEnv()

	for numCase, testCase := range testCasesSignInUp {
		reqBody, err := json.Marshal(testCase.RequestBody)
		if err != nil {
			t.Errorf("internal error: error while unmarshalling JSON: %s", err)
		}
		reqBodyReader := bytes.NewReader(reqBody)

		req := httptest.NewRequest(testCase.Method, testCase.Path, reqBodyReader)
		req.Header.Add("Content-Type", "application/json")

		w := httptest.NewRecorder()
		testEnv.Router.ServeHTTP(w, req)
		require.Equal(t, testCase.StatusCode, w.Code, fmt.Sprintf("TestApiSignInUp %d, test case %v, wrong status", numCase, testCase))
	}
}
