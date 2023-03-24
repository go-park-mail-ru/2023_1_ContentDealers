package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/testenv"
	"github.com/stretchr/testify/require"
)

var testCasesWithCookie = []testenv.TestCase{
	{
		// signup с куки: запрещено
		Path:   "/user/signup",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    testenv.TestUser.Email,
			"password": testenv.TestUser.Password,
		},
		WithCookie: true,
		StatusCode: 403,
	},
	{
		// signin с куки: запрещено
		Path:   "/user/signin",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    testenv.TestUser.Email,
			"password": testenv.TestUser.Password,
		},
		WithCookie: true,
		StatusCode: 403,
	},
	{
		// profile с куки: разрешено
		Path:       "/user/profile",
		Method:     "GET",
		WithCookie: true,
		StatusCode: 200,
	},
	{
		// для selections кука опциональна
		Path:       "/selections",
		Method:     "GET",
		WithCookie: true,
		StatusCode: 200,
	},
	{
		// для selections кука опциональна
		Path:       "/selections",
		Method:     "GET",
		WithCookie: false,
		StatusCode: 200,
	},
	{
		// profile без куки: запрещено
		Path:       "/user/profile",
		Method:     "GET",
		WithCookie: false,
		StatusCode: 401,
	},
	{
		// logout без куки: запрещено
		Path:       "/user/logout",
		Method:     "POST",
		WithCookie: false,
		StatusCode: 401,
	},
	{
		// logout с куки: разрешено
		Path:       "/user/logout",
		Method:     "POST",
		WithCookie: true,
		StatusCode: 200,
	},
	{
		// profile с "протухшей" кукой: запрещено
		Path:       "/user/profile",
		Method:     "GET",
		WithCookie: true,
		StatusCode: 401,
	},
	{
		// signin с "протухшей" кукой: на сервере она удалена, поэтому сервер просто выдаст новую
		Path:   "/user/signin",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    testenv.TestUser.Email,
			"password": testenv.TestUser.Password,
		},
		WithCookie: true,
		StatusCode: 200,
	},
}

func TestApiCookie(t *testing.T) {
	testEnv := testenv.NewTestEnv()
	// регистрация
	reqBody, err := json.Marshal(testenv.TestUser)
	if err != nil {
		t.Errorf("internal error: error while unmarshalling JSON: %s", err)
	}
	reqBodyReader := bytes.NewReader(reqBody)
	req := httptest.NewRequest("POST", "/user/signup", reqBodyReader)
	req.Header.Add("Content-Type", "application/json")

	w := httptest.NewRecorder()
	testEnv.Router.ServeHTTP(w, req)
	require.Equal(t, http.StatusCreated, w.Code, fmt.Sprintf("TestApiCookie signup, wrong status %d, expected %d", w.Code, http.StatusCreated))

	// авторизация

	reqBodyReader = bytes.NewReader(reqBody)
	req = httptest.NewRequest("POST", "/user/signin", reqBodyReader)
	req.Header.Add("Content-Type", "application/json")
	w = httptest.NewRecorder()
	testEnv.Router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code, fmt.Sprintf("TestApiCookie signin, wrong status %d, expected %d", w.Code, http.StatusOK))
	cookie := w.Result().Header.Get("Set-Cookie")

	for numCase, testCase := range testCasesWithCookie {
		var reqBodyReader io.Reader // reqBodyReader = nil
		if testCase.RequestBody != nil {
			reqBodyReader = bytes.NewReader(reqBody)
		}
		req = httptest.NewRequest(testCase.Method, testCase.Path, reqBodyReader)
		if testCase.WithCookie {
			req.Header.Add("Cookie", cookie)
		}
		if reqBodyReader != nil {
			req.Header.Add("Content-Type", "application/json")
		}
		w = httptest.NewRecorder()
		testEnv.Router.ServeHTTP(w, req)
		require.Equal(t, testCase.StatusCode, w.Code, fmt.Sprintf("TestApiCookie [%d], wrong status %d, expected %d", numCase, w.Code, testCase.StatusCode))
	}
}
