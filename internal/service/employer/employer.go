package employer

import (
	"context"
	"fmt"
	"time"

	"jobot/internal/repository"
	"jobot/internal/service/models"

	"github.com/google/uuid"
)

type EmployerService struct {
	employerRepository repository.EmployerRepository
}

func NewEmployerService(employerRepository repository.EmployerRepository) *EmployerService {
	return &EmployerService{employerRepository: employerRepository}
}

func (s *EmployerService) CreateEmployer(ctx context.Context, employer *models.Employer) error {
	// Генерируем UUID для employer
	employer.EmployerID = uuid.New()
	employer.CreatedAt = time.Now()
	employer.UpdatedAt = time.Now()

	return s.employerRepository.CreateEmployer(ctx, employer)
}

func (s *EmployerService) GetEmployer(ctx context.Context, id uuid.UUID) (*models.Employer, error) {
	employer, err := s.employerRepository.GetEmployer(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get employer by ID: %w", err)
	}

	return employer, nil
}

func (s *EmployerService) GetEmployerByUserID(ctx context.Context, userID uuid.UUID) (*models.Employer, error) {
	employer, err := s.employerRepository.GetEmployerByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get employer by user ID: %w", err)
	}

	return employer, nil
}

func (s *EmployerService) UpdateEmployer(ctx context.Context, req *models.EmployerUpdateRequest, id uuid.UUID) error {
	getEmployer, err := s.GetEmployer(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get employer: %w", err)
	}

	// Обновляем только переданные поля
	if req.CompanyName != nil {
		getEmployer.CompanyName = *req.CompanyName
	}

	if req.CompanyDescription != nil {
		getEmployer.CompanyDescription = *req.CompanyDescription
	}

	if req.CompanyWebsite != nil {
		getEmployer.CompanyWebsite = *req.CompanyWebsite
	}

	if req.CompanyLocation != nil {
		getEmployer.CompanyLocation = *req.CompanyLocation
	}

	if req.CompanySize != nil {
		getEmployer.CompanySize = *req.CompanySize
	}

	getEmployer.UpdatedAt = time.Now()

	err = s.employerRepository.UpdateEmployer(ctx, getEmployer)
	if err != nil {
		return fmt.Errorf("failed to update employer: %w", err)
	}

	return nil
}

func (s *EmployerService) DeleteEmployer(ctx context.Context, id uuid.UUID) error {
	err := s.employerRepository.DeleteEmployer(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete employer: %w", err)
	}

	return nil
}
