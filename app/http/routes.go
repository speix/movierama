package http

import (
	"net/http"

	"github.com/speix/movierama/app/http/handlers"

	"github.com/speix/movierama/app/data"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string               // Route identifier by name
	Method      string               // Http Method GET|POST|PUT|DELETE
	Pattern     string               // RegEx for request routing
	HandlerFunc http.HandlerFunc     // Handler function for this route
	MiddleWares []mux.MiddlewareFunc // Middleware function for this route.
}

type Routes []Route

func createRoutes(storage *data.Storage) Routes {

	return Routes{
		Route{
			"User Registration Route",
			"POST",
			"/users",
			handlers.UserHandler{UserService: storage, SessionService: storage}.Register,
			[]mux.MiddlewareFunc{MWHandler{}.userRegister},
		},
		Route{
			"User Login Route",
			"POST",
			"/users/session",
			handlers.UserHandler{UserService: storage, SessionService: storage}.Login,
			[]mux.MiddlewareFunc{MWHandler{}.userLogin},
		},
		Route{
			"User Logout Route",
			"DELETE",
			"/users/session",
			handlers.UserHandler{SessionService: storage}.Logout,
			nil,
		},
		Route{
			"User Session Status",
			"GET",
			"/users/session",
			handlers.UserHandler{SessionService: storage}.Status,
			nil,
		},
		Route{
			"Add Movie Route",
			"POST",
			"/movies",
			handlers.MovieHandler{MovieService: storage, SessionService: storage}.Add,
			[]mux.MiddlewareFunc{MWHandler{SessionService: storage}.auth, MWHandler{}.movieAdd},
		},
		Route{
			"Add Vote Route",
			"POST",
			"/movies/{movie_id:[0-9]+}/vote/{positive:[0-1]}",
			handlers.VoteHandler{VoteService: storage, SessionService: storage}.Vote,
			[]mux.MiddlewareFunc{MWHandler{SessionService: storage}.auth},
		},
		Route{
			"Retract Vote Route",
			"DELETE",
			"/movies/{movie_id:[0-9]+}/vote",
			handlers.VoteHandler{VoteService: storage, SessionService: storage}.Retract,
			[]mux.MiddlewareFunc{MWHandler{SessionService: storage}.auth},
		},
		Route{
			"Sort Movies Route",
			"GET",
			"/movies/{sort_by}/{user:[0-9]+}",
			handlers.MovieHandler{MovieService: storage, SessionService: storage}.Sort,
			nil,
		},
		Route{
			"Sort Movies Route No User",
			"GET",
			"/movies/{sort_by}",
			handlers.MovieHandler{MovieService: storage, SessionService: storage}.Sort,
			nil,
		},
	}
}
