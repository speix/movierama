package data

import (
	"testing"

	"github.com/speix/movierama/app"
)

func TestStorage_CreateSession(t *testing.T) {

	storage := &Storage{}
	sessions := make(app.Sessions)
	storage.Sessions = &sessions

	user := app.User{UserID: 1, Email: "test@test.com", Name: "Username"}

	key := storage.CreateSession(&user)

	if user.Name != sessions[key].UserName {
		t.Errorf("Expected %v got %v", user.Name, sessions[key].UserName)
	}

}

func TestStorage_GetRemoveSession(t *testing.T) {

	storage := &Storage{}
	sessions := make(app.Sessions)
	storage.Sessions = &sessions

	user := app.User{UserID: 1, Email: "test@test.com", Name: "Username"}
	sessions["key"] = app.Session{UserID: user.UserID, UserName: user.Name}

	storage.RemoveSession("key")

	session := storage.GetSession("key")

	if session.UserID != 0 {
		t.Errorf("Expected %v got %v", 0, session.UserID)
	}

}

func TestStorage_SetCookie(t *testing.T) {

	storage := &Storage{}
	cookie := storage.SetCookie("key", 30)

	if cookie.Value != "key" {
		t.Errorf("Expected %v got %v", "key", cookie.Value)
	}

}
