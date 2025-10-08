package repository

import (
	"context"
	"jobot/internal/service/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, userService *models.User) error
	GetUser(ctx context.Context, id uuid.UUID) (*models.User, error)
	UpdateUser(ctx context.Context, userService *models.User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type EmployeeRepository interface {
	CreateEmployee(ctx context.Context, employeeService *models.Employee) error
	GetEmployee(ctx context.Context, id uuid.UUID) (*models.Employee, error)
	GetEmployeeByUserID(ctx context.Context, userID uuid.UUID) (*models.Employee, error)
	UpdateEmployee(ctx context.Context, employeeService *models.Employee) error
	DeleteEmployee(ctx context.Context, id uuid.UUID) error
}

type EmployerRepository interface {
	CreateEmployer(ctx context.Context, employerService *models.Employer) error
	GetEmployer(ctx context.Context, id uuid.UUID) (*models.Employer, error)
	GetEmployerByUserID(ctx context.Context, userID uuid.UUID) (*models.Employer, error)
	UpdateEmployer(ctx context.Context, employerService *models.Employer) error
	DeleteEmployer(ctx context.Context, id uuid.UUID) error
}

type ResumeRepository interface {
	CreateResume(ctx context.Context, resumeService *models.Resume) error
	GetResume(ctx context.Context, id uuid.UUID) (*models.Resume, error)
	GetResumeByEmployeeID(ctx context.Context, employeeID uuid.UUID) (*models.Resume, error)
	UpdateResume(ctx context.Context, resumeService *models.Resume) error
	DeleteResume(ctx context.Context, id uuid.UUID) error
}

type VacancyRepository interface {
	CreateVacancy(ctx context.Context, vacancyService *models.Vacancy) error
	GetVacancy(ctx context.Context, id uuid.UUID) (*models.Vacancy, error)
	GetVacancyList(ctx context.Context) (*models.VacancyList, error)
	GetVacanciesByEmployer(ctx context.Context, employerID uuid.UUID) (*models.EmployerVacancyList, error)
	UpdateVacancy(ctx context.Context, vacancyService *models.Vacancy) error
	DeleteVacancy(ctx context.Context, id uuid.UUID) error
}

type ReactionRepository interface {
	CreateReaction(ctx context.Context, reactionService *models.Reaction) error
	GetReaction(ctx context.Context, id uuid.UUID) (*models.Reaction, error)
	GetReactionsByEmployee(ctx context.Context, employeeID uuid.UUID) (*models.EmployeeReactionList, error)
	DeleteReaction(ctx context.Context, id uuid.UUID) error
}
