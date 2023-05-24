package test

import (
	"fmt"
	"time"

	config "github.com/go-park-mail-ru/2023_1_ContentDealers/config"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/csrf"
	middlewareCSRF "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/csrf/middleware"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user"
	middlewareUser "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user/middleware"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/middleware"
	csrfUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/csrf"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	sessionDomain "github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/domain"
	"github.com/gorilla/mux"
)

type TestCase struct {
	Path         string
	Method       string
	WithCookie   bool
	WithCSRF     bool
	RequestBody  map[string]interface{}
	ResponseBody map[string]interface{}
	StatusCode   int
	MockBehavior func(user *MockUserGateway, session *MockSessionGateway)
	Extra        interface{}
}

var (
	TestUser = domain.User{
		ID:           1,
		Email:        "111@mail.ru",
		PasswordHash: "passwd111",
		AvatarURL:    "media/avatars/default_avatar.jpg",
	}
	UpdatedEmail = "222@mail.ru"
	TestSession  = sessionDomain.Session{
		ID:        "test_uuid",
		UserID:    1,
		ExpiresAt: time.Date(2024, time.January, 1, 12, 34, 56, 0, time.UTC),
	}
	CSRFSalt = "filmiumsecret123"
)

// userGateway Mock
// sessionGateway Mock

type TestRouter struct {
	Router *mux.Router
}

func NewTestRouter(userGate UserGateway, sessionGate SessionGateway, userUsecase UserUsecase) (*TestRouter, error) {
	cfgGeneral, err := config.GetCfg("config_test.yml")
	if err != nil {
		return nil, fmt.Errorf("Fail to parse config yml file: %w", err)
	}
	cfg := cfgGeneral.ApiGateway
	logger, err := logging.NewLogger(cfg.Logging, "api-gateway")
	if err != nil {
		return nil, fmt.Errorf("Fail to initialization logger: %w", err)
	}

	CSRFUseCase, err := csrfUseCase.NewUseCase(CSRFSalt, logger)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	userHandler := user.NewHandler(userGate, userUsecase, sessionGate, logger, cfg.Avatar)
	csrfHandler := csrf.NewHandler(CSRFUseCase, logger, cfg.CSRF)

	authMiddleware := middlewareUser.NewAuth(sessionGate, logger)
	CSRFMiddleware := middlewareCSRF.NewCSRF(*CSRFUseCase, logger, cfg.CSRF)
	generalMiddleware := middleware.NewGeneral(logger)

	router := mux.NewRouter()

	router.Use(generalMiddleware.AccessLog)
	// после AccessLog, потому что AccessLog навешивает декоратор на ResponseWriter
	router.Use(generalMiddleware.Metrics)
	router.Use(generalMiddleware.Panic)
	router.Use(generalMiddleware.SetContentTypeJSON)

	authRouter := router.Methods("GET", "POST").Subrouter()
	unAuthRouter := router.Methods("GET", "POST").Subrouter()

	authRouter.Use(authMiddleware.RequireAuth)
	authRouter.Use(CSRFMiddleware.RequireCSRF)
	unAuthRouter.Use(authMiddleware.RequireUnAuth)

	unAuthRouter.HandleFunc("/user/signin", userHandler.SignIn).Methods("POST")
	unAuthRouter.HandleFunc("/user/signup", userHandler.SignUp).Methods("POST")

	authRouter.HandleFunc("/user/logout", userHandler.Logout).Methods("POST")
	authRouter.HandleFunc("/user/profile", userHandler.Info).Methods("GET")
	authRouter.HandleFunc("/user/csrf", csrfHandler.GetCSRF).Methods("GET")

	authRouter.HandleFunc("/user/update", userHandler.Update).Methods("POST")

	authRouter.HandleFunc("/user/avatar/update", userHandler.UpdateAvatar).Methods("POST")
	authRouter.HandleFunc("/user/avatar/delete", userHandler.DeleteAvatar).Methods("POST")

	return &TestRouter{
		Router: router,
	}, nil

}
