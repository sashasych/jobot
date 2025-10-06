package models

import "time"

// ResumeCreateRequest - DTO для создания резюме
// EmployeeID - ID сотрудника
// TgFileID - ID файла в Telegram

type ResumeCreateRequest struct {
	EmployeeID string `json:"employee_id" validate:"required"`
	TgFileID   string `json:"tg_file_id" validate:"required"`
}

// ResumeResponse - DTO для получения резюме
// EmployeeID - ID сотрудника
// TgFileID - ID файла в Telegram
// CreatedAt - Дата создания
// UpdatedAt - Дата обновления

type ResumeResponse struct {
	ResumeID   string    `json:"resume_id"`
	EmployeeID string    `json:"employee_id"`
	TgFileID   string    `json:"tg_file_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ResumeUpdateRequest - DTO для обновления резюме
// EmployeeID - ID сотрудника
// TgFileID - ID файла в Telegram

type ResumeUpdateRequest struct {
	ResumeID string  `json:"resume_id" validate:"required"`
	TgFileID *string `json:"tg_file_id,omitempty"`
}
