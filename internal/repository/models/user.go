package models

import (
	"time"

	"github.com/google/uuid"
)

// User - базовая модель пользователя для аутентификации
type User struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"-" db:"password_hash"` // не возвращаем в JSON
	IsActive     bool      `json:"is_active" db:"is_active"`
	FirstName    string    `json:"first_name" db:"first_name"`
	LastName     string    `json:"last_name" db:"last_name"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// UserRole - роль пользователя в системе
type UserRole struct {
	ID        uuid.UUID `json:"id" db:"id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	Role      string    `json:"role" db:"role"` // "employee", "employer", "admin"
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
