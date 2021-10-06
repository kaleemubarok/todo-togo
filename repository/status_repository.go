package repository

import (
	"github.com/jmoiron/sqlx"
	"todo-togo/entity"
)

type IStatusRepo interface {
	SelectAllStatus() ([]entity.Status, error)
}

func NewStatusRepo(dbParam *sqlx.DB) IStatusRepo {
	return &SRepo{
		db: dbParam,
	}
}

func (t *SRepo) SelectAllStatus() ([]entity.Status, error) {
	var res []entity.Status
	err := t.db.Select(&res,"SELECT * FROM status")
	if err != nil {
		return nil, err
	}

	return res, nil
}