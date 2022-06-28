package handler

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	ApplicationService "github.com/sergeygardner/meal-planner-api/application/service/builder"
	DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/domain/service"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence"
	InfrastructureService "github.com/sergeygardner/meal-planner-api/infrastructure/service/repository"
	"time"
)

var (
	errorIngredientCreate = errors.New("ingredient has not created by provided data")
	errorIngredientExists = errors.New("ingredient has not created by provided data")
	errorIngredientInfo   = errors.New("ingredient cannot be showed by provided data")
)

func IngredientCreate(userId *uuid.UUID, ingredientDTO *DomainEntity.Ingredient) (*DomainEntity.Ingredient, error) {
	ingredientRepository := InfrastructureService.GetFactoryRepository().GetIngredientRepository()
	criteria := ingredientRepository.GetCriteria().GetCriteriaByName(&ingredientDTO.Name, nil)
	criteria = ingredientRepository.GetCriteria().GetCriteriaByUserId(userId, criteria)
	ingredientFindOne, errorIngredientFindOne := ingredientRepository.FindOne(criteria)

	if errorIngredientFindOne == nil {
		return nil, errorIngredientCreate
	} else if ingredientFindOne != nil {
		return nil, errorIngredientExists
	} else {
		ingredientDTO.UserId = *userId
		ingredientDTO.DateInsert = time.Now().UTC()
		ingredientDTO.DateUpdate = time.Now().UTC()

		ingredient, errorIngredientInsertOne := ingredientRepository.InsertOne(prepareIngredientRepositoryInsert(ingredientDTO))

		if errorIngredientInsertOne != nil {
			return nil, errors.Wrapf(errorIngredientInsertOne, "an error occurred while creating a ingredient in the database by privided data %s", ingredientDTO)
		} else {
			return ingredient, nil
		}
	}
}

func IngredientsInfo(userId *uuid.UUID, criteria *persistence.Criteria) ([]*DomainEntity.Ingredient, error) {
	return ApplicationService.BuildIngredientEntities(nil, userId, criteria)
}

func IngredientInfo(id *uuid.UUID, userId *uuid.UUID, criteria *persistence.Criteria) (*DomainEntity.Ingredient, error) {
	return getIngredientEntity(id, userId, criteria)
}

func IngredientUpdate(id *uuid.UUID, userId *uuid.UUID, ingredientDTO *DomainEntity.Ingredient) (*DomainEntity.Ingredient, error) {
	ingredientRepository := InfrastructureService.GetFactoryRepository().GetIngredientRepository()
	ingredient, errorIngredient := getIngredientEntity(id, userId, nil)

	if errorIngredient != nil {
		return nil, errors.Wrapf(errorIngredient, "an error occurred while updating a ingredient by privided data id=%s,userId=%s,entityId=%v,criteria=%v", id, userId, nil, nil)
	}

	ingredientDTO.Id = *id
	ingredientDTO.UserId = *userId
	ingredientDTO.DateInsert = ingredient.DateInsert
	ingredientDTO.DateUpdate = time.Now().UTC()

	ingredientUpdated, errorIngredientUpdated := service.Update(ingredient, ingredientDTO)

	if errorIngredientUpdated != nil {
		return nil, errors.Wrapf(errorIngredientUpdated, "an error occurred while updating a ingredient by privided data %s", ingredientDTO)
	}

	restoredIngredientUpdated, okRestoredIngredientUpdated := ingredientUpdated.Interface().(*DomainEntity.Ingredient)

	if !okRestoredIngredientUpdated {
		return nil, errors.Wrapf(errorUpdateRestored, "an error occurred while restoring updated a ingredient by privided data %s", ingredientUpdated)
	}

	updateOne, errorUpdateOne := ingredientRepository.UpdateOne(
		ingredientRepository.GetCriteria().GetCriteriaById(&restoredIngredientUpdated.Id, nil),
		restoredIngredientUpdated,
	)

	if errorUpdateOne != nil {
		return nil, errors.Wrapf(errorUpdateOne, "an error occurred while updating a ingredient entity in the database %s", restoredIngredientUpdated)
	}

	return updateOne, nil
}

func IngredientDelete(id *uuid.UUID, userId *uuid.UUID) (bool, error) {
	ingredientRepository := InfrastructureService.GetFactoryRepository().GetIngredientRepository()

	criteria := ingredientRepository.GetCriteria().GetCriteriaByUserId(userId, nil)
	criteria = ingredientRepository.GetCriteria().GetCriteriaById(id, criteria)

	return ingredientRepository.DeleteOne(criteria)
}

func getIngredientEntity(id *uuid.UUID, userId *uuid.UUID, criteria *persistence.Criteria) (*DomainEntity.Ingredient, error) {
	ingredientEntities, errorIngredientEntities := ApplicationService.BuildIngredientEntities(id, userId, criteria)
	if errorIngredientEntities != nil {
		return nil, errors.Wrapf(errorIngredientEntities, "an error occurred while getting a ingredient by privided data id=%s,userId=%s,criteria=%v", id, userId, criteria)
	} else if len(ingredientEntities) == 0 {
		return nil, errorIngredientInfo
	}
	return ingredientEntities[0], nil
}
