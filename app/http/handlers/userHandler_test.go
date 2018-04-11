package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/speix/movierama/app/data"

	"github.com/speix/movierama/app"
	"github.com/speix/movierama/app/config"
)

func TestUserHandler_RegisterUserExists(t *testing.T) {

	database := config.NewDB("../../../movies.db")
	sessions := make(app.Sessions)

	storage := &data.Storage{
		Database: database,
		Sessions: &sessions,
	}

	expected := 409
	user := new(app.User)
	user.Email = "spei@supergramm.com"
	h := &UserHandler{UserService: storage, SessionService: storage}

	handler := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user", user)
		h.Register(w, r.WithContext(ctx))
	}

	req := httptest.NewRequest("POST", "/", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	if expected != resp.StatusCode {
		t.Errorf("Expected status code %v got %v", expected, resp.StatusCode)
	}
}
