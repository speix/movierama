package helpers

import (
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/speix/movierama/app"
)

// AggregateMovie summarizes texts and groups data
// before sending the response back to the client
func AggregateMovie(movie *app.Movie, user *app.User) *app.Movie {

	likes, _ := strconv.Atoi(movie.Likes)
	hates, _ := strconv.Atoi(movie.Hates)

	movie.CanRetractVote = movie.Voted.Valid
	if movie.CanRetractVote {
		if movie.Voted.Bool {
			movie.RetractAction = "Unlike"
		} else {
			movie.RetractAction = "Unhate"
		}
	}

	movie.Votes = likes + hates
	movie.Likes = votesToText(movie.Likes, "like")
	movie.Hates = votesToText(movie.Hates, "hate")
	movie.Created = daysToText(movie.Created)
	movie.User = user

	return movie
}

// DaysToText converts an actual date string
// to a "days elapsed from today" string
func daysToText(date string) string {

	format := "2006-01-02 15:04:05"
	dbDate, _ := time.Parse(format, date)
	diff := time.Now().Sub(dbDate)
	days := math.Ceil(diff.Hours() / 24)

	switch days {
	case 0:
		return "today"
	case 1:
		return strconv.FormatFloat(days, 'f', 0, 32) + " day ago"
	default:
		return strconv.FormatFloat(days, 'f', 0, 32) + " days ago"
	}

}

func votesToText(sum, vote string) string {

	if sum == "0" {
		return strings.Title(vote)
	}

	return sum
}
