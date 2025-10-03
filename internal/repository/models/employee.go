package models

import (
	"time"

	"github.com/google/uuid"
)

// Employee - профиль соискателя (работника)
type Employee struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	UserID      uuid.UUID  `json:"user_id" db:"user_id"` // FK -> users.id
	FirstName   string     `json:"first_name" db:"first_name"`
	LastName    string     `json:"last_name" db:"last_name"`
	Phone       string     `json:"phone" db:"phone"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty" db:"date_of_birth"`
	Address     string     `json:"address" db:"address"`
	City        string     `json:"city" db:"city"`
	Country     string     `json:"country" db:"country"`
	Avatar      string     `json:"avatar,omitempty" db:"avatar"` // URL к фото
	Bio         string     `json:"bio" db:"bio"`                 // Краткая информация о себе
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}

// Resume - структура резюме на английском языке
type Resume struct {
	ID         uuid.UUID `json:"id" db:"id"`
	EmployeeID uuid.UUID `json:"employee_id" db:"employee_id"`

	// Основная информация
	Title     string `json:"title" db:"title"`         // "Senior Go Developer"
	Summary   string `json:"summary" db:"summary"`     // Professional summary
	Objective string `json:"objective" db:"objective"` // Career objective

	// Контактная информация
	ContactInfo ContactInfo `json:"contact_info" db:"contact_info"`

	// Структурированные данные
	Skills         Skills          `json:"skills" db:"skills"`
	Experience     Experiences     `json:"experience" db:"experience"`
	Education      []Education     `json:"education" db:"education"`
	Certifications []Certification `json:"certifications" db:"certifications"`
	Projects       []Project       `json:"projects" db:"projects"`
	Publications   []Publication   `json:"publications" db:"publications"`
	Languages      []Language      `json:"languages" db:"languages"`
	Awards         []Award         `json:"awards" db:"awards"`

	// Массивы строк
	Tags     []string `json:"tags" db:"tags"`
	Links    []string `json:"links" db:"links"`
	Keywords []string `json:"keywords" db:"keywords"`

	// Метаданные
	Language  string    `json:"language" db:"language"` // "en", "ru", etc.
	Country   string    `json:"country" db:"country"`   // "US", "RU", etc.
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// ContactInfo - контактная информация
type ContactInfo struct {
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Location  string `json:"location"` // "New York, NY" or "Moscow, Russia"
	LinkedIn  string `json:"linkedin"`
	GitHub    string `json:"github"`
	Website   string `json:"website"`
	Portfolio string `json:"portfolio"`
}

// Skill - навык с уровнем владения
type Skill struct {
	Name     string `json:"name"`      // "Go", "PostgreSQL", "Docker"
	Level    string `json:"level"`     // "Beginner", "Intermediate", "Advanced", "Expert"
	Years    int    `json:"years"`     // Years of experience
	Category string `json:"category"`  // "Programming Languages", "Databases", "Tools"
	LastUsed string `json:"last_used"` // "2024-01-01"
}

type Skills []Skill

// Experience - опыт работы
type Experience struct {
	Company          string   `json:"company"`
	Position         string   `json:"position"`
	Location         string   `json:"location"`
	StartDate        string   `json:"start_date"` // "2020-01-01"
	EndDate          *string  `json:"end_date,omitempty"`
	Current          bool     `json:"current"`
	Description      string   `json:"description"`
	Responsibilities []string `json:"responsibilities"`
	Achievements     []string `json:"achievements"`
	Technologies     []string `json:"technologies"`
}

type Experiences []Experience

// Education - образование
type Education struct {
	Institution     string   `json:"institution"`
	Degree          string   `json:"degree"` // "Bachelor of Science", "Master of Science"
	Field           string   `json:"field"`  // "Computer Science", "Software Engineering"
	StartDate       string   `json:"start_date"`
	EndDate         *string  `json:"end_date,omitempty"`
	GPA             *float64 `json:"gpa,omitempty"`
	Location        string   `json:"location"`
	RelevantCourses []string `json:"relevant_courses"`
}

// Certification - сертификация
type Certification struct {
	Name         string  `json:"name"`
	Issuer       string  `json:"issuer"` // "AWS", "Google", "Microsoft"
	IssueDate    string  `json:"issue_date"`
	ExpiryDate   *string `json:"expiry_date,omitempty"`
	CredentialID string  `json:"credential_id"`
	URL          string  `json:"url"`
}

// Project - проект
type Project struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Technologies []string `json:"technologies"`
	StartDate    string   `json:"start_date"`
	EndDate      *string  `json:"end_date,omitempty"`
	URL          string   `json:"url"`
	GitHub       string   `json:"github"`
	Role         string   `json:"role"` // "Lead Developer", "Contributor"
	TeamSize     int      `json:"team_size"`
}

// Publication - публикация
type Publication struct {
	Title           string   `json:"title"`
	Authors         []string `json:"authors"`
	Journal         string   `json:"journal"`
	PublicationDate string   `json:"publication_date"`
	URL             string   `json:"url"`
	DOI             string   `json:"doi"`
}

// Language - язык
type Language struct {
	Name        string `json:"name"`        // "English", "Russian", "Spanish"
	Proficiency string `json:"proficiency"` // "Native", "Fluent", "Intermediate", "Basic"
	Speaking    string `json:"speaking"`
	Reading     string `json:"reading"`
	Writing     string `json:"writing"`
}

// Award - награда
type Award struct {
	Name        string `json:"name"`
	Issuer      string `json:"issuer"`
	Date        string `json:"date"`
	Description string `json:"description"`
	URL         string `json:"url"`
}
