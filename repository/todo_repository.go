package repository

import (
	"github.com/jmoiron/sqlx"
	"log"
	"todo-togo/entity"
)

type sRepo struct {
	db *sqlx.DB
}

type iTodoRepo interface {
	CreateTodo(todo entity.Todo) (entity.Todo, error)
	UpdateTodo(todo entity.Todo) (entity.Todo, error)
	SelectTodo(todo entity.Todo) (entity.Todo, error)
	SelectAllTodo() ([]entity.Todo, error)
	DeleteTodo(todo entity.Todo) error
}

func NewTodoRepo(dbParam *sqlx.DB) iTodoRepo {
	return &sRepo{
		db: dbParam,
	}
}

func (t *sRepo) CreateTodo(todo entity.Todo) (entity.Todo, error)  {
	res, err := t.db.Exec("INSERT INTO todo (title, \"desc\", due_date, pic, status) VALUES ($1, $2, $3, $4, $5)",
		todo.Title,todo.Description,todo.DueDate,todo.PIC,todo.Status)

	if err != nil {
		log.Println("error on db exec:",err.Error())
		return entity.Todo{}, err
	}

	id,err := res.LastInsertId()
	if err != nil {
		log.Println("error on get last insertedID",err.Error())
	}
	todo.ID= int(id)

	return todo, nil
}

func (t *sRepo) UpdateTodo(todo entity.Todo) (entity.Todo, error) {
	_, err := t.db.Exec("UPDATE todo set title=$1, \"desc\"=$2, due_date=$3, pic=$4, status=$5 WHERE id=$6",
		todo.Title, todo.Description, todo.DueDate, todo.PIC, todo.Status, todo.ID)

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (t *sRepo) SelectTodo(todo entity.Todo) (entity.Todo, error) {
	res := entity.Todo{}
	err := t.db.Get(&res, "SELECT t.*, s.desc status_desc, u.name pic_desc FROM todo t " +
		"INNER JOIN status s ON s.ID=t.status " +
		"INNER JOIN user u ON u.ID=t.pic " +
		"where t.id=$1",todo.ID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (t *sRepo) SelectAllTodo() ([]entity.Todo, error) {
	var res []entity.Todo
	err := t.db.Select(&res,"SELECT t.*, s.desc status_desc, u.name pic_desc FROM todo t " +
		"INNER JOIN status s ON s.ID=t.status " +
		"INNER JOIN user u ON u.ID=t.pic ")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *sRepo) DeleteTodo(todo entity.Todo) error {
	_, err := t.db.Exec("DELETE FROM todo WHERE id=$1",todo.ID)
	if err != nil {
		return err
	}

	return nil
}