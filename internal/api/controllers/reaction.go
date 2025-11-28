package controllers

import (
	"net/http"

	"jobot/internal/api/converter"
	"jobot/internal/api/models"
	"jobot/internal/service"
	"jobot/pkg/logger"
)

const (
	ReactionIDPathValue = "ReactionID"
)

type ReactionController struct {
	reactionService service.ReactionService
	BaseController
}

func NewReactionController(reactionService service.ReactionService) *ReactionController {
	return &ReactionController{reactionService: reactionService}
}

func (c *ReactionController) CreateReaction(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("create_reaction")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start create reaction request")

	req := &models.ReactionCreateRequest{}

	err := c.ReadRequestBody(r, req)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	serviceReaction, err := converter.ReactionCreateRequestToServiceReaction(req)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	createdReaction, err := c.reactionService.CreateReaction(ctx, serviceReaction)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	response := converter.ServiceReactionToReactionResponse(createdReaction)
	response.Reaction = req.Reaction

	c.JSONSimpleSuccess(w, http.StatusCreated, response)

	log.Info("Create reaction request completed")
}

func (c *ReactionController) GetReaction(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("get_reaction")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start get reaction request")

	reactionUUID, err := c.GetUUIDFromPath(r, ReactionIDPathValue)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	reaction, err := c.reactionService.GetReaction(ctx, reactionUUID)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, converter.ServiceReactionToReactionResponse(reaction))

	log.Info("Get reaction request completed")
}

func (c *ReactionController) GetEmployeeReactions(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("get_employee_reactions")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start get employee reactions request")

	employeeUUID, err := c.GetUUIDFromPath(r, EmployeeIDPathValue)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	reactionList, err := c.reactionService.GetEmployeeReactions(ctx, employeeUUID)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, converter.ServiceEmployeeReactionListToReactionEmployeeListResponse(reactionList))

	log.Info("Get employee reactions request completed")
}

func (c *ReactionController) DeleteReaction(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context()).Named("delete_reaction")
	ctx := logger.ContextWithLogger(r.Context(), log)

	log.Info("Start delete reaction request")

	reactionUUID, err := c.GetUUIDFromPath(r, ReactionIDPathValue)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = c.reactionService.DeleteReaction(ctx, reactionUUID)
	if err != nil {
		c.JSONSimpleError(w, err.Error(), http.StatusInternalServerError)

		return
	}

	c.JSONSimpleSuccess(w, http.StatusOK, nil)

	log.Info("Delete reaction request completed")
}
