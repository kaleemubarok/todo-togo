package service

import (
	"reflect"
	"testing"
	"todo-togo/model"
	"todo-togo/repository"
)

func TestUserService_GetAllUser(t *testing.T) {
	type fields struct {
		repo repository.IUserRepo
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*model.UserReqResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserService{
				repo: tt.fields.repo,
			}
			got, err := u.GetAllUser()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.GetAllUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.GetAllUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
