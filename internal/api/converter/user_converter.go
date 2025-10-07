package converter

import (
	apiModels "jobot/internal/api/models"
	serviceModels "jobot/internal/service/models"
)

// API → Service конвертеры

// UserCreateRequestToServiceUser конвертирует API запрос в сервисную модель User
func UserCreateRequestToServiceUser(req *apiModels.UserCreateRequest) *serviceModels.User {
	return &serviceModels.User{
		TgUserName: req.TgUserName,
		TgChatID:   req.TgChatID,
		IsActive:   req.IsActive,
		IsPremium:  req.IsPremium,
		Role:       req.Role,
	}
}

// UserUpdateRequestToServiceUser конвертирует API запрос обновления в сервисную модель
func UserUpdateRequestToServiceUserUpdateRequest(req *apiModels.UserUpdateRequest) (*serviceModels.UserUpdateRequest, error) {
	updateUser := &serviceModels.UserUpdateRequest{}

	if req.TgUserName != nil {
		updateUser.TgUserName = req.TgUserName
	}

	if req.IsActive != nil {
		updateUser.IsActive = req.IsActive
	}

	if req.TgChatID != nil {
		updateUser.TgChatID = req.TgChatID
	}

	if req.IsPremium != nil {
		updateUser.IsPremium = req.IsPremium
	}

	if req.Role != nil {
		updateUser.Role = req.Role
	}

	return updateUser, nil
}

// Service → API конвертеры

// ServiceUserToUserResponse конвертирует сервисную модель в API ответ
func ServiceUserToUserResponse(user *serviceModels.User) *apiModels.UserResponse {
	return &apiModels.UserResponse{
		ID:         user.ID.String(),
		TgUserName: user.TgUserName,
		TgChatID:   user.TgChatID,
		IsActive:   user.IsActive,
		IsPremium:  user.IsPremium,
		Role:       user.Role,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
}
