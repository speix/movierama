package data

import (
	"net/http"
	"time"

	"github.com/twinj/uuid"

	"github.com/speix/movierama/app"
)

/*
	SessionService implementation of CreateSession/RemoveSession/GetSession/SetCookie endpoints.
*/

func (s *Storage) CreateSession(u *app.User) string {
	key := uuid.NewV4().String()
	(*s.Sessions)[key] = app.Session{UserID: u.UserID, UserName: u.Name, LoggedIn: true, Created: time.Now()}
	return key
}

func (s *Storage) RemoveSession(key string) {
	delete(*s.Sessions, key)
}

func (s *Storage) GetSession(key string) app.Session {
	return (*s.Sessions)[key]
}

func (s *Storage) SetCookie(key string, days int) *http.Cookie {
	return &http.Cookie{
		Name:     "mrcookie",
		Value:    key,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(days) * 24 * time.Hour),
	}
}
