package service

import (
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
	"todo-togo/model"
	"todo-togo/repository"
)

type todoServiceMock struct {
	todo,status,user   mock.Mock
}


func Test_todoService_AddTodo(t1 *testing.T) {
	todo := model.CreateTodoRequest{
		ID:          0,
		Title:       "Title test",
		Description: "title test desc details",
		DueDate:     "2021-10-20",
		PIC:         1,
		Status:      1,
	}


}

func Test_todoService_DeleteTodo(t1 *testing.T) {
	type fields struct {
		todo   repository.ITodoRepo
		status repository.IStatusRepo
		user   repository.IStatusRepo
	}
	type args struct {
		todo model.CreateTodoRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := todoService{
				todo:   tt.fields.todo,
				status: tt.fields.status,
				user:   tt.fields.user,
			}
			if err := t.DeleteTodo(tt.args.todo); (err != nil) != tt.wantErr {
				t1.Errorf("DeleteTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_todoService_GetAllTodo(t1 *testing.T) {
	type fields struct {
		todo   repository.ITodoRepo
		status repository.IStatusRepo
		user   repository.IStatusRepo
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*model.CreateTodoResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := todoService{
				todo:   tt.fields.todo,
				status: tt.fields.status,
				user:   tt.fields.user,
			}
			got, err := t.GetAllTodo()
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetAllTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetAllTodo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todoService_GetTodo(t1 *testing.T) {
	type fields struct {
		todo   repository.ITodoRepo
		status repository.IStatusRepo
		user   repository.IStatusRepo
	}
	type args struct {
		todo model.CreateTodoRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.CreateTodoResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := todoService{
				todo:   tt.fields.todo,
				status: tt.fields.status,
				user:   tt.fields.user,
			}
			got, err := t.GetTodo(tt.args.todo)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetTodo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todoService_ModifyTodo(t1 *testing.T) {
	type fields struct {
		todo   repository.ITodoRepo
		status repository.IStatusRepo
		user   repository.IStatusRepo
	}
	type args struct {
		todo model.CreateTodoRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.CreateTodoResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := todoService{
				todo:   tt.fields.todo,
				status: tt.fields.status,
				user:   tt.fields.user,
			}
			got, err := t.ModifyTodo(tt.args.todo)
			if (err != nil) != tt.wantErr {
				t1.Errorf("ModifyTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("ModifyTodo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
