package models

import (
	"time"
)

// UserCreateRequest - DTO для создания пользователя (API → Service)
type UserCreateRequest struct {
	TgUserName string `json:"tg_user_name" validate:"required"`
	TgChatID   string `json:"tg_chat_id" validate:"required"`
	IsActive   bool   `json:"is_active" validate:"required"`
	IsPremium  bool   `json:"is_premium" validate:"required"`
	Role       string `json:"role" validate:"required,oneof=employee employer"`
}

// UserUpdateRequest - DTO для обновления пользователя (API → Service)
type UserUpdateRequest struct {
	TgUserName *string `json:"tg_user_name,omitempty"`
	TgChatID   *string `json:"tg_chat_id,omitempty"`
	IsActive   *bool   `json:"is_active,omitempty"`
	IsPremium  *bool   `json:"is_premium,omitempty"`
	Role       *string `json:"role" validate:"omitempty,oneof=employee employer"`
}

// UserResponse - DTO для ответа API (Service → API)
type UserResponse struct {
	ID         string    `json:"id"`
	TgUserName string    `json:"tg_user_name"`
	TgChatID   string    `json:"tg_chat_id"`
	IsActive   bool      `json:"is_active"`
	IsPremium  bool      `json:"is_premium"`
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
