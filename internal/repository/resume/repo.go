package resume

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
	ErrResumeNotFound      = errors.New("resume not found")
	ErrResumeAlreadyExists = errors.New("resume already exists")
)

type ResumeRepository struct {
	db *pgxpool.Pool
}

func NewResumeRepository(db *pgxpool.Pool) *ResumeRepository {
	return &ResumeRepository{db: db}
}

// CreateResume создает новое резюме в БД
func (r *ResumeRepository) CreateResume(ctx context.Context, resume *models.Resume) error {
	query := `
		INSERT INTO resumes (resume_id, employee_id, tg_file_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(ctx, query,
		resume.ResumeID,
		resume.EmployeeID,
		resume.TgFileID,
		resume.CreatedAt,
		resume.UpdatedAt,
	)

	if err != nil {
		if err.Error() == "duplicate key value violates unique constraint" {
			return ErrResumeAlreadyExists
		}

		return fmt.Errorf("failed to create resume: %w", err)
	}

	return nil
}

// GetResume получает резюме по ID
func (r *ResumeRepository) GetResume(ctx context.Context, id uuid.UUID) (*models.Resume, error) {
	query := `
		SELECT resume_id, employee_id, tg_file_id, created_at, updated_at
		FROM resumes
		WHERE resume_id = $1
	`

	resume := &models.Resume{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&resume.ResumeID,
		&resume.EmployeeID,
		&resume.TgFileID,
		&resume.CreatedAt,
		&resume.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrResumeNotFound
		}
		return nil, fmt.Errorf("failed to get resume by id: %w", err)
	}

	return resume, nil
}

// GetResumeByEmployeeID получает резюме по Employee ID
func (r *ResumeRepository) GetResumeByEmployeeID(ctx context.Context, employeeID uuid.UUID) (*models.Resume, error) {
	query := `
		SELECT resume_id, employee_id, tg_file_id, created_at, updated_at
		FROM resumes
		WHERE employee_id = $1
	`

	resume := &models.Resume{}
	err := r.db.QueryRow(ctx, query, employeeID).Scan(
		&resume.ResumeID,
		&resume.EmployeeID,
		&resume.TgFileID,
		&resume.CreatedAt,
		&resume.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrResumeNotFound
		}
		return nil, fmt.Errorf("failed to get resume by employee id: %w", err)
	}

	return resume, nil
}

// UpdateResume обновляет данные резюме
func (r *ResumeRepository) UpdateResume(ctx context.Context, resume *models.Resume) error {
	query := `
		UPDATE resumes
		SET tg_file_id = $2, updated_at = $3
		WHERE resume_id = $1
	`

	result, err := r.db.Exec(ctx, query,
		resume.ResumeID,
		resume.TgFileID,
		resume.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to update resume: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrResumeNotFound
	}

	return nil
}

// DeleteResume удаляет резюме
func (r *ResumeRepository) DeleteResume(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM resumes WHERE resume_id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete resume: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrResumeNotFound
	}

	return nil
}
