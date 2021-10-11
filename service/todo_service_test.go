package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"todo-togo/entity"
	"todo-togo/model"
	"todo-togo/repository/mocks"
)

func TestNewTodoService(t *testing.T) {
	mockTodoService := new(mocks.ITodoRepo)
	mockTodoService.On("NewTodoService", mockTodoService).Return()
}

func Test_todoService_AddTodo(t *testing.T) {
	mockTodoService := new(mocks.ITodoRepo)
	var mockNewTodo model.CreateTodoRequest

	mockNewTodo=model.CreateTodoRequest{
		Title:       "New Todo",
		Description: "New Desc",
		DueDate:     "2021-10-20",
		PIC:         1,
		Status:      1,
	}

	prepareTodo:= entity.Todo{
		Title:       mockNewTodo.Title,
		Description: mockNewTodo.Description,
		DueDate:     mockNewTodo.DueDate,
		PIC:         mockNewTodo.PIC,
		Status:      mockNewTodo.Status,
	}

	t.Run("map-empty", func(t *testing.T) {
		ResetAllStatDummy()
		NewTodoService(mockTodoService)
		u := NewTodoService(mockTodoService)
		res,err:=u.AddTodo(mockNewTodo)

		assert.Nil(t, res)
		assert.Error(t, err)
		mockTodoService.AssertExpectations(t)

	})

	t.Run("success", func(t *testing.T) {
		FillAllStatDummy()

		//mockTodoService.On("SelectAllStatus").Return(mockListSuccessStatus, nil)
		mockTodoService.On("CreateTodo", prepareTodo).Return(&entity.Todo{}, nil).Once()
		u := NewTodoService(mockTodoService)
		res, err := u.AddTodo(mockNewTodo)

		assert.Nil(t, err)
		assert.NotNil(t, res)

		mockTodoService.AssertExpectations(t)
	})

	t.Run("error-createTodo", func(t *testing.T) {
		FillAllStatDummy()

		//mockTodoService.On("SelectAllStatus").Return(mockListSuccessStatus, nil)
		mockTodoService.On("CreateTodo", entity.Todo{}).Return(nil, errors.New("error returned")).Once()
		u := NewTodoService(mockTodoService)
		res, err := u.AddTodo(model.CreateTodoRequest{})

		assert.Nil(t, res)
		assert.NotNil(t, err)

		mockTodoService.AssertExpectations(t)
	})
}

func Test_todoService_DeleteTodo(t *testing.T) {
	mockTodoService := new(mocks.ITodoRepo)
	var mockDeletedTodo model.CreateTodoRequest

	mockDeletedTodo =model.CreateTodoRequest{ID: 1}

	deletedItem := entity.Todo{ID: mockDeletedTodo.ID}

	t.Run("error-select", func(t *testing.T) {
		FillAllStatDummy()
		NewTodoService(mockTodoService)
		mockTodoService.On("SelectTodo",deletedItem).Return(&entity.Todo{},errors.New("")).Once()
		u := NewTodoService(mockTodoService)
		err:=u.DeleteTodo(mockDeletedTodo)

		assert.Error(t, err)
		mockTodoService.AssertExpectations(t)

	})

	t.Run("error-deleteTodo", func(t *testing.T) {
		FillAllStatDummy()

		mockTodoService.On("SelectTodo",deletedItem).Return(&deletedItem,nil).Once()
		mockTodoService.On("DeleteTodo", deletedItem).Return(errors.New("error returned")).Once()
		u := NewTodoService(mockTodoService)
		err := u.DeleteTodo(mockDeletedTodo)

		assert.NotNil(t, err)
		mockTodoService.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		FillAllStatDummy()

		mockTodoService.On("SelectTodo",deletedItem).Return(&entity.Todo{},nil).Once()
		mockTodoService.On("DeleteTodo", deletedItem).Return(nil).Once()
		u := NewTodoService(mockTodoService)
		err := u.DeleteTodo(mockDeletedTodo)

		assert.Nil(t, err)

		mockTodoService.AssertExpectations(t)
	})


}

func Test_todoService_GetAllTodo(t *testing.T) {
	mockTodoService := new(mocks.ITodoRepo)

	t.Run("error-select", func(t *testing.T) {

		mockTodoService.On("SelectAllTodo").Return(nil, errors.New("error on selectTodo")).Once()
		u := NewTodoService(mockTodoService)
		_, err := u.GetAllTodo()

		assert.NotNil(t, err)
		mockTodoService.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		var todos = []*entity.Todo{
			{ID: 1, Title: "1", Description: "1", DueDate: "2020-10-10", PIC: 1, PICName: "1", Status: 1, StatusDesc: "New"},
			{ID: 2, Title: "2", Description: "2", DueDate: "2020-10-10", PIC: 2, PICName: "2", Status: 2, StatusDesc: "OnGoing"},
		}
		mockTodoService.On("SelectAllTodo").Return(todos, nil).Once()
		u := NewTodoService(mockTodoService)
		res, err := u.GetAllTodo()

		assert.Nil(t, err)
		assert.NotNil(t, res)

		mockTodoService.AssertExpectations(t)
	})
}

func Test_todoService_GetTodo(t *testing.T) {
	mockTodoService := new(mocks.ITodoRepo)
	var mockRequestedID model.CreateTodoRequest

	mockRequestedID =model.CreateTodoRequest{ID: 1}

	requestedID := entity.Todo{ID: mockRequestedID.ID}

	t.Run("error-select", func(t *testing.T) {

		mockTodoService.On("SelectTodo", requestedID).Return(nil, errors.New("error on selectTodo")).Once()
		u := NewTodoService(mockTodoService)
		_, err := u.GetTodo(mockRequestedID)

		assert.NotNil(t, err)
		mockTodoService.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		mockTodoService.On("SelectTodo", requestedID).Return(&entity.Todo{}, nil).Once()
		u := NewTodoService(mockTodoService)
		res, err := u.GetTodo(mockRequestedID)

		assert.Nil(t, err)
		assert.NotNil(t, res)

		mockTodoService.AssertExpectations(t)
	})
}

func Test_todoService_ModifyTodo(t *testing.T) {
	mockTodoService := new(mocks.ITodoRepo)
	var mockModifyTodo model.CreateTodoRequest

	mockModifyTodo =model.CreateTodoRequest{
		Title:       "New Todo",
		Description: "New Desc",
		DueDate:     "2021-10-20",
		PIC:         1,
		Status:      1,
	}

	modifiedTodo := entity.Todo{
		Title:       mockModifyTodo.Title,
		Description: mockModifyTodo.Description,
		DueDate:     mockModifyTodo.DueDate,
		PIC:         mockModifyTodo.PIC,
		Status:      mockModifyTodo.Status,
	}

	t.Run("map-empty", func(t *testing.T) {
		ResetAllStatDummy()
		NewTodoService(mockTodoService)
		u := NewTodoService(mockTodoService)
		res,err:=u.ModifyTodo(mockModifyTodo)

		assert.Nil(t, res)
		assert.Error(t, err)
		mockTodoService.AssertExpectations(t)
	})

	t.Run("error-select", func(t *testing.T) {
		FillAllStatDummy()

		mockTodoService.On("SelectTodo", modifiedTodo).Return(nil, errors.New("error on selectTodo")).Once()
		u := NewTodoService(mockTodoService)
		_, err := u.ModifyTodo(mockModifyTodo)

		assert.NotNil(t, err)

		mockTodoService.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		FillAllStatDummy()

		mockTodoService.On("SelectTodo", modifiedTodo).Return(&entity.Todo{}, nil).Once()
		mockTodoService.On("UpdateTodo", modifiedTodo).Return(&entity.Todo{}, nil).Once()
		u := NewTodoService(mockTodoService)
		res, err := u.ModifyTodo(mockModifyTodo)

		assert.Nil(t, err)
		assert.NotNil(t, res)

		mockTodoService.AssertExpectations(t)
	})

	t.Run("error-updateTodo", func(t *testing.T) {
		FillAllStatDummy()

		mockTodoService.On("SelectTodo", modifiedTodo).Return(&entity.Todo{}, nil).Once()
		mockTodoService.On("UpdateTodo", modifiedTodo).Return(nil, errors.New("")).Once()
		u := NewTodoService(mockTodoService)
		res, err := u.ModifyTodo(mockModifyTodo)

		assert.Nil(t, res)
		assert.NotNil(t, err)

		mockTodoService.AssertExpectations(t)
	})
}
