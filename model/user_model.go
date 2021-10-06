package model

type UserReqResponse struct {
	UserID   int    `json:"id" example:""`
	Name     string `json:"name" example:"Ibu Badu"`
	Email    string `json:"email" example:"surat@mail.com"`
	Salt     string `json:"salt" example:""`
	Password string `json:"password" example:"password123"`
}

type LoginInput struct {
	Email    string `json:"email" example:"badu.naik.gajah@gmail.com"`
	Password string `json:"password" example:"rahasiaIbuBADU"`
}