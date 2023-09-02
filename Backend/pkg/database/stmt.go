package database

import "database/sql"

var (
	createGame        *sql.Stmt
	userExist         *sql.Stmt
	createUser        *sql.Stmt
	updateUserGuesses *sql.Stmt
	userWin           *sql.Stmt
	createVote        *sql.Stmt
	getWinners        *sql.Stmt

	// User
	GetUser      *sql.Stmt
	GetUserGames *sql.Stmt
	DeleteUser   *sql.Stmt

	// Leaderboard
	LeaderboardAsc  *sql.Stmt
	LeaderboardDesc *sql.Stmt
)

func prepareStmts() (err error) {
	createGame, err = DB.Prepare("INSERT INTO Game DEFAULT VALUES")
	if err != nil {
		return
	}

	userExist, err = DB.Prepare("SELECT EXISTS(SELECT id FROM Users WHERE id = ?)")
	if err != nil {
		return
	}

	createUser, err = DB.Prepare("INSERT INTO Users VALUES (?, ?, default, default, default)")
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

	GetUser, err = DB.Prepare("SELECT id, correct, vote FROM Game JOIN (SELECT game, vote FROM Votes WHERE player = ? LIMIT 50 ORDER BY game desc) V on Game.id = V.game")
	if err != nil {
		return
	}

	DeleteUser, err = DB.Prepare("DELETE FROM Users WHERE id = ?")
	if err != nil {
		return
	}

	return
}
