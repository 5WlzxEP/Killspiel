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
 		latest tinyint unsigned default 0
 	)`)
	if err != nil {
		return err
	}

	// Games
	_, err = tx.Exec(`CREATE TABLE IF NOT EXISTS Game (
    	id integer primary key AUTOINCREMENT,
    	time timestamp default CURRENT_TIMESTAMP
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
    	foreign key (game) REFERENCES Game(id),
    	foreign key (player) REFERENCES Users(id) 
 	)`)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

var (
	createGame        *sql.Stmt
	userExist         *sql.Stmt
	createUser        *sql.Stmt
	updateUserGuesses *sql.Stmt
	userWin           *sql.Stmt
	createVote        *sql.Stmt
	getWinners        *sql.Stmt
)

func prepareStmts() (err error) {
	createGame, err = DB.Prepare("INSERT INTO Game VALUES ()")
	if err != nil {
		return
	}

	userExist, err = DB.Prepare("SELECT EXISTS(SELECT id FROM Users WHERE id = ?)")
	if err != nil {
		return
	}

	createUser, err = DB.Prepare("INSERT INTO Users VALUES (?, ?)")
	if err != nil {
		return
	}

	updateUserGuesses, err = DB.Prepare("UPDATE Users SET guesses = guesses + 1 WHERE id = ?")
	if err != nil {
		return
	}

	userWin, err = DB.Prepare("UPDATE Users SET points = points + 1 WHERE id = ?")
	if err != nil {
		return
	}

	createVote, err = DB.Prepare("INSERT INTO Votes VALUES (?, ?, ?)")
	if err != nil {
		return
	}

	getWinners, err = DB.Prepare("SELECT * from Users WHERE id in (SELECT player FROM Votes WHERE game = ? AND abs(vote - ?) <= 1e-6)")
	if err != nil {
		return
	}

	return
}
