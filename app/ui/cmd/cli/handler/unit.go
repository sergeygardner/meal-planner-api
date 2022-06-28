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
	unitDTO                 *DomainEntity.Unit
	unitId                  *uuid.UUID
	statusUnitDeleteSuccess = "the recipe unit has been deleted successful"
	statusUnitDeleteError   = errors.New("the recipe unit has not been deleted")
)

func unitsInfo(_ string) (int, error) {
	units, errorUnits := handler.UnitsInfo(nil)

	if errorUnits != nil {
		return StatusError, errorUnits
	} else {
		if units == nil {
			units = []*DomainEntity.Unit{}
		}

		printTable("Unit", units, DomainEntity.Unit{})

		return StatusOk, nil
	}
}

func unitCreate(message string) (int, error) {
	if unitDTO == nil {
		unitDTO = &DomainEntity.Unit{}
		showDialogMessage("input name for Unit")
	} else if unitDTO.Name == "" {
		unitDTO.Name = message
		showDialogMessage("input status for Unit. choose from (%v,%v)", kind.UnitStatusUnPublished, kind.UnitStatusPublished)
	} else if unitDTO.Status == "" {
		unitDTO.Status = kind.UnitStatus(message)

		unit, errorUnit := handler.UnitCreate(unitDTO)

		unitDTO = nil

		if errorUnit != nil {
			return StatusError, errorUnit
		} else {
			printTable("Unit", []*DomainEntity.Unit{unit}, DomainEntity.Unit{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func unitInfo(message string) (int, error) {
	var (
		unitIdValue uuid.UUID
		errorUnitId error
	)

	if message == "UnitInfo" {
		showDialogMessage("input id for Unit")

		return StatusContinue, nil
	}

	if unitId == nil {
		unitIdValue, errorUnitId = uuid.Parse(message)

		unitId = &unitIdValue
	} else {
		errorUnitId = nil
	}

	if errorUnitId != nil {
		return StatusError, errorUnitId
	} else {
		unit, errorUnit := handler.UnitInfo(unitId, nil)

		unitId = nil

		if errorUnit != nil {
			return StatusError, errorUnit
		} else {
			printTable("Unit", []*DomainEntity.Unit{unit}, DomainEntity.Unit{})

			return StatusOk, nil
		}
	}
}

func unitUpdate(message string) (int, error) {
	var (
		unitIdValue uuid.UUID
		errorUnitId error
	)

	if message == "UnitUpdate" {
		showDialogMessage("input id for Unit")

		return StatusContinue, nil
	}

	if unitId == nil {
		unitIdValue, errorUnitId = uuid.Parse(message)

		unitId = &unitIdValue
	} else {
		errorUnitId = nil
	}

	if errorUnitId != nil {
		return StatusError, errorUnitId
	} else if unitDTO == nil {
		unitDTO = &DomainEntity.Unit{}
		unitDTO.Id = *unitId
		unitDTO.DateUpdate = time.Now().UTC()
		showDialogMessage("input name for Unit")
	} else if unitDTO.Name == "" {
		unitDTO.Name = message
		showDialogMessage("input status for Unit. choose from (%v,%v)", kind.UnitStatusUnPublished, kind.UnitStatusPublished)
	} else if unitDTO.Status == "" {
		unitDTO.Status = kind.UnitStatus(message)

		unit, errorUnit := handler.UnitUpdate(unitId, unitDTO)

		unitId = nil

		if errorUnit != nil {
			return StatusError, errorUnit
		} else {
			printTable("Unit", []*DomainEntity.Unit{unit}, DomainEntity.Unit{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func unitDelete(message string) (int, error) {
	var (
		unitIdValue uuid.UUID
		errorUnitId error
	)

	if message == "UnitDelete" {
		showDialogMessage("input id for Unit")

		return StatusContinue, nil
	}

	if unitId == nil {
		unitIdValue, errorUnitId = uuid.Parse(message)

		unitId = &unitIdValue
	} else {
		errorUnitId = nil
	}

	if errorUnitId != nil {
		return StatusError, errorUnitId
	} else {
		unitDeleteStatus, errorUnitDeleteStatus := handler.UnitDelete(unitId)

		unitId = nil

		if errorUnitDeleteStatus != nil {
			return StatusError, errorUnitDeleteStatus
		} else if unitDeleteStatus {
			showInfoMessage(statusUnitDeleteSuccess)

			return StatusOk, nil
		} else {
			return StatusError, statusUnitDeleteError
		}
	}
}
