package repository

import (
	"github.com/jmoiron/sqlx"
	"log"
	"todo-togo/entity"
)

type IUserRepo interface {
	SelectUser(u entity.User) (*entity.User, error)
	SelectAllUser() ([]*entity.User, error)
	AddUser(u entity.User) (*entity.User, error)
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
	err := t.db.Get(&res, "SELECT * FROM user WHERE id=$1 OR email=$2",u.UserID,u.Email)
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
	_, err := t.db.Exec("UPDATE user set name=$1, password=$2, email=$3 WHERE id=$4",u.Name,u.Password,u.Email,u.UserID)
	if err != nil {
		return nil, err
	}

	return &u, err
}

func (t *SRepo) AddUser(u entity.User) (*entity.User, error) {
	res, err := t.db.Exec("INSERT INTO user (name, email, password, salt) VALUES ($1,$2,$3,$4)",u.Name,u.Email,u.Password,u.Salt)
	if err != nil {
		return nil, err
	}

	id,err := res.LastInsertId()
	if err != nil {
		log.Println("error on get last insertedID",err.Error())
	}
	u.UserID= int(id)

	return &u, err
}

func (t *SRepo) DeleteUser(u entity.User) error {
	_, err := t.db.Exec("DELETE FROM user WHERE id=$1",u.UserID)
	if err != nil {
		return err
	}

	return nil
}
