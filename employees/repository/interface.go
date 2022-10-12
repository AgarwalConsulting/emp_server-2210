package repository

import (
	"context"

	"algogrit.com/emp_server/entities"
)

type EmployeeRepository interface {
	ListAll(ctx context.Context) ([]entities.Employee, error)
	Save(ctx context.Context, newEmp entities.Employee) (*entities.Employee, error)
}
