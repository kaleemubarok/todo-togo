package service

import (
	"errors"
	"log"
	"strconv"
	"todo-togo/entity"
	"todo-togo/model"
	"todo-togo/repository"
)

type ITodoService interface {
	AddTodo(todo model.CreateTodoRequest) (*model.CreateTodoResponse, error)
	ModifyTodo(todo model.CreateTodoRequest) (*model.CreateTodoResponse, error)
	GetTodo(todo model.CreateTodoRequest) (*model.CreateTodoResponse, error)
	GetAllTodo() ([]*model.CreateTodoResponse, error)
	DeleteTodo(todo model.CreateTodoRequest) error
}

type todoService struct {
	todo   repository.ITodoRepo
	status repository.IStatusRepo
	user   repository.IStatusRepo
}

func NewTodoService(repo *repository.ITodoRepo) ITodoService {
	return todoService{todo: *repo}
}

func (t todoService) AddTodo(todo model.CreateTodoRequest) (*model.CreateTodoResponse, error) {
	if len(MapStatus) == 0 {
		log.Println("status->", todo.Status, len(MapStatus[todo.Status]))
		return nil, errors.New("error on AddTodo service, wrong status code/id")
	}

	//TODO validate for user id if correct already

	prepareTodo := entity.Todo{
		Title:       todo.Title,
		Description: todo.Description,
		DueDate:     todo.DueDate,
		PIC:         todo.PIC,
		Status:      todo.Status,
	}

	newTodo, err := t.todo.CreateTodo(prepareTodo)
	if err != nil {
		return nil, err
	}

	return &model.CreateTodoResponse{
		ID:          newTodo.ID,
		Title:       newTodo.Title,
		Description: newTodo.Description,
		DueDate:     newTodo.DueDate,
		PIC:         strconv.Itoa(newTodo.PIC),
		Status:      MapStatus[newTodo.Status],
	}, nil
}

func (t todoService) ModifyTodo(todo model.CreateTodoRequest) (*model.CreateTodoResponse, error) {
	if len(MapStatus) == 0 {
		log.Println("status->", todo.Status, len(MapStatus[todo.Status]))
		return nil, errors.New("error on AddTodo service, wrong status code/id")
	}

	//TODO validate for user id if correct already

	prepareTodo := entity.Todo{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		DueDate:     todo.DueDate,
		PIC:         todo.PIC,
		Status:      todo.Status,
	}

	//check if data exist
	_, err := t.todo.SelectTodo(prepareTodo)
	if err != nil {
		return nil, err
	}

	newTodo, err := t.todo.UpdateTodo(prepareTodo)
	if err != nil {
		return nil, err
	}

	return &model.CreateTodoResponse{
		ID:          newTodo.ID,
		Title:       newTodo.Title,
		Description: newTodo.Description,
		DueDate:     newTodo.DueDate,
		PIC:         strconv.Itoa(newTodo.PIC),
		Status:      MapStatus[newTodo.Status],
	}, nil
}

func (t todoService) GetTodo(todo model.CreateTodoRequest) (*model.CreateTodoResponse, error) {
	requestedID := entity.Todo{
		ID: todo.ID,
	}
	todoItem, err := t.todo.SelectTodo(requestedID)
	if err != nil {
		return nil, err
	}

	res := model.CreateTodoResponse{
		ID:          todoItem.ID,
		Title:       todoItem.Title,
		Description: todoItem.Description,
		DueDate:     todoItem.DueDate,
		PIC:         todoItem.PICName,
		Status:      todoItem.StatusDesc,
	}

	return &res, nil
}

func (t todoService) GetAllTodo() ([]*model.CreateTodoResponse, error) {
	todos, err := t.todo.SelectAllTodo()
	if err != nil {
		return nil, err
	}

	var res []*model.CreateTodoResponse
	for _, todo := range todos {
		res = append(res, &model.CreateTodoResponse{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			DueDate:     todo.DueDate,
			PIC:         todo.PICName,
			Status:      todo.StatusDesc,
		})
	}

	return res, nil
}

func (t todoService) DeleteTodo(todo model.CreateTodoRequest) error {
	deletedItem := entity.Todo{
		ID: todo.ID,
	}

	_, err := t.todo.SelectTodo(deletedItem)
	if err != nil {
		return err
	}

	err = t.todo.DeleteTodo(deletedItem)
	if err != nil {
		return err
	}

	return nil
}
