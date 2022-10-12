package repository_test

import (
	"context"
	"sync"
	"testing"

	"algogrit.com/emp_server/employees/repository"
	"algogrit.com/emp_server/entities"
	"github.com/stretchr/testify/assert"
)

func TestConsistency(t *testing.T) {
	testCtx, _ := context.WithCancel(context.Background())

	sut := repository.NewInMem()

	originalEmps, err := sut.ListAll(testCtx)

	assert.Nil(t, err)
	assert.NotNil(t, originalEmps)
	assert.Equal(t, 3, len(originalEmps))

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			newEmp := entities.Employee{1, "Gaurav", "LnD", 1001}
			_, err := sut.Save(testCtx, newEmp)

			assert.Nil(t, err)

			_, err = sut.ListAll(testCtx)
			assert.Nil(t, err)
		}()
	}

	wg.Wait()

	updatedEmps, err := sut.ListAll(testCtx)

	assert.Nil(t, err)
	assert.NotNil(t, updatedEmps)
	assert.Equal(t, 103, len(updatedEmps))
}
