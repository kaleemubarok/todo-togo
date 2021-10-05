package entity

type Todo struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"desc"`
	DueDate     string `db:"due_date"`
	PIC         int    `db:"pic"`
	PICName     string `db:"pic_desc"`
	Status      int    `db:"status"`
	StatusDesc  string `db:"status_desc"`
}
