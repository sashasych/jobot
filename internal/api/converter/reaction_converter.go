package converter

import (
	apiModels "jobot/internal/api/models"
	serviceModels "jobot/internal/service/models"

	"github.com/google/uuid"
)

// API → Service конвертеры

// ReactionCreateRequestToServiceReaction конвертирует API запрос в сервисную модель Reaction
func ReactionCreateRequestToServiceReaction(req *apiModels.ReactionCreateRequest) (*serviceModels.Reaction, error) {
	employeeID, err := uuid.Parse(req.EmployeeID)
	if err != nil {
		return nil, err
	}

	vacancyID, err := uuid.Parse(req.VacansieID)
	if err != nil {
		return nil, err
	}

	return &serviceModels.Reaction{
		EmployeeID: employeeID,
		VacancyID:  vacancyID,
	}, nil
}

// Service → API конвертеры

// ServiceReactionToReactionResponse конвертирует сервисную модель в API ответ
func ServiceReactionToReactionResponse(reaction *serviceModels.Reaction) *apiModels.ReactionResponse {
	return &apiModels.ReactionResponse{
		ReactionID: reaction.ID.String(),
		EmployeeID: reaction.EmployeeID.String(),
		VacansieID: reaction.VacancyID.String(),
		CreatedAt:  reaction.CreatedAt,
	}
}

// ServiceEmployeeReactionListToReactionEmployeeListResponse конвертирует список реакций сотрудника в API ответ
func ServiceEmployeeReactionListToReactionEmployeeListResponse(reactionList *serviceModels.EmployeeReactionList) *apiModels.ReactionEmployeeListResponse {
	reactionIDs := make([]string, 0, len(reactionList.Reactions))
	for _, reaction := range reactionList.Reactions {
		reactionIDs = append(reactionIDs, reaction.ID.String())
	}

	return &apiModels.ReactionEmployeeListResponse{
		ReactionsIDs: reactionIDs,
		EmployeeID:   reactionList.EmployeeID.String(),
	}
}
