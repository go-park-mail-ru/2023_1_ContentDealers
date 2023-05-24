package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"testing"

	domainSession "github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/domain"
	domainUser "github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/domain"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var Cookie string
var CSRFToken string

var testCasesAuthorization = []TestCase{
	{
		// регистрация: успех
		Path:   "/user/signup",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    TestUser.Email,
			"password": TestUser.PasswordHash,
		},
		StatusCode: 200,
		MockBehavior: func(userGate *MockUserGateway, sessionGate *MockSessionGateway) {
			userReq := domainUser.User{
				Email:        TestUser.Email,
				PasswordHash: TestUser.PasswordHash,
				AvatarURL:    TestUser.AvatarURL,
			}
			userRes := domainUser.User{
				ID:           TestUser.ID,
				Email:        TestUser.Email,
				PasswordHash: TestUser.PasswordHash,
				AvatarURL:    TestUser.AvatarURL,
			}
			userGate.EXPECT().Register(gomock.Any(), userReq).Return(userRes, nil)
		},
	},
	{
		// вход: успех
		Path:   "/user/signin",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    TestUser.Email,
			"password": TestUser.PasswordHash,
		},
		StatusCode: 200,
		MockBehavior: func(userGate *MockUserGateway, sessionGate *MockSessionGateway) {
			userReq := domainUser.User{
				Email:        TestUser.Email,
				PasswordHash: TestUser.PasswordHash,
			}
			userRes := domainUser.User{
				ID:           TestUser.ID,
				Email:        TestUser.Email,
				PasswordHash: TestUser.PasswordHash,
				AvatarURL:    TestUser.AvatarURL,
			}
			userReqForCreateSess := domainUser.User{
				ID:           TestUser.ID,
				Email:        TestUser.Email,
				PasswordHash: TestUser.PasswordHash,
				AvatarURL:    TestUser.AvatarURL,
			}
			sessionRes := domainSession.Session{
				ID:        TestSession.ID,
				UserID:    TestSession.UserID,
				ExpiresAt: TestSession.ExpiresAt,
			}
			userGate.EXPECT().Auth(gomock.Any(), userReq).Return(userRes, nil)

			sessionGate.EXPECT().Create(gomock.Any(), userReqForCreateSess).Return(sessionRes, nil)
		},
	},
}

var testCasesCSRF = []TestCase{
	{
		// получение csrf без куки
		Path:        "/user/csrf",
		Method:      "GET",
		RequestBody: nil,
		StatusCode:  400,
		WithCookie:  false,
		MockBehavior: func(userGate *MockUserGateway, sessionGate *MockSessionGateway) {

			sessionRes := domainSession.Session{
				ID:        TestSession.ID,
				UserID:    TestSession.UserID,
				ExpiresAt: TestSession.ExpiresAt,
			}

			sessionGate.EXPECT().Get(gomock.Any(), sessionRes.ID).Return(sessionRes, nil)
		},
	},
	{
		// получение csrf с куки
		Path:         "/user/csrf",
		Method:       "GET",
		RequestBody:  nil,
		StatusCode:   200,
		WithCookie:   true,
		MockBehavior: func(userGate *MockUserGateway, sessionGate *MockSessionGateway) {},
	},
}

var testCasesAccountActions = []TestCase{
	{
		// посмотр данных
		Path:        "/user/profile",
		Method:      "GET",
		RequestBody: nil,
		StatusCode:  200,
		WithCookie:  true,
		MockBehavior: func(userGate *MockUserGateway, sessionGate *MockSessionGateway) {
			sessionRes := domainSession.Session{
				ID:        TestSession.ID,
				UserID:    TestSession.UserID,
				ExpiresAt: TestSession.ExpiresAt,
			}
			userRes := domainUser.User{
				ID:           TestUser.ID,
				Email:        UpdatedEmail,
				PasswordHash: TestUser.PasswordHash,
			}

			sessionGate.EXPECT().Get(gomock.Any(), TestSession.ID).Return(sessionRes, nil)

			userGate.EXPECT().GetByID(gomock.Any(), sessionRes.UserID).Return(userRes, nil)
		},
	},
	{
		// обновление данных
		Path:   "/user/update",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    UpdatedEmail,
			"password": TestUser.PasswordHash,
		},
		StatusCode: 200,
		WithCookie: true,
		WithCSRF:   true,
		MockBehavior: func(userGate *MockUserGateway, sessionGate *MockSessionGateway) {
			sessionRes := domainSession.Session{
				ID:        TestSession.ID,
				UserID:    TestSession.UserID,
				ExpiresAt: TestSession.ExpiresAt,
			}
			userRes := domainUser.User{
				ID:           TestUser.ID,
				Email:        UpdatedEmail,
				PasswordHash: TestUser.PasswordHash,
			}

			sessionGate.EXPECT().Get(gomock.Any(), TestSession.ID).Return(sessionRes, nil)

			userGate.EXPECT().GetByID(gomock.Any(), sessionRes.UserID).Return(userRes, nil)

			userGate.EXPECT().Update(gomock.Any(), userRes).Return(nil)
		},
	},
	{
		// обновление данных бек токена
		Path:   "/user/update",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    UpdatedEmail,
			"password": TestUser.PasswordHash,
		},
		StatusCode: 400,
		WithCookie: true,
		WithCSRF:   false,
		MockBehavior: func(userGate *MockUserGateway, sessionGate *MockSessionGateway) {
			sessionRes := domainSession.Session{
				ID:        TestSession.ID,
				UserID:    TestSession.UserID,
				ExpiresAt: TestSession.ExpiresAt,
			}

			sessionGate.EXPECT().Get(gomock.Any(), TestSession.ID).Return(sessionRes, nil)
		},
	},

	// выход
	{
		Path:        "/user/logout",
		Method:      "POST",
		RequestBody: nil,
		StatusCode:  200,
		WithCookie:  true,
		WithCSRF:    true,
		MockBehavior: func(userGate *MockUserGateway, sessionGate *MockSessionGateway) {
			sessionRes := domainSession.Session{
				ID:        TestSession.ID,
				UserID:    TestSession.UserID,
				ExpiresAt: TestSession.ExpiresAt,
			}
			sessionGate.EXPECT().Get(gomock.Any(), TestSession.ID).Return(sessionRes, nil)
			sessionGate.EXPECT().Delete(gomock.Any(), TestSession.ID).Return(nil)
		},
	},
}

var testCasesUploadAvatar = []TestCase{
	{
		// обновление данных бек токена
		Path:       "/user/avatar/update",
		Method:     "POST",
		StatusCode: 200,
		WithCookie: true,
		WithCSRF:   true,
		MockBehavior: func(userGate *MockUserGateway, sessionGate *MockSessionGateway) {
			sessionRes := domainSession.Session{
				ID:        TestSession.ID,
				UserID:    TestSession.UserID,
				ExpiresAt: TestSession.ExpiresAt,
			}
			userRes := domainUser.User{
				ID:           TestUser.ID,
				Email:        UpdatedEmail,
				PasswordHash: TestUser.PasswordHash,
			}

			sessionGate.EXPECT().Get(gomock.Any(), TestSession.ID).Return(sessionRes, nil)

			userGate.EXPECT().GetByID(gomock.Any(), sessionRes.UserID).Return(userRes, nil)

			userGate.EXPECT().UpdateAvatar(gomock.Any(), userRes, gomock.Any()).Return(userRes, nil)
		},
	},
}

var testCasesDeleteAvatar = []TestCase{
	{
		Path:        "/user/avatar/delete",
		Method:      "POST",
		RequestBody: nil,
		StatusCode:  200,
		WithCookie:  true,
		WithCSRF:    true,
		MockBehavior: func(userGate *MockUserGateway, sessionGate *MockSessionGateway) {
			sessionRes := domainSession.Session{
				ID:        TestSession.ID,
				UserID:    TestSession.UserID,
				ExpiresAt: TestSession.ExpiresAt,
			}
			userRes := domainUser.User{
				ID:           TestUser.ID,
				Email:        UpdatedEmail,
				PasswordHash: TestUser.PasswordHash,
			}
			sessionGate.EXPECT().Get(gomock.Any(), TestSession.ID).Return(sessionRes, nil)
			userGate.EXPECT().GetByID(gomock.Any(), sessionRes.UserID).Return(userRes, nil)
			userGate.EXPECT().DeleteAvatar(gomock.Any(), userRes).Return(nil)
		},
	},
}

func TestUserAccountHandlers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userGatewayMock := NewMockUserGateway(ctrl)
	sessionGatewayMock := NewMockSessionGateway(ctrl)
	userUsecase := NewMockUserUsecase(ctrl)

	testRouter, err := NewTestRouter(userGatewayMock, sessionGatewayMock, userUsecase)
	if err != nil {
		t.Errorf("fail to init router: %v", err)
	}

	var w *httptest.ResponseRecorder

	for numCase, testCase := range testCasesAuthorization {
		var reqBodyReader io.Reader
		if testCase.RequestBody != nil {
			reqBody, err := json.Marshal(testCase.RequestBody)
			if err != nil {
				t.Errorf("internal error: error while unmarshalling JSON: %s", err)
			}
			reqBodyReader = bytes.NewReader(reqBody)
		}
		req := httptest.NewRequest(testCase.Method, testCase.Path, reqBodyReader)
		req.Header.Add("Content-Type", "application/json")

		w = httptest.NewRecorder()

		testCase.MockBehavior(userGatewayMock, sessionGatewayMock)

		testRouter.Router.ServeHTTP(w, req)

		require.Equal(t, testCase.StatusCode, w.Code, fmt.Sprintf("TestApiCookie [%d], wrong status %d, expected %d", numCase, w.Code, testCase.StatusCode))
	}

	// получение куки
	Cookie = w.Result().Header.Get("Set-Cookie")

	for numCase, testCase := range testCasesCSRF {
		req := httptest.NewRequest(testCase.Method, testCase.Path, nil)
		req.Header.Add("Content-Type", "application/json")
		if testCase.WithCookie {
			req.Header.Add("Cookie", Cookie)
		}

		testCase.MockBehavior(userGatewayMock, sessionGatewayMock)

		w = httptest.NewRecorder()
		testRouter.Router.ServeHTTP(w, req)

		require.Equal(t, testCase.StatusCode, w.Code, fmt.Sprintf("TestApiCookie [%d], wrong status %d, expected %d", numCase, w.Code, testCase.StatusCode))
	}

	// получение токена
	csrfMap := map[string]interface{}{}
	decoder := json.NewDecoder(w.Body)
	err = decoder.Decode(&csrfMap)
	if err != nil {
		t.Errorf("fail to decode csrf: %v", err)
	}
	body := csrfMap["body"].(map[string]interface{})
	CSRFToken = body["csrf-token"].(string)

	for numCase, testCase := range testCasesAccountActions {
		var reqBodyReader io.Reader
		if testCase.RequestBody != nil {
			reqBody, err := json.Marshal(testCase.RequestBody)
			if err != nil {
				t.Errorf("internal error: error while unmarshalling JSON: %s", err)
			}
			reqBodyReader = bytes.NewReader(reqBody)
		}
		req := httptest.NewRequest(testCase.Method, testCase.Path, reqBodyReader)
		req.Header.Add("Content-Type", "application/json")
		if testCase.WithCookie {
			req.Header.Add("Cookie", Cookie)
		}
		if testCase.WithCSRF {
			req.Header.Add("csrf-token", CSRFToken)
		}

		w := httptest.NewRecorder()

		testCase.MockBehavior(userGatewayMock, sessionGatewayMock)

		testRouter.Router.ServeHTTP(w, req)

		require.Equal(t, testCase.StatusCode, w.Code, fmt.Sprintf("TestApiCookie [%d], wrong status %d, expected %d", numCase, w.Code, testCase.StatusCode))
	}

	// отправка аватарки с content-type: multipart/from-data
	for numCase, testCase := range testCasesUploadAvatar {
		body := new(bytes.Buffer)
		mw := multipart.NewWriter(body)

		file, err := os.Open("avatar.jpg")
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		writer, err := mw.CreateFormFile("avatar", "avatar.jpg")
		if err != nil {
			t.Fatal(err)
		}

		if _, err := io.Copy(writer, file); err != nil {
			t.Fatal(err)
		}

		mw.Close()

		req := httptest.NewRequest(testCase.Method, testCase.Path, body)
		req.Header.Add("Content-Type", mw.FormDataContentType())
		if testCase.WithCookie {
			req.Header.Add("Cookie", Cookie)
		}
		if testCase.WithCSRF {
			req.Header.Add("csrf-token", CSRFToken)
		}

		w := httptest.NewRecorder()

		testCase.MockBehavior(userGatewayMock, sessionGatewayMock)

		testRouter.Router.ServeHTTP(w, req)

		require.Equal(t, testCase.StatusCode, w.Code, fmt.Sprintf("TestApiCookie [%d], wrong status %d, expected %d", numCase, w.Code, testCase.StatusCode))
	}

	// удаление аватарки
	for numCase, testCase := range testCasesDeleteAvatar {
		req := httptest.NewRequest(testCase.Method, testCase.Path, nil)
		req.Header.Add("Content-Type", "application/json")
		if testCase.WithCookie {
			req.Header.Add("Cookie", Cookie)
		}
		if testCase.WithCSRF {
			req.Header.Add("csrf-token", CSRFToken)
		}

		testCase.MockBehavior(userGatewayMock, sessionGatewayMock)

		w = httptest.NewRecorder()
		testRouter.Router.ServeHTTP(w, req)

		require.Equal(t, testCase.StatusCode, w.Code, fmt.Sprintf("TestApiCookie [%d], wrong status %d, expected %d", numCase, w.Code, testCase.StatusCode))
	}
}
