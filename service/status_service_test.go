package service

import (
	"errors"
	"testing"
	"todo-togo/entity"
	"todo-togo/repository/mocks"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
)

func NewStatusServiceMock(repo *mocks.IStatusRepo) IStatusService {
	return &StatusService{stat: repo}
}

func TestNewStatusService(t *testing.T) {
	mockStatusRepo := new(mocks.IStatusRepo)
	mockStatusRepo.On("NewStatusService",mockStatusRepo).Return()
}

func TestStatusService_PrepareAllStatus(t *testing.T) {
	mockStatusRepo := new(mocks.IStatusRepo)
	var mockStatus entity.Status
	err := faker.FakeData(&mockStatus)
	assert.NoError(t, err)

	mockListStatus := make([]entity.Status, 0)
	mockListStatus = append(mockListStatus, mockStatus)
	mockStatusRepo.On("SelectAllStatus").Return(mockListStatus, nil)

	t.Run("success", func(t *testing.T) {
		mockListSuccessStatus := append(mockListStatus, mockStatus)
		mockStatusRepo.On("SelectAllStatus").Return(mockListSuccessStatus, nil)
		u := NewStatusService(mockStatusRepo)
		u.PrepareAllStatus()

		assert.NotNil(t, MapStatus)

		mockStatusRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockStatusRepo.On("SelectAllStatus").Return(nil, errors.New("error mock"))
		u := NewStatusService(mockStatusRepo)
		u.PrepareAllStatus()

		mockStatusRepo.AssertExpectations(t)
	})

}

func FillAllStatDummy()  {
	MapStatus = map[int]string{
		1:"New",
		2:"OnGoing",
		3:"Done",
		4:"Deleted",
	}
}

func ResetAllStatDummy()  {
	MapStatus = make(map[int]string,0)
}