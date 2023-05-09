package setup

import (
	"io"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/content"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/csrf"
	middlewareCSRF "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/csrf/middleware"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/favorites"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/genre"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/person"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/rating"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/search"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/selection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user"
	middlewareUser "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user/middleware"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/middleware"
	csrfUseCase "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/csrf"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, `{"status":404}`)
	w.Header().Set("Content-Type", "application/json")
}

type SettingsRouter struct {
	AllowedOrigins   []string
	FavHandler       favorites.Handler
	RateHandler      rating.Handler
	UserHandler      user.Handler
	CSRFHandler      csrf.Handler
	SelectionHandler selection.Handler
	ContentHandler   content.Handler
	PersonHandler    person.Handler
	SessionGateway   SessionGateway
	CSRFUseCase      csrfUseCase.UseCase
	SearchHandler    search.Handler
	GenreHandler     genre.Handler
	Logger           logging.Logger
	CSRFConfig       csrf.CSRFConfig
}

type FakeLogger struct {
}

func (f FakeLogger) Printf(string, ...interface{}) {
}

func Routes(s *SettingsRouter) *mux.Router {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   s.AllowedOrigins,
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PATCH"},
		AllowCredentials: true,
		Debug:            true,
	})
	corsMiddleware.Log = FakeLogger{}
	authMiddleware := middlewareUser.NewAuth(s.SessionGateway, s.Logger)
	CSRFMiddleware := middlewareCSRF.NewCSRF(s.CSRFUseCase, s.Logger, s.CSRFConfig)
	generalMiddleware := middleware.NewGeneral(s.Logger)

	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(NotFound)

	router.Use(generalMiddleware.AccessLog)
	// после AccessLog, потому что AccessLog навешивает декоратор на ResponseWriter
	router.Use(generalMiddleware.Metrics)
	router.Use(generalMiddleware.Panic)
	router.Use(corsMiddleware.Handler)
	router.Use(generalMiddleware.SetContentTypeJSON)

	authRouter := router.Methods("GET", "POST").Subrouter()
	unAuthRouter := router.Methods("GET", "POST").Subrouter()

	authRouter.Use(authMiddleware.RequireAuth)
	authRouter.Use(CSRFMiddleware.RequireCSRF)
	unAuthRouter.Use(authMiddleware.RequireUnAuth)

	router.Handle("/metrics", promhttp.Handler())

	router.HandleFunc("/selections", s.SelectionHandler.GetAll)
	router.HandleFunc("/selections/{id:[0-9]+}", s.SelectionHandler.GetByID)
	router.HandleFunc("/persons/{id:[0-9]+}", s.PersonHandler.GetByID)
	router.HandleFunc("/films/{content_id:[0-9]+}", s.ContentHandler.GetFilmByContentID)
	router.HandleFunc("/series/{content_id:[0-9]+}", s.ContentHandler.GetSeriesByContentID)
	router.HandleFunc("/search", s.SearchHandler.Search)
	router.HandleFunc("/genres", s.GenreHandler.GetAll)
	router.HandleFunc("/genres/{id:[0-9]+}", s.GenreHandler.GetContentByID)

	unAuthRouter.HandleFunc("/user/signin", s.UserHandler.SignIn).Methods("POST")
	unAuthRouter.HandleFunc("/user/signup", s.UserHandler.SignUp).Methods("POST")

	authRouter.HandleFunc("/user/logout", s.UserHandler.Logout).Methods("POST")
	authRouter.HandleFunc("/user/profile", s.UserHandler.Info).Methods("GET")
	authRouter.HandleFunc("/user/csrf", s.CSRFHandler.GetCSRF).Methods("GET")

	authRouter.HandleFunc("/user/update", s.UserHandler.Update).Methods("POST")

	authRouter.HandleFunc("/favorites/content", s.FavHandler.GetFavContent).Methods("GET")
	authRouter.HandleFunc("/favorites/content/{id:[0-9]+}/has", s.FavHandler.HasFavContent).Methods("GET")
	authRouter.HandleFunc("/favorites/content/add", s.FavHandler.AddFavContent).Methods("POST")
	authRouter.HandleFunc("/favorites/content/delete", s.FavHandler.DeleteFavContent).Methods("POST")

	authRouter.HandleFunc("/rating", s.RateHandler.GetRatingByUser).Methods("GET")
	authRouter.HandleFunc("/rating/content/{id:[0-9]+}/has", s.RateHandler.HasRating).Methods("GET")
	authRouter.HandleFunc("/rating/add", s.RateHandler.AddRating).Methods("POST")
	authRouter.HandleFunc("/rating/delete", s.RateHandler.DeleteRating).Methods("POST")

	authRouter.HandleFunc("/user/avatar/update", s.UserHandler.UpdateAvatar).Methods("POST")
	authRouter.HandleFunc("/user/avatar/delete", s.UserHandler.DeleteAvatar).Methods("POST")

	return router
}
