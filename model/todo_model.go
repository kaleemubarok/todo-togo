package model

type CreateTodoRequest struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	PIC         int    `json:"pic"`
	Status      int    `json:"status"`
}

type CreateTodoResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	PIC         string `json:"pic"`
	Status      string `json:"status"`
}
