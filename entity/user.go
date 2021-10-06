package entity

type User struct {
	UserID   int    `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Salt     string `db:"salt"`
	Password string `db:"password"`
}
