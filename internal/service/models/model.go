package models

import (
	"time"

	"github.com/google/uuid"
)

// User - модель пользователя в сервисном слое
type User struct {
	ID         uuid.UUID `json:"id"`
	TgUserName string    `json:"tg_user_name"`
	TgChatID   string    `json:"tg_chat_id"`
	IsActive   bool      `json:"is_active"`
	IsPremium  bool      `json:"is_premium"`
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// UserUpdateRequest - модель для обновления пользователя
type UserUpdateRequest struct {
	TgUserName *string `json:"tg_user_name"`
	TgChatID   *string `json:"tg_chat_id"`
	IsActive   *bool   `json:"is_active"`
	IsPremium  *bool   `json:"is_premium"`
	Role       *string `json:"role"`
}

// Employee - модель сотрудника
type Employee struct {
	EmployeeID string    `json:"employee_id"`
	UserID     string    `json:"user_id"`
	Tags       []string  `json:"tags"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// EmployeeUpdateRequest - модель для обновления сотрудника
type EmployeeUpdateRequest struct {
	Tags *[]string `json:"tags"`
}

// Employer - модель работодателя
type Employer struct {
	EmployerID         string    `json:"employer_id"`
	UserID             string    `json:"user_id"`
	CompanyName        string    `json:"company_name"`
	CompanyDescription string    `json:"company_description"`
	CompanyWebsite     string    `json:"company_website"`
	CompanyLocation    string    `json:"company_location"`
	CompanySize        string    `json:"company_size"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

// EmployerUpdateRequest - модель для обновления работодателя
type EmployerUpdateRequest struct {
	CompanyName        *string `json:"company_name"`
	CompanyDescription *string `json:"company_description"`
	CompanyWebsite     *string `json:"company_website"`
	CompanyLocation    *string `json:"company_location"`
	CompanySize        *string `json:"company_size"`
}

// Resume - модель резюме
type Resume struct {
	ResumeID   uuid.UUID `json:"resume_id"`
	EmployeeID uuid.UUID `json:"employee_id"`
	TgFileID   string    `json:"tg_file_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ResumeUpdateRequest - модель для обновления резюме
type ResumeUpdateRequest struct {
	TgFileID *string `json:"tg_file_id"`
}

// Vacancy - модель вакансии
type Vacancy struct {
	VacansieID  string    `json:"vacansie_id"`
	EmployerID  string    `json:"employer_id"`
	Tags        []string  `json:"tags"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Salary      string    `json:"salary"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// VacancyUpdateRequest - модель для обновления вакансии
type VacancyUpdateRequest struct {
	Tags        *[]string `json:"tags"`
	Title       *string   `json:"title"`
	Description *string   `json:"description"`
	Location    *string   `json:"location"`
	Salary      *string   `json:"salary"`
}

// EmployerVacancyList - модель списка вакансий работодателя
type EmployerVacancyList struct {
	Vacansies  []Vacancy `json:"vacansies"`
	EmployerID uuid.UUID `json:"employer_id"`
}

type VacancyList struct {
	Vacansies []Vacancy `json:"vacansies"`
}

// Reaction - модель реакции
type Reaction struct {
	ID         uuid.UUID `json:"id"`
	EmployeeID uuid.UUID `json:"employee_id"`
	VacancyID  uuid.UUID `json:"vacancy_id"`
	CreatedAt  time.Time `json:"created_at"`
}

// EmployeeReactionList - модель списка реакций сотрудника
type EmployeeReactionList struct {
	Reactions  []Reaction `json:"reactions"`
	EmployeeID uuid.UUID  `json:"employee_id"`
}
