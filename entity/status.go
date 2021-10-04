package entity

type Status struct {
	StatusID  int    `db:"id"`
	StatusTxt string `db:"desc"`
}
