package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"jobot/internal/api/converter"
	"jobot/internal/api/models"
	repo "jobot/internal/repository/user"
	"jobot/internal/service"
	"jobot/pkg/logger"

	"github.com/google/uuid"
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

	user := &models.UserCreateRequest{}

	err := c.ReadRequestBody(r, user)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = c.userService.CreateUser(ctx, converter.UserCreateRequestToServiceUser(user))
	if err != nil {
		c.handleUserServiceError(w, err)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusCreated, user)

	log.Info("Create user request completed")
}

/*
func (c *UserController) GetListUser(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("get_list_user")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start get list user request")

	users, err := c.userService.GetListUser(ctx)
	if err != nil {
		c.handleUserServiceError(w, err)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, users)
}
*/

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("get_user")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start get user request")

	userID, err := getUserUUID(r)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	user, err := c.userService.GetUserByID(ctx, userID)
	if err != nil {
		c.handleUserServiceError(w, err)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, user)

	log.Info("Get user request completed")
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("update_user")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start update user request")

	user := &models.UserUpdateRequest{}
	err := c.ReadRequestBody(r, user)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	updateUser, err := converter.UserUpdateRequestToServiceUser(user)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = c.userService.UpdateUser(ctx, updateUser)
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

	userID, err := getUserUUID(r)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = c.userService.DeleteUser(ctx, userID)
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

func getUserUUID(r *http.Request) (uuid.UUID, error) {
	//id, err := uuid.Parse(r.URL.Query().Get("UserID"))
	id, err := uuid.Parse(r.PathValue("UserID"))
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user id: %w", err)
	}

	return id, nil
}
