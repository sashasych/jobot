package controllers

import (
	"net/http"

	"jobot/internal/api/converter"
	"jobot/internal/api/models"
	"jobot/internal/service"
	"jobot/pkg/logger"
)

const (
	ResumeIDPathValue = "ResumeID"
)

type ResumeController struct {
	resumeService service.ResumeService
	BaseController
}

func NewResumeController(resumeService service.ResumeService) *ResumeController {
	return &ResumeController{resumeService: resumeService}
}

func (c *ResumeController) CreateResume(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("create_resume")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start create resume request")

	req := &models.ResumeCreateRequest{}

	err := c.ReadRequestBody(r, req)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	serviceResume, err := converter.ResumeCreateRequestToServiceResume(req)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	createdResume, err := c.resumeService.CreateResume(ctx, serviceResume)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusCreated, converter.ServiceResumeToResumeResponse(createdResume))

	log.Info("Create resume request completed")
}

func (c *ResumeController) GetResume(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("get_resume")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start get resume request")

	resumeUUID, err := c.GetUUIDFromPath(r, ResumeIDPathValue)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	resume, err := c.resumeService.GetResume(ctx, resumeUUID)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, converter.ServiceResumeToResumeResponse(resume))

	log.Info("Get resume request completed")
}

func (c *ResumeController) GetResumeByEmployeeID(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("get_resume")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start get resume request")

	employeeUUID, err := c.GetUUIDFromPath(r, EmployeeIDPathValue)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	resume, err := c.resumeService.GetResumeByEmployeeID(ctx, employeeUUID)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, converter.ServiceResumeToResumeResponse(resume))

	log.Info("Get resume request completed")
}

func (c *ResumeController) UpdateResume(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("update_resume")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start update resume request")

	resumeUUID, err := c.GetUUIDFromPath(r, ResumeIDPathValue)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	resume := &models.ResumeUpdateRequest{}

	err = c.ReadRequestBody(r, resume)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	updateResume, err := converter.ResumeUpdateRequestToServiceResumeUpdateRequest(resume)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = c.resumeService.UpdateResume(ctx, updateResume, resumeUUID)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, resume)

	log.Info("Update resume request completed")
}

func (c *ResumeController) DeleteResume(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("delete_resume")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start delete resume request")

	resumeUUID, err := c.GetUUIDFromPath(r, ResumeIDPathValue)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = c.resumeService.DeleteResume(ctx, resumeUUID)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, nil)

	log.Info("Delete resume request completed")
}
