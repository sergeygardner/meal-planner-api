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
	"reflect"
	"time"
)

var (
	errorRecipeCategoryCreate = errors.New("recipe category has not created by provided data")
	errorRecipeCategoryExists = errors.New("recipe category has not created by provided data")
	errorRecipeCategoryInfo   = errors.New("recipe category cannot be showed by provided data")
)

func RecipeCategoryCreate(userId *uuid.UUID, entityId *uuid.UUID, recipeCategoryDTO *DomainEntity.RecipeCategory) (*DomainAggregate.RecipeCategory, error) {
	recipeCategoryRepository := InfrastructureService.GetFactoryRepository().GetRecipeCategoryRepository()
	categoryRepository := InfrastructureService.GetFactoryRepository().GetCategoryRepository()
	criteria := recipeCategoryRepository.GetCriteria().GetCriteriaByUserId(userId, nil)
	criteria = recipeCategoryRepository.GetCriteria().GetCriteriaByEntityId(entityId, criteria)

	category, errorCategory := categoryRepository.FindOne(categoryRepository.GetCriteria().GetCriteriaById(&recipeCategoryDTO.DeriveId, nil))

	if errorCategory != nil {
		return nil, errorCategory
	}

	criteria = recipeCategoryRepository.GetCriteria().GetCriteriaByDeriveId(&category.Id, criteria)

	recipeCategoryFindOne, errorRecipeCategoryFindOne := recipeCategoryRepository.FindOne(criteria)

	if errorRecipeCategoryFindOne == nil {
		return nil, errorRecipeCategoryCreate
	} else if recipeCategoryFindOne != nil {
		return nil, errorRecipeCategoryExists
	} else {
		recipeCategoryDTO.UserId = *userId
		recipeCategoryDTO.EntityId = *entityId
		recipeCategoryDTO.DateInsert = time.Now().UTC()
		recipeCategoryDTO.DateUpdate = time.Now().UTC()

		recipeCategory, errorRecipeCategoriesInsertOne := recipeCategoryRepository.InsertOne(prepareRecipeCategoryRepositoryInsert(recipeCategoryDTO))

		if errorRecipeCategoriesInsertOne != nil {
			return nil, errors.Wrapf(errorRecipeCategoriesInsertOne, "an error occurred while creating a recipe category in the database by privided data %s", recipeCategoryDTO)
		} else {
			return getCategoryAggregate(&recipeCategory.Id, &recipeCategory.UserId, &recipeCategory.EntityId, nil)
		}
	}
}

func RecipeCategoriesInfo(userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) ([]*DomainAggregate.RecipeCategory, error) {
	return ApplicationService.BuildRecipeCategoriesAggregate(nil, userId, entityId, criteria)
}

func RecipeCategoryInfo(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) (*DomainAggregate.RecipeCategory, error) {
	return getCategoryAggregate(id, userId, entityId, criteria)
}

func RecipeCategoryUpdate(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, recipeCategoryDTO *DomainEntity.RecipeCategory) (*DomainAggregate.RecipeCategory, error) {
	var criteria *persistence.Criteria

	recipeCategoryRepository := InfrastructureService.GetFactoryRepository().GetRecipeCategoryRepository()
	categoryRepository := InfrastructureService.GetFactoryRepository().GetCategoryRepository()
	reflectRecipeCategoryDTO := reflect.ValueOf(*recipeCategoryDTO)
	deriveId := reflectRecipeCategoryDTO.FieldByName("DeriveId")

	if !deriveId.IsZero() {
		category, errorCategory := categoryRepository.FindOne(categoryRepository.GetCriteria().GetCriteriaById(&recipeCategoryDTO.DeriveId, nil))

		if errorCategory != nil {
			return nil, errorCategory
		}

		criteria = recipeCategoryRepository.GetCriteria().GetCriteriaByDeriveId(&category.Id, nil)
	}

	recipeCategoryAggregate, errorRecipeCategoryAggregate := getCategoryAggregate(id, userId, entityId, criteria)

	if errorRecipeCategoryAggregate != nil {
		return nil, errors.Wrapf(errorRecipeCategoryAggregate, "an error occurred while updating a recipe category by privided data id=%s,userId=%s,entityId=%v,criteria=%v", id, userId, nil, nil)
	}

	recipeCategoryDTO.Id = *id
	recipeCategoryDTO.UserId = *userId
	recipeCategoryDTO.EntityId = *entityId
	recipeCategoryDTO.DateInsert = recipeCategoryAggregate.Entity.DateInsert
	recipeCategoryDTO.DateUpdate = time.Now().UTC()

	recipeCategoryEntityUpdated, errorRecipeCategoryEntityUpdate := service.Update(recipeCategoryAggregate.Entity, recipeCategoryDTO)

	if errorRecipeCategoryEntityUpdate != nil {
		return nil, errors.Wrapf(errorRecipeCategoryEntityUpdate, "an error occurred while updating a recipe category by privided data %s", recipeCategoryDTO)
	}

	restoredRecipeCategoryEntityUpdated, okRestoredRecipeCategoryEntityUpdated := recipeCategoryEntityUpdated.Interface().(*DomainEntity.RecipeCategory)

	if !okRestoredRecipeCategoryEntityUpdated {
		return nil, errors.Wrapf(errorUpdateRestored, "an error occurred while restoring updated a recipe category by privided data %s", restoredRecipeCategoryEntityUpdated)
	}

	updateOne, errorUpdateOne := recipeCategoryRepository.UpdateOne(
		recipeCategoryRepository.GetCriteria().GetCriteriaById(&restoredRecipeCategoryEntityUpdated.Id, nil),
		restoredRecipeCategoryEntityUpdated,
	)

	if errorUpdateOne != nil {
		return nil, errors.Wrapf(errorUpdateOne, "an error occurred while updating a recipe category entity in the database %s", restoredRecipeCategoryEntityUpdated)
	}

	recipeCategoryAggregate.Entity = updateOne

	return recipeCategoryAggregate, nil
}

func RecipeCategoryDelete(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID) (bool, error) {
	recipeCategoryRepository := InfrastructureService.GetFactoryRepository().GetRecipeCategoryRepository()

	criteria := recipeCategoryRepository.GetCriteria().GetCriteriaById(id, nil)
	criteria = recipeCategoryRepository.GetCriteria().GetCriteriaByUserId(userId, criteria)
	criteria = recipeCategoryRepository.GetCriteria().GetCriteriaByEntityId(entityId, criteria)

	return recipeCategoryRepository.DeleteOne(criteria)
}

func getCategoryAggregate(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) (*DomainAggregate.RecipeCategory, error) {
	recipeCategoriesAggregate, errorRecipeCategoriesAggregate := ApplicationService.BuildRecipeCategoriesAggregate(id, userId, entityId, criteria)
	if errorRecipeCategoriesAggregate != nil {
		return nil, errors.Wrapf(errorRecipeCategoriesAggregate, "an error occurred while getting a recipe category by privided data id=%s,userId=%s,entityId=%s,criteria=%v", id, userId, entityId, criteria)
	} else if len(recipeCategoriesAggregate) == 0 {
		return nil, errorRecipeCategoryInfo
	}
	return recipeCategoriesAggregate[0], nil
}
