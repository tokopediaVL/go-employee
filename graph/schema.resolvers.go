package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"github.com/tkp-VL037/go-employee/db"
	gm "github.com/tkp-VL037/go-employee/graph/model"
	"github.com/tkp-VL037/go-employee/model"
	"gorm.io/gorm"
)

// AddEmployee is the resolver for the addEmployee field.
func (r *mutationResolver) AddEmployee(ctx context.Context, input *gm.NewEmployee) (*gm.EmployeeResponse, error) {
	panic(fmt.Errorf("not implemented: AddEmployee - addEmployee"))
}

// UpdateEmployeeDetail is the resolver for the updateEmployeeDetail field.
func (r *mutationResolver) UpdateEmployeeDetail(ctx context.Context, input *gm.UpdateEmployee) (*gm.EmployeeResponse, error) {
	panic(fmt.Errorf("not implemented: UpdateEmployeeDetail - updateEmployeeDetail"))
}

// DeleteEmployee is the resolver for the deleteEmployee field.
func (r *mutationResolver) DeleteEmployee(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteEmployee - deleteEmployee"))
}

// GetEmployees is the resolver for the getEmployees field.
func (r *queryResolver) GetEmployees(ctx context.Context) ([]*gm.EmployeeResponse, error) {
	var employees []*model.Employee

	result := db.DB.Preload("Statistic").Find(&employees)
	if result.Error != nil {
		return nil, result.Error
	}

	employeesStatsResponse := make([]*gm.EmployeeResponse, len(employees))
	for i, es := range employees {
		employeesStatsResponse[i] = &gm.EmployeeResponse{
			ID:        es.ID,
			Name:      es.Name,
			Age:       es.Age,
			Position:  es.Position,
			ViewCount: int(es.Statistic.ViewCount),
		}
	}

	return employeesStatsResponse, nil
}

// GetEmployeeDetail is the resolver for the getEmployeeDetail field.
func (r *queryResolver) GetEmployeeDetail(ctx context.Context, id string) (*gm.EmployeeResponse, error) {
	var employee *model.Employee
	if err := db.DB.Preload("Statistic").First(&employee, "id = ?", id).Error; err != nil {
		return nil, err
	}

	err := db.DB.Model(&model.Statistic{}).Where("employee_id = ?", id).
		UpdateColumn("view_count", gorm.Expr("view_count + ?", 1)).Error
	if err != nil {
		return nil, err
	}

	return &gm.EmployeeResponse{
		ID:        employee.ID,
		Name:      employee.Name,
		Age:       employee.Age,
		Position:  employee.Position,
		ViewCount: int(employee.Statistic.ViewCount),
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
