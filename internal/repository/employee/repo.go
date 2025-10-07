package employee

import (
	"context"
	"errors"
	"fmt"

	"jobot/internal/service/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrEmployeeNotFound      = errors.New("employee not found")
	ErrEmployeeAlreadyExists = errors.New("employee already exists")
)

type EmployeeRepository struct {
	db *pgxpool.Pool
}

func NewEmployeeRepository(db *pgxpool.Pool) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

// CreateEmployee создает нового сотрудника в БД
func (r *EmployeeRepository) CreateEmployee(ctx context.Context, employee *models.Employee) error {
	query := `
		INSERT INTO employees (employee_id, user_id, tags, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(ctx, query,
		employee.EmployeeID,
		employee.UserID,
		employee.Tags,
		employee.CreatedAt,
		employee.UpdatedAt,
	)

	if err != nil {
		if err.Error() == "duplicate key value violates unique constraint" {
			return ErrEmployeeAlreadyExists
		}

		return fmt.Errorf("failed to create employee: %w", err)
	}

	return nil
}

// GetEmployee получает сотрудника по ID
func (r *EmployeeRepository) GetEmployee(ctx context.Context, id uuid.UUID) (*models.Employee, error) {
	query := `
		SELECT employee_id, user_id, tags, created_at, updated_at
		FROM employees
		WHERE employee_id = $1
	`

	employee := &models.Employee{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&employee.EmployeeID,
		&employee.UserID,
		&employee.Tags,
		&employee.CreatedAt,
		&employee.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrEmployeeNotFound
		}
		return nil, fmt.Errorf("failed to get employee by id: %w", err)
	}

	return employee, nil
}

// GetEmployeeByUserID получает сотрудника по User ID
func (r *EmployeeRepository) GetEmployeeByUserID(ctx context.Context, userID uuid.UUID) (*models.Employee, error) {
	query := `
		SELECT employee_id, user_id, tags, created_at, updated_at
		FROM employees
		WHERE user_id = $1
	`

	employee := &models.Employee{}
	err := r.db.QueryRow(ctx, query, userID).Scan(
		&employee.EmployeeID,
		&employee.UserID,
		&employee.Tags,
		&employee.CreatedAt,
		&employee.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrEmployeeNotFound
		}
		return nil, fmt.Errorf("failed to get employee by user id: %w", err)
	}

	return employee, nil
}

// UpdateEmployee обновляет данные сотрудника
func (r *EmployeeRepository) UpdateEmployee(ctx context.Context, employee *models.Employee) error {
	query := `
		UPDATE employees
		SET tags = $2, updated_at = $3
		WHERE employee_id = $1
	`

	result, err := r.db.Exec(ctx, query,
		employee.EmployeeID,
		employee.Tags,
		employee.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to update employee: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrEmployeeNotFound
	}

	return nil
}

// DeleteEmployee удаляет сотрудника
func (r *EmployeeRepository) DeleteEmployee(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM employees WHERE employee_id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete employee: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrEmployeeNotFound
	}

	return nil
}
