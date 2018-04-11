package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/speix/movierama/app"
	"github.com/speix/movierama/app/config"
	"github.com/speix/movierama/app/data"
)

func TestVoteHandler_RetractNoMovieID(t *testing.T) {

	expected := 409
	database := config.NewDB("../../../movies.db")
	sessions := make(app.Sessions)

	storage := &data.Storage{
		Database: database,
		Sessions: &sessions,
	}

	h := &VoteHandler{VoteService: storage, SessionService: storage}

	handler := func(w http.ResponseWriter, r *http.Request) {
		h.Retract(w, r)
	}

	req := httptest.NewRequest("DELETE", "/", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	if expected != resp.StatusCode {
		t.Errorf("Expected status code %v got %v", expected, resp.StatusCode)
	}
}
