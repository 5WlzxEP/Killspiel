package database

import (
	"database/sql"
	_ "github.com/glebarez/go-sqlite"
	"path"
)

var DB *sql.DB

func Init(configDir string) (err error) {
	DB, err = sql.Open("sqlite", path.Join(configDir, "db.sqlite"))
	if err != nil {
		return
	}
	err = CreateTables()
	if err != nil {
		return
	}
	err = prepareStmts()
	return
}

func CreateTables() error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Users
	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS Users (
    	id int primary key unique,
 		name varchar(255) unique,
 		guesses int default 1,
 		points int default 0,
 		latest smallint default 0
 	)`)
	if err != nil {
		return err
	}

	_, err = tx.Exec("CREATE UNIQUE INDEX IF NOT EXISTS Usernames ON Users(name)")
	if err != nil {
		return err
	}

	// Games
	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS Game (
    	id integer primary key AUTOINCREMENT,
    	correct float default null,
    	"precision" float default 0.01,
    	userCount integer default 0,
    	correctCount integer default 0,
    	time timestamp default CURRENT_TIMESTAMP,
    	info text,
    	verteilung text
 	)`)
	if err != nil {
		return err
	}

	// Votes
	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS Votes (
    	game int,
    	player int,
    	vote float,
    	PRIMARY KEY (game, player),
    	foreign key (game) REFERENCES Game(id) ON DELETE CASCADE,
    	foreign key (player) REFERENCES Users(id) ON DELETE CASCADE 
 	)`)
	if err != nil {
		return err
	}

	_, err = tx.Exec("CREATE INDEX IF NOT EXISTS Votesplayers ON Votes(player)")
	if err != nil {
		return err
	}

	_, err = tx.Exec("CREATE INDEX IF NOT EXISTS Votesgame ON Votes(game)")
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
