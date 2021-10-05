package service

import (
	"todo-togo/entity"
	"todo-togo/model"
	"todo-togo/repository"
)

type UserService struct {
	repo repository.IUserRepo
}

func (u UserService) GetAllUser() ([]*model.UserReqResponse, error) {
	res, err := u.repo.SelectAllUser()
	if err != nil {
		return nil, err
	}

	var users []*model.UserReqResponse
	for _, r := range res {
		users = append(users, &model.UserReqResponse{
			UserID:   r.UserID,
			Name:     r.Name,
			Salt:     r.Salt,
			Password: r.Password,
		})
	}

	return users, nil
}

func (u UserService) GetUser(usr model.UserReqResponse) (*model.UserReqResponse, error) {
	user := entity.User{
		UserID:   usr.UserID,
	}
	r, err := u.repo.SelectUser(user)
	if err != nil {
		return nil, err
	}

	res := model.UserReqResponse{
			UserID:   r.UserID,
			Name:     r.Name,
			Salt:     r.Salt,
			Password: r.Password,
	}

	return &res, nil
}

func (u UserService) UpdateUser(usr model.UserReqResponse) (*model.UserReqResponse, error) {
	user := entity.User{
		UserID:   usr.UserID,
		Password: usr.Password,
		Name: usr.Name,
		Salt: usr.Salt,
	}

	_, err := u.repo.SelectUser(user)
	if err != nil {
		return nil, err
	}

	res, err := u.repo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	updatedUser := model.UserReqResponse{
		UserID:   res.UserID,
		Name:     res.Name,
		Salt:     res.Salt,
		Password: res.Password,
	}

	return &updatedUser, nil
}

func (u UserService) DeleteUser(usr model.UserReqResponse) error {
	id := entity.User{UserID: usr.UserID}
	err := u.repo.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}

type IUserService interface {
	GetAllUser() ([]*model.UserReqResponse, error)
	GetUser(u model.UserReqResponse) (*model.UserReqResponse, error)
	UpdateUser(u model.UserReqResponse) (*model.UserReqResponse, error)
	DeleteUser(u model.UserReqResponse) error
}

func NewUserService(repo *repository.IUserRepo) IUserService  {
	return UserService{repo: *repo}
}