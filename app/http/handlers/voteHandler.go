package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/speix/movierama/app"
	"github.com/speix/movierama/app/errors"
)

// VoteHandler
// wraps a SessionService for session handling for users
// wraps a VoteService for database interaction with votes
// wraps and ApiResponseService for client responses
type VoteHandler struct {
	app.SessionService
	app.VoteService
	errors.ApiResponseService
}

// Vote POST request saves a vote into the database given a user.
func (h VoteHandler) Vote(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	movieID, err := strconv.Atoi(params["movie_id"])
	if err != nil {
		h.Respond(w, http.StatusConflict, err.Error())
		return
	}

	positive, err := strconv.Atoi(params["positive"])
	if err != nil {
		h.Respond(w, http.StatusConflict, err.Error())
		return
	}

	ck, err := r.Cookie("mrcookie")
	if err != nil {
		h.Respond(w, http.StatusConflict, err.Error())
		return
	}

	session := h.SessionService.GetSession(ck.Value)

	movie, err := h.VoteService.Vote(movieID, session.UserID, positive)
	if err != nil {
		h.Respond(w, http.StatusConflict, err.Error())
		return
	}

	data, err := json.Marshal(&movie)
	if err != nil {
		h.Respond(w, http.StatusConflict, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

// Retract DELETE request removes a vote given a movie.
func (h VoteHandler) Retract(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	movieID, err := strconv.Atoi(params["movie_id"])
	if err != nil {
		h.Respond(w, http.StatusConflict, err.Error())
		return
	}

	ck, err := r.Cookie("mrcookie")
	if err != nil {
		h.Respond(w, http.StatusConflict, err.Error())
		return
	}

	session := h.SessionService.GetSession(ck.Value)

	movie, err := h.VoteService.Retract(movieID, session.UserID)
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
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
