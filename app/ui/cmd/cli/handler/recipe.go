package handler

import (
	"errors"
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/application/handler"
	DomainAggregate "github.com/sergeygardner/meal-planner-api/domain/aggregate"
	DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	"reflect"
	"time"
)

var (
	recipeDTO                 *DomainEntity.Recipe
	recipeId                  *uuid.UUID
	statusRecipeDeleteSuccess = "the recipe has been deleted successful"
	statusRecipeDeleteError   = errors.New("the recipe has not been deleted")
)

func recipesInfo(_ string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	recipes, errorRecipes := handler.RecipesInfo(&token.UserId, nil)

	if errorRecipes != nil {
		return StatusError, errorRecipes
	} else {
		if recipes == nil {
			recipes = []*DomainAggregate.Recipe{}
		}

		printTable("RecipeAggregate", recipes, DomainAggregate.Recipe{})

		return StatusOk, nil
	}
}

func recipeCreate(message string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if recipeDTO == nil {
		recipeDTO = &DomainEntity.Recipe{}
		recipeDTO.UserId = token.UserId
		showDialogMessage("input start time for recipe")
	} else if reflect.ValueOf(plannerDTO.StartTime).IsZero() {
		parsedDate, errorParsedDate := time.Parse(time.RFC3339, message)

		if errorParsedDate != nil {
			return StatusError, errorParsedDate
		}

		plannerDTO.StartTime = parsedDate
		showDialogMessage("input end time for recipe")
	} else if reflect.ValueOf(plannerDTO.EndTime).IsZero() {
		parsedDate, errorParsedDate := time.Parse(time.RFC3339, message)

		if errorParsedDate != nil {
			return StatusError, errorParsedDate
		}

		plannerDTO.EndTime = parsedDate
		showDialogMessage("input name for recipe")
	} else if plannerDTO.Name == "" {
		plannerDTO.Name = message
		showDialogMessage("input status for recipe. choose from (%v,%v)", kind.RecipeStatusUnPublished, kind.RecipeStatusPublished)
	} else if recipeDTO.Status == "" {
		recipeDTO.Status = kind.RecipeStatus(message)

		recipe, errorRecipe := handler.RecipeCreate(&token.UserId, recipeDTO)

		recipeDTO = nil

		if errorRecipe != nil {
			return StatusError, errorRecipe
		} else {
			printTable("RecipeAggregate", []*DomainAggregate.Recipe{recipe}, DomainAggregate.Recipe{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func recipeInfo(message string) (int, error) {
	var (
		recipeIdValue uuid.UUID
		errorRecipeId error
	)

	if message == "RecipeInfo" {
		showDialogMessage("input id for Recipe")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if recipeId == nil {
		recipeIdValue, errorRecipeId = uuid.Parse(message)

		recipeId = &recipeIdValue
	} else {
		errorRecipeId = nil
	}

	if errorRecipeId != nil {
		return StatusError, errorRecipeId
	} else {
		recipe, errorRecipe := handler.RecipeInfo(recipeId, &token.UserId, nil)

		recipeId = nil

		if errorRecipe != nil {
			return StatusError, errorRecipe
		} else {
			printTable("RecipeAggregate", []*DomainAggregate.Recipe{recipe}, DomainAggregate.Recipe{})

			return StatusOk, nil
		}
	}
}

func recipeUpdate(message string) (int, error) {
	var (
		recipeIdValue uuid.UUID
		errorRecipeId error
	)

	if message == "RecipeUpdate" {
		showDialogMessage("input id for Recipe")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if recipeId == nil {
		recipeIdValue, errorRecipeId = uuid.Parse(message)

		recipeId = &recipeIdValue
	} else {
		errorRecipeId = nil
	}

	if errorRecipeId != nil {
		return StatusError, errorRecipeId
	} else if recipeDTO == nil {
		recipeDTO = &DomainEntity.Recipe{}
		recipeDTO.Id = *recipeId
		recipeDTO.UserId = token.UserId
		recipeDTO.DateUpdate = time.Now().UTC()
		showDialogMessage("input start time for recipe")
	} else if reflect.ValueOf(plannerDTO.StartTime).IsZero() {
		parsedDate, errorParsedDate := time.Parse(time.RFC3339, message)

		if errorParsedDate != nil {
			return StatusError, errorParsedDate
		}

		plannerDTO.StartTime = parsedDate
		showDialogMessage("input end time for recipe")
	} else if reflect.ValueOf(plannerDTO.EndTime).IsZero() {
		parsedDate, errorParsedDate := time.Parse(time.RFC3339, message)

		if errorParsedDate != nil {
			return StatusError, errorParsedDate
		}

		plannerDTO.EndTime = parsedDate
		showDialogMessage("input name for recipe")
	} else if plannerDTO.Name == "" {
		plannerDTO.Name = message
		showDialogMessage("input status for recipe. choose from (%v,%v)", kind.RecipeStatusUnPublished, kind.RecipeStatusPublished)
	} else if recipeDTO.Status == "" {
		recipeDTO.Status = kind.RecipeStatus(message)

		recipe, errorRecipe := handler.RecipeUpdate(recipeId, &token.UserId, recipeDTO)

		recipeId = nil

		if errorRecipe != nil {
			return StatusError, errorRecipe
		} else {
			printTable("RecipeAggregate", []*DomainAggregate.Recipe{recipe}, DomainAggregate.Recipe{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func recipeDelete(message string) (int, error) {
	var (
		recipeIdValue uuid.UUID
		errorRecipeId error
	)

	if message == "RecipeDelete" {
		showDialogMessage("input id for Recipe")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if recipeId == nil {
		recipeIdValue, errorRecipeId = uuid.Parse(message)

		recipeId = &recipeIdValue
	} else {
		errorRecipeId = nil
	}

	if errorRecipeId != nil {
		return StatusError, errorRecipeId
	} else {
		recipeDeleteStatus, errorRecipeDeleteStatus := handler.RecipeDelete(recipeId, &token.UserId)

		recipeId = nil

		if errorRecipeDeleteStatus != nil {
			return StatusError, errorRecipeDeleteStatus
		} else if recipeDeleteStatus {
			showInfoMessage(statusRecipeDeleteSuccess)

			return StatusOk, nil
		} else {
			return StatusError, statusRecipeDeleteError
		}
	}
}
