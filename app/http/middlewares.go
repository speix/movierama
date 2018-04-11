package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/speix/movierama/app/errors"

	"github.com/speix/movierama/app"
	"github.com/speix/movierama/app/helpers"
)

type MWHandler struct {
	app.SessionService
	errors.ApiResponseService
}

// userRegister middleware handles input errors on user registration.
// If no error occurs, it forwards the request along with a context to the next HandlerFunc.
func (mh MWHandler) userRegister(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user := new(app.User)

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			mh.Respond(w, http.StatusBadRequest, err.Error())
			return
		}

		if !helpers.ValidEmail(user.Email) {
			mh.Respond(w, http.StatusBadRequest, " User email is invalid!")
			return
		}

		if len(user.Name) == 0 {
			mh.Respond(w, http.StatusBadRequest, " User name is empty!")
			return
		}

		if len(user.Password) < 5 {
			mh.Respond(w, http.StatusBadRequest, " Password should be at least 6 characters long!")
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

// userLogin middleware handles input errors on user login process.
// If no error occurs, it forwards the request along with a context to the next HandlerFunc.
func (mh MWHandler) userLogin(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user := new(app.User)

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			mh.Respond(w, http.StatusBadRequest, err.Error())
			return
		}

		if !helpers.ValidEmail(user.Email) {
			mh.Respond(w, http.StatusBadRequest, " User email is invalid!")
			return
		}

		if len(user.Password) == 0 {
			mh.Respond(w, http.StatusBadRequest, " User password is empty!")
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

// movieAdd middleware handles input errors on movie addition process.
// If no error occurs, it forwards the request along with a context to the next HandlerFunc.
func (mh MWHandler) movieAdd(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		movie := new(app.Movie)
		movie.User = new(app.User)

		err := json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			mh.Respond(w, http.StatusBadRequest, err.Error())
			return
		}

		if len(movie.Title) == 0 {
			mh.Respond(w, http.StatusBadRequest, " Movie title is empty!")
			return
		}

		if len(movie.Description) == 0 {
			mh.Respond(w, http.StatusBadRequest, " Movie description is empty!")
			return
		}

		ctx := context.WithValue(r.Context(), "movie", movie)

		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

// auth middleware handles the authentication process of each endpoint that needs authenticated permission.
// Responds with 401 Unauthorized header if there's no cookie set or no valid record set on the Sessions storage.
func (mh MWHandler) auth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, err := r.Cookie("mrcookie")
		if err != nil {
			mh.Respond(w, http.StatusUnauthorized, "You must log in to complete this action!")
			return
		}

		next.ServeHTTP(w, r)
	})
}
