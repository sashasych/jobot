package service

import (
	"context"
	"jobot/internal/service/models"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, id uuid.UUID) (*models.User, error)
	UpdateUser(ctx context.Context, req *models.UserUpdateRequest, id uuid.UUID) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type EmployeeService interface {
	CreateEmployee(ctx context.Context, employee *models.Employee) error
	GetEmployeeByID(ctx context.Context, id uuid.UUID) (*models.Employee, error)
	UpdateEmployee(ctx context.Context, req *models.EmployeeUpdateRequest, id uuid.UUID) error
	DeleteEmployee(ctx context.Context, id uuid.UUID) error
}

type ResumeService interface {
	CreateResume(ctx context.Context, resume *models.Resume) error
	GetResumeByID(ctx context.Context, id uuid.UUID) (*models.Resume, error)
	UpdateResume(ctx context.Context, req *models.ResumeUpdateRequest, id uuid.UUID) error
	DeleteResume(ctx context.Context, id uuid.UUID) error
}

type EmployerService interface {
	CreateEmployer(ctx context.Context, employer *models.Employer) error
	GetEmployerByID(ctx context.Context, id uuid.UUID) (*models.Employer, error)
	UpdateEmployer(ctx context.Context, req *models.EmployerUpdateRequest, id uuid.UUID) error
	DeleteEmployer(ctx context.Context, id uuid.UUID) error
}

type VacancyService interface {
	CreateVacancy(ctx context.Context, vacancy models.Vacancy) error
	GetVacancyByID(ctx context.Context, id uuid.UUID) (models.Vacancy, error)
	GetVacancyList(ctx context.Context) (models.VacancyList, error)
	UpdateVacancy(ctx context.Context, req models.VacancyUpdateRequest, id uuid.UUID) error
	DeleteVacancy(ctx context.Context, id uuid.UUID) error
}

type ReactionService interface {
	CreateReaction(ctx context.Context, reaction models.Reaction) error
	GetReactionByID(ctx context.Context, id uuid.UUID) (models.Reaction, error)
	DeleteReaction(ctx context.Context, id uuid.UUID) error
}
