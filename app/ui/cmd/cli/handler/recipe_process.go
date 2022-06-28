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
	recipeProcessDTO                 *DomainEntity.RecipeProcess
	recipeProcessId                  *uuid.UUID
	statusRecipeProcessDeleteSuccess = "the recipe process has been deleted successful"
	statusRecipeProcessDeleteError   = errors.New("the recipe process has not been deleted")
)

func recipeProcessesInfo(_ string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else {
		recipeProcesses, errorRecipeProcesses := handler.RecipeProcessesInfo(&token.UserId, parentId, nil)

		if errorRecipeProcesses != nil {
			return StatusError, errorRecipeProcesses
		} else {
			if recipeProcesses == nil {
				recipeProcesses = []*DomainAggregate.RecipeProcess{}
			}

			printTable("RecipeProcessAggregate", recipeProcesses, DomainAggregate.RecipeProcess{})

			return StatusOk, nil
		}
	}
}

func recipeProcessCreate(message string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if recipeProcessDTO == nil {
		recipeProcessDTO = &DomainEntity.RecipeProcess{}
		recipeProcessDTO.EntityId = *parentId
		recipeProcessDTO.UserId = token.UserId
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
		showDialogMessage("input status for Planner Recipe. choose from (%v,%v)", kind.RecipeProcessStatusUnPublished, kind.RecipeProcessStatusPublished)
	} else if recipeProcessDTO.Status == "" {
		recipeProcessDTO.Status = kind.RecipeProcessStatus(message)

		recipeProcess, errorRecipeProcess := handler.RecipeProcessCreate(&token.UserId, parentId, recipeProcessDTO)

		recipeProcessDTO = nil

		if errorRecipeProcess != nil {
			return StatusError, errorRecipeProcess
		} else {
			printTable("RecipeProcessAggregate", []*DomainAggregate.RecipeProcess{recipeProcess}, DomainAggregate.RecipeProcess{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func recipeProcessInfo(message string) (int, error) {
	var (
		recipeProcessIdValue uuid.UUID
		errorRecipeProcessId error
	)

	if message == "RecipeProcessInfo" {
		showDialogMessage("input id for RecipeProcess")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if recipeProcessId == nil {
		recipeProcessIdValue, errorRecipeProcessId = uuid.Parse(message)

		recipeProcessId = &recipeProcessIdValue
	} else {
		errorRecipeProcessId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorRecipeProcessId != nil {
		return StatusError, errorRecipeProcessId
	} else {
		recipeProcess, errorRecipeProcess := handler.RecipeProcessInfo(recipeProcessId, &token.UserId, parentId, nil)

		recipeProcessId = nil

		if errorRecipeProcess != nil {
			return StatusError, errorRecipeProcess
		} else {
			printTable("RecipeProcessAggregate", []*DomainAggregate.RecipeProcess{recipeProcess}, DomainAggregate.RecipeProcess{})

			return StatusOk, nil
		}
	}
}

func recipeProcessUpdate(message string) (int, error) {
	var (
		recipeProcessIdValue uuid.UUID
		errorRecipeProcessId error
	)

	if message == "RecipeProcessUpdate" {
		showDialogMessage("input id for RecipeProcess")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if recipeProcessId == nil {
		recipeProcessIdValue, errorRecipeProcessId = uuid.Parse(message)

		recipeProcessId = &recipeProcessIdValue
	} else {
		errorRecipeProcessId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorRecipeProcessId != nil {
		return StatusError, errorRecipeProcessId
	} else if recipeProcessDTO == nil {
		recipeProcessDTO = &DomainEntity.RecipeProcess{}
		recipeProcessDTO.Id = *recipeProcessId
		recipeProcessDTO.EntityId = *parentId
		recipeProcessDTO.UserId = token.UserId
		recipeProcessDTO.DateUpdate = time.Now().UTC()
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
		showDialogMessage("input status for Planner Recipe. choose from (%v,%v)", kind.RecipeProcessStatusUnPublished, kind.RecipeProcessStatusPublished)
	} else if recipeProcessDTO.Status == "" {
		recipeProcessDTO.Status = kind.RecipeProcessStatus(message)

		recipeProcess, errorRecipeProcess := handler.RecipeProcessUpdate(recipeProcessId, &token.UserId, parentId, recipeProcessDTO)

		recipeProcessId = nil

		if errorRecipeProcess != nil {
			return StatusError, errorRecipeProcess
		} else {
			printTable("RecipeProcessAggregate", []*DomainAggregate.RecipeProcess{recipeProcess}, DomainAggregate.RecipeProcess{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func recipeProcessDelete(message string) (int, error) {
	var (
		recipeProcessIdValue uuid.UUID
		errorRecipeProcessId error
	)

	if message == "RecipeProcessDelete" {
		showDialogMessage("input id for RecipeProcess")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if recipeProcessId == nil {
		recipeProcessIdValue, errorRecipeProcessId = uuid.Parse(message)

		recipeProcessId = &recipeProcessIdValue
	} else {
		errorRecipeProcessId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "recipe_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorRecipeProcessId != nil {
		return StatusError, errorRecipeProcessId
	} else {
		recipeProcessDeleteStatus, errorRecipeProcessDeleteStatus := handler.RecipeProcessDelete(recipeProcessId, &token.UserId, parentId)

		recipeProcessId = nil

		if errorRecipeProcessDeleteStatus != nil {
			return StatusError, errorRecipeProcessDeleteStatus
		} else if recipeProcessDeleteStatus {
			showInfoMessage(statusRecipeProcessDeleteSuccess)

			return StatusOk, nil
		} else {
			return StatusError, statusRecipeProcessDeleteError
		}
	}
}
