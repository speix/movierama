package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/speix/movierama/app"
)

func TestMovieHandler_AddSession(t *testing.T) {

	expected := 409
	movie := new(app.Movie)
	h := &MovieHandler{}

	handler := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "movie", movie)
		h.Add(w, r.WithContext(ctx))
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
