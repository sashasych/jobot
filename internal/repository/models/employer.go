package models

import (
	"time"

	"github.com/google/uuid"
)

// Employer - профиль работодателя (представителя компании)
type Employer struct {
	ID        uuid.UUID `json:"id" db:"id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`       // FK -> users.id
	CompanyID uuid.UUID `json:"company_id" db:"company_id"` // FK -> companies.id
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Position  string    `json:"position" db:"position"` // Должность в компании
	Phone     string    `json:"phone" db:"phone"`
	Avatar    string    `json:"avatar,omitempty" db:"avatar"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Company - компания работодателя
type Company struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Industry    string    `json:"industry" db:"industry"`         // "IT", "Finance", "Healthcare"
	CompanySize string    `json:"company_size" db:"company_size"` // "1-10", "11-50", "51-200", "201-500", "500+"
	Website     string    `json:"website" db:"website"`
	Logo        string    `json:"logo,omitempty" db:"logo"` // URL к логотипу
	Address     string    `json:"address" db:"address"`
	City        string    `json:"city" db:"city"`
	Country     string    `json:"country" db:"country"`
	Founded     *int      `json:"founded,omitempty" db:"founded"` // Год основания
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// JobPosting - вакансия от работодателя
type JobPosting struct {
	ID               uuid.UUID `json:"id" db:"id"`
	EmployerID       uuid.UUID `json:"employer_id" db:"employer_id"` // FK -> employers.id
	CompanyID        uuid.UUID `json:"company_id" db:"company_id"`   // FK -> companies.id
	Title            string    `json:"title" db:"title"`
	Description      string    `json:"description" db:"description"`
	Requirements     []string  `json:"requirements" db:"requirements"`
	Responsibilities []string  `json:"responsibilities" db:"responsibilities"`
	Skills           []string  `json:"skills" db:"skills"`

	// Условия работы
	EmploymentType string `json:"employment_type" db:"employment_type"` // "Full-time", "Part-time", "Contract", "Freelance"
	WorkLocation   string `json:"work_location" db:"work_location"`     // "Remote", "On-site", "Hybrid"
	Experience     string `json:"experience" db:"experience"`           // "Junior", "Middle", "Senior"

	// Зарплата
	SalaryMin      *int   `json:"salary_min,omitempty" db:"salary_min"`
	SalaryMax      *int   `json:"salary_max,omitempty" db:"salary_max"`
	SalaryCurrency string `json:"salary_currency" db:"salary_currency"` // "USD", "EUR", "RUB"
	SalaryPeriod   string `json:"salary_period" db:"salary_period"`     // "hour", "month", "year"

	// Локация
	City    string `json:"city" db:"city"`
	Country string `json:"country" db:"country"`

	// Метаданные
	Status            string     `json:"status" db:"status"` // "active", "closed", "draft"
	ViewsCount        int        `json:"views_count" db:"views_count"`
	ApplicationsCount int        `json:"applications_count" db:"applications_count"`
	ExpiresAt         *time.Time `json:"expires_at,omitempty" db:"expires_at"`
	CreatedAt         time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at" db:"updated_at"`
}

// JobApplication - отклик на вакансию
type JobApplication struct {
	ID           uuid.UUID `json:"id" db:"id"`
	JobPostingID uuid.UUID `json:"job_posting_id" db:"job_posting_id"` // FK -> job_postings.id
	EmployeeID   uuid.UUID `json:"employee_id" db:"employee_id"`       // FK -> employees.id
	ResumeID     uuid.UUID `json:"resume_id" db:"resume_id"`           // FK -> resumes.id
	CoverLetter  string    `json:"cover_letter" db:"cover_letter"`
	Status       string    `json:"status" db:"status"` // "pending", "reviewed", "interview", "rejected", "accepted"
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
