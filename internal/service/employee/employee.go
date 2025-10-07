package employee

import (
	"context"
	"fmt"
	"time"

	"jobot/internal/repository"
	"jobot/internal/service/models"

	"github.com/google/uuid"
)

type EmployeeService struct {
	employeeRepository repository.EmployeeRepository
}

func NewEmployeeService(employeeRepository repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{employeeRepository: employeeRepository}
}

func (s *EmployeeService) CreateEmployee(ctx context.Context, employee *models.Employee) error {
	// Генерируем UUID для employee
	employee.EmployeeID = uuid.New()
	employee.CreatedAt = time.Now()
	employee.UpdatedAt = time.Now()

	return s.employeeRepository.CreateEmployee(ctx, employee)
}

func (s *EmployeeService) GetEmployee(ctx context.Context, id uuid.UUID) (*models.Employee, error) {
	employee, err := s.employeeRepository.GetEmployee(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get employee by ID: %w", err)
	}

	return employee, nil
}

func (s *EmployeeService) GetEmployeeByUserID(ctx context.Context, userID uuid.UUID) (*models.Employee, error) {
	employee, err := s.employeeRepository.GetEmployeeByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get employee by user ID: %w", err)
	}

	return employee, nil
}

func (s *EmployeeService) UpdateEmployee(ctx context.Context, req *models.EmployeeUpdateRequest, id uuid.UUID) error {
	getEmployee, err := s.GetEmployee(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get employee by user ID: %w", err)
	}

	// Обновляем только переданные поля
	if req.Tags != nil {
		getEmployee.Tags = *req.Tags
	}

	getEmployee.UpdatedAt = time.Now()

	err = s.employeeRepository.UpdateEmployee(ctx, getEmployee)
	if err != nil {
		return fmt.Errorf("failed to update employee: %w", err)
	}

	return nil
}

func (s *EmployeeService) DeleteEmployee(ctx context.Context, id uuid.UUID) error {
	err := s.employeeRepository.DeleteEmployee(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete employee: %w", err)
	}

	return nil
}
