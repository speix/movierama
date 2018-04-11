package main

import (
	"log"

	"github.com/speix/movierama/app"

	"github.com/speix/movierama/app/config"
	"github.com/speix/movierama/app/data"
	"github.com/speix/movierama/app/http"
)

// Main function initializes a Database object and a Sessions hashmap
// and wraps them inside a Storage object starting the http server.
func main() {

	database := config.NewDB("./movies.db")
	sessions := make(app.Sessions)

	defer database.DB.Close()

	//config.MigrateDatabase(database)

	server := http.AppServer(
		&data.Storage{
			Database: database,
			Sessions: &sessions,
		},
	)

	log.Fatal(server.ListenAndServe())
}
