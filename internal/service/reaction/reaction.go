package reaction

import (
	"context"
	"fmt"
	"time"

	"jobot/internal/repository"
	"jobot/internal/service/models"

	"github.com/google/uuid"
)

type ReactionService struct {
	reactionRepository repository.ReactionRepository
}

func NewReactionService(reactionRepository repository.ReactionRepository) *ReactionService {
	return &ReactionService{reactionRepository: reactionRepository}
}

func (s *ReactionService) CreateReaction(ctx context.Context, reaction *models.Reaction) (*models.Reaction, error) {
	reaction.ID = uuid.New()
	reaction.CreatedAt = time.Now()

	if err := s.reactionRepository.CreateReaction(ctx, reaction); err != nil {
		return nil, fmt.Errorf("failed to create reaction: %w", err)
	}

	return reaction, nil
}

func (s *ReactionService) GetReaction(ctx context.Context, id uuid.UUID) (*models.Reaction, error) {
	reaction, err := s.reactionRepository.GetReaction(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get reaction by ID: %w", err)
	}

	return reaction, nil
}

func (s *ReactionService) GetEmployeeReactions(ctx context.Context, employeeID uuid.UUID) (*models.EmployeeReactionList, error) {
	reactionList, err := s.reactionRepository.GetReactionsByEmployee(ctx, employeeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get reactions by employee ID: %w", err)
	}

	return reactionList, nil
}

func (s *ReactionService) DeleteReaction(ctx context.Context, id uuid.UUID) error {
	err := s.reactionRepository.DeleteReaction(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete reaction: %w", err)
	}

	return nil
}
