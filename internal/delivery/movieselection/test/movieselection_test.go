package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/setup"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/testenv"
	"github.com/stretchr/testify/require"
)

var testCasesMovies = []testenv.TestCase{
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
}

func TestApiMovies(t *testing.T) {
	testEnv := testenv.NewTestEnv()

	for numCase, testCase := range testCasesMovies {
		req := httptest.NewRequest(testCase.Method, testCase.Path, nil)
		w := httptest.NewRecorder()
		testEnv.Router.ServeHTTP(w, req)
		require.Equal(t, testCase.StatusCode, w.Code, fmt.Sprintf("TestApiMovies %d, test case %v, wrong status", numCase, testCase))

		expectedResBody, err := json.Marshal(testCase.ResponseBody)
		if err != nil {
			t.Errorf("internal error: error while unmarshalling JSON: %s", err)
		}
		resBody, err := io.ReadAll(w.Body)
		if err != nil {
			t.Errorf("error while reading response body: %s", err)
		}
		// сравниваем []byte
		require.Equal(t, resBody, expectedResBody, fmt.Sprintf("TestApiMovies %d, test case %v, wrong body", numCase, testCase))
	}
}
