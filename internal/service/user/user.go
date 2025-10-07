package user

import (
	"context"
	"fmt"
	"jobot/internal/repository"
	"jobot/internal/service/models"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return s.userRepository.CreateUser(ctx, user)
}

func (s *UserService) GetUser(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user, err := s.userRepository.GetUser(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *models.UserUpdateRequest, id uuid.UUID) error {
	getUser, err := s.GetUser(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	if req.TgUserName != nil {
		getUser.TgUserName = *req.TgUserName
	}

	if req.TgChatID != nil {
		getUser.TgChatID = *req.TgChatID
	}

	if req.IsActive != nil {
		getUser.IsActive = *req.IsActive
	}

	if req.IsPremium != nil {
		getUser.IsPremium = *req.IsPremium
	}

	if req.Role != nil {
		getUser.Role = *req.Role
	}

	getUser.UpdatedAt = time.Now()

	err = s.userRepository.UpdateUser(ctx, getUser)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := s.userRepository.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
