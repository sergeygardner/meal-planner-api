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
	recipeCategoryDTO                 *DomainEntity.RecipeCategory
	recipeCategoryId                  *uuid.UUID
	statusRecipeCategoryDeleteSuccess = "the recipe category has been deleted successful"
	statusRecipeCategoryDeleteError   = errors.New("the recipe category has not been deleted")
)

func recipeCategoriesInfo(_ string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else {
		recipeCategories, errorRecipeCategories := handler.RecipeCategoriesInfo(&token.UserId, parentId, nil)

		if errorRecipeCategories != nil {
			return StatusError, errorRecipeCategories
		} else {
			if recipeCategories == nil {
				recipeCategories = []*DomainAggregate.RecipeCategory{}
			}

			printTable("RecipeCategoryAggregate", recipeCategories, DomainAggregate.RecipeCategory{})

			return StatusOk, nil
		}
	}
}

func recipeCategoryCreate(message string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if recipeCategoryDTO == nil {
		recipeCategoryDTO = &DomainEntity.RecipeCategory{}
		recipeCategoryDTO.EntityId = *parentId
		recipeCategoryDTO.UserId = token.UserId
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
		showDialogMessage("input status for Planner Recipe. choose from (%v,%v)", kind.RecipeCategoryStatusUnPublished, kind.RecipeCategoryStatusPublished)
	} else if recipeCategoryDTO.Status == "" {
		recipeCategoryDTO.Status = kind.RecipeCategoryStatus(message)

		recipeCategory, errorRecipeCategory := handler.RecipeCategoryCreate(&token.UserId, parentId, recipeCategoryDTO)

		recipeCategoryDTO = nil

		if errorRecipeCategory != nil {
			return StatusError, errorRecipeCategory
		} else {
			printTable("RecipeCategoryAggregate", []*DomainAggregate.RecipeCategory{recipeCategory}, DomainAggregate.RecipeCategory{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func recipeCategoryInfo(message string) (int, error) {
	var (
		recipeCategoryIdValue uuid.UUID
		errorRecipeCategoryId error
	)

	if message == "RecipeCategoryInfo" {
		showDialogMessage("input id for RecipeCategory")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if recipeCategoryId == nil {
		recipeCategoryIdValue, errorRecipeCategoryId = uuid.Parse(message)

		recipeCategoryId = &recipeCategoryIdValue
	} else {
		errorRecipeCategoryId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorRecipeCategoryId != nil {
		return StatusError, errorRecipeCategoryId
	} else {
		recipeCategory, errorRecipeCategory := handler.RecipeCategoryInfo(recipeCategoryId, &token.UserId, parentId, nil)

		recipeCategoryId = nil

		if errorRecipeCategory != nil {
			return StatusError, errorRecipeCategory
		} else {
			printTable("RecipeCategoryAggregate", []*DomainAggregate.RecipeCategory{recipeCategory}, DomainAggregate.RecipeCategory{})

			return StatusOk, nil
		}
	}
}

func recipeCategoryUpdate(message string) (int, error) {
	var (
		recipeCategoryIdValue uuid.UUID
		errorRecipeCategoryId error
	)

	if message == "RecipeCategoryUpdate" {
		showDialogMessage("input id for RecipeCategory")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if recipeCategoryId == nil {
		recipeCategoryIdValue, errorRecipeCategoryId = uuid.Parse(message)

		recipeCategoryId = &recipeCategoryIdValue
	} else {
		errorRecipeCategoryId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorRecipeCategoryId != nil {
		return StatusError, errorRecipeCategoryId
	} else if recipeCategoryDTO == nil {
		recipeCategoryDTO = &DomainEntity.RecipeCategory{}
		recipeCategoryDTO.Id = *recipeCategoryId
		recipeCategoryDTO.EntityId = *parentId
		recipeCategoryDTO.UserId = token.UserId
		recipeCategoryDTO.DateUpdate = time.Now().UTC()
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
		showDialogMessage("input status for Planner Recipe. choose from (%v,%v)", kind.RecipeCategoryStatusUnPublished, kind.RecipeCategoryStatusPublished)
	} else if recipeCategoryDTO.Status == "" {
		recipeCategoryDTO.Status = kind.RecipeCategoryStatus(message)

		recipeCategory, errorRecipeCategory := handler.RecipeCategoryUpdate(recipeCategoryId, &token.UserId, parentId, recipeCategoryDTO)

		recipeCategoryId = nil

		if errorRecipeCategory != nil {
			return StatusError, errorRecipeCategory
		} else {
			printTable("RecipeCategoryAggregate", []*DomainAggregate.RecipeCategory{recipeCategory}, DomainAggregate.RecipeCategory{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func recipeCategoryDelete(message string) (int, error) {
	var (
		recipeCategoryIdValue uuid.UUID
		errorRecipeCategoryId error
	)

	if message == "RecipeCategoryDelete" {
		showDialogMessage("input id for RecipeCategory")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if recipeCategoryId == nil {
		recipeCategoryIdValue, errorRecipeCategoryId = uuid.Parse(message)

		recipeCategoryId = &recipeCategoryIdValue
	} else {
		errorRecipeCategoryId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorRecipeCategoryId != nil {
		return StatusError, errorRecipeCategoryId
	} else {
		recipeCategoryDeleteStatus, errorRecipeCategoryDeleteStatus := handler.RecipeCategoryDelete(recipeCategoryId, &token.UserId, parentId)

		recipeCategoryId = nil

		if errorRecipeCategoryDeleteStatus != nil {
			return StatusError, errorRecipeCategoryDeleteStatus
		} else if recipeCategoryDeleteStatus {
			showInfoMessage(statusRecipeCategoryDeleteSuccess)

			return StatusOk, nil
		} else {
			return StatusError, statusRecipeCategoryDeleteError
		}
	}
}
