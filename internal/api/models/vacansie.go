package models

import "time"

type VacansieCreateRequest struct {
	EmployerID  string   `json:"employer_id" validate:"required"`
	Tags        []string `json:"tags" validate:"required"`
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Location    string   `json:"location" validate:"required"`
	Salary      string   `json:"salary" validate:"required"`
}

type VacansieUpdateRequest struct {
	VacansieID  string    `json:"vacansie_id" validate:"required"`
	Tags        *[]string `json:"tags,omitempty"`
	Title       *string   `json:"title,omitempty"`
	Description *string   `json:"description,omitempty"`
	Location    *string   `json:"location,omitempty"`
	Salary      *string   `json:"salary,omitempty"`
}

type VacansieResponse struct {
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

type VacansieEmployerListResponse struct {
	Vacansies  []VacansieResponse `json:"vacansies"`
	EmployerID string             `json:"employer_id"`
}

// TODO: добавить pagination, search, filter
type VacansieListResponse struct {
	Vacansies []VacansieResponse `json:"vacansies"`
}
