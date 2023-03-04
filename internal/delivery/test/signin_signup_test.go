package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/test_env"
	"github.com/stretchr/testify/require"
)

var testCasesSignInUp = []test_env.TestCase{
	{
		// регистрация: успех
		Path:   "/signup",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    test_env.TestUser.Email,
			"password": test_env.TestUser.Password,
		},
		StatusCode: 201,
	},
	{
		// регистрация: пользователь уже существует
		Path:   "/signup",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    test_env.TestUser.Email,
			"password": test_env.TestUser.Password,
		},
		StatusCode: 409,
	},
	{
		// вход: несуществующий пользователь
		Path:   "/signin",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    "NotExist@mail.ru",
			"password": test_env.TestUser.Password,
		},
		StatusCode: 404,
	},
	{
		// вход: неверный пароль
		Path:   "/signin",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    test_env.TestUser.Email,
			"password": "wrongPassword",
		},
		StatusCode: 404,
	},
	{
		// вход: успех
		Path:   "/signin",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    test_env.TestUser.Email,
			"password": test_env.TestUser.Password,
		},
		StatusCode: 200,
	},
	{
		// вход: повторно можно авторизироваться, если в запросе не было кук
		Path:   "/signin",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    test_env.TestUser.Email,
			"password": test_env.TestUser.Password,
		},
		StatusCode: 200,
	},
	{
		// регистрация: невалидная почта
		Path:   "/signup",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    "roma.mail.ru",
			"password": test_env.TestUser.Password,
		},
		StatusCode: 400,
	},
	{
		// регистрация: короткий пароль
		Path:   "/signup",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    test_env.TestUser.Email,
			"password": "1",
		},
		StatusCode: 400,
	},
	{
		// регистрация: отсутствие пароля
		Path:   "/signup",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email": test_env.TestUser.Email,
		},
		StatusCode: 400,
	},
}

func TestApiSignInUp(t *testing.T) {
	testEnv := test_env.NewTestEnv()

	for numCase, testCase := range testCasesSignInUp {
		reqBody, err := json.Marshal(testCase.RequestBody)
		if err != nil {
			t.Errorf("internal error: error while unmarshalling JSON: %s", err)
		}
		reqBodyReader := bytes.NewReader(reqBody)

		req := httptest.NewRequest(testCase.Method, testCase.Path, reqBodyReader)
		w := httptest.NewRecorder()
		testEnv.Router.ServeHTTP(w, req)
		require.Equal(t, testCase.StatusCode, w.Code, fmt.Sprintf("TestApiSignInUp %d, test case %v, wrong status", numCase, testCase))
	}
}
