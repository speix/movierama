package app

/*
	VoteService abstraction is the contract for voting	manipulation of each movie by each user.
	Votes can be either positive or negative, represented by Positive attribute either being 0 or 1.
	Retract removes a given vote while Vote adds one.
*/

type Vote struct {
	VoteID   int    `json:"vote_id" db:"vote_id"`
	MovieID  int    `json:"movie_id" db:"movie_id"`
	UserID   int    `json:"user_id" db:"user_id"`
	Positive int    `db:"positive"`
	Created  string `db:"created"`
}

type VoteService interface {
	Retract(movieID, userID int) (*Movie, error)
	Vote(movieID, userID, positive int) (*Movie, error)
}
