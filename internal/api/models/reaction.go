package models

import "time"

type ReactionCreateRequest struct {
	EmployeeID string `json:"employee_id" validate:"required"`
	VacansieID string `json:"vacansie_id" validate:"required"`
	Reaction   string `json:"reaction" validate:"required,oneof=like dislike"`
}

type ReactionResponse struct {
	ReactionID string    `json:"reaction_id"`
	EmployeeID string    `json:"employee_id"`
	VacansieID string    `json:"vacansie_id"`
	Reaction   string    `json:"reaction"`
	CreatedAt  time.Time `json:"created_at"`
}

type ReactionEmployeeListResponse struct {
	ReactionsIDs []string `json:"reactions_ids"`
	EmployeeID   string   `json:"employee_id"`
}
