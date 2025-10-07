package employer

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
	ErrEmployerNotFound      = errors.New("employer not found")
	ErrEmployerAlreadyExists = errors.New("employer already exists")
)

type EmployerRepository struct {
	db *pgxpool.Pool
}

func NewEmployerRepository(db *pgxpool.Pool) *EmployerRepository {
	return &EmployerRepository{db: db}
}

// CreateEmployer создает нового работодателя в БД
func (r *EmployerRepository) CreateEmployer(ctx context.Context, employer *models.Employer) error {
	query := `
		INSERT INTO employers (employer_id, user_id, company_name, company_description, company_website, company_location, company_size, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.Exec(ctx, query,
		employer.EmployerID,
		employer.UserID,
		employer.CompanyName,
		employer.CompanyDescription,
		employer.CompanyWebsite,
		employer.CompanyLocation,
		employer.CompanySize,
		employer.CreatedAt,
		employer.UpdatedAt,
	)

	if err != nil {
		if err.Error() == "duplicate key value violates unique constraint" {
			return ErrEmployerAlreadyExists
		}

		return fmt.Errorf("failed to create employer: %w", err)
	}

	return nil
}

// GetEmployer получает работодателя по ID
func (r *EmployerRepository) GetEmployer(ctx context.Context, id uuid.UUID) (*models.Employer, error) {
	query := `
		SELECT employer_id, user_id, company_name, company_description, company_website, company_location, company_size, created_at, updated_at
		FROM employers
		WHERE employer_id = $1
	`

	employer := &models.Employer{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&employer.EmployerID,
		&employer.UserID,
		&employer.CompanyName,
		&employer.CompanyDescription,
		&employer.CompanyWebsite,
		&employer.CompanyLocation,
		&employer.CompanySize,
		&employer.CreatedAt,
		&employer.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrEmployerNotFound
		}
		return nil, fmt.Errorf("failed to get employer by id: %w", err)
	}

	return employer, nil
}

// GetEmployerByUserID получает работодателя по User ID
func (r *EmployerRepository) GetEmployerByUserID(ctx context.Context, userID uuid.UUID) (*models.Employer, error) {
	query := `
		SELECT employer_id, user_id, company_name, company_description, company_website, company_location, company_size, created_at, updated_at
		FROM employers
		WHERE user_id = $1
	`

	employer := &models.Employer{}
	err := r.db.QueryRow(ctx, query, userID).Scan(
		&employer.EmployerID,
		&employer.UserID,
		&employer.CompanyName,
		&employer.CompanyDescription,
		&employer.CompanyWebsite,
		&employer.CompanyLocation,
		&employer.CompanySize,
		&employer.CreatedAt,
		&employer.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrEmployerNotFound
		}
		return nil, fmt.Errorf("failed to get employer by user id: %w", err)
	}

	return employer, nil
}

// UpdateEmployer обновляет данные работодателя
func (r *EmployerRepository) UpdateEmployer(ctx context.Context, employer *models.Employer) error {
	query := `
		UPDATE employers
		SET company_name = $2, company_description = $3, company_website = $4, company_location = $5, company_size = $6, updated_at = $7
		WHERE employer_id = $1
	`

	result, err := r.db.Exec(ctx, query,
		employer.EmployerID,
		employer.CompanyName,
		employer.CompanyDescription,
		employer.CompanyWebsite,
		employer.CompanyLocation,
		employer.CompanySize,
		employer.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to update employer: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrEmployerNotFound
	}

	return nil
}

// DeleteEmployer удаляет работодателя
func (r *EmployerRepository) DeleteEmployer(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM employers WHERE employer_id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete employer: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrEmployerNotFound
	}

	return nil
}
