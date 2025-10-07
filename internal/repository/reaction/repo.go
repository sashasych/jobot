package reaction

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
	ErrReactionNotFound      = errors.New("reaction not found")
	ErrReactionAlreadyExists = errors.New("reaction already exists")
)

type ReactionRepository struct {
	db *pgxpool.Pool
}

func NewReactionRepository(db *pgxpool.Pool) *ReactionRepository {
	return &ReactionRepository{db: db}
}

// CreateReaction создает новую реакцию в БД
func (r *ReactionRepository) CreateReaction(ctx context.Context, reaction *models.Reaction) error {
	query := `
		INSERT INTO reactions (id, employee_id, vacancy_id, created_at)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(ctx, query,
		reaction.ID,
		reaction.EmployeeID,
		reaction.VacancyID,
		reaction.CreatedAt,
	)

	if err != nil {
		if err.Error() == "duplicate key value violates unique constraint" {
			return ErrReactionAlreadyExists
		}

		return fmt.Errorf("failed to create reaction: %w", err)
	}

	return nil
}

// GetReaction получает реакцию по ID
func (r *ReactionRepository) GetReaction(ctx context.Context, id uuid.UUID) (*models.Reaction, error) {
	query := `
		SELECT id, employee_id, vacancy_id, created_at
		FROM reactions
		WHERE id = $1
	`

	reaction := &models.Reaction{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&reaction.ID,
		&reaction.EmployeeID,
		&reaction.VacancyID,
		&reaction.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrReactionNotFound
		}
		return nil, fmt.Errorf("failed to get reaction by id: %w", err)
	}

	return reaction, nil
}

// GetReactionsByEmployee получает реакции сотрудника
func (r *ReactionRepository) GetReactionsByEmployee(ctx context.Context, employeeID uuid.UUID) (*models.EmployeeReactionList, error) {
	query := `
		SELECT id, employee_id, vacancy_id, created_at
		FROM reactions
		WHERE employee_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(ctx, query, employeeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get reactions by employee: %w", err)
	}
	defer rows.Close()

	reactions := make([]models.Reaction, 0)
	for rows.Next() {
		var reaction models.Reaction
		err := rows.Scan(
			&reaction.ID,
			&reaction.EmployeeID,
			&reaction.VacancyID,
			&reaction.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan reaction: %w", err)
		}
		reactions = append(reactions, reaction)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating reactions: %w", err)
	}

	return &models.EmployeeReactionList{
		Reactions:  reactions,
		EmployeeID: employeeID,
	}, nil
}

// DeleteReaction удаляет реакцию
func (r *ReactionRepository) DeleteReaction(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM reactions WHERE id = $1`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete reaction: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrReactionNotFound
	}

	return nil
}
