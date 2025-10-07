package models

import "time"

// При создании сотрудника, мы создаем запись в таблице employees
// и связываем ее с пользователем по внешнему ключу user_id
// и запись в таблице resumes, связываем ее с сотрудником по внешнему ключу employee_id

// EmployeeCreateRequest - DTO для создания сотрудника
// EmployeeID - ID сотрудника
// Tags - Теги сотрудника

type EmployeeCreateRequest struct {
	UserID string   `json:"user_id" validate:"required"`
	Tags   []string `json:"tags" validate:"required"`
}

// EmployeeResponse - DTO для получения сотрудника
// EmployeeID - ID сотрудника
// UserID - ID пользователя, внешний ключ
// Tags - Теги сотрудника
// CreatedAt - Дата создания
// UpdatedAt - Дата обновления

type EmployeeUpdateRequest struct {
	Tags *[]string `json:"tags,omitempty"`
}

type EmployeeResponse struct {
	EmployeeID string    `json:"employee_id"`
	UserID     string    `json:"user_id"`
	Tags       []string  `json:"tags"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
