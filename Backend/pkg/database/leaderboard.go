package database

import "context"

func GetUserCount(ctx context.Context) (count int, err error) {
	err = DB.QueryRowContext(ctx, "SELECT COUNT(*) FROM Users").Scan(&count)
	if err != nil {
		return
	}
	return
}
