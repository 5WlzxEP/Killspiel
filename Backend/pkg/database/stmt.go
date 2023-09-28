package database

import "database/sql"

var (
	CreateGame        *sql.Stmt
	UserExist         *sql.Stmt
	CreateUser        *sql.Stmt
	UpdateUserGuesses *sql.Stmt
	CreateVote        *sql.Stmt
	UpdateUser        *sql.Stmt
	// SetGameCorrect correct, gameId
	SetGameCorrect *sql.Stmt

	// User
	GetUser      *sql.Stmt
	GetUserGames *sql.Stmt
	DeleteUser   *sql.Stmt
	SearchUser   *sql.Stmt

	// Leaderboard
	LeaderboardAsc  *sql.Stmt
	LeaderboardDesc *sql.Stmt

	// Game
	GetGame           *sql.Stmt
	SetGameVerteilung *sql.Stmt
	GetPlayersByVote  *sql.Stmt
)

func prepareStmts() (err error) {
	CreateGame, err = DB.Prepare("INSERT INTO Game (userCount, info) VALUES (?, ?)")
	if err != nil {
		return
	}

	SetGameVerteilung, err = DB.Prepare("UPDATE Game SET verteilung = ? WHERE id = ?")
	if err != nil {
		return err
	}

	UserExist, err = DB.Prepare("SELECT EXISTS(SELECT id FROM Users WHERE id = ?)")
	if err != nil {
		return
	}

	CreateUser, err = DB.Prepare("INSERT INTO Users (id, name) VALUES (?, ?)")
	if err != nil {
		return
	}

	UpdateUserGuesses, err = DB.Prepare("UPDATE Users SET guesses = guesses + 1 WHERE id = ?")
	if err != nil {
		return
	}

	CreateVote, err = DB.Prepare("INSERT INTO Votes VALUES (?, ?, ?)")
	if err != nil {
		return
	}

	LeaderboardAsc, err = DB.Prepare(`SELECT id, name, guesses, points, latest FROM Users ORDER BY 
                                                    CASE ? 
                                                        WHEN 1 THEN points 
                                                        WHEN 2 THEN name 
                                                        WHEN 3 THEN guesses
													END 
													LIMIT ? OFFSET ?`)
	if err != nil {
		return err
	}

	LeaderboardDesc, err = DB.Prepare(`SELECT id, name, guesses, points, latest FROM Users ORDER BY 
                                                    CASE ? 
                                                        WHEN 1 THEN points 
                                                        WHEN 2 THEN name 
                                                        WHEN 3 THEN guesses
													END 
													DESC LIMIT ? OFFSET ?`)
	if err != nil {
		return err
	}

	GetUser, err = DB.Prepare("SELECT name, points, guesses, latest FROM Users WHERE id = ?")
	if err != nil {
		return
	}

	// DO NOT CHANGE LIMIT, or when changed also change historyPool size in User/user.go
	GetUserGames, err = DB.Prepare("SELECT id, correct, vote, time, game.precision FROM Game JOIN (SELECT game, vote FROM Votes WHERE player = ? ORDER BY game desc LIMIT 25) V on Game.id = V.game")
	if err != nil {
		return
	}

	DeleteUser, err = DB.Prepare("DELETE FROM Users WHERE id = ?")
	if err != nil {
		return
	}

	SearchUser, err = DB.Prepare("SELECT id, name, points FROM Users WHERE name like ?")
	if err != nil {
		return err
	}

	UpdateUser, err = DB.Prepare("UPDATE Users SET latest = ((latest << 1) + ?) % 256, points = points + ? WHERE id = ?;")
	if err != nil {
		return err
	}

	SetGameCorrect, err = DB.Prepare("UPDATE Game SET correct = ?, precision = ?, correctCount = ? WHERE id = ?;")
	if err != nil {
		return err
	}

	GetGame, err = DB.Prepare("SELECT correct, time, correctCount, userCount, game.precision, verteilung FROM Game WHERE id = ?")
	if err != nil {
		return err
	}

	GetPlayersByVote, err = DB.Prepare("SELECT id, name FROM Users JOIN (SELECT player FROM Votes WHERE game = ? and vote = ?) as Vp on Users.id = Vp.player")
	if err != nil {
		return err
	}

	return
}
