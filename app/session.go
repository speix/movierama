package app

import (
	"net/http"
	"time"
)

/*
	SessionService abstraction provides persistence to the app and across requests.
	Sessions are implemented using hashmaps with uuid as keys and Session objects for data.
*/

type Session struct {
	UserID   int       `json:"user_id"`
	UserName string    `json:"user_name"`
	LoggedIn bool      `json:"logged_in"`
	Created  time.Time `json:"-"`
}

type Sessions map[string]Session

type SessionService interface {
	CreateSession(u *User) string
	RemoveSession(key string)
	GetSession(key string) Session
	SetCookie(key string, days int) *http.Cookie
}
