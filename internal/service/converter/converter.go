package converter

import (
	repoModels "jobot/internal/repository/models"
	srvModels "jobot/internal/service/models"

	"golang.org/x/crypto/bcrypt"
)

// Service → Repository конвертеры

// ServiceUserToRepositoryUser конвертирует сервисную модель пользователя в репозиторную
func ServiceUserToRepositoryUser(user *srvModels.User) *repoModels.User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	return &repoModels.User{
		ID:           user.ID,
		Email:        user.Email,
		PasswordHash: string(hashedPassword),
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		IsActive:     user.IsActive,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

// ServiceUserRoleToRepositoryUserRole конвертирует сервисную роль в репозиторную
func ServiceUserRoleToRepositoryUserRole(role *srvModels.UserRole) *repoModels.UserRole {
	return &repoModels.UserRole{
		ID:        role.ID,
		UserID:    role.UserID,
		Role:      role.Role,
		CreatedAt: role.CreatedAt,
	}
}

// Repository → Service конвертеры

// RepositoryUserToServiceUser конвертирует репозиторную модель пользователя в сервисную
func RepositoryUserToServiceUser(user *repoModels.User) *srvModels.User {
	// uncrypt password
	password, _ := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)

	return &srvModels.User{
		ID:        user.ID,
		Email:     user.Email,
		Password:  string(password),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// RepositoryUserRoleToServiceUserRole конвертирует репозиторную роль в сервисную
func RepositoryUserRoleToServiceUserRole(role *repoModels.UserRole) *srvModels.UserRole {
	return &srvModels.UserRole{
		ID:        role.ID,
		UserID:    role.UserID,
		Role:      role.Role,
		CreatedAt: role.CreatedAt,
	}
}

// RepositoryUserRolesToServiceUserRoles конвертирует массив ролей из репозитория в сервис
func RepositoryUserRolesToServiceUserRoles(roles []repoModels.UserRole) []srvModels.UserRole {
	serviceRoles := make([]srvModels.UserRole, len(roles))
	for i, role := range roles {
		serviceRoles[i] = *RepositoryUserRoleToServiceUserRole(&role)
	}
	return serviceRoles
}

// RepositoryUserWithRolesToServiceUserWithRoles конвертирует пользователя с ролями
func RepositoryUserWithRolesToServiceUserWithRoles(user *repoModels.User, roles []repoModels.UserRole) *srvModels.UserWithRoles {
	return &srvModels.UserWithRoles{
		User:  *RepositoryUserToServiceUser(user),
		Roles: RepositoryUserRolesToServiceUserRoles(roles),
	}
}
