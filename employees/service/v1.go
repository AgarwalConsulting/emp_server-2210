package service

import (
	"context"

	"algogrit.com/emp_server/employees/repository"
	"algogrit.com/emp_server/entities"
)

type v1Svc struct {
	repo repository.EmployeeRepository
}

func (svc *v1Svc) Index(ctx context.Context) ([]entities.Employee, error) {
	return svc.repo.ListAll(ctx)
}

func (svc *v1Svc) Create(ctx context.Context, newEmp entities.Employee) (*entities.Employee, error) {
	return svc.repo.Save(ctx, newEmp)
}

func NewV1(repo repository.EmployeeRepository) EmployeeService {
	return &v1Svc{repo}
}
