package data

import (
	"database/sql"
	"errors"
	"time"

	"github.com/mattn/go-sqlite3"

	"github.com/speix/movierama/app"
	"github.com/speix/movierama/app/helpers"
)

/*
	UserService implementation of Login/Register endpoints.
*/

func (s *Storage) Login(email, password string) (*app.User, error) {

	user := new(app.User)
	pwd := helpers.HashPassword(email, password)

	err := s.DB.Get(user, "select * from user where email = $1 and password = $2", email, pwd)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(" Wrong login credentials")
		} else {
			return nil, err
		}
	}

	return user, nil
}

func (s *Storage) Register(u *app.User) (*app.User, error) {

	sql := "insert into user (email, password, name, created) values($1, $2, $3, $4)"
	stmt, err := s.DB.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.Email, helpers.HashPassword(u.Email, u.Password), u.Name, time.Now().String())
	if err != nil {
		if err.(sqlite3.Error).Code.Error() == "constraint failed" {
			return nil, errors.New(" User exists!")
		}
		return nil, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	u.UserID = int(userID)

	return u, nil
}
