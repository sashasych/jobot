package controllers

import (
	"errors"
	"net/http"

	"jobot/internal/api/converter"
	"jobot/internal/api/models"
	repo "jobot/internal/repository/user"
	"jobot/internal/service"
	"jobot/pkg/logger"
)

const (
	UserIDPathValue = "UserID"
)

type UserController struct {
	userService service.UserService
	BaseController
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("create_user")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start create user request")

	req := &models.UserCreateRequest{}

	err := c.ReadRequestBody(r, req)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	serviceUser := converter.UserCreateRequestToServiceUser(req)

	createdUser, err := c.userService.CreateUser(ctx, serviceUser)
	if err != nil {
		c.handleUserServiceError(w, err)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusCreated, createdUser)

	log.Info("Create user request completed")
}

// TODO: Конвертер
func (c *UserController) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("get_user_profile")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start get user profile request")

	userUUID, err := c.GetUUIDFromPath(r, UserIDPathValue)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	profile, err := c.userService.GetUserProfile(ctx, userUUID)
	if err != nil {
		c.handleUserServiceError(w, err)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, profile)

	log.Info("Get user profile request completed")
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("get_user")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start get user request")

	userUUID, err := c.GetUUIDFromPath(r, UserIDPathValue)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	user, err := c.userService.GetUser(ctx, userUUID)
	if err != nil {
		c.handleUserServiceError(w, err)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, converter.ServiceUserToUserResponse(user))

	log.Info("Get user request completed")
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("update_user")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start update user request")

	userUUID, err := c.GetUUIDFromPath(r, UserIDPathValue)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	user := &models.UserUpdateRequest{}

	err = c.ReadRequestBody(r, user)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	updateUser, err := converter.UserUpdateRequestToServiceUserUpdateRequest(user)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = c.userService.UpdateUser(ctx, updateUser, userUUID)
	if err != nil {
		c.handleUserServiceError(w, err)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, user)

	log.Info("Update user request completed")
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("delete_user")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start delete user request")

	userUUID, err := c.GetUUIDFromPath(r, "UserID")
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = c.userService.DeleteUser(ctx, userUUID)
	if err != nil {
		c.handleUserServiceError(w, err)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, nil)

	log.Info("Delete user request completed")
}

func (c *UserController) handleUserServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, repo.ErrUserNotFound):
		c.JSONSimpleError(w, err.Error(), http.StatusNotFound)
	case errors.Is(err, repo.ErrUserAlreadyExists):
		c.JSONSimpleError(w, err.Error(), http.StatusConflict)
	default:
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)
	}
}
