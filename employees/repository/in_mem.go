package repository

import (
	"context"
	"sync"

	"algogrit.com/emp_server/entities"
)

type inmem struct {
	employees []entities.Employee
	mut       sync.RWMutex
}

func (repo *inmem) ListAll(ctx context.Context) ([]entities.Employee, error) {
	repo.mut.RLock()
	defer repo.mut.RUnlock()

	return repo.employees, nil
}

func (repo *inmem) Save(ctx context.Context, newEmp entities.Employee) (*entities.Employee, error) {
	repo.mut.Lock()
	defer repo.mut.Unlock()

	newEmp.ID = len(repo.employees) + 1

	repo.employees = append(repo.employees, newEmp)

	return &newEmp, nil
}

func NewInMem() EmployeeRepository {
	var employees = []entities.Employee{
		{1, "Gaurav", "LnD", 1001},
		{2, "Senthil", "Cloud", 10002},
		{3, "Sonali", "SRE", 10010},
	}

	return &inmem{employees: employees}
}
