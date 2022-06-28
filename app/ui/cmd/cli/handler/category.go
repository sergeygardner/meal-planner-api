package handler

import (
	"errors"
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/application/handler"
	DomainAggregate "github.com/sergeygardner/meal-planner-api/domain/aggregate"
	DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	"time"
)

var (
	categoryDTO                 *DomainEntity.Category
	categoryId                  *uuid.UUID
	statusCategoryDeleteSuccess = "the recipe category has been deleted successful"
	statusCategoryDeleteError   = errors.New("the recipe category has not been deleted")
)

func categoriesInfo(_ string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	categories, errorCategories := handler.CategoriesInfo(&token.UserId, nil)

	if errorCategories != nil {
		return StatusError, errorCategories
	} else {
		if categories == nil {
			categories = []*DomainAggregate.Category{}
		}

		printTable("CategoryAggregate", categories, DomainAggregate.Category{})

		return StatusOk, nil
	}
}

func categoryCreate(message string) (int, error) {
	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if categoryDTO == nil {
		categoryDTO = &DomainEntity.Category{}
		categoryDTO.UserId = token.UserId
		showDialogMessage("input name for Category")
	} else if categoryDTO.Name == "" {
		categoryDTO.Name = message
		showDialogMessage("input status for Category. choose from (%v,%v)", kind.CategoryStatusUnPublished, kind.CategoryStatusPublished)
	} else if categoryDTO.Status == "" {
		categoryDTO.Status = kind.CategoryStatus(message)

		category, errorCategory := handler.CategoryCreate(&token.UserId, categoryDTO)

		categoryDTO = nil

		if errorCategory != nil {
			return StatusError, errorCategory
		} else {
			printTable("CategoryAggregate", []*DomainAggregate.Category{category}, DomainAggregate.Category{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func categoryInfo(message string) (int, error) {
	var (
		categoryIdValue uuid.UUID
		errorCategoryId error
	)

	if message == "CategoryInfo" {
		showDialogMessage("input id for Category")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if categoryId == nil {
		categoryIdValue, errorCategoryId = uuid.Parse(message)

		categoryId = &categoryIdValue
	} else {
		errorCategoryId = nil
	}

	if errorCategoryId != nil {
		return StatusError, errorCategoryId
	} else {
		category, errorCategory := handler.CategoryInfo(categoryId, &token.UserId, nil)

		categoryId = nil

		if errorCategory != nil {
			return StatusError, errorCategory
		} else {
			printTable("CategoryAggregate", []*DomainAggregate.Category{category}, DomainAggregate.Category{})

			return StatusOk, nil
		}
	}
}

func categoryUpdate(message string) (int, error) {
	var (
		categoryIdValue uuid.UUID
		errorCategoryId error
	)

	if message == "CategoryUpdate" {
		showDialogMessage("input id for Category")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if categoryId == nil {
		categoryIdValue, errorCategoryId = uuid.Parse(message)

		categoryId = &categoryIdValue
	} else {
		errorCategoryId = nil
	}

	if errorCategoryId != nil {
		return StatusError, errorCategoryId
	} else if categoryDTO == nil {
		categoryDTO = &DomainEntity.Category{}
		categoryDTO.Id = *categoryId
		categoryDTO.UserId = token.UserId
		categoryDTO.DateUpdate = time.Now().UTC()
		showDialogMessage("input name for Category")
	} else if categoryDTO.Name == "" {
		categoryDTO.Name = message
		showDialogMessage("input status for Category. choose from (%v,%v)", kind.CategoryStatusUnPublished, kind.CategoryStatusPublished)
	} else if categoryDTO.Status == "" {
		categoryDTO.Status = kind.CategoryStatus(message)

		category, errorCategory := handler.CategoryUpdate(categoryId, &token.UserId, categoryDTO)

		categoryId = nil

		if errorCategory != nil {
			return StatusError, errorCategory
		} else {
			printTable("CategoryAggregate", []*DomainAggregate.Category{category}, DomainAggregate.Category{})

			return StatusOk, nil
		}
	}

	return StatusContinue, nil
}

func categoryDelete(message string) (int, error) {
	var (
		categoryIdValue uuid.UUID
		errorCategoryId error
	)

	if message == "CategoryDelete" {
		showDialogMessage("input id for Category")

		return StatusContinue, nil
	}

	token, errorExtractClaimsFromContext := tokenFromContext()

	if errorExtractClaimsFromContext != nil {
		return StatusError, errorExtractClaimsFromContext
	}

	if categoryId == nil {
		categoryIdValue, errorCategoryId = uuid.Parse(message)

		categoryId = &categoryIdValue
	} else {
		errorCategoryId = nil
	}

	if errorCategoryId != nil {
		return StatusError, errorCategoryId
	} else {
		categoryDeleteStatus, errorCategoryDeleteStatus := handler.CategoryDelete(categoryId, &token.UserId)

		categoryId = nil

		if errorCategoryDeleteStatus != nil {
			return StatusError, errorCategoryDeleteStatus
		} else if categoryDeleteStatus {
			showInfoMessage(statusCategoryDeleteSuccess)

			return StatusOk, nil
		} else {
			return StatusError, statusCategoryDeleteError
		}
	}
}
