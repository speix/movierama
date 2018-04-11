package data

import (
	"github.com/speix/movierama/app"
	"github.com/speix/movierama/app/config"
)

type Storage struct {
	*config.Database
	*app.Sessions
}
