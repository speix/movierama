package config

import (
	"os"
	"time"

	"github.com/speix/movierama/app/helpers"
)

var schema = `
CREATE TABLE user (
    user_id integer PRIMARY KEY,
    email text,
    password text,
    name text,
	created integer,
	CONSTRAINT unique_email UNIQUE (email)
);

CREATE TABLE movie (
    movie_id integer PRIMARY KEY,
    user_id text NULL,
    title text,
	description text,
	created integer,
	FOREIGN KEY(user_id) REFERENCES user(user_id)
);

CREATE TABLE vote (
    vote_id integer PRIMARY KEY,
    movie_id integer,
	user_id integer,
    positive integer,
	created integer,
	FOREIGN KEY(movie_id) REFERENCES movie(movie_id),
	FOREIGN KEY(user_id) REFERENCES user(user_id)
)`

func MigrateDatabase(database *Database) {

	os.Create("./app/data/movies.db")

	db := database.DB

	db.MustExec(schema)

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO user (email, password, name, created) VALUES ($1, $2, $3, $4)",
		"john@doe.com",
		helpers.HashPassword("john@doe.com", "123123"),
		"John Doe",
		time.Now())
	tx.MustExec("INSERT INTO user (email, password, name, created) VALUES ($1, $2, $3, $4)",
		"jack@white.com",
		helpers.HashPassword("jack@white.com", "123123"),
		"Jack White",
		time.Now())
	tx.MustExec("INSERT INTO user (email, password, name, created) VALUES ($1, $2, $3, $4)",
		"jenn@jameson.com",
		helpers.HashPassword("jenn@jameson.com", "123123"),
		"Jenn Jameson",
		time.Now())
	tx.MustExec("INSERT INTO user (email, password, name, created) VALUES ($1, $2, $3, $4)",
		"spei@supergramm.com",
		helpers.HashPassword("spei@supergramm.com", "123123"),
		"Stathis Peioglou",
		time.Now())

	tx.MustExec("INSERT INTO movie(user_id, title, description, created) VALUES ($1, $2, $3, $4)",
		4,
		"In the Name of the Father",
		"Gerry Conlon is shown in Belfast stripping lead from roofs of houses when security forces home in on the district with armoured cars, and a riot breaks out. Gerry's father, Giuseppe Conlon, later saves him from IRA punishment, and he is sent off to London to stay with his aunt, Anne Maguire, for his own good. Instead, he finds a squat, to explore, as he puts it, \"free love and dope\". In October 1974, Gerry happens to gain entry to a prostitute's flat and he steals the £700 he finds there and chats briefly with a man sitting in a park. On that evening in Guildford, southwest of London, there is an explosion at a pub that kills four off-duty soldiers and a civilian, and wounds sixty-five others.",
		time.Now())
	tx.MustExec("INSERT INTO movie(user_id, title, description, created) VALUES ($1, $2, $3, $4)",
		1,
		"The Shawshank Redemption",
		"In 1947 Portland, Maine, banker Andy Dufresne is convicted of murdering his wife and her lover, and sentenced to two consecutive life sentences at the Shawshank State Penitentiary. He is befriended by contraband smuggler, Ellis \"Red\" Redding, an inmate serving a life sentence. Red procures a rock hammer, and later a large poster of Rita Hayworth for Andy. Working in the prison laundry, Andy is regularly assaulted and raped by \"the Sisters\" and their leader, Bogs.",
		time.Now())
	tx.MustExec("INSERT INTO movie(user_id, title, description, created) VALUES ($1, $2, $3, $4)",
		1,
		"Twelve Monkeys",
		"A deadly virus released in 1996 wipes out almost all of humanity, forcing survivors to live underground. A group known as the Army of the Twelve Monkeys is believed to have released the virus. In 2035, James Cole is a prisoner living in a subterranean compound beneath the ruins of Philadelphia. Cole is selected to be trained and sent back in time to find the original virus in order to help scientists develop a cure. Meanwhile, Cole is troubled by recurring dreams involving a foot chase and shooting at an airport.",
		time.Now())
	tx.MustExec("INSERT INTO movie(user_id, title, description, created) VALUES ($1, $2, $3, $4)",
		2,
		"Inception",
		"Dominick \"Dom\" Cobb and Arthur are \"extractors\", who perform corporate espionage using an experimental military technology to infiltrate the subconscious of their targets and extract valuable information through a shared dream world. Their latest target, Japanese businessman Saito, reveals that he arranged their mission himself to test Cobb for a seemingly impossible job: planting an idea in a person's subconscious, or \"inception\". To break up the energy conglomerate of ailing competitor Maurice Fischer, Saito wants Cobb to convince Fischer's son and heir, Robert, to dissolve his father's company. In return, Saito promises to use his influence to clear Cobb of a murder charge, allowing Cobb to return home to his children.",
		time.Now())
	tx.MustExec("INSERT INTO movie(user_id, title, description, created) VALUES ($1, $2, $3, $4)",
		3,
		"Return of the Jedi",
		"Luke Skywalker initiates a plan to rescue Han Solo from the crime lord Jabba the Hutt with the help of Princess Leia, Lando Calrissian, Chewbacca, C-3PO, and R2-D2. Leia infiltrates Jabba's palace on Tatooine, disguised as a bounty hunter, with Chewbacca as her prisoner. Leia releases Han from his carbonite prison, but she is captured and enslaved. Luke arrives soon afterward, but is also captured after a tense standoff.",
		time.Now())
	tx.Commit()

}
