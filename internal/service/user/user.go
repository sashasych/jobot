package user

import (
	"context"
	"errors"
	"fmt"
	"jobot/internal/repository"
	"jobot/internal/service/models"
	"time"

	"github.com/google/uuid"
)

var (
	ErrUserRoleNotFound = errors.New("user role not found")
)

type UserService struct {
	userRepository     repository.UserRepository
	employeeRepository repository.EmployeeRepository
	employerRepository repository.EmployerRepository
}

func NewUserService(userRepository repository.UserRepository, employeeRepository repository.EmployeeRepository, employerRepository repository.EmployerRepository) *UserService {
	return &UserService{userRepository: userRepository, employeeRepository: employeeRepository, employerRepository: employerRepository}
}

func (s *UserService) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	user.ID = uuid.New()
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	if err := s.userRepository.CreateUser(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
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

func (s *UserService) GetUserProfile(ctx context.Context, id uuid.UUID) (*models.UserProfile, error) {
	user, err := s.GetUser(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	switch user.Role {
	case "employee":
		employee, err := s.employeeRepository.GetEmployeeByUserID(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("failed to get employee: %w", err)
		}

		return &models.UserProfile{User: user, Employee: employee}, nil
	case "employer":
		employer, err := s.employerRepository.GetEmployerByUserID(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("failed to get employer: %w", err)
		}

		return &models.UserProfile{User: user, Employer: employer}, nil
	default:
		return nil, ErrUserRoleNotFound
	}
}
