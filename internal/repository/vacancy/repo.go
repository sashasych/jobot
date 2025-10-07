package vacancy

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
	ErrVacancyNotFound      = errors.New("vacancy not found")
	ErrVacancyAlreadyExists = errors.New("vacancy already exists")
)

type VacancyRepository struct {
	db *pgxpool.Pool
}

func NewVacancyRepository(db *pgxpool.Pool) *VacancyRepository {
	return &VacancyRepository{db: db}
}

// CreateVacancy создает новую вакансию в БД
func (r *VacancyRepository) CreateVacancy(ctx context.Context, vacancy *models.Vacancy) error {
	query := `
		INSERT INTO vacancies (vacansie_id, employer_id, tags, title, description, location, salary, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.Exec(ctx, query,
		vacancy.VacansieID,
		vacancy.EmployerID,
		vacancy.Tags,
		vacancy.Title,
		vacancy.Description,
		vacancy.Location,
		vacancy.Salary,
		vacancy.CreatedAt,
		vacancy.UpdatedAt,
	)

	if err != nil {
		if err.Error() == "duplicate key value violates unique constraint" {
			return ErrVacancyAlreadyExists
		}

		return fmt.Errorf("failed to create vacancy: %w", err)
	}

	return nil
}

// GetVacancy получает вакансию по ID
func (r *VacancyRepository) GetVacancy(ctx context.Context, id uuid.UUID) (*models.Vacancy, error) {
	query := `
		SELECT vacansie_id, employer_id, tags, title, description, location, salary, created_at, updated_at
		FROM vacancies
		WHERE vacansie_id = $1
	`

	vacancy := &models.Vacancy{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&vacancy.VacansieID,
		&vacancy.EmployerID,
		&vacancy.Tags,
		&vacancy.Title,
		&vacancy.Description,
		&vacancy.Location,
		&vacancy.Salary,
		&vacancy.CreatedAt,
		&vacancy.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrVacancyNotFound
		}
		return nil, fmt.Errorf("failed to get vacancy by id: %w", err)
	}

	return vacancy, nil
}

// GetVacancyList получает список всех вакансий
func (r *VacancyRepository) GetVacancyList(ctx context.Context) (*models.VacancyList, error) {
	query := `
		SELECT vacansie_id, employer_id, tags, title, description, location, salary, created_at, updated_at
		FROM vacancies
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get vacancy list: %w", err)
	}
	defer rows.Close()

	vacancies := make([]models.Vacancy, 0)
	for rows.Next() {
		var vacancy models.Vacancy
		err := rows.Scan(
			&vacancy.VacansieID,
			&vacancy.EmployerID,
			&vacancy.Tags,
			&vacancy.Title,
			&vacancy.Description,
			&vacancy.Location,
			&vacancy.Salary,
			&vacancy.CreatedAt,
			&vacancy.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan vacancy: %w", err)
		}
		vacancies = append(vacancies, vacancy)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating vacancies: %w", err)
	}

	return &models.VacancyList{Vacansies: vacancies}, nil
}

// GetVacanciesByEmployer получает вакансии работодателя
func (r *VacancyRepository) GetVacanciesByEmployer(ctx context.Context, employerID uuid.UUID) (*models.EmployerVacancyList, error) {
	query := `
		SELECT vacansie_id, employer_id, tags, title, description, location, salary, created_at, updated_at
		FROM vacancies
		WHERE employer_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query, employerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get vacancies by employer: %w", err)
	}
	defer rows.Close()

	vacancies := make([]models.Vacancy, 0)
	for rows.Next() {
		var vacancy models.Vacancy
		err := rows.Scan(
			&vacancy.VacansieID,
			&vacancy.EmployerID,
			&vacancy.Tags,
			&vacancy.Title,
			&vacancy.Description,
			&vacancy.Location,
			&vacancy.Salary,
			&vacancy.CreatedAt,
			&vacancy.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan vacancy: %w", err)
		}
		vacancies = append(vacancies, vacancy)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating vacancies: %w", err)
	}

	return &models.EmployerVacancyList{
		Vacansies:  vacancies,
		EmployerID: employerID,
	}, nil
}

// UpdateVacancy обновляет данные вакансии
func (r *VacancyRepository) UpdateVacancy(ctx context.Context, vacancy *models.Vacancy) error {
	query := `
		UPDATE vacancies
		SET tags = $2, title = $3, description = $4, location = $5, salary = $6, updated_at = $7
		WHERE vacansie_id = $1
	`

	result, err := r.db.Exec(ctx, query,
		vacancy.VacansieID,
		vacancy.Tags,
		vacancy.Title,
		vacancy.Description,
		vacancy.Location,
		vacancy.Salary,
		vacancy.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to update vacancy: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrVacancyNotFound
	}

	return nil
}

// DeleteVacancy удаляет вакансию
func (r *VacancyRepository) DeleteVacancy(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM vacancies WHERE vacansie_id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete vacancy: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrVacancyNotFound
	}

	return nil
}
