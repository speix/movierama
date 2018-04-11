package data

import (
	"errors"
	"time"

	"github.com/speix/movierama/app"
)

/*
	VoteService implementation of Vote/Retract endpoints.
*/

func (s *Storage) Vote(movieID, userID, positive int) (*app.Movie, error) {

	check := new(struct{ Result int })

	// Check if current user is the owner of the movie
	ownerSql := `select exists(
					select 1 from movie m
					where m.movie_id = $1 and m.user_id = $2
				) as result`
	err := s.DB.Get(check, ownerSql, movieID, userID)
	if err != nil {
		return nil, err
	}

	if check.Result == 1 {
		return nil, errors.New(" You cannot vote for your own movies.")
	}

	// Check if current user already voted for this movie matching positive attribute
	votedSql := `select exists(
					select 1 from vote v
				  	where v.movie_id = $1 and v.user_id = $2 and v.positive = $3
				) as result`
	err = s.DB.Get(check, votedSql, movieID, userID, positive)
	if err != nil {
		return nil, err
	}

	if check.Result == 1 {
		return nil, errors.New(" You have already voted for this movie.")
	}

	// Insert new vote if movie_id/user_id do not exist, update positive value otherwise
	sql := `insert or replace into vote (vote_id, movie_id, positive, created, user_id)
			  values ((select vote_id from vote where movie_id = $1 and user_id = $2), $3, $4, $5, $6)`
	stmt, err := s.DB.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(movieID, userID, movieID, positive, time.Now(), userID)
	if err != nil {
		return nil, err
	}

	movie, err := s.Get(movieID, userID)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (s *Storage) Retract(movieID, userID int) (*app.Movie, error) {

	sql := "delete from vote where movie_id = $1 and user_id = $2"
	_, err := s.DB.Exec(sql, movieID, userID)
	if err != nil {
		return nil, err
	}

	movie, err := s.Get(movieID, userID)
	if err != nil {
		return nil, err
	}

	return movie, nil
}
