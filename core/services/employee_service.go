package services

import (
	"ams-service/core/entities"
	"ams-service/middlewares"
	"fmt"
)

var EMPLOYEE_LOG_PREFIX string = "employee_service.go"

type EmployeeRepository interface {
	GetEmployeeByID(request entities.GetEmployeeByIdRequest) (entities.Employee, error)
	RegisterEmployee(request entities.RegisterEmployeeRequest) error
}

type EmployeeService struct {
	repo EmployeeRepository
}

func NewEmployeeService(repo EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) GetEmployeeByID(request entities.GetEmployeeByIdRequest) (entities.Employee, error) {
	employee, err := s.repo.GetEmployeeByID(request)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error getting employee by ID: %v", EMPLOYEE_LOG_PREFIX, err))
		return entities.Employee{}, err
	}
	return employee, nil
}

func (s *EmployeeService) RegisterEmployee(request entities.RegisterEmployeeRequest) error {
	err := s.repo.RegisterEmployee(request)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error registering employee: %v", EMPLOYEE_LOG_PREFIX, err))
		return err
	}
	return nil
}
