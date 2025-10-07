package controllers

import (
	"net/http"

	"jobot/internal/api/converter"
	"jobot/internal/api/models"
	"jobot/internal/service"
	"jobot/pkg/logger"
)

const (
	EmployeeIDPathValue = "EmployeeID"
)

type EmployeeController struct {
	employeeService service.EmployeeService
	BaseController
}

func NewEmployeeController(employeeService service.EmployeeService) *EmployeeController {
	return &EmployeeController{employeeService: employeeService}
}

func (c *EmployeeController) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("create_employee")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start create employee request")

	employee := &models.EmployeeCreateRequest{}

	err := c.ReadRequestBody(r, employee)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	serviceEmployee, err := converter.EmployeeCreateRequestToServiceEmployee(employee)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = c.employeeService.CreateEmployee(ctx, serviceEmployee)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusCreated, employee)

	log.Info("Create employee request completed")
}

func (c *EmployeeController) GetEmployee(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("get_employee")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start get employee request")

	employeeUUID, err := c.GetUUIDFromPath(r, EmployeeIDPathValue)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	employee, err := c.employeeService.GetEmployeeByUserID(ctx, employeeUUID)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, converter.ServiceEmployeeToEmployeeResponse(employee))

	log.Info("Get employee request completed")
}

func (c *EmployeeController) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("update_employee")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start update employee request")

	employeeUUID, err := c.GetUUIDFromPath(r, EmployeeIDPathValue)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	employee := &models.EmployeeUpdateRequest{}

	err = c.ReadRequestBody(r, employee)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	updateEmployee, err := converter.EmployeeUpdateRequestToServiceEmployeeUpdateRequest(employee)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = c.employeeService.UpdateEmployee(ctx, updateEmployee, employeeUUID)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, employee)

	log.Info("Update employee request completed")
}

func (c *EmployeeController) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("delete_employee")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start delete employee request")

	employeeUUID, err := c.GetUUIDFromPath(r, EmployeeIDPathValue)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = c.employeeService.DeleteEmployee(ctx, employeeUUID)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, nil)

	log.Info("Delete employee request completed")
}
