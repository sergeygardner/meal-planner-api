package handler

import (
	"errors"
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/application/handler"
	DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	"time"
)

var (
	ingredientDTO                 *DomainEntity.Ingredient
	ingredientId                  *uuid.UUID
	statusIngredientDeleteSuccess = "the recipe alt name has been deleted successful"
	statusIngredientDeleteError   = errors.New("the recipe alt name has not been deleted")
)

func ingredientsInfo(_ string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	ingredients, errorIngredients := handler.IngredientsInfo(&token.UserId, nil)

	if errorIngredients != nil {
		return StatusError, errorIngredients
	} else {
		if ingredients == nil {
			ingredients = []*DomainEntity.Ingredient{}
		}

		printTable("Ingredient", ingredients, DomainEntity.Ingredient{})

		return StatusOk, nil
	}
}

func ingredientCreate(message string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if ingredientDTO == nil {
		ingredientDTO = &DomainEntity.Ingredient{}
		ingredientDTO.UserId = token.UserId
		showDialogMessage("input name for Ingredient")
	} else if ingredientDTO.Name == "" {
		ingredientDTO.Name = message
		showDialogMessage("input status for Ingredient. choose from (%v,%v)", kind.IngredientStatusUnPublished, kind.IngredientStatusPublished)
	} else if ingredientDTO.Status == "" {
		ingredientDTO.Status = kind.IngredientStatus(message)

		ingredient, errorIngredient := handler.IngredientCreate(&token.UserId, ingredientDTO)

		ingredientDTO = nil

		if errorIngredient != nil {
			return StatusError, errorIngredient
		} else {
			printTable("Ingredient", []*DomainEntity.Ingredient{ingredient}, DomainEntity.Ingredient{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func ingredientInfo(message string) (int, error) {
	var (
		ingredientIdValue uuid.UUID
		errorIngredientId error
	)

	if message == "IngredientInfo" {
		showDialogMessage("input id for Ingredient")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if ingredientId == nil {
		ingredientIdValue, errorIngredientId = uuid.Parse(message)

		ingredientId = &ingredientIdValue
	} else {
		errorIngredientId = nil
	}

	if errorIngredientId != nil {
		return StatusError, errorIngredientId
	} else {
		ingredient, errorIngredient := handler.IngredientInfo(ingredientId, &token.UserId, nil)

		ingredientId = nil

		if errorIngredient != nil {
			return StatusError, errorIngredient
		} else {
			printTable("Ingredient", []*DomainEntity.Ingredient{ingredient}, DomainEntity.Ingredient{})

			return StatusOk, nil
		}
	}
}

func ingredientUpdate(message string) (int, error) {
	var (
		ingredientIdValue uuid.UUID
		errorIngredientId error
	)

	if message == "IngredientUpdate" {
		showDialogMessage("input id for Ingredient")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if ingredientId == nil {
		ingredientIdValue, errorIngredientId = uuid.Parse(message)

		ingredientId = &ingredientIdValue
	} else {
		errorIngredientId = nil
	}

	if errorIngredientId != nil {
		return StatusError, errorIngredientId
	} else if ingredientDTO == nil {
		ingredientDTO = &DomainEntity.Ingredient{}
		ingredientDTO.Id = *ingredientId
		ingredientDTO.UserId = token.UserId
		ingredientDTO.DateUpdate = time.Now().UTC()
		showDialogMessage("input name for Ingredient")
	} else if ingredientDTO.Name == "" {
		ingredientDTO.Name = message
		showDialogMessage("input status for Ingredient. choose from (%v,%v)", kind.IngredientStatusUnPublished, kind.IngredientStatusPublished)
	} else if ingredientDTO.Status == "" {
		ingredientDTO.Status = kind.IngredientStatus(message)

		ingredient, errorIngredient := handler.IngredientUpdate(ingredientId, &token.UserId, ingredientDTO)

		ingredientId = nil

		if errorIngredient != nil {
			return StatusError, errorIngredient
		} else {
			printTable("Ingredient", []*DomainEntity.Ingredient{ingredient}, DomainEntity.Ingredient{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func ingredientDelete(message string) (int, error) {
	var (
		ingredientIdValue uuid.UUID
		errorIngredientId error
	)

	if message == "IngredientDelete" {
		showDialogMessage("input id for Ingredient")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if ingredientId == nil {
		ingredientIdValue, errorIngredientId = uuid.Parse(message)

		ingredientId = &ingredientIdValue
	} else {
		errorIngredientId = nil
	}

	if errorIngredientId != nil {
		return StatusError, errorIngredientId
	} else {
		ingredientDeleteStatus, errorIngredientDeleteStatus := handler.IngredientDelete(ingredientId, &token.UserId)

		ingredientId = nil

		if errorIngredientDeleteStatus != nil {
			return StatusError, errorIngredientDeleteStatus
		} else if ingredientDeleteStatus {
			showInfoMessage(statusIngredientDeleteSuccess)

			return StatusOk, nil
		} else {
			return StatusError, statusIngredientDeleteError
		}
	}
}
