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

func TestStatusService_PrepareAllStatus(t *testing.T) {
	mockStatusRepo := new(mocks.IStatusRepo)
	var mockStatus entity.Status
	err := faker.FakeData(&mockStatus)
	assert.NoError(t, err)

	//errvar allStat []entity.Status

	mockListStatus := make([]entity.Status, 0)
	mockListStatus = append(mockListStatus, mockStatus)
	////
	//okay := mockStatusRepo.IsMethodCallable(t, "SelectAllStatus")
	//u := NewStatusServiceMock(mockStatusRepo)
	//u.PrepareAllStatus()

	assert.NotNil(t, mockStatusRepo.On("SelectAllStatus"))
	assert.NoError(t, err)

	//assert.Equal(t, true, okay)
	// mockStatusRepo.AssertCalled(t, "Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64"))
}

func TestNewStatusService(t *testing.T) {
	mockStatusRepo := new(mocks.IStatusRepo)
	//var mockStatus entity.Status
	//err := faker.FakeData(&mockStatus)
	//assert.NoError(t, err)

	mockStatusRepo.On("NewStatusService",mockStatusRepo).Return()
}

func TestStatusService_PrepareAllStatus1(t *testing.T) {
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
		//assert.Equal(t, cursorExpected, nextCursor)
		//assert.NotEmpty(t, nextCursor)
		//assert.NoError(t, err)
		//assert.Len(t, list, len(mockListStatus))
		assert.NotNil(t, MapStatus)

		mockStatusRepo.AssertExpectations(t)
		//mockAuthorrepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockStatusRepo.On("SelectAllStatus").Return(nil, errors.New("error mock"))
		u := NewStatusService(mockStatusRepo)
		u.PrepareAllStatus()
		//assert.Nil(t, MapStatus)
		mockStatusRepo.AssertExpectations(t)
	})

}