package repository

import (
	"context"

	"algogrit.com/emp_server/entities"
)

//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE.go -package $GOPACKAGE

type EmployeeRepository interface {
	ListAll(ctx context.Context) ([]entities.Employee, error)
	Save(ctx context.Context, newEmp entities.Employee) (*entities.Employee, error)
}
