package repository

import (
	"context"

	"algogrit.com/emp_server/entities"
)

type inmem struct {
	employees []entities.Employee
}

func (repo *inmem) ListAll(ctx context.Context) ([]entities.Employee, error) {
	return repo.employees, nil
}

func (repo *inmem) Save(ctx context.Context, newEmp entities.Employee) (*entities.Employee, error) {
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
