package handler

import (
	"errors"
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/application/handler"
	DomainAggregate "github.com/sergeygardner/meal-planner-api/domain/aggregate"
	DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	UiService "github.com/sergeygardner/meal-planner-api/ui/service"
	"reflect"
	"time"
)

var (
	plannerRecipeDTO                 *DomainEntity.PlannerRecipe
	plannerRecipeId                  *uuid.UUID
	statusPlannerRecipeDeleteSuccess = "the recipe plannerRecipe has been deleted successful"
	statusPlannerRecipeDeleteError   = errors.New("the recipe plannerRecipe has not been deleted")
)

func plannerRecipesInfo(_ string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else {
		plannerRecipes, errorPlannerRecipes := handler.PlannerRecipesInfo(&token.UserId, parentId, nil)

		if errorPlannerRecipes != nil {
			return StatusError, errorPlannerRecipes
		} else {
			if plannerRecipes == nil {
				plannerRecipes = []*DomainAggregate.PlannerRecipe{}
			}

			printTable("PlannerRecipeAggregate", plannerRecipes, DomainAggregate.PlannerRecipe{})

			return StatusOk, nil
		}
	}
}

func plannerRecipeCreate(message string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if plannerRecipeDTO == nil {
		plannerRecipeDTO = &DomainEntity.PlannerRecipe{}
		plannerRecipeDTO.EntityId = *parentId
		plannerRecipeDTO.UserId = token.UserId
		showDialogMessage("input start time for Planner Recipe")
	} else if reflect.ValueOf(plannerDTO.StartTime).IsZero() {
		parsedDate, errorParsedDate := time.Parse(time.RFC3339, message)

		if errorParsedDate != nil {
			return StatusError, errorParsedDate
		}

		plannerDTO.StartTime = parsedDate
		showDialogMessage("input end time for Planner Recipe")
	} else if reflect.ValueOf(plannerDTO.EndTime).IsZero() {
		parsedDate, errorParsedDate := time.Parse(time.RFC3339, message)

		if errorParsedDate != nil {
			return StatusError, errorParsedDate
		}

		plannerDTO.EndTime = parsedDate
		showDialogMessage("input name for Planner Recipe")
	} else if plannerDTO.Name == "" {
		plannerDTO.Name = message
		showDialogMessage("input status for Planner Recipe. choose from (%v,%v)", kind.PlannerRecipeStatusInActive, kind.PlannerRecipeStatusActive)
	} else if plannerRecipeDTO.Status == "" {
		plannerRecipeDTO.Status = kind.PlannerRecipeStatus(message)

		plannerRecipe, errorPlannerRecipe := handler.PlannerRecipeCreate(&token.UserId, parentId, plannerRecipeDTO)

		plannerRecipeDTO = nil

		if errorPlannerRecipe != nil {
			return StatusError, errorPlannerRecipe
		} else {
			printTable("PlannerRecipeAggregate", []*DomainAggregate.PlannerRecipe{plannerRecipe}, DomainAggregate.PlannerRecipe{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func plannerRecipeInfo(message string) (int, error) {
	var (
		plannerRecipeIdValue uuid.UUID
		errorPlannerRecipeId error
	)

	if message == "PlannerRecipeInfo" {
		showDialogMessage("input id for PlannerRecipe")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if plannerRecipeId == nil {
		plannerRecipeIdValue, errorPlannerRecipeId = uuid.Parse(message)

		plannerRecipeId = &plannerRecipeIdValue
	} else {
		errorPlannerRecipeId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorPlannerRecipeId != nil {
		return StatusError, errorPlannerRecipeId
	} else {
		plannerRecipe, errorPlannerRecipe := handler.PlannerRecipeInfo(plannerRecipeId, &token.UserId, parentId, nil)

		plannerRecipeId = nil

		if errorPlannerRecipe != nil {
			return StatusError, errorPlannerRecipe
		} else {
			printTable("PlannerRecipeAggregate", []*DomainAggregate.PlannerRecipe{plannerRecipe}, DomainAggregate.PlannerRecipe{})

			return StatusOk, nil
		}
	}
}

func plannerRecipeUpdate(message string) (int, error) {
	var (
		plannerRecipeIdValue uuid.UUID
		errorPlannerRecipeId error
	)

	if message == "PlannerRecipeUpdate" {
		showDialogMessage("input id for PlannerRecipe")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if plannerRecipeId == nil {
		plannerRecipeIdValue, errorPlannerRecipeId = uuid.Parse(message)

		plannerRecipeId = &plannerRecipeIdValue
	} else {
		errorPlannerRecipeId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorPlannerRecipeId != nil {
		return StatusError, errorPlannerRecipeId
	} else if plannerRecipeDTO == nil {
		plannerRecipeDTO = &DomainEntity.PlannerRecipe{}
		plannerRecipeDTO.Id = *plannerRecipeId
		plannerRecipeDTO.EntityId = *parentId
		plannerRecipeDTO.UserId = token.UserId
		plannerRecipeDTO.DateUpdate = time.Now().UTC()
		showDialogMessage("input start time for Planner Recipe")
	} else if reflect.ValueOf(plannerDTO.StartTime).IsZero() {
		parsedDate, errorParsedDate := time.Parse(time.RFC3339, message)

		if errorParsedDate != nil {
			return StatusError, errorParsedDate
		}

		plannerDTO.StartTime = parsedDate
		showDialogMessage("input end time for Planner Recipe")
	} else if reflect.ValueOf(plannerDTO.EndTime).IsZero() {
		parsedDate, errorParsedDate := time.Parse(time.RFC3339, message)

		if errorParsedDate != nil {
			return StatusError, errorParsedDate
		}

		plannerDTO.EndTime = parsedDate
		showDialogMessage("input name for Planner Recipe")
	} else if plannerDTO.Name == "" {
		plannerDTO.Name = message
		showDialogMessage("input status for Planner Recipe. choose from (%v,%v)", kind.PlannerRecipeStatusInActive, kind.PlannerRecipeStatusActive)
	} else if plannerRecipeDTO.Status == "" {
		plannerRecipeDTO.Status = kind.PlannerRecipeStatus(message)

		plannerRecipe, errorPlannerRecipe := handler.PlannerRecipeUpdate(plannerRecipeId, &token.UserId, parentId, plannerRecipeDTO)

		plannerRecipeId = nil

		if errorPlannerRecipe != nil {
			return StatusError, errorPlannerRecipe
		} else {
			printTable("PlannerRecipeAggregate", []*DomainAggregate.PlannerRecipe{plannerRecipe}, DomainAggregate.PlannerRecipe{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func plannerRecipeDelete(message string) (int, error) {
	var (
		plannerRecipeIdValue uuid.UUID
		errorPlannerRecipeId error
	)

	if message == "PlannerRecipeDelete" {
		showDialogMessage("input id for PlannerRecipe")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if plannerRecipeId == nil {
		plannerRecipeIdValue, errorPlannerRecipeId = uuid.Parse(message)

		plannerRecipeId = &plannerRecipeIdValue
	} else {
		errorPlannerRecipeId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorPlannerRecipeId != nil {
		return StatusError, errorPlannerRecipeId
	} else {
		plannerRecipeDeleteStatus, errorPlannerRecipeDeleteStatus := handler.PlannerRecipeDelete(plannerRecipeId, &token.UserId, parentId)

		plannerRecipeId = nil

		if errorPlannerRecipeDeleteStatus != nil {
			return StatusError, errorPlannerRecipeDeleteStatus
		} else if plannerRecipeDeleteStatus {
			showInfoMessage(statusPlannerRecipeDeleteSuccess)

			return StatusOk, nil
		} else {
			return StatusError, statusPlannerRecipeDeleteError
		}
	}
}
