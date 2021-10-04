package entity

type User struct {
	UserID int    `db:"id"`
	Name   string `db:"name"`
}
