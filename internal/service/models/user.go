package models

import (
	"time"

	"github.com/google/uuid"
)

// User - модель пользователя в сервисном слое
type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // не возвращаем в JSON
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRole - модель роли пользователя в сервисном слое
type UserRole struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

// UserWithRoles - пользователь с ролями
type UserWithRoles struct {
	User  User       `json:"user"`
	Roles []UserRole `json:"roles"`
}
