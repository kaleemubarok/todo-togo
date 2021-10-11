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
		mockTodoService.On("CreateTodo", prepareTodo).Return(&entity.Todo{}, nil)
		u := NewTodoService(mockTodoService)
		u.AddTodo(mockNewTodo)
		res, err := mockTodoService.CreateTodo(prepareTodo)

		assert.Nil(t, err)
		assert.NotNil(t, res)

		mockTodoService.AssertExpectations(t)
	})

	t.Run("error-createTodo", func(t *testing.T) {
		FillAllStatDummy()

		//mockTodoService.On("SelectAllStatus").Return(mockListSuccessStatus, nil)
		mockTodoService.On("CreateTodo", entity.Todo{}).Return(nil, errors.New("error returned"))
		u := NewTodoService(mockTodoService)
		u.AddTodo(model.CreateTodoRequest{})
		res, err := mockTodoService.CreateTodo(entity.Todo{})

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

//func Test_todoService_GetAllTodo(t1 *testing.T) {
//	type fields struct {
//		todo   repository.ITodoRepo
//		status repository.IStatusRepo
//		user   repository.IStatusRepo
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		want    []*model.CreateTodoResponse
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t1.Run(tt.name, func(t1 *testing.T) {
//			t := todoService{
//				todo:   tt.fields.todo,
//				status: tt.fields.status,
//				user:   tt.fields.user,
//			}
//			got, err := t.GetAllTodo()
//			if (err != nil) != tt.wantErr {
//				t1.Errorf("GetAllTodo() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t1.Errorf("GetAllTodo() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_todoService_GetTodo(t1 *testing.T) {
//	type fields struct {
//		todo   repository.ITodoRepo
//		status repository.IStatusRepo
//		user   repository.IStatusRepo
//	}
//	type args struct {
//		todo model.CreateTodoRequest
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *model.CreateTodoResponse
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t1.Run(tt.name, func(t1 *testing.T) {
//			t := todoService{
//				todo:   tt.fields.todo,
//				status: tt.fields.status,
//				user:   tt.fields.user,
//			}
//			got, err := t.GetTodo(tt.args.todo)
//			if (err != nil) != tt.wantErr {
//				t1.Errorf("GetTodo() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t1.Errorf("GetTodo() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_todoService_ModifyTodo(t *testing.T) {

}
