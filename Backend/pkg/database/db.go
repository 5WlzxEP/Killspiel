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
	return
}

func CreateTables() error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("CREATE TABLE IF NOT EXISTS USERS (id int )")
	if err != nil {
		return err
	}

	return nil
}
