package controllers

import (
	"net/http"

	"jobot/internal/api/converter"
	"jobot/internal/api/models"
	"jobot/internal/service"
	"jobot/pkg/logger"
)

const (
	VacancyIDPathValue = "VacancyID"
)

type VacancyController struct {
	vacancyService service.VacancyService
	BaseController
}

func NewVacancyController(vacancyService service.VacancyService) *VacancyController {
	return &VacancyController{vacancyService: vacancyService}
}

func (c *VacancyController) CreateVacancy(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("create_vacancy")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start create vacancy request")

	vacancy := &models.VacansieCreateRequest{}

	err := c.ReadRequestBody(r, vacancy)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	serviceVacancy, err := converter.VacancyCreateRequestToServiceVacancy(vacancy)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = c.vacancyService.CreateVacancy(ctx, serviceVacancy)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusCreated, vacancy)

	log.Info("Create vacancy request completed")
}

func (c *VacancyController) GetVacancy(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("get_vacancy")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start get vacancy request")

	vacancyUUID, err := c.GetUUIDFromPath(r, VacancyIDPathValue)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	vacancy, err := c.vacancyService.GetVacancyByID(ctx, vacancyUUID)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, converter.ServiceVacancyToVacancyResponse(vacancy))

	log.Info("Get vacancy request completed")
}

func (c *VacancyController) GetVacancyList(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("get_vacancy_list")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start get vacancy list request")

	vacancyList, err := c.vacancyService.GetVacancyList(ctx)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, converter.ServiceVacancyListToVacancyListResponse(vacancyList))

	log.Info("Get vacancy list request completed")
}

func (c *VacancyController) GetEmployerVacancies(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("get_employer_vacancies")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start get employer vacancies request")

	employerUUID, err := c.GetUUIDFromPath(r, EmployerIDPathValue)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	vacancyList, err := c.vacancyService.GetEmployerVacancies(ctx, employerUUID)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, converter.ServiceEmployerVacancyListToEmployerVacancyListResponse(vacancyList))

	log.Info("Get employer vacancies request completed")
}

func (c *VacancyController) UpdateVacancy(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("update_vacancy")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start update vacancy request")

	vacancyUUID, err := c.GetUUIDFromPath(r, VacancyIDPathValue)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	vacancy := &models.VacansieUpdateRequest{}

	err = c.ReadRequestBody(r, vacancy)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	updateVacancy, err := converter.VacancyUpdateRequestToServiceVacancyUpdateRequest(vacancy)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = c.vacancyService.UpdateVacancy(ctx, updateVacancy, vacancyUUID)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, vacancy)

	log.Info("Update vacancy request completed")
}

func (c *VacancyController) DeleteVacancy(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("delete_vacancy")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start delete vacancy request")

	vacancyUUID, err := c.GetUUIDFromPath(r, VacancyIDPathValue)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = c.vacancyService.DeleteVacancy(ctx, vacancyUUID)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, nil)

	log.Info("Delete vacancy request completed")
}
