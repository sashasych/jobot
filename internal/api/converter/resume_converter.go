package converter

import (
	apiModels "jobot/internal/api/models"
	serviceModels "jobot/internal/service/models"

	"github.com/google/uuid"
)

// API → Service конвертеры

// ResumeCreateRequestToServiceResume конвертирует API запрос в сервисную модель Resume
func ResumeCreateRequestToServiceResume(req *apiModels.ResumeCreateRequest) (*serviceModels.Resume, error) {
	employeeID, err := uuid.Parse(req.EmployeeID)
	if err != nil {
		return nil, err
	}

	return &serviceModels.Resume{
		EmployeeID: employeeID,
		TgFileID:   req.TgFileID,
	}, nil
}

// ResumeUpdateRequestToServiceResumeUpdateRequest конвертирует API запрос обновления в сервисную модель
func ResumeUpdateRequestToServiceResumeUpdateRequest(req *apiModels.ResumeUpdateRequest) (*serviceModels.ResumeUpdateRequest, error) {
	updateResume := &serviceModels.ResumeUpdateRequest{}

	if req.TgFileID != nil {
		updateResume.TgFileID = req.TgFileID
	}

	return updateResume, nil
}

// Service → API конвертеры

// ServiceResumeToResumeResponse конвертирует сервисную модель в API ответ
func ServiceResumeToResumeResponse(resume *serviceModels.Resume) *apiModels.ResumeResponse {
	return &apiModels.ResumeResponse{
		ResumeID:   resume.ResumeID.String(),
		EmployeeID: resume.EmployeeID.String(),
		TgFileID:   resume.TgFileID,
		CreatedAt:  resume.CreatedAt,
		UpdatedAt:  resume.UpdatedAt,
	}
}
