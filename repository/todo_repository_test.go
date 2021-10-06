package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
	"reflect"
	"regexp"
	"testing"
	"todo-togo/entity"
)



func TestSRepo_CreateTodo(t1 *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t1.Fatalf("an error '%s' was not expected when opening a stub" +
			"database connection", err)
	}
	defer db.Close()
	rows := sqlxmock.NewRows([]string{"id", "desc"}).
		AddRow(1,"New").
		AddRow(2,"OnGoing").
		AddRow(3,"Done").
		AddRow(4,"Deleted")

	query := regexp.QuoteMeta("SELECT * FROM status")
	mock.ExpectQuery(query).WillReturnRows(rows)
	a := NewStatusRepo(db)
	res, err := a.SelectAllStatus()
	assert.NoError(t2, err)
	assert.NotNil(t2, res)
}

func TestSRepo_DeleteTodo(t1 *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		todo entity.Todo
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
			t := &SRepo{
				db: tt.fields.db,
			}
			if err := t.DeleteTodo(tt.args.todo); (err != nil) != tt.wantErr {
				t1.Errorf("DeleteTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSRepo_SelectAllTodo(t1 *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*entity.Todo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &SRepo{
				db: tt.fields.db,
			}
			got, err := t.SelectAllTodo()
			if (err != nil) != tt.wantErr {
				t1.Errorf("SelectAllTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("SelectAllTodo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSRepo_SelectTodo(t1 *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		todo entity.Todo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Todo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &SRepo{
				db: tt.fields.db,
			}
			got, err := t.SelectTodo(tt.args.todo)
			if (err != nil) != tt.wantErr {
				t1.Errorf("SelectTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("SelectTodo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSRepo_UpdateTodo(t1 *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		todo entity.Todo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Todo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &SRepo{
				db: tt.fields.db,
			}
			got, err := t.UpdateTodo(tt.args.todo)
			if (err != nil) != tt.wantErr {
				t1.Errorf("UpdateTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("UpdateTodo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
