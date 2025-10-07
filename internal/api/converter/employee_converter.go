package converter

import (
	apiModels "jobot/internal/api/models"
	serviceModels "jobot/internal/service/models"

	"github.com/google/uuid"
)

// API → Service конвертеры

// EmployeeCreateRequestToServiceEmployee конвертирует API запрос в сервисную модель Employee
func EmployeeCreateRequestToServiceEmployee(req *apiModels.EmployeeCreateRequest) (*serviceModels.Employee, error) {
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, err
	}

	return &serviceModels.Employee{
		UserID: userID,
		Tags:   req.Tags,
	}, nil
}

// EmployeeUpdateRequestToServiceEmployeeUpdateRequest конвертирует API запрос обновления в сервисную модель
func EmployeeUpdateRequestToServiceEmployeeUpdateRequest(req *apiModels.EmployeeUpdateRequest) (*serviceModels.EmployeeUpdateRequest, error) {
	updateEmployee := &serviceModels.EmployeeUpdateRequest{}

	if req.Tags != nil {
		updateEmployee.Tags = req.Tags
	}

	return updateEmployee, nil
}

// Service → API конвертеры

// ServiceEmployeeToEmployeeResponse конвертирует сервисную модель в API ответ
func ServiceEmployeeToEmployeeResponse(employee *serviceModels.Employee) *apiModels.EmployeeResponse {
	return &apiModels.EmployeeResponse{
		EmployeeID: employee.EmployeeID.String(),
		UserID:     employee.UserID.String(),
		Tags:       employee.Tags,
		CreatedAt:  employee.CreatedAt,
		UpdatedAt:  employee.UpdatedAt,
	}
}
