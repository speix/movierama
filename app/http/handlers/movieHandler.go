package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/speix/movierama/app"
	"github.com/speix/movierama/app/errors"
)

// MovieHandler
// wraps a SessionService for session handling for users
// wraps a MovieService for database interaction with movies
// wraps and ApiResponseService for client responses
type MovieHandler struct {
	app.SessionService
	app.MovieService
	errors.ApiResponseService
}

// Add POST request saves a movie into the database given the connected user.
func (h MovieHandler) Add(w http.ResponseWriter, r *http.Request) {

	movie := r.Context().Value("movie").(*app.Movie)

	ck, err := r.Cookie("mrcookie")
	if err != nil {
		h.Respond(w, http.StatusConflict, err.Error())
		return
	}

	session := h.SessionService.GetSession(ck.Value)
	movie.User.UserID = session.UserID
	movie.User.Name = session.UserName

	err = h.MovieService.Add(movie)
	if err != nil {
		h.Respond(w, http.StatusConflict, err.Error())
		return
	}

	data, err := json.Marshal(movie)
	if err != nil {
		h.Respond(w, http.StatusConflict, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

// Sort GET request sorts the movies given a sorting rule and an optional user.
func (h MovieHandler) Sort(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	sessionUserID := "null"

	ck, err := r.Cookie("mrcookie")
	if err == nil {
		session := h.SessionService.GetSession(ck.Value)
		sessionUserID = strconv.Itoa(session.UserID)
	}

	movies, _ := h.MovieService.Sort(params["user"], sessionUserID, params["sort_by"])

	if len(movies) == 0 {
		h.Respond(w, http.StatusNotFound, "You haven't submitted any movies.")
		return
	}

	data, err := json.Marshal(movies)
	if err != nil {
		h.Respond(w, http.StatusConflict, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
