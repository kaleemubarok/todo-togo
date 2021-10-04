package repository

import (
	"github.com/jmoiron/sqlx"
	"todo-togo/entity"
)

type iUserRepo interface {
	SelectUser(status entity.User) (entity.User, error)
}

func NewUserRepo(dbParam *sqlx.DB) iUserRepo {
	return &sRepo{
		db: dbParam,
	}
}

func (t *sRepo) SelectUser(u entity.User) (entity.User, error) {
	res := entity.User{}
	err := t.db.Get(&res, "SELECT * FROM user WHERE id=$1",u.UserID)
	if err != nil {
		return res, err
	}

	return res, nil
}
