package repository

import (
	"github.com/jmoiron/sqlx"
	"todo-togo/entity"
)

type IUserRepo interface {
	SelectUser(u entity.User) (*entity.User, error)
	SelectAllUser() ([]*entity.User, error)
	UpdateUser(u entity.User) (*entity.User, error)
	DeleteUser(u entity.User) error
}

func NewUserRepo(dbParam *sqlx.DB) IUserRepo {
	return &SRepo{
		db: dbParam,
	}
}

func (t *SRepo) SelectUser(u entity.User) (*entity.User, error) {
	res := entity.User{}
	err := t.db.Get(&res, "SELECT * FROM user WHERE id=$1",u.UserID)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (t *SRepo) SelectAllUser() ([]*entity.User, error) {
	var res []*entity.User
	err := t.db.Get(&res, "SELECT * FROM user")
	if err != nil {
		return nil, err
	}

	return res, err
}

func (t *SRepo) UpdateUser(u entity.User) (*entity.User, error) {
	_, err := t.db.Exec("UPDATE user set name=$1, password=$2 WHERE id=$3",u.Name,u.Password,u.UserID)
	if err != nil {
		return nil, err
	}

	return &u, err
}

func (t *SRepo) DeleteUser(u entity.User) error {
	_, err := t.db.Exec("DELETE FROM user WHERE id=$1",u.UserID)
	if err != nil {
		return err
	}

	return nil
}
