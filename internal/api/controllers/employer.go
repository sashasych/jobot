package controllers

import (
	"net/http"

	"jobot/internal/api/converter"
	"jobot/internal/api/models"
	"jobot/internal/service"
	"jobot/pkg/logger"
)

const (
	EmployerIDPathValue = "EmployerID"
)

type EmployerController struct {
	employerService service.EmployerService
	BaseController
}

func NewEmployerController(employerService service.EmployerService) *EmployerController {
	return &EmployerController{employerService: employerService}
}

func (c *EmployerController) CreateEmployer(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("create_employer")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start create employer request")

	employer := &models.EmployerCreateRequest{}

	err := c.ReadRequestBody(r, employer)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	serviceEmployer, err := converter.EmployerCreateRequestToServiceEmployer(employer)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = c.employerService.CreateEmployer(ctx, serviceEmployer)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusCreated, employer)

	log.Info("Create employer request completed")
}

func (c *EmployerController) GetEmployer(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("get_employer")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start get employer request")

	employerUUID, err := c.GetUUIDFromPath(r, EmployerIDPathValue)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	employer, err := c.employerService.GetEmployer(ctx, employerUUID)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, converter.ServiceEmployerToEmployerResponse(employer))

	log.Info("Get employer request completed")
}

func (c *EmployerController) GetEmployerByUserID(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("get_employer")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start get employer request")

	userUUID, err := c.GetUUIDFromPath(r, UserIDPathValue)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	employer, err := c.employerService.GetEmployerByUserID(ctx, userUUID)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, converter.ServiceEmployerToEmployerResponse(employer))

	log.Info("Get employer request completed")
}

func (c *EmployerController) UpdateEmployer(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("update_employer")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start update employer request")

	employerUUID, err := c.GetUUIDFromPath(r, EmployerIDPathValue)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	employer := &models.EmployerUpdateRequest{}

	err = c.ReadRequestBody(r, employer)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	updateEmployer, err := converter.EmployerUpdateRequestToServiceEmployerUpdateRequest(employer)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = c.employerService.UpdateEmployer(ctx, updateEmployer, employerUUID)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, employer)

	log.Info("Update employer request completed")
}

func (c *EmployerController) DeleteEmployer(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("delete_employer")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start delete employer request")

	employerUUID, err := c.GetUUIDFromPath(r, EmployerIDPathValue)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = c.employerService.DeleteEmployer(ctx, employerUUID)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, nil)

	log.Info("Delete employer request completed")
}

