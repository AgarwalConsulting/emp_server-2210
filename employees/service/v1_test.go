package service_test

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"algogrit.com/emp_server/employees/repository"
	"algogrit.com/emp_server/employees/service"
	"algogrit.com/emp_server/entities"
)

func TestIndex(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockEmployeeRepository(ctrl)

	sut := service.NewV1(mockRepo)

	expectedEmps := []entities.Employee{
		{1, "Gaurav", "LnD", 1001},
	}

	mockRepo.EXPECT().ListAll(ctx).Return(expectedEmps, nil)

	employees, err := sut.Index(ctx)

	assert.Nil(t, err)

	assert.NotNil(t, employees)
	assert.NotEqual(t, 0, len(employees))

	assert.Equal(t, expectedEmps, employees)
}
