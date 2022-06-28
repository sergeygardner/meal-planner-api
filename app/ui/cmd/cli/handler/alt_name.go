package handler

import (
	"errors"
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/application/handler"
	DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	UiService "github.com/sergeygardner/meal-planner-api/ui/service"
	"time"
)

var (
	altNameDTO                 *DomainEntity.AltName
	altNameId                  *uuid.UUID
	statusAltNameDeleteSuccess = "the recipe alt name has been deleted successful"
	statusAltNameDeleteError   = errors.New("the recipe alt name has not been deleted")
)

func altNamesInfo(_ string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "alt_name_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else {
		altNames, errorAltNames := handler.AltNamesInfo(&token.UserId, parentId, nil)

		if errorAltNames != nil {
			return StatusError, errorAltNames
		} else {
			if altNames == nil {
				altNames = []*DomainEntity.AltName{}
			}

			printTable("AltName", altNames, DomainEntity.AltName{})

			return StatusOk, nil
		}
	}
}

func altNameCreate(message string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "alt_name_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if altNameDTO == nil {
		altNameDTO = &DomainEntity.AltName{}
		altNameDTO.EntityId = *parentId
		altNameDTO.UserId = token.UserId
		showDialogMessage("input name for AltName")
	} else if altNameDTO.Name == "" {
		altNameDTO.Name = message
		showDialogMessage("input status for AltName. choose from (%v,%v)", kind.AltNameStatusUnPublished, kind.AltNameStatusPublished)
	} else if altNameDTO.Status == "" {
		altNameDTO.Status = kind.AltNameStatus(message)

		altName, errorAltName := handler.AltNameCreate(&token.UserId, parentId, altNameDTO)

		altNameDTO = nil

		if errorAltName != nil {
			return StatusError, errorAltName
		} else {
			printTable("AltName", []*DomainEntity.AltName{altName}, DomainEntity.AltName{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func altNameInfo(message string) (int, error) {
	var (
		altNameIdValue uuid.UUID
		errorAltNameId error
	)

	if message == "AltNameInfo" {
		showDialogMessage("input id for AltName")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if altNameId == nil {
		altNameIdValue, errorAltNameId = uuid.Parse(message)

		altNameId = &altNameIdValue
	} else {
		errorAltNameId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "alt_name_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorAltNameId != nil {
		return StatusError, errorAltNameId
	} else {
		altName, errorAltName := handler.AltNameInfo(altNameId, &token.UserId, parentId, nil)

		altNameId = nil

		if errorAltName != nil {
			return StatusError, errorAltName
		} else {
			printTable("AltName", []*DomainEntity.AltName{altName}, DomainEntity.AltName{})

			return StatusOk, nil
		}
	}
}

func altNameUpdate(message string) (int, error) {
	var (
		altNameIdValue uuid.UUID
		errorAltNameId error
	)

	if message == "AltNameUpdate" {
		showDialogMessage("input id for AltName")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if altNameId == nil {
		altNameIdValue, errorAltNameId = uuid.Parse(message)

		altNameId = &altNameIdValue
	} else {
		errorAltNameId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "alt_name_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorAltNameId != nil {
		return StatusError, errorAltNameId
	} else if altNameDTO == nil {
		altNameDTO = &DomainEntity.AltName{}
		altNameDTO.Id = *altNameId
		altNameDTO.EntityId = *parentId
		altNameDTO.UserId = token.UserId
		altNameDTO.DateUpdate = time.Now().UTC()
		showDialogMessage("input name for AltName")
	} else if altNameDTO.Name == "" {
		altNameDTO.Name = message
		showDialogMessage("input status for AltName. choose from (%v,%v)", kind.AltNameStatusUnPublished, kind.AltNameStatusPublished)
	} else if altNameDTO.Status == "" {
		altNameDTO.Status = kind.AltNameStatus(message)

		altName, errorAltName := handler.AltNameUpdate(altNameId, &token.UserId, parentId, altNameDTO)

		altNameId = nil

		if errorAltName != nil {
			return StatusError, errorAltName
		} else {
			printTable("AltName", []*DomainEntity.AltName{altName}, DomainEntity.AltName{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func altNameDelete(message string) (int, error) {
	var (
		altNameIdValue uuid.UUID
		errorAltNameId error
	)

	if message == "AltNameDelete" {
		showDialogMessage("input id for AltName")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if altNameId == nil {
		altNameIdValue, errorAltNameId = uuid.Parse(message)

		altNameId = &altNameIdValue
	} else {
		errorAltNameId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "alt_name_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorAltNameId != nil {
		return StatusError, errorAltNameId
	} else {
		altNameDeleteStatus, errorAltNameDeleteStatus := handler.AltNameDelete(altNameId, &token.UserId, parentId)

		altNameId = nil

		if errorAltNameDeleteStatus != nil {
			return StatusError, errorAltNameDeleteStatus
		} else if altNameDeleteStatus {
			showInfoMessage(statusAltNameDeleteSuccess)

			return StatusOk, nil
		} else {
			return StatusError, statusAltNameDeleteError
		}
	}
}
