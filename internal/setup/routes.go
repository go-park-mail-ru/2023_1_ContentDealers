package setup

import (
	"io"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/csrf"
	middlewareCSRF "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/csrf/middleware"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/movieselection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user/middleware"
	csrfUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/csrf"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	_ "github.com/go-park-mail-ru/2023_1_ContentDealers/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, `{"status":404}`)
	w.Header().Set("Content-Type", "application/json")
}

type SettingsRouter struct {
	AllowedOrigins        []string
	UserHandler           user.Handler
	CSRFHandler           csrf.Handler
	MovieSelectionHandler movieselection.Handler
	SessionUseCase        SessionUseCase
	CSRFUseCase           csrfUseCase.CSRF
	Logger                logging.Logger
}

func Routes(s *SettingsRouter) *mux.Router {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   s.AllowedOrigins,
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PATCH"},
		AllowCredentials: true,
		Debug:            true,
	})
	corsMiddleware.Log = s.Logger
	authMiddleware := middleware.NewAuth(s.SessionUseCase)
	CSRFMiddleware := middlewareCSRF.NewCSRF(s.CSRFUseCase)

	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(NotFound)

	authRouter := router.Methods("GET", "POST").Subrouter()
	unAuthRouter := router.Methods("GET", "POST").Subrouter()

	router.Use(corsMiddleware.Handler)
	router.Use(middleware.SetContentTypeJSON)
	authRouter.Use(authMiddleware.RequireAuth)
	authRouter.Use(CSRFMiddleware.RequireCSRF)
	unAuthRouter.Use(authMiddleware.RequireUnAuth)

	router.HandleFunc("/selections", s.MovieSelectionHandler.GetAll)
	router.HandleFunc("/selections/{id:[0-9]+}", s.MovieSelectionHandler.GetByID)

	unAuthRouter.HandleFunc("/user/signin", s.UserHandler.SignIn).Methods("POST")
	unAuthRouter.HandleFunc("/user/signup", s.UserHandler.SignUp).Methods("POST")

	authRouter.HandleFunc("/user/logout", s.UserHandler.Logout).Methods("POST")
	authRouter.HandleFunc("/user/profile", s.UserHandler.Info).Methods("GET")
	// только авторизированные могут запрашивать токен
	authRouter.HandleFunc("/user/csrf", s.CSRFHandler.GetCSRF).Methods("GET")

	// TODO: PATCH в постмане выдавал 405 Method not allowed
	authRouter.HandleFunc("/user/avatar/upload", s.UserHandler.UpdateAvatar).Methods("POST")
	// authRouter.HandleFunc("/user/update/profile", s.UserHandler.UploadAvatar).Methods("GET")

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(httpSwagger.URL("./swagger/doc.json")))

	return router
}
