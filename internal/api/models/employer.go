package models

import "time"

type EmployerCreateRequest struct {
	UserID             string `json:"user_id" validate:"required"`
	CompanyName        string `json:"company_name" validate:"required"`
	CompanyDescription string `json:"company_description" validate:"required"`
	CompanyWebsite     string `json:"company_website" validate:"required"`
	CompanyLocation    string `json:"company_location" validate:"required"`
	CompanySize        string `json:"company_size" validate:"required"`
}

type EmployerUpdateRequest struct {
	EmployerID         string  `json:"employer_id" validate:"required"`
	CompanyName        *string `json:"company_name,omitempty"`
	CompanyDescription *string `json:"company_description,omitempty"`
	CompanyWebsite     *string `json:"company_website,omitempty"`
	CompanyLocation    *string `json:"company_location,omitempty"`
	CompanySize        *string `json:"company_size,omitempty"`
}

type EmployerResponse struct {
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
