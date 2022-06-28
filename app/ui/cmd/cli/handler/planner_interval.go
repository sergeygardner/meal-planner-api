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
	plannerIntervalDTO                 *DomainEntity.PlannerInterval
	plannerIntervalId                  *uuid.UUID
	statusPlannerIntervalDeleteSuccess = "the recipe plannerInterval has been deleted successful"
	statusPlannerIntervalDeleteError   = errors.New("the recipe plannerInterval has not been deleted")
)

func plannerIntervalsInfo(_ string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "interval_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else {
		plannerIntervals, errorPlannerIntervals := handler.PlannerIntervalsInfo(&token.UserId, parentId, nil)

		if errorPlannerIntervals != nil {
			return StatusError, errorPlannerIntervals
		} else {
			if plannerIntervals == nil {
				plannerIntervals = []*DomainAggregate.PlannerInterval{}
			}

			printTable("PlannerIntervalAggregate", plannerIntervals, DomainAggregate.PlannerInterval{})

			return StatusOk, nil
		}
	}
}

func plannerIntervalCreate(message string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "interval_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if plannerIntervalDTO == nil {
		plannerIntervalDTO = &DomainEntity.PlannerInterval{}
		plannerIntervalDTO.EntityId = *parentId
		plannerIntervalDTO.UserId = token.UserId
		showDialogMessage("input start time for Planner Interval")
	} else if reflect.ValueOf(plannerDTO.StartTime).IsZero() {
		parsedDate, errorParsedDate := time.Parse(time.RFC3339, message)

		if errorParsedDate != nil {
			return StatusError, errorParsedDate
		}

		plannerDTO.StartTime = parsedDate
		showDialogMessage("input end time for Planner Interval")
	} else if reflect.ValueOf(plannerDTO.EndTime).IsZero() {
		parsedDate, errorParsedDate := time.Parse(time.RFC3339, message)

		if errorParsedDate != nil {
			return StatusError, errorParsedDate
		}

		plannerDTO.EndTime = parsedDate
		showDialogMessage("input name for Planner Interval")
	} else if plannerDTO.Name == "" {
		plannerDTO.Name = message
		showDialogMessage("input status for Planner Interval. choose from (%v,%v)", kind.PlannerIntervalStatusInActive, kind.PlannerIntervalStatusActive)
	} else if plannerIntervalDTO.Status == "" {
		plannerIntervalDTO.Status = kind.PlannerIntervalStatus(message)

		plannerInterval, errorPlannerInterval := handler.PlannerIntervalCreate(&token.UserId, parentId, plannerIntervalDTO)

		plannerIntervalDTO = nil

		if errorPlannerInterval != nil {
			return StatusError, errorPlannerInterval
		} else {
			printTable("PlannerIntervalAggregate", []*DomainAggregate.PlannerInterval{plannerInterval}, DomainAggregate.PlannerInterval{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func plannerIntervalInfo(message string) (int, error) {
	var (
		plannerIntervalIdValue uuid.UUID
		errorPlannerIntervalId error
	)

	if message == "PlannerIntervalInfo" {
		showDialogMessage("input id for PlannerInterval")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if plannerIntervalId == nil {
		plannerIntervalIdValue, errorPlannerIntervalId = uuid.Parse(message)

		plannerIntervalId = &plannerIntervalIdValue
	} else {
		errorPlannerIntervalId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "interval_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorPlannerIntervalId != nil {
		return StatusError, errorPlannerIntervalId
	} else {
		plannerInterval, errorPlannerInterval := handler.PlannerIntervalInfo(plannerIntervalId, &token.UserId, parentId, nil)

		plannerIntervalId = nil

		if errorPlannerInterval != nil {
			return StatusError, errorPlannerInterval
		} else {
			printTable("PlannerIntervalAggregate", []*DomainAggregate.PlannerInterval{plannerInterval}, DomainAggregate.PlannerInterval{})

			return StatusOk, nil
		}
	}
}

func plannerIntervalUpdate(message string) (int, error) {
	var (
		plannerIntervalIdValue uuid.UUID
		errorPlannerIntervalId error
	)

	if message == "PlannerIntervalUpdate" {
		showDialogMessage("input id for PlannerInterval")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if plannerIntervalId == nil {
		plannerIntervalIdValue, errorPlannerIntervalId = uuid.Parse(message)

		plannerIntervalId = &plannerIntervalIdValue
	} else {
		errorPlannerIntervalId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "interval_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorPlannerIntervalId != nil {
		return StatusError, errorPlannerIntervalId
	} else if plannerIntervalDTO == nil {
		plannerIntervalDTO = &DomainEntity.PlannerInterval{}
		plannerIntervalDTO.Id = *plannerIntervalId
		plannerIntervalDTO.EntityId = *parentId
		plannerIntervalDTO.UserId = token.UserId
		plannerIntervalDTO.DateUpdate = time.Now().UTC()
		showDialogMessage("input start time for Planner Interval")
	} else if reflect.ValueOf(plannerDTO.StartTime).IsZero() {
		parsedDate, errorParsedDate := time.Parse(time.RFC3339, message)

		if errorParsedDate != nil {
			return StatusError, errorParsedDate
		}

		plannerDTO.StartTime = parsedDate
		showDialogMessage("input end time for Planner Interval")
	} else if reflect.ValueOf(plannerDTO.EndTime).IsZero() {
		parsedDate, errorParsedDate := time.Parse(time.RFC3339, message)

		if errorParsedDate != nil {
			return StatusError, errorParsedDate
		}

		plannerDTO.EndTime = parsedDate
		showDialogMessage("input name for Planner Interval")
	} else if plannerDTO.Name == "" {
		plannerDTO.Name = message
		showDialogMessage("input status for Planner Interval. choose from (%v,%v)", kind.PlannerIntervalStatusInActive, kind.PlannerIntervalStatusActive)
	} else if plannerIntervalDTO.Status == "" {
		plannerIntervalDTO.Status = kind.PlannerIntervalStatus(message)

		plannerInterval, errorPlannerInterval := handler.PlannerIntervalUpdate(plannerIntervalId, &token.UserId, parentId, plannerIntervalDTO)

		plannerIntervalId = nil

		if errorPlannerInterval != nil {
			return StatusError, errorPlannerInterval
		} else {
			printTable("PlannerIntervalAggregate", []*DomainAggregate.PlannerInterval{plannerInterval}, DomainAggregate.PlannerInterval{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func plannerIntervalDelete(message string) (int, error) {
	var (
		plannerIntervalIdValue uuid.UUID
		errorPlannerIntervalId error
	)

	if message == "PlannerIntervalDelete" {
		showDialogMessage("input id for PlannerInterval")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if plannerIntervalId == nil {
		plannerIntervalIdValue, errorPlannerIntervalId = uuid.Parse(message)

		plannerIntervalId = &plannerIntervalIdValue
	} else {
		errorPlannerIntervalId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "interval_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorPlannerIntervalId != nil {
		return StatusError, errorPlannerIntervalId
	} else {
		plannerIntervalDeleteStatus, errorPlannerIntervalDeleteStatus := handler.PlannerIntervalDelete(plannerIntervalId, &token.UserId, parentId)

		plannerIntervalId = nil

		if errorPlannerIntervalDeleteStatus != nil {
			return StatusError, errorPlannerIntervalDeleteStatus
		} else if plannerIntervalDeleteStatus {
			showInfoMessage(statusPlannerIntervalDeleteSuccess)

			return StatusOk, nil
		} else {
			return StatusError, statusPlannerIntervalDeleteError
		}
	}
}
