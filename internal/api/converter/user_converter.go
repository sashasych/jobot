package converter

import (
	"fmt"
	apiModels "jobot/internal/api/models"
	serviceModels "jobot/internal/service/models"

	"github.com/google/uuid"
)

// API → Service конвертеры

// UserCreateRequestToServiceUser конвертирует API запрос в сервисную модель
func UserCreateRequestToServiceUser(req *apiModels.UserCreateRequest) *serviceModels.User {
	return &serviceModels.User{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		// PasswordHash будет установлен в сервисе после хеширования
		IsActive: true, // по умолчанию активен
	}
}

// UserUpdateRequestToServiceUser конвертирует API запрос обновления в сервисную модель
func UserUpdateRequestToServiceUser(req *apiModels.UserUpdateRequest) (*serviceModels.User, error) {
	userID, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	updateUser := &serviceModels.User{
		ID: userID,
	}

	if req.Email != nil {
		updateUser.Email = *req.Email
	}

	if req.IsActive != nil {
		updateUser.IsActive = *req.IsActive
	}

	if req.FirstName != nil {
		updateUser.FirstName = *req.FirstName
	}

	if req.LastName != nil {
		updateUser.LastName = *req.LastName
	}

	return updateUser, nil
}

/*
// UserRoleRequestToServiceUserRole конвертирует API запрос роли в сервисную модель
func UserRoleRequestToServiceUserRole(req *apiModels.UserRoleRequest) *serviceModels.UserRole {
	return &serviceModels.UserRole{
		UserID: req.UserID,
		Role:   req.Role,
	}
}
*/

// Service → API конвертеры

// ServiceUserToUserResponse конвертирует сервисную модель в API ответ
func ServiceUserToUserResponse(user *serviceModels.User) *apiModels.UserResponse {
	return &apiModels.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

/*
// ServiceUserRoleToUserRoleResponse конвертирует сервисную роль в API ответ
func ServiceUserRoleToUserRoleResponse(role *serviceModels.UserRole) *apiModels.UserRoleResponse {
	return &apiModels.UserRoleResponse{
		ID:        role.ID,
		UserID:    role.UserID,
		Role:      role.Role,
		CreatedAt: role.CreatedAt,
	}
}

// ServiceUserRolesToUserRoleResponses конвертирует массив ролей
func ServiceUserRolesToUserRoleResponses(roles []serviceModels.UserRole) []apiModels.UserRoleResponse {
	responses := make([]apiModels.UserRoleResponse, len(roles))
	for i, role := range roles {
		responses[i] = *ServiceUserRoleToUserRoleResponse(&role)
	}
	return responses
}

// ServiceUserWithRolesToUserResponse конвертирует пользователя с ролями в API ответ
func ServiceUserWithRolesToUserResponse(userWithRoles *serviceModels.UserWithRoles) *apiModels.UserResponse {
	response := ServiceUserToUserResponse(&userWithRoles.User)
	// Здесь можно добавить роли в ответ, если нужно
	return response
}

// Вспомогательные функции

// GenerateUserID генерирует новый UUID для пользователя
func GenerateUserID() uuid.UUID {
	return uuid.New()
}
*/

// ValidateUserRole проверяет валидность роли
func ValidateUserRole(role string) bool {
	validRoles := []string{"employee", "employer", "admin"}
	for _, validRole := range validRoles {
		if role == validRole {
			return true
		}
	}
	return false
}
