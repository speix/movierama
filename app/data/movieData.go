package data

import (
	"fmt"
	"time"

	"github.com/speix/movierama/app"
	"github.com/speix/movierama/app/helpers"
)

/*
	MovieService implementation of Add/Sort/Get endpoints.
*/

func (s *Storage) Add(m *app.Movie) error {

	sql := "insert into movie (user_id, title, description, created) values($1, $2, $3, $4)"
	stmt, err := s.DB.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(m.User.UserID, m.Title, m.Description, time.Now())
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Sort(userID, sessionUserID, by string) ([]*app.Movie, error) {

	movies := make([]*app.Movie, 0)
	sql := ""

	if len(userID) == 0 {

		rawSql := `select
					  m.movie_id, m.user_id, m.title, m.description,
					  datetime(m.created, 'localtime') as created,
					  (select count(vote_id) from vote where m.movie_id = vote.movie_id and positive = 1) as likes,
					  (select count(vote_id) from vote where m.movie_id = vote.movie_id and positive = 0) as hates,
					  case m.user_id when %s then 1 else 0 end as is_owner,
  					  (select positive from vote where m.movie_id = vote.movie_id and vote.user_id = %s) as voted
					from movie m
					order by %s desc`
		sql = fmt.Sprintf(rawSql, sessionUserID, sessionUserID, by)

	} else {

		rawSql := `select
					  m.movie_id, m.user_id, m.title, m.description,
					  datetime(m.created, 'localtime') as created,
					  (select count(vote_id) from vote where m.movie_id = vote.movie_id and positive = 1) as likes,
					  (select count(vote_id) from vote where m.movie_id = vote.movie_id and positive = 0) as hates,
					  case m.user_id when %s then 1 else 0 end as is_owner,
  					  (select positive from vote where m.movie_id = vote.movie_id and vote.user_id = %s) as voted
					from movie m
					  where m.user_id = %s
					order by %s desc`
		sql = fmt.Sprintf(rawSql, sessionUserID, sessionUserID, userID, by)

	}

	err := s.DB.Select(&movies, sql)
	if err != nil {
		return nil, err
	}

	for m := range movies {

		user := new(app.User)
		err := s.DB.Get(user, "select name from user where user_id = $1", movies[m].UserID)
		if err != nil {
			return nil, err
		}

		movies[m] = helpers.AggregateMovie(movies[m], user)
	}

	return movies, nil
}

func (s *Storage) Get(movieID, userID int) (*app.Movie, error) {

	movie := new(app.Movie)

	sql := `select
			  m.movie_id, m.user_id, m.title, m.description,
			  datetime(m.created, 'localtime') as created,
			  (select count(vote_id) from vote where m.movie_id = vote.movie_id and positive = 1) as likes,
			  (select count(vote_id) from vote where m.movie_id = vote.movie_id and positive = 0) as hates,
			  case m.user_id when $1 then 1 else 0 end as is_owner,
  			  (select positive from vote where m.movie_id = vote.movie_id and vote.user_id = $2) as voted
			from movie m
			where m.movie_id = $3`
	err := s.DB.Get(movie, sql, userID, userID, movieID)
	if err != nil {
		return nil, err
	}

	user := new(app.User)
	err = s.DB.Get(user, "select name from user where user_id = $1", movie.UserID)
	if err != nil {
		return nil, err
	}

	movie = helpers.AggregateMovie(movie, user)

	return movie, nil
}
