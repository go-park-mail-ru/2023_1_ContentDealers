package setup

import (
	"io"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/movieselection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user/csrf"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user/middleware"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/setup/logger"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	_ "github.com/go-park-mail-ru/2023_1_ContentDealers/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"status":404}`)
}

type SettingsRouter struct {
	AllowedOrigins        []string
	UserHandler           user.Handler
	MovieSelectionHandler movieselection.Handler
	SessionUseCase        SessionUseCase
	CryptToken            csrf.CryptToken
	Logger                logger.Logger
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
	CSRFMiddleware := middleware.NewCSRF(s.CryptToken)

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
	authRouter.HandleFunc("/user/csrf", s.UserHandler.GetCSRF).Methods("GET")

	// TODO: PATCH в постмане выдавал 405 Method not allowed
	authRouter.HandleFunc("/user/avatar/upload", s.UserHandler.UpdateAvatar).Methods("POST")
	// authRouter.HandleFunc("/user/update/profile", s.UserHandler.UploadAvatar).Methods("GET")

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(httpSwagger.URL("./swagger/doc.json")))

	return router
}
