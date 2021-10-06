package model

type JsonInternalErrorReturn struct {
	Status bool   `json:"status" example:"false"`
	Error  string `json:"error" example:"Error to get all todos"`
}

type JsonBadRequestErrorReturn struct {
	Status bool   `json:"status" example:"false"`
	Error  string `json:"error" example:"Error to parse updated todos"`
}

type JsonSuccessTodosReturn struct {
	Status bool                 `json:"status" example:"true"`
	Todos  []CreateTodoResponse `json:"todos"`
}

type JsonDeleteTodosReturn struct {
	Status bool   `json:"status" example:"true"`
	Todos  string `json:"todos" example:"nil"`
}

type JsonSuccessLoginReturn struct {
	Status bool   `json:"status" example:"true"`
	Token  string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzM3MzEwMzQsInVzZXJfaWQiOjQsInVzZXJfbmFtZSI6Ikpva28gQW51YXIzIn0.16vwI9kQQj7ER4yEoZDUezm-il3YZdBYr_XAoasWoj0"`
}

type JsonSuccessUserReturn struct {
	Status bool              `json:"status" example:"true"`
	User   []UserReqResponse `json:"user"`
}

type JsonDeleteUserReturn struct {
	Status bool   `json:"status" example:"true"`
	User   string `json:"user"" example:"nil"`
}
