package handler

import (
	"errors"
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/application/handler"
	DomainAggregate "github.com/sergeygardner/meal-planner-api/domain/aggregate"
	DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	UiService "github.com/sergeygardner/meal-planner-api/ui/service"
	"strconv"
	"time"
)

var (
	pictureDTO                 *DomainEntity.Picture
	pictureId                  *uuid.UUID
	statusPictureDeleteSuccess = "the recipe picture has been deleted successful"
	statusPictureDeleteError   = errors.New("the recipe picture has not been deleted")
)

func picturesInfo(_ string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "picture_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else {
		pictures, errorPictures := handler.PicturesInfo(&token.UserId, parentId, nil)

		if errorPictures != nil {
			return StatusError, errorPictures
		} else {
			if pictures == nil {
				pictures = []*DomainAggregate.Picture{}
			}

			printTable("PictureAggregate", pictures, DomainAggregate.Picture{})

			return StatusOk, nil
		}
	}
}

func pictureCreate(message string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "picture_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if pictureDTO == nil {
		pictureDTO = &DomainEntity.Picture{}
		pictureDTO.EntityId = *parentId
		pictureDTO.UserId = token.UserId
		showDialogMessage("input name for Picture")
	} else if pictureDTO.Name == "" {
		pictureDTO.Name = message
		showDialogMessage("input url for Picture. With protocol.")
	} else if pictureDTO.URL == "" {
		pictureDTO.URL = message
		showDialogMessage("input width for Picture. Integer.")
	} else if pictureDTO.Width == 0 {
		width, errorParseInt := strconv.ParseInt(message, 10, 64)

		if errorParseInt != nil {
			return StatusError, errorParseInt
		}

		pictureDTO.Width = width
		showDialogMessage("input height for Picture. Integer.")
	} else if pictureDTO.Height == 0 {
		height, errorParseInt := strconv.ParseInt(message, 10, 64)

		if errorParseInt != nil {
			return StatusError, errorParseInt
		}

		pictureDTO.Height = height
		showDialogMessage("input size for Picture. Integer.")
	} else if pictureDTO.Size == 0 {
		size, errorParseInt := strconv.ParseInt(message, 10, 64)

		if errorParseInt != nil {
			return StatusError, errorParseInt
		}

		pictureDTO.Size = size
		showDialogMessage("input type for Picture e.g. image/png, image/jpg etc")
	} else if pictureDTO.Type == "" {
		pictureDTO.Type = message
		showDialogMessage("input status for Picture. choose from (%v,%v)", kind.PictureStatusUnPublished, kind.PictureStatusPublished)
	} else if pictureDTO.Status == "" {
		pictureDTO.Status = kind.PictureStatus(message)

		picture, errorPicture := handler.PictureCreate(&token.UserId, parentId, pictureDTO)

		pictureDTO = nil

		if errorPicture != nil {
			return StatusError, errorPicture
		} else {
			printTable("PictureAggregate", []*DomainAggregate.Picture{picture}, DomainAggregate.Picture{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func pictureInfo(message string) (int, error) {
	var (
		pictureIdValue uuid.UUID
		errorPictureId error
	)

	if message == "PictureInfo" {
		showDialogMessage("input id for Picture")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if pictureId == nil {
		pictureIdValue, errorPictureId = uuid.Parse(message)

		pictureId = &pictureIdValue
	} else {
		errorPictureId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "picture_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorPictureId != nil {
		return StatusError, errorPictureId
	} else {
		picture, errorPicture := handler.PictureInfo(pictureId, &token.UserId, parentId, nil)

		pictureId = nil

		if errorPicture != nil {
			return StatusError, errorPicture
		} else {
			printTable("PictureAggregate", []*DomainAggregate.Picture{picture}, DomainAggregate.Picture{})

			return StatusOk, nil
		}
	}
}

func pictureUpdate(message string) (int, error) {
	var (
		pictureIdValue uuid.UUID
		errorPictureId error
	)

	if message == "PictureUpdate" {
		showDialogMessage("input id for Picture")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if pictureId == nil {
		pictureIdValue, errorPictureId = uuid.Parse(message)

		pictureId = &pictureIdValue
	} else {
		errorPictureId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "picture_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorPictureId != nil {
		return StatusError, errorPictureId
	} else if pictureDTO == nil {
		pictureDTO = &DomainEntity.Picture{}
		pictureDTO.Id = *pictureId
		pictureDTO.EntityId = *parentId
		pictureDTO.UserId = token.UserId
		pictureDTO.DateUpdate = time.Now().UTC()
		showDialogMessage("input name for Picture")
	} else if pictureDTO.Name == "" {
		pictureDTO.Name = message
		showDialogMessage("input url for Picture. With protocol.")
	} else if pictureDTO.URL == "" {
		pictureDTO.URL = message
		showDialogMessage("input width for Picture. Integer.")
	} else if pictureDTO.Width == 0 {
		width, errorParseInt := strconv.ParseInt(message, 10, 64)

		if errorParseInt != nil {
			return StatusError, errorParseInt
		}

		pictureDTO.Width = width
		showDialogMessage("input height for Picture. Integer.")
	} else if pictureDTO.Height == 0 {
		height, errorParseInt := strconv.ParseInt(message, 10, 64)

		if errorParseInt != nil {
			return StatusError, errorParseInt
		}

		pictureDTO.Height = height
		showDialogMessage("input size for Picture. Integer.")
	} else if pictureDTO.Size == 0 {
		size, errorParseInt := strconv.ParseInt(message, 10, 64)

		if errorParseInt != nil {
			return StatusError, errorParseInt
		}

		pictureDTO.Size = size
		showDialogMessage("input type for Picture e.g. image/png, image/jpg etc")
	} else if pictureDTO.Type == "" {
		pictureDTO.Type = message
		showDialogMessage("input status for Picture. choose from (%v,%v)", kind.PictureStatusUnPublished, kind.PictureStatusPublished)
	} else if pictureDTO.Status == "" {
		pictureDTO.Status = kind.PictureStatus(message)

		picture, errorPicture := handler.PictureUpdate(pictureId, &token.UserId, parentId, pictureDTO)

		pictureId = nil

		if errorPicture != nil {
			return StatusError, errorPicture
		} else {
			printTable("PictureAggregate", []*DomainAggregate.Picture{picture}, DomainAggregate.Picture{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func pictureDelete(message string) (int, error) {
	var (
		pictureIdValue uuid.UUID
		errorPictureId error
	)

	if message == "PictureDelete" {
		showDialogMessage("input id for Picture")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if pictureId == nil {
		pictureIdValue, errorPictureId = uuid.Parse(message)

		pictureId = &pictureIdValue
	} else {
		errorPictureId = nil
	}

	parentId, errorParentId := UiService.GetParentId(parentIdKeys, parentIdValues, "picture_id")

	if errorParentId != nil {
		return StatusError, errorParentId
	} else if errorPictureId != nil {
		return StatusError, errorPictureId
	} else {
		pictureDeleteStatus, errorPictureDeleteStatus := handler.PictureDelete(pictureId, &token.UserId, parentId)

		pictureId = nil

		if errorPictureDeleteStatus != nil {
			return StatusError, errorPictureDeleteStatus
		} else if pictureDeleteStatus {
			showInfoMessage(statusPictureDeleteSuccess)

			return StatusOk, nil
		} else {
			return StatusError, statusPictureDeleteError
		}
	}
}
