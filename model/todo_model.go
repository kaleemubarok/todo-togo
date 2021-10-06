package model

type CreateTodoRequest struct {
	ID          int    `json:"id"`
	Title       string `json:"title" example:"Clean up source code"`
	Description string `json:"description" example:"Clean all leftover comment and unnecessary code"`
	DueDate     string `json:"due_date" example:"2021-12-30"`
	PIC         int    `json:"pic" example:"2"`
	Status      int    `json:"status" example:"1"`
}

type CreateTodoResponse struct {
	ID          int    `json:"id" example:"1"`
	Title       string `json:"title" example:"Clean up source code"`
	Description string `json:"description" example:"Clean all leftover comment and unnecessary code"`
	DueDate     string `json:"due_date" example:"2021-12-30"`
	PIC         string `json:"pic" example:"Bambang Prstoyo"`
	Status      string `json:"status" example:"New"`
}
