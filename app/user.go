package app

/*
	UserService abstraction provides a user the ability to Login or Register.
*/
type User struct {
	UserID   int    `json:"user_id" db:"user_id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Name     string `json:"name"  db:"name"`
	Created  string `json:"created" db:"created"`
}

type UserService interface {
	Login(email, password string) (*User, error)
	Register(u *User) (*User, error)
}
