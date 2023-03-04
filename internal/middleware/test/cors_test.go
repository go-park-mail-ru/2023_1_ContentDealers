package test

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/test_env"
	"github.com/stretchr/testify/require"
)

func TestCors(t *testing.T) {
	testEnv := test_env.NewTestEnv()

	// запрос с заголовком Origin
	req := httptest.NewRequest("POST", "/selections", nil)
	req.Header.Add("Origin", test_env.TestOrigin)
	w := httptest.NewRecorder()
	testEnv.Router.ServeHTTP(w, req)
	resAllowOrigin := w.Result().Header.Get("Access-Control-Allow-Origin")
	resAllowCredentials := w.Result().Header.Get("Access-Control-Allow-Credentials")

	require.Equal(t, test_env.TestOrigin, resAllowOrigin, fmt.Sprintf("TestCors, wrong header origin %s, expected %s", resAllowOrigin, test_env.TestOrigin))
	require.Equal(t, "true", resAllowCredentials, fmt.Sprintf("TestCors, wrong header origin %s, expected %s", resAllowOrigin, test_env.TestOrigin))

	// запрос без заголовка Origin

	req = httptest.NewRequest("POST", "/selections", nil)
	w = httptest.NewRecorder()
	testEnv.Router.ServeHTTP(w, req)

	resAllowOrigin = w.Result().Header.Get("Access-Control-Allow-Origin")
	resAllowCredentials = w.Result().Header.Get("Access-Control-Allow-Credentials")

	require.Equal(t, "", resAllowOrigin, fmt.Sprintf("TestCors, wrong header origin %s, expected %s", resAllowOrigin, test_env.TestOrigin))
	require.Equal(t, "", resAllowCredentials, fmt.Sprintf("TestCors, wrong header origin %s, expected %s", resAllowOrigin, test_env.TestOrigin))

}
