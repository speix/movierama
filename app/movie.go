package app

import "database/sql"

/*
	MovieService abstraction provides 3 Endpoints
	for Adding Getting and Sorting movies.
*/

type Movie struct {
	MovieId        int          `json:"movie_id" db:"movie_id"`
	UserID         int          `json:"user_id" db:"user_id"`
	User           *User        `json:"user"`
	Title          string       `json:"title" db:"title"`
	Description    string       `json:"description" db:"description"`
	Likes          string       `json:"likes" db:"likes"`
	Hates          string       `json:"hates" db:"hates"`
	Votes          int          `json:"votes"`
	IsOwner        bool         `json:"is_owner" db:"is_owner"`
	Voted          sql.NullBool `db:"voted"`
	CanRetractVote bool         `json:"can_retract"`
	RetractAction  string       `json:"retract_action"`
	Created        string       `json:"created" db:"created"`
}

type MovieService interface {
	Add(m *Movie) error
	Get(movieID, userID int) (*Movie, error)
	Sort(userID, sessionUserID, by string) ([]*Movie, error)
}
