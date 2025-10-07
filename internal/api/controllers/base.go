package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type BaseController struct {
}

func NewBaseController() *BaseController {
	return &BaseController{}
}

/*
func (c *BaseController) WriteSuccessResponse(w http.ResponseWriter, r *http.Request, statusCode int, message string, data interface{}) {
	response := SuccessResponse{
		Message: message,
		Data:    data,
	}
	json.NewEncoder(w).Encode(response)
}

func (c *BaseController) WriteErrorResponse(w http.ResponseWriter, r *http.Request, statusCode int, error, message string) {
	response := ErrorResponse{
		Error:   error,
		Message: message,
		Code:    statusCode,
	}
	json.NewEncoder(w).Encode(response)
}

func (c *BaseController) WriteResponse(w http.ResponseWriter, status int, v interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func (c *BaseController) WriteError(w http.ResponseWriter, status int, err error) {
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

func (c *BaseController) WriteSuccess(w http.ResponseWriter, status int, v interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
*/

// Базовая структура успешного ответа
type SuccessResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
	Meta    *Meta       `json:"meta,omitempty"`
}

// Структура для ошибок
type ErrorResponse struct {
	Error     string      `json:"error"`
	Message   string      `json:"message"`
	Code      int         `json:"code,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
	Details   interface{} `json:"details,omitempty"`
}

// Мета-информация для пагинации и т.д.
type Meta struct {
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"per_page,omitempty"`
	Total      int `json:"total,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
}

func (c *BaseController) ReadRequestBody(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

// Успешный ответ
func (c *BaseController) JSONSuccess(w http.ResponseWriter, data interface{}, message string, code int, meta *Meta) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := SuccessResponse{
		Data:    data,
		Message: message,
		Meta:    meta,
	}

	json.NewEncoder(w).Encode(response)
}

// Ответ с ошибкой
func (c *BaseController) JSONError(w http.ResponseWriter, errorMsg string, message string, code int, details interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := ErrorResponse{
		Error:     errorMsg,
		Message:   message,
		Code:      code,
		Timestamp: time.Now(),
		Details:   details,
	}

	json.NewEncoder(w).Encode(response)
}

// Простой успешный ответ
func (c *BaseController) JSONSimpleSuccess(w http.ResponseWriter, code int, data interface{}) {
	c.JSONSuccess(w, data, "", code, nil)
}

// Простая ошибка
func (c *BaseController) JSONSimpleError(w http.ResponseWriter, message string, code int) {
	c.JSONError(w, http.StatusText(code), message, code, nil)
}

func (c *BaseController) GetUUIDFromPath(r *http.Request, value string) (uuid.UUID, error) {
	id, err := uuid.Parse(r.PathValue(value))
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user id: %w", err)
	}

	return id, nil
}
