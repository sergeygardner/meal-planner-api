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
	plannerDTO                 *DomainEntity.Planner
	plannerId                  *uuid.UUID
	statusPlannerDeleteSuccess = "the recipe planner has been deleted successful"
	statusPlannerDeleteError   = errors.New("the recipe planner has not been deleted")
)

func plannersInfo(_ string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	planners, errorPlanners := handler.PlannersInfo(&token.UserId, nil)

	if errorPlanners != nil {
		return StatusError, errorPlanners
	} else {
		if planners == nil {
			planners = []*DomainAggregate.Planner{}
		}

		printTable("PlannerAggregate", planners, DomainAggregate.Planner{})

		return StatusOk, nil
	}
}

func plannerCreate(message string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if plannerDTO == nil {
		plannerDTO = &DomainEntity.Planner{}
		plannerDTO.UserId = token.UserId
		showDialogMessage("input start time for Planner")
	} else if reflect.ValueOf(plannerDTO.StartTime).IsZero() {
		parsedDate, errorParsedDate := time.Parse(time.RFC3339, message)

		if errorParsedDate != nil {
			return StatusError, errorParsedDate
		}

		plannerDTO.StartTime = parsedDate
		showDialogMessage("input end time for Planner")
	} else if reflect.ValueOf(plannerDTO.EndTime).IsZero() {
		parsedDate, errorParsedDate := time.Parse(time.RFC3339, message)

		if errorParsedDate != nil {
			return StatusError, errorParsedDate
		}

		plannerDTO.EndTime = parsedDate
		showDialogMessage("input name for Planner")
	} else if plannerDTO.Name == "" {
		plannerDTO.Name = message
		showDialogMessage("input status for Planner. choose from (%v,%v)", kind.PlannerStatusInActive, kind.PlannerStatusActive)
	} else if plannerDTO.Status == "" {
		plannerDTO.Status = kind.PlannerStatus(message)

		planner, errorPlanner := handler.PlannerCreate(&token.UserId, plannerDTO)

		plannerDTO = nil

		if errorPlanner != nil {
			return StatusError, errorPlanner
		} else {
			printTable("PlannerAggregate", []*DomainAggregate.Planner{planner}, DomainAggregate.Planner{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func plannerInfo(message string) (int, error) {
	var (
		plannerIdValue uuid.UUID
		errorPlannerId error
	)

	if message == "PlannerInfo" {
		showDialogMessage("input id for Planner")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if plannerId == nil {
		plannerIdValue, errorPlannerId = uuid.Parse(message)

		plannerId = &plannerIdValue
	} else {
		errorPlannerId = nil
	}

	if errorPlannerId != nil {
		return StatusError, errorPlannerId
	} else {
		planner, errorPlanner := handler.PlannerInfo(plannerId, &token.UserId, nil)

		plannerId = nil

		if errorPlanner != nil {
			return StatusError, errorPlanner
		} else {
			printTable("PlannerAggregate", []*DomainAggregate.Planner{planner}, DomainAggregate.Planner{})

			return StatusOk, nil
		}
	}
}

func plannerUpdate(message string) (int, error) {
	var (
		plannerIdValue uuid.UUID
		errorPlannerId error
	)

	if message == "PlannerUpdate" {
		showDialogMessage("input id for Planner")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if plannerId == nil {
		plannerIdValue, errorPlannerId = uuid.Parse(message)

		plannerId = &plannerIdValue
	} else {
		errorPlannerId = nil
	}

	if errorPlannerId != nil {
		return StatusError, errorPlannerId
	} else if plannerDTO == nil {
		plannerDTO = &DomainEntity.Planner{}
		plannerDTO.Id = *plannerId
		plannerDTO.UserId = token.UserId
		plannerDTO.DateUpdate = time.Now().UTC()
		showDialogMessage("input name for Planner")
	} else if plannerDTO.Name == "" {
		plannerDTO.Name = message
		showDialogMessage("input status for Planner. choose from (%v,%v)", kind.PlannerStatusInActive, kind.PlannerStatusActive)
	} else if plannerDTO.Status == "" {
		plannerDTO.Status = kind.PlannerStatus(message)

		planner, errorPlanner := handler.PlannerUpdate(plannerId, &token.UserId, plannerDTO)

		plannerId = nil

		if errorPlanner != nil {
			return StatusError, errorPlanner
		} else {
			printTable("PlannerAggregate", []*DomainAggregate.Planner{planner}, DomainAggregate.Planner{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func plannerDelete(message string) (int, error) {
	var (
		plannerIdValue uuid.UUID
		errorPlannerId error
	)

	if message == "PlannerDelete" {
		showDialogMessage("input id for Planner")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if plannerId == nil {
		plannerIdValue, errorPlannerId = uuid.Parse(message)

		plannerId = &plannerIdValue
	} else {
		errorPlannerId = nil
	}

	if errorPlannerId != nil {
		return StatusError, errorPlannerId
	} else {
		plannerDeleteStatus, errorPlannerDeleteStatus := handler.PlannerDelete(plannerId, &token.UserId)

		plannerId = nil

		if errorPlannerDeleteStatus != nil {
			return StatusError, errorPlannerDeleteStatus
		} else if plannerDeleteStatus {
			showInfoMessage(statusPlannerDeleteSuccess)

			return StatusOk, nil
		} else {
			return StatusError, statusPlannerDeleteError
		}
	}
}
