package converter

import (
	apiModels "jobot/internal/api/models"
	serviceModels "jobot/internal/service/models"

	"github.com/google/uuid"
)

// API → Service конвертеры

// VacancyCreateRequestToServiceVacancy конвертирует API запрос в сервисную модель Vacancy
func VacancyCreateRequestToServiceVacancy(req *apiModels.VacansieCreateRequest) (*serviceModels.Vacancy, error) {
	employerID, err := uuid.Parse(req.EmployerID)
	if err != nil {
		return nil, err
	}

	return &serviceModels.Vacancy{
		EmployerID:  employerID,
		Tags:        req.Tags,
		Title:       req.Title,
		Description: req.Description,
		Location:    req.Location,
		Salary:      req.Salary,
	}, nil
}

// Service → API конвертеры

// ServiceVacancyToVacancyResponse конвертирует сервисную модель в API ответ
func ServiceVacancyToVacancyResponse(vacancy *serviceModels.Vacancy) *apiModels.VacansieResponse {
	return &apiModels.VacansieResponse{
		VacansieID:  vacancy.VacansieID.String(),
		EmployerID:  vacancy.EmployerID.String(),
		Tags:        vacancy.Tags,
		Title:       vacancy.Title,
		Description: vacancy.Description,
		Location:    vacancy.Location,
		Salary:      vacancy.Salary,
		CreatedAt:   vacancy.CreatedAt,
		UpdatedAt:   vacancy.UpdatedAt,
	}
}

// VacancyUpdateRequestToServiceVacancyUpdateRequest конвертирует API запрос обновления в сервисную модель
func VacancyUpdateRequestToServiceVacancyUpdateRequest(req *apiModels.VacansieUpdateRequest) (*serviceModels.VacancyUpdateRequest, error) {
	updateVacancy := &serviceModels.VacancyUpdateRequest{}

	if req.Tags != nil {
		updateVacancy.Tags = req.Tags
	}

	if req.Title != nil {
		updateVacancy.Title = req.Title
	}

	if req.Description != nil {
		updateVacancy.Description = req.Description
	}

	if req.Location != nil {
		updateVacancy.Location = req.Location
	}

	if req.Salary != nil {
		updateVacancy.Salary = req.Salary
	}

	return updateVacancy, nil
}

// ServiceVacancyListToVacancyListResponse конвертирует список вакансий в API ответ
func ServiceVacancyListToVacancyListResponse(vacancyList *serviceModels.VacancyList) *apiModels.VacansieListResponse {
	vacancies := make([]apiModels.VacansieResponse, 0, len(vacancyList.Vacansies))
	for _, vacancy := range vacancyList.Vacansies {
		vacancies = append(vacancies, *ServiceVacancyToVacancyResponse(&vacancy))
	}

	return &apiModels.VacansieListResponse{
		Vacansies: vacancies,
	}
}

// ServiceEmployerVacancyListToEmployerVacancyListResponse конвертирует список вакансий работодателя в API ответ
func ServiceEmployerVacancyListToEmployerVacancyListResponse(vacancyList *serviceModels.EmployerVacancyList) *apiModels.VacansieEmployerListResponse {
	vacancies := make([]apiModels.VacansieResponse, 0, len(vacancyList.Vacansies))
	for _, vacancy := range vacancyList.Vacansies {
		vacancies = append(vacancies, *ServiceVacancyToVacancyResponse(&vacancy))
	}

	return &apiModels.VacansieEmployerListResponse{
		Vacansies:  vacancies,
		EmployerID: vacancyList.EmployerID.String(),
	}
}
