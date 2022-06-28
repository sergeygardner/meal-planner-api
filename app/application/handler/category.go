package handler

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	ApplicationService "github.com/sergeygardner/meal-planner-api/application/service/builder"
	DomainAggregate "github.com/sergeygardner/meal-planner-api/domain/aggregate"
	DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/domain/service"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence"
	InfrastructureService "github.com/sergeygardner/meal-planner-api/infrastructure/service/repository"
	"time"
)

var (
	errorCategoryCreate = errors.New("category has not created by provided data")
	errorCategoryExists = errors.New("category has not created by provided data")
	errorCategoryInfo   = errors.New("category cannot be showed by provided data")
)

func CategoryCreate(userId *uuid.UUID, categoryDTO *DomainEntity.Category) (*DomainAggregate.Category, error) {
	categoryRepository := InfrastructureService.GetFactoryRepository().GetCategoryRepository()
	criteria := categoryRepository.GetCriteria().GetCriteriaByName(&categoryDTO.Name, nil)
	criteria = categoryRepository.GetCriteria().GetCriteriaByUserId(userId, criteria)
	categoryFindOne, errorCategoryFindOne := categoryRepository.FindOne(criteria)

	if errorCategoryFindOne == nil {
		return nil, errorCategoryCreate
	} else if categoryFindOne != nil {
		return nil, errorCategoryExists
	} else {
		categoryDTO.UserId = *userId
		categoryDTO.DateInsert = time.Now().UTC()
		categoryDTO.DateUpdate = time.Now().UTC()

		category, errorCategoryInsertOne := categoryRepository.InsertOne(prepareCategoryRepositoryInsert(categoryDTO))

		if errorCategoryInsertOne != nil {
			return nil, errors.Wrapf(errorCategoryInsertOne, "an error occurred while creating a category in the database by privided data %s", categoryDTO)
		} else {
			return getCategoryEntity(&category.Id, &category.UserId, nil)
		}
	}
}

func CategoriesInfo(userId *uuid.UUID, criteria *persistence.Criteria) ([]*DomainAggregate.Category, error) {
	return ApplicationService.BuildCategoryAggregate(nil, userId, criteria)
}

func CategoryInfo(id *uuid.UUID, userId *uuid.UUID, criteria *persistence.Criteria) (*DomainAggregate.Category, error) {
	return getCategoryEntity(id, userId, criteria)
}

func CategoryUpdate(id *uuid.UUID, userId *uuid.UUID, categoryDTO *DomainEntity.Category) (*DomainAggregate.Category, error) {
	categoryRepository := InfrastructureService.GetFactoryRepository().GetCategoryRepository()
	category, errorCategory := getCategoryEntity(id, userId, nil)

	if errorCategory != nil {
		return nil, errors.Wrapf(errorCategory, "an error occurred while updating a category by privided data id=%s,userId=%s,entityId=%v,criteria=%v", id, userId, nil, nil)
	}

	categoryDTO.Id = *id
	categoryDTO.UserId = *userId
	categoryDTO.DateInsert = category.Entity.DateInsert
	categoryDTO.DateUpdate = time.Now().UTC()

	categoryUpdated, errorCategoryUpdated := service.Update(category.Entity, categoryDTO)

	if errorCategoryUpdated != nil {
		return nil, errors.Wrapf(errorCategoryUpdated, "an error occurred while updating a category by privided data %s", categoryDTO)
	}

	restoredCategoryUpdated, okRestoredCategoryUpdated := categoryUpdated.Interface().(*DomainEntity.Category)

	if !okRestoredCategoryUpdated {
		return nil, errors.Wrapf(errorUpdateRestored, "an error occurred while restoring updated a category by privided data %s", categoryUpdated)
	}

	updateOne, errorUpdateOne := categoryRepository.UpdateOne(
		categoryRepository.GetCriteria().GetCriteriaById(&restoredCategoryUpdated.Id, nil),
		restoredCategoryUpdated,
	)

	if errorUpdateOne != nil {
		return nil, errors.Wrapf(errorUpdateOne, "an error occurred while updating a category entity in the database %s", restoredCategoryUpdated)
	}

	return getCategoryEntity(&updateOne.Id, &updateOne.UserId, nil)
}

func CategoryDelete(id *uuid.UUID, userId *uuid.UUID) (bool, error) {
	categoryRepository := InfrastructureService.GetFactoryRepository().GetCategoryRepository()

	criteria := categoryRepository.GetCriteria().GetCriteriaByUserId(userId, nil)
	criteria = categoryRepository.GetCriteria().GetCriteriaById(id, criteria)

	return categoryRepository.DeleteOne(criteria)
}

func getCategoryEntity(id *uuid.UUID, userId *uuid.UUID, criteria *persistence.Criteria) (*DomainAggregate.Category, error) {
	categoryAggregates, errorCategoryAggregates := ApplicationService.BuildCategoryAggregate(id, userId, criteria)
	if errorCategoryAggregates != nil {
		return nil, errors.Wrapf(errorCategoryAggregates, "an error occurred while getting a category by privided data id=%s,userId=%s,criteria=%v", id, userId, criteria)
	} else if len(categoryAggregates) == 0 {
		return nil, errorCategoryInfo
	}
	return categoryAggregates[0], nil
}
