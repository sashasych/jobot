package api

import (
	"net/http"
)

// Main controller
type Controller struct {
	UserController
	EmployeeController
	ResumeController
	EmployerController
	VacancyController
	ReactionController
}

// Controller interfaces

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type EmployeeController interface {
	CreateEmployee(w http.ResponseWriter, r *http.Request)
	GetEmployee(w http.ResponseWriter, r *http.Request)
	UpdateEmployee(w http.ResponseWriter, r *http.Request)
	DeleteEmployee(w http.ResponseWriter, r *http.Request)
}

type ResumeController interface {
	CreateResume(w http.ResponseWriter, r *http.Request)
	GetResume(w http.ResponseWriter, r *http.Request)
	UpdateResume(w http.ResponseWriter, r *http.Request)
	DeleteResume(w http.ResponseWriter, r *http.Request)
}

type EmployerController interface {
	CreateEmployer(w http.ResponseWriter, r *http.Request)
	GetEmployer(w http.ResponseWriter, r *http.Request)
	UpdateEmployer(w http.ResponseWriter, r *http.Request)
	DeleteEmployer(w http.ResponseWriter, r *http.Request)
}

type VacancyController interface {
	CreateVacancy(w http.ResponseWriter, r *http.Request)
	GetVacancy(w http.ResponseWriter, r *http.Request)
	GetVacancyList(w http.ResponseWriter, r *http.Request)
	GetEmployerVacancies(w http.ResponseWriter, r *http.Request)
	UpdateVacancy(w http.ResponseWriter, r *http.Request)
	DeleteVacancy(w http.ResponseWriter, r *http.Request)
}

type ReactionController interface {
	CreateReaction(w http.ResponseWriter, r *http.Request)
	GetReaction(w http.ResponseWriter, r *http.Request)
	GetEmployeeReactions(w http.ResponseWriter, r *http.Request)
	DeleteReaction(w http.ResponseWriter, r *http.Request)
}
