package resume

import (
	"context"
	"fmt"
	"time"

	"jobot/internal/repository"
	"jobot/internal/service/models"

	"github.com/google/uuid"
)

type ResumeService struct {
	resumeRepository repository.ResumeRepository
}

func NewResumeService(resumeRepository repository.ResumeRepository) *ResumeService {
	return &ResumeService{resumeRepository: resumeRepository}
}

func (s *ResumeService) CreateResume(ctx context.Context, resume *models.Resume) (*models.Resume, error) {
	resume.ResumeID = uuid.New()
	now := time.Now()
	resume.CreatedAt = now
	resume.UpdatedAt = now

	if err := s.resumeRepository.CreateResume(ctx, resume); err != nil {
		return nil, fmt.Errorf("failed to create resume: %w", err)
	}

	return resume, nil
}

func (s *ResumeService) GetResume(ctx context.Context, id uuid.UUID) (*models.Resume, error) {
	resume, err := s.resumeRepository.GetResume(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get resume: %w", err)
	}

	return resume, nil
}

func (s *ResumeService) GetResumeByEmployeeID(ctx context.Context, employeeID uuid.UUID) (*models.Resume, error) {
	resume, err := s.resumeRepository.GetResumeByEmployeeID(ctx, employeeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get resume by employee ID: %w", err)
	}

	return resume, nil
}

func (s *ResumeService) UpdateResume(ctx context.Context, req *models.ResumeUpdateRequest, id uuid.UUID) error {
	getResume, err := s.resumeRepository.GetResume(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get resume: %w", err)
	}

	// Обновляем только переданные поля
	if req.TgFileID != nil {
		getResume.TgFileID = *req.TgFileID
	}

	getResume.UpdatedAt = time.Now()

	err = s.resumeRepository.UpdateResume(ctx, getResume)
	if err != nil {
		return fmt.Errorf("failed to update resume: %w", err)
	}

	return nil
}

func (s *ResumeService) DeleteResume(ctx context.Context, id uuid.UUID) error {
	err := s.resumeRepository.DeleteResume(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete resume: %w", err)
	}

	return nil
}
