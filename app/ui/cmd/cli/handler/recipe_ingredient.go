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
	recipeIngredientDTO                 *DomainEntity.RecipeIngredient
	recipeIngredientId                  *uuid.UUID
	statusRecipeIngredientDeleteSuccess = "the recipe ingredient has been deleted successful"
	statusRecipeIngredientDeleteError   = errors.New("the recipe ingredient has not been deleted")
)

func recipeIngredientsInfo(_ string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else {
		recipeIngredients, errorRecipeIngredients := handler.RecipeIngredientsInfo(&token.UserId, parentId, nil)

		if errorRecipeIngredients != nil {
			return StatusError, errorRecipeIngredients
		} else {
			if recipeIngredients == nil {
				recipeIngredients = []*DomainAggregate.RecipeIngredient{}
			}

			printTable("RecipeIngredientAggregate", recipeIngredients, DomainAggregate.RecipeIngredient{})

			return StatusOk, nil
		}
	}
}

func recipeIngredientCreate(message string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if recipeIngredientDTO == nil {
		recipeIngredientDTO = &DomainEntity.RecipeIngredient{}
		recipeIngredientDTO.EntityId = *parentId
		recipeIngredientDTO.UserId = token.UserId
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
		showDialogMessage("input status for Planner Recipe. choose from (%v,%v)", kind.RecipeIngredientStatusUnPublished, kind.RecipeIngredientStatusPublished)
	} else if recipeIngredientDTO.Status == "" {
		recipeIngredientDTO.Status = kind.RecipeIngredientStatus(message)

		recipeIngredient, errorRecipeIngredient := handler.RecipeIngredientCreate(&token.UserId, parentId, recipeIngredientDTO)

		recipeIngredientDTO = nil

		if errorRecipeIngredient != nil {
			return StatusError, errorRecipeIngredient
		} else {
			printTable("RecipeIngredientAggregate", []*DomainAggregate.RecipeIngredient{recipeIngredient}, DomainAggregate.RecipeIngredient{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func recipeIngredientInfo(message string) (int, error) {
	var (
		recipeIngredientIdValue uuid.UUID
		errorRecipeIngredientId error
	)

	if message == "RecipeIngredientInfo" {
		showDialogMessage("input id for RecipeIngredient")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if recipeIngredientId == nil {
		recipeIngredientIdValue, errorRecipeIngredientId = uuid.Parse(message)

		recipeIngredientId = &recipeIngredientIdValue
	} else {
		errorRecipeIngredientId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorRecipeIngredientId != nil {
		return StatusError, errorRecipeIngredientId
	} else {
		recipeIngredient, errorRecipeIngredient := handler.RecipeIngredientInfo(recipeIngredientId, &token.UserId, parentId, nil)

		recipeIngredientId = nil

		if errorRecipeIngredient != nil {
			return StatusError, errorRecipeIngredient
		} else {
			printTable("RecipeIngredientAggregate", []*DomainAggregate.RecipeIngredient{recipeIngredient}, DomainAggregate.RecipeIngredient{})

			return StatusOk, nil
		}
	}
}

func recipeIngredientUpdate(message string) (int, error) {
	var (
		recipeIngredientIdValue uuid.UUID
		errorRecipeIngredientId error
	)

	if message == "RecipeIngredientUpdate" {
		showDialogMessage("input id for RecipeIngredient")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if recipeIngredientId == nil {
		recipeIngredientIdValue, errorRecipeIngredientId = uuid.Parse(message)

		recipeIngredientId = &recipeIngredientIdValue
	} else {
		errorRecipeIngredientId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorRecipeIngredientId != nil {
		return StatusError, errorRecipeIngredientId
	} else if recipeIngredientDTO == nil {
		recipeIngredientDTO = &DomainEntity.RecipeIngredient{}
		recipeIngredientDTO.Id = *recipeIngredientId
		recipeIngredientDTO.EntityId = *parentId
		recipeIngredientDTO.UserId = token.UserId
		recipeIngredientDTO.DateUpdate = time.Now().UTC()
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
		showDialogMessage("input status for Planner Recipe. choose from (%v,%v)", kind.RecipeIngredientStatusUnPublished, kind.RecipeIngredientStatusPublished)
	} else if recipeIngredientDTO.Status == "" {
		recipeIngredientDTO.Status = kind.RecipeIngredientStatus(message)

		recipeIngredient, errorRecipeIngredient := handler.RecipeIngredientUpdate(recipeIngredientId, &token.UserId, parentId, recipeIngredientDTO)

		recipeIngredientId = nil

		if errorRecipeIngredient != nil {
			return StatusError, errorRecipeIngredient
		} else {
			printTable("RecipeIngredientAggregate", []*DomainAggregate.RecipeIngredient{recipeIngredient}, DomainAggregate.RecipeIngredient{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func recipeIngredientDelete(message string) (int, error) {
	var (
		recipeIngredientIdValue uuid.UUID
		errorRecipeIngredientId error
	)

	if message == "RecipeIngredientDelete" {
		showDialogMessage("input id for RecipeIngredient")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if recipeIngredientId == nil {
		recipeIngredientIdValue, errorRecipeIngredientId = uuid.Parse(message)

		recipeIngredientId = &recipeIngredientIdValue
	} else {
		errorRecipeIngredientId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorRecipeIngredientId != nil {
		return StatusError, errorRecipeIngredientId
	} else {
		recipeIngredientDeleteStatus, errorRecipeIngredientDeleteStatus := handler.RecipeIngredientDelete(recipeIngredientId, &token.UserId, parentId)

		recipeIngredientId = nil

		if errorRecipeIngredientDeleteStatus != nil {
			return StatusError, errorRecipeIngredientDeleteStatus
		} else if recipeIngredientDeleteStatus {
			showInfoMessage(statusRecipeIngredientDeleteSuccess)

			return StatusOk, nil
		} else {
			return StatusError, statusRecipeIngredientDeleteError
		}
	}
}
