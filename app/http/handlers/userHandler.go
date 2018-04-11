package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/speix/movierama/app"
	"github.com/speix/movierama/app/errors"
)

// UserHandler
// wraps a SessionService for session handling for users
// wraps a UserService for database interaction with users
// wraps and ApiResponseService for client responses
type UserHandler struct {
	app.SessionService
	app.UserService
	errors.ApiResponseService
}

// Login POST request logs a user into the system.
func (h UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	userCtx := r.Context().Value("user").(*app.User)

	user, err := h.UserService.Login(userCtx.Email, userCtx.Password)
	if err != nil {
		h.Respond(w, http.StatusConflict, err.Error())
		return
	}

	key := h.SessionService.CreateSession(user)
	cookie := h.SessionService.SetCookie(key, 30)

	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// Register POST request registers a user to the system.
func (h UserHandler) Register(w http.ResponseWriter, r *http.Request) {

	userCtx := r.Context().Value("user").(*app.User)

	user, err := h.UserService.Register(userCtx)
	if err != nil {
		h.Respond(w, http.StatusConflict, err.Error())
		return
	}

	key := h.SessionService.CreateSession(user)
	cookie := h.SessionService.SetCookie(key, 30)

	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h UserHandler) Logout(w http.ResponseWriter, r *http.Request) {

	ck, err := r.Cookie("mrcookie")
	if err != nil {
		h.Respond(w, http.StatusConflict, err.Error())
		return
	}

	h.SessionService.RemoveSession(ck.Value)
	cookie := h.SessionService.SetCookie("", 0)

	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h UserHandler) Status(w http.ResponseWriter, r *http.Request) {

	ck, err := r.Cookie("mrcookie")
	if err != nil {
		h.Respond(w, http.StatusConflict, err.Error())
		return
	}
	session := h.SessionService.GetSession(ck.Value)

	data, err := json.Marshal(app.Session{UserID: session.UserID, UserName: session.UserName, LoggedIn: true})
	if err != nil {
		h.Respond(w, http.StatusConflict, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
