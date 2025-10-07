package vacancy

import (
	"context"
	"fmt"
	"time"

	"jobot/internal/repository"
	"jobot/internal/service/models"

	"github.com/google/uuid"
)

type VacancyService struct {
	vacancyRepository repository.VacancyRepository
}

func NewVacancyService(vacancyRepository repository.VacancyRepository) *VacancyService {
	return &VacancyService{vacancyRepository: vacancyRepository}
}

func (s *VacancyService) CreateVacancy(ctx context.Context, vacancy *models.Vacancy) error {
	// Генерируем UUID для vacancy
	vacancy.VacansieID = uuid.New()
	vacancy.CreatedAt = time.Now()
	vacancy.UpdatedAt = time.Now()

	return s.vacancyRepository.CreateVacancy(ctx, vacancy)
}

func (s *VacancyService) GetVacancyByID(ctx context.Context, id uuid.UUID) (*models.Vacancy, error) {
	vacancy, err := s.vacancyRepository.GetVacancy(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get vacancy by ID: %w", err)
	}

	return vacancy, nil
}

func (s *VacancyService) GetEmployerVacancies(ctx context.Context, employerID uuid.UUID) (*models.EmployerVacancyList, error) {
	vacancyList, err := s.vacancyRepository.GetVacanciesByEmployer(ctx, employerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get vacancies by employer ID: %w", err)
	}

	return vacancyList, nil
}

func (s *VacancyService) GetVacancyList(ctx context.Context) (*models.VacancyList, error) {
	vacancyList, err := s.vacancyRepository.GetVacancyList(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get vacancies list: %w", err)
	}

	return vacancyList, nil
}

func (s *VacancyService) UpdateVacancy(ctx context.Context, req *models.VacancyUpdateRequest, id uuid.UUID) error {
	getVacancy, err := s.vacancyRepository.GetVacancy(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get vacancy: %w", err)
	}

	// Обновляем только переданные поля
	if req.Tags != nil {
		getVacancy.Tags = *req.Tags
	}

	if req.Title != nil {
		getVacancy.Title = *req.Title
	}

	if req.Description != nil {
		getVacancy.Description = *req.Description
	}

	if req.Location != nil {
		getVacancy.Location = *req.Location
	}

	if req.Salary != nil {
		getVacancy.Salary = *req.Salary
	}

	getVacancy.UpdatedAt = time.Now()

	err = s.vacancyRepository.UpdateVacancy(ctx, getVacancy)
	if err != nil {
		return fmt.Errorf("failed to update vacancy: %w", err)
	}

	return nil
}

func (s *VacancyService) DeleteVacancy(ctx context.Context, id uuid.UUID) error {
	err := s.vacancyRepository.DeleteVacancy(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete vacancy: %w", err)
	}

	return nil
}
