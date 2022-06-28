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
	errorRecipeIngredientCreate = errors.New("recipe ingredient has not created by provided data")
	errorRecipeIngredientExists = errors.New("recipe ingredient has not created by provided data")
	errorRecipeIngredientInfo   = errors.New("recipe ingredient cannot be showed by provided data")
)

func RecipeIngredientCreate(userId *uuid.UUID, entityId *uuid.UUID, recipeIngredientDTO *DomainEntity.RecipeIngredient) (*DomainAggregate.RecipeIngredient, error) {
	recipeIngredientRepository := InfrastructureService.GetFactoryRepository().GetRecipeIngredientRepository()
	criteria := recipeIngredientRepository.GetCriteria().GetCriteriaByUserId(userId, nil)
	criteria = recipeIngredientRepository.GetCriteria().GetCriteriaByEntityId(entityId, criteria)
	criteria = recipeIngredientRepository.GetCriteria().GetCriteriaByDeriveId(&recipeIngredientDTO.DeriveId, criteria)
	recipeIngredientFindOne, errorRecipeIngredientFindOne := recipeIngredientRepository.FindOne(criteria)

	if errorRecipeIngredientFindOne == nil {
		return nil, errorRecipeIngredientCreate
	} else if recipeIngredientFindOne != nil {
		return nil, errorRecipeIngredientExists
	} else {
		recipeIngredientDTO.UserId = *userId
		recipeIngredientDTO.EntityId = *entityId
		recipeIngredientDTO.DateInsert = time.Now().UTC()
		recipeIngredientDTO.DateUpdate = time.Now().UTC()

		recipeIngredient, errorRecipeIngredientsInsertOne := recipeIngredientRepository.InsertOne(prepareRecipeIngredientRepositoryInsert(recipeIngredientDTO))

		if errorRecipeIngredientsInsertOne != nil {
			return nil, errors.Wrapf(errorRecipeIngredientsInsertOne, "an error occurred while creating a recipe ingredient in the database by privided data %s", recipeIngredientDTO)
		} else {
			return getIngredientAggregate(&recipeIngredient.Id, &recipeIngredient.UserId, &recipeIngredient.EntityId, nil)
		}
	}
}

func RecipeIngredientsInfo(userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) ([]*DomainAggregate.RecipeIngredient, error) {
	return ApplicationService.BuildRecipeIngredientsAggregate(nil, userId, entityId, criteria)
}

func RecipeIngredientInfo(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) (*DomainAggregate.RecipeIngredient, error) {
	return getIngredientAggregate(id, userId, entityId, criteria)
}

func RecipeIngredientUpdate(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, recipeIngredientDTO *DomainEntity.RecipeIngredient) (*DomainAggregate.RecipeIngredient, error) {
	recipeIngredientAggregate, errorRecipeIngredientAggregate := getIngredientAggregate(id, userId, entityId, nil)
	recipeIngredientRepository := InfrastructureService.GetFactoryRepository().GetRecipeIngredientRepository()

	if errorRecipeIngredientAggregate != nil {
		return nil, errors.Wrapf(errorRecipeIngredientAggregate, "an error occurred while updating a recipe ingredient by privided data id=%s,userId=%s,entityId=%v,criteria=%v", id, userId, nil, nil)
	}

	recipeIngredientDTO.Id = *id
	recipeIngredientDTO.UserId = *userId
	recipeIngredientDTO.EntityId = *entityId
	recipeIngredientDTO.DateInsert = recipeIngredientAggregate.Entity.DateInsert
	recipeIngredientDTO.DateUpdate = time.Now().UTC()

	recipeIngredientEntityUpdated, errorRecipeIngredientEntityUpdate := service.Update(recipeIngredientAggregate.Entity, recipeIngredientDTO)

	if errorRecipeIngredientEntityUpdate != nil {
		return nil, errors.Wrapf(errorRecipeIngredientEntityUpdate, "an error occurred while updating a recipe ingredient by privided data %s", recipeIngredientDTO)
	}

	restoredRecipeIngredientEntityUpdated, okRestoredRecipeIngredientEntityUpdated := recipeIngredientEntityUpdated.Interface().(*DomainEntity.RecipeIngredient)

	if !okRestoredRecipeIngredientEntityUpdated {
		return nil, errors.Wrapf(errorUpdateRestored, "an error occurred while restoring updated a recipe ingredient by privided data %s", restoredRecipeIngredientEntityUpdated)
	}

	updateOne, errorUpdateOne := recipeIngredientRepository.UpdateOne(
		recipeIngredientRepository.GetCriteria().GetCriteriaById(&restoredRecipeIngredientEntityUpdated.Id, nil),
		restoredRecipeIngredientEntityUpdated,
	)

	if errorUpdateOne != nil {
		return nil, errors.Wrapf(errorUpdateOne, "an error occurred while updating a recipe ingredient entity in the database %s", restoredRecipeIngredientEntityUpdated)
	}

	recipeIngredientAggregate.Entity = updateOne

	return recipeIngredientAggregate, nil
}

func RecipeIngredientDelete(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID) (bool, error) {
	recipeIngredientRepository := InfrastructureService.GetFactoryRepository().GetRecipeIngredientRepository()

	criteria := recipeIngredientRepository.GetCriteria().GetCriteriaById(id, nil)
	criteria = recipeIngredientRepository.GetCriteria().GetCriteriaByUserId(userId, criteria)
	criteria = recipeIngredientRepository.GetCriteria().GetCriteriaByEntityId(entityId, criteria)

	return recipeIngredientRepository.DeleteOne(criteria)
}

func getIngredientAggregate(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) (*DomainAggregate.RecipeIngredient, error) {
	recipeIngredientsAggregate, errorRecipeIngredientsAggregate := ApplicationService.BuildRecipeIngredientsAggregate(id, userId, entityId, criteria)
	if errorRecipeIngredientsAggregate != nil {
		return nil, errors.Wrapf(errorRecipeIngredientsAggregate, "an error occurred while getting a recipe ingredient by privided data id=%s,userId=%s,entityId=%s,criteria=%v", id, userId, entityId, criteria)
	} else if len(recipeIngredientsAggregate) == 0 {
		return nil, errorRecipeIngredientInfo
	}
	return recipeIngredientsAggregate[0], nil
}
