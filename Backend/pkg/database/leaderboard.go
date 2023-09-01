package database

import "database/sql"

func GetUserCount() (count int, err error) {
	err = DB.QueryRow("SELECT COUNT(*) FROM Users").Scan(&count)
	if err != nil {
		return
	}
	return
}

var (
	LeaderboardAsc  *sql.Stmt
	LeaderboardDesc *sql.Stmt
)
