package converter

import (
	apiModels "jobot/internal/api/models"
	serviceModels "jobot/internal/service/models"

	"github.com/google/uuid"
)

// API → Service конвертеры

// EmployerCreateRequestToServiceEmployer конвертирует API запрос в сервисную модель Employer
func EmployerCreateRequestToServiceEmployer(req *apiModels.EmployerCreateRequest) (*serviceModels.Employer, error) {
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, err
	}

	return &serviceModels.Employer{
		UserID:             userID,
		CompanyName:        req.CompanyName,
		CompanyDescription: req.CompanyDescription,
		CompanyWebsite:     req.CompanyWebsite,
		CompanyLocation:    req.CompanyLocation,
		CompanySize:        req.CompanySize,
	}, nil
}

// EmployerUpdateRequestToServiceEmployerUpdateRequest конвертирует API запрос обновления в сервисную модель
func EmployerUpdateRequestToServiceEmployerUpdateRequest(req *apiModels.EmployerUpdateRequest) (*serviceModels.EmployerUpdateRequest, error) {
	updateEmployer := &serviceModels.EmployerUpdateRequest{}

	if req.CompanyName != nil {
		updateEmployer.CompanyName = req.CompanyName
	}

	if req.CompanyDescription != nil {
		updateEmployer.CompanyDescription = req.CompanyDescription
	}

	if req.CompanyWebsite != nil {
		updateEmployer.CompanyWebsite = req.CompanyWebsite
	}

	if req.CompanyLocation != nil {
		updateEmployer.CompanyLocation = req.CompanyLocation
	}

	if req.CompanySize != nil {
		updateEmployer.CompanySize = req.CompanySize
	}

	return updateEmployer, nil
}

// Service → API конвертеры

// ServiceEmployerToEmployerResponse конвертирует сервисную модель в API ответ
func ServiceEmployerToEmployerResponse(employer *serviceModels.Employer) *apiModels.EmployerResponse {
	return &apiModels.EmployerResponse{
		EmployerID:         employer.EmployerID.String(),
		UserID:             employer.UserID.String(),
		CompanyName:        employer.CompanyName,
		CompanyDescription: employer.CompanyDescription,
		CompanyWebsite:     employer.CompanyWebsite,
		CompanyLocation:    employer.CompanyLocation,
		CompanySize:        employer.CompanySize,
		CreatedAt:          employer.CreatedAt,
		UpdatedAt:          employer.UpdatedAt,
	}
}
