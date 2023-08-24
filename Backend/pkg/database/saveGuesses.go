package database

import (
	"database/sql"
)

// TODO moves types into own package
type Guess struct {
	Name  string
	Guess float64
}

func SaveGuesses(m *map[int]Guess) (err error) {
	tx, err := DB.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	result, err := tx.Stmt(createGame).Exec("")
	if err != nil {
		return
	}
	var gameId int64
	if gameId, err = result.LastInsertId(); err != nil {
		var rows *sql.Rows
		rows, err = tx.Query("SELECT last_insert_rowid()")
		if err != nil {
			return
		}
		for rows.Next() {
			err = rows.Scan(&gameId)
			if err != nil {
				return
			}
		}
	}

	for id, guess := range *m {
		rows, _ := tx.Stmt(userExist).Query(id)
		rows.Next()
		var exists bool
		_ = rows.Scan(&exists)
		rows.Close()

		if !exists {
			_, err = tx.Stmt(createUser).Exec(id, guess.Name)
			if err != nil {
				return
			}
		} else {
			_, err = tx.Stmt(updateUserGuesses).Exec(id)
			if err != nil {
				return
			}
		}

		_, err = tx.Stmt(createVote).Exec(gameId, id, guess.Guess)
		if err != nil {
			return
		}

	}

	err = tx.Commit()
	if err != nil {
		return
	}

	return
}
