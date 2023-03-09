package test

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/testenv"
	"github.com/stretchr/testify/require"
)

func TestCors(t *testing.T) {
	testEnv := testenv.NewTestEnv()

	// запрос с заголовком Origin
	req := httptest.NewRequest("POST", "/selections", nil)
	req.Header.Add("Origin", testenv.TestOrigin)
	w := httptest.NewRecorder()
	testEnv.Router.ServeHTTP(w, req)
	resAllowOrigin := w.Result().Header.Get("Access-Control-Allow-Origin")
	resAllowCredentials := w.Result().Header.Get("Access-Control-Allow-Credentials")

	require.Equal(t, testenv.TestOrigin, resAllowOrigin,
		fmt.Sprintf("TestCors, wrong header origin %s, expected %s", resAllowOrigin, testenv.TestOrigin))
	require.Equal(t, "true", resAllowCredentials,
		fmt.Sprintf("TestCors, wrong header origin %s, expected %s", resAllowOrigin, testenv.TestOrigin))

	// запрос без заголовка Origin

	req = httptest.NewRequest("POST", "/selections", nil)
	w = httptest.NewRecorder()
	testEnv.Router.ServeHTTP(w, req)

	resAllowOrigin = w.Result().Header.Get("Access-Control-Allow-Origin")
	resAllowCredentials = w.Result().Header.Get("Access-Control-Allow-Credentials")

	require.Equal(t, "", resAllowOrigin,
		fmt.Sprintf("TestCors, wrong header origin %s, expected %s", resAllowOrigin, testenv.TestOrigin))
	require.Equal(t, "", resAllowCredentials,
		fmt.Sprintf("TestCors, wrong header origin %s, expected %s", resAllowOrigin, testenv.TestOrigin))

}
