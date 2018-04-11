package http

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"github.com/speix/movierama/app/data"
)

// AppServer wraps application's http server
// given the routes and a storage object.
func AppServer(storage *data.Storage) *http.Server {

	routes := createRoutes(storage)

	return &http.Server{
		Handler:      newRouter(routes),
		Addr:         ":" + os.Getenv("MR_SERVER_PORT"),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

}

// NewRouter wraps default mux.Router
// to easily add routes through routes.go.
func newRouter(routes Routes) *mux.Router {

	// Initialize http Router
	router := mux.NewRouter().StrictSlash(true)

	// Parse routes
	for _, route := range routes {

		var handler http.Handler
		handler = route.HandlerFunc

		// Register middlewares if any
		if route.MiddleWares != nil {
			for m := range route.MiddleWares {
				handler = route.MiddleWares[m](handler)
			}
		}

		// Register routes to be matched and dispatched to each handler.
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	// Serve static routes
	index := http.StripPrefix("/", http.FileServer(http.Dir("./static")))
	router.PathPrefix("/").Handler(index)

	return router
}
