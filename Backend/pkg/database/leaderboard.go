package database

func GetUserCount() (count int, err error) {
	err = DB.QueryRow("SELECT COUNT(*) FROM Users").Scan(&count)
	if err != nil {
		return
	}
	return
}
