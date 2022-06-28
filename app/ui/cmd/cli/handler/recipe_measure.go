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
	recipeMeasureDTO                 *DomainEntity.RecipeMeasure
	recipeMeasureId                  *uuid.UUID
	statusRecipeMeasureDeleteSuccess = "the recipe measure has been deleted successful"
	statusRecipeMeasureDeleteError   = errors.New("the recipe measure has not been deleted")
)

func recipeMeasuresInfo(_ string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else {
		recipeMeasures, errorRecipeMeasures := handler.RecipeMeasuresInfo(&token.UserId, parentId, nil)

		if errorRecipeMeasures != nil {
			return StatusError, errorRecipeMeasures
		} else {
			if recipeMeasures == nil {
				recipeMeasures = []*DomainAggregate.RecipeMeasure{}
			}

			printTable("RecipeMeasureAggregate", recipeMeasures, DomainAggregate.RecipeMeasure{})

			return StatusOk, nil
		}
	}
}

func recipeMeasureCreate(message string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if recipeMeasureDTO == nil {
		recipeMeasureDTO = &DomainEntity.RecipeMeasure{}
		recipeMeasureDTO.EntityId = *parentId
		recipeMeasureDTO.UserId = token.UserId
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
		showDialogMessage("input status for Planner Recipe. choose from (%v,%v)", kind.RecipeMeasureStatusUnPublished, kind.RecipeMeasureStatusPublished)
	} else if recipeMeasureDTO.Status == "" {
		recipeMeasureDTO.Status = kind.RecipeMeasureStatus(message)

		recipeMeasure, errorRecipeMeasure := handler.RecipeMeasureCreate(&token.UserId, parentId, recipeMeasureDTO)

		recipeMeasureDTO = nil

		if errorRecipeMeasure != nil {
			return StatusError, errorRecipeMeasure
		} else {
			printTable("RecipeMeasureAggregate", []*DomainAggregate.RecipeMeasure{recipeMeasure}, DomainAggregate.RecipeMeasure{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func recipeMeasureInfo(message string) (int, error) {
	var (
		recipeMeasureIdValue uuid.UUID
		errorRecipeMeasureId error
	)

	if message == "RecipeMeasureInfo" {
		showDialogMessage("input id for RecipeMeasure")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if recipeMeasureId == nil {
		recipeMeasureIdValue, errorRecipeMeasureId = uuid.Parse(message)

		recipeMeasureId = &recipeMeasureIdValue
	} else {
		errorRecipeMeasureId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorRecipeMeasureId != nil {
		return StatusError, errorRecipeMeasureId
	} else {
		recipeMeasure, errorRecipeMeasure := handler.RecipeMeasureInfo(recipeMeasureId, &token.UserId, parentId, nil)

		recipeMeasureId = nil

		if errorRecipeMeasure != nil {
			return StatusError, errorRecipeMeasure
		} else {
			printTable("RecipeMeasureAggregate", []*DomainAggregate.RecipeMeasure{recipeMeasure}, DomainAggregate.RecipeMeasure{})

			return StatusOk, nil
		}
	}
}

func recipeMeasureUpdate(message string) (int, error) {
	var (
		recipeMeasureIdValue uuid.UUID
		errorRecipeMeasureId error
	)

	if message == "RecipeMeasureUpdate" {
		showDialogMessage("input id for RecipeMeasure")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if recipeMeasureId == nil {
		recipeMeasureIdValue, errorRecipeMeasureId = uuid.Parse(message)

		recipeMeasureId = &recipeMeasureIdValue
	} else {
		errorRecipeMeasureId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorRecipeMeasureId != nil {
		return StatusError, errorRecipeMeasureId
	} else if recipeMeasureDTO == nil {
		recipeMeasureDTO = &DomainEntity.RecipeMeasure{}
		recipeMeasureDTO.Id = *recipeMeasureId
		recipeMeasureDTO.EntityId = *parentId
		recipeMeasureDTO.UserId = token.UserId
		recipeMeasureDTO.DateUpdate = time.Now().UTC()
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
		showDialogMessage("input status for Planner Recipe. choose from (%v,%v)", kind.RecipeMeasureStatusUnPublished, kind.RecipeMeasureStatusPublished)
	} else if recipeMeasureDTO.Status == "" {
		recipeMeasureDTO.Status = kind.RecipeMeasureStatus(message)

		recipeMeasure, errorRecipeMeasure := handler.RecipeMeasureUpdate(recipeMeasureId, &token.UserId, parentId, recipeMeasureDTO)

		recipeMeasureId = nil

		if errorRecipeMeasure != nil {
			return StatusError, errorRecipeMeasure
		} else {
			printTable("RecipeMeasureAggregate", []*DomainAggregate.RecipeMeasure{recipeMeasure}, DomainAggregate.RecipeMeasure{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func recipeMeasureDelete(message string) (int, error) {
	var (
		recipeMeasureIdValue uuid.UUID
		errorRecipeMeasureId error
	)

	if message == "RecipeMeasureDelete" {
		showDialogMessage("input id for RecipeMeasure")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if recipeMeasureId == nil {
		recipeMeasureIdValue, errorRecipeMeasureId = uuid.Parse(message)

		recipeMeasureId = &recipeMeasureIdValue
	} else {
		errorRecipeMeasureId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorRecipeMeasureId != nil {
		return StatusError, errorRecipeMeasureId
	} else {
		recipeMeasureDeleteStatus, errorRecipeMeasureDeleteStatus := handler.RecipeMeasureDelete(recipeMeasureId, &token.UserId, parentId)

		recipeMeasureId = nil

		if errorRecipeMeasureDeleteStatus != nil {
			return StatusError, errorRecipeMeasureDeleteStatus
		} else if recipeMeasureDeleteStatus {
			showInfoMessage(statusRecipeMeasureDeleteSuccess)

			return StatusOk, nil
		} else {
			return StatusError, statusRecipeMeasureDeleteError
		}
	}
}
