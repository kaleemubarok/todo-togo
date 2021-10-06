package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"todo-togo/repository/mocks"
)

func NewStatusServiceMock(repo *mocks.IStatusRepo) IStatusService  {
	return &StatusService{stat: repo}
}

func TestStatusService_PrepareAllStatus(t *testing.T) {
	mockStatusService := NewStatusServiceMock(&mocks.IStatusRepo{})

	assert.error(t, mockStatusService.PrepareAllStatus)
}
