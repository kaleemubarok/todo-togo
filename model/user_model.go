package model

type UserReqResponse struct {
	UserID   int    `json:"id"`
	Name     string `json:"name"`
	Salt     string `json:"salt"`
	Password string `json:"password"`
}
