package handler

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	ApplicationServiceBuilder "github.com/sergeygardner/meal-planner-api/application/service/builder"
	DomainAggregate "github.com/sergeygardner/meal-planner-api/domain/aggregate"
	DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/domain/service"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence"
	InfrastructureService "github.com/sergeygardner/meal-planner-api/infrastructure/service/repository"
	"time"
)

var (
	errorRecipeCreate = errors.New("recipe has not created by provided data [1]")
	errorRecipeExists = errors.New("recipe has not created by provided data [2]")
	errorRecipeInfo   = errors.New("recipe cannot be showed by provided data")
	errorRecipeDelete = errors.New("recipe cannot be delete by provided data")
)

func RecipeCreate(userId *uuid.UUID, recipeDTO *DomainEntity.Recipe) (*DomainAggregate.Recipe, error) {
	recipesRepository := InfrastructureService.GetFactoryRepository().GetRecipeRepository()
	criteria := recipesRepository.GetCriteria().GetCriteriaByName(&recipeDTO.Name, nil)
	criteria = recipesRepository.GetCriteria().GetCriteriaByUserId(userId, criteria)
	recipeFindOne, errorRecipesFindOne := recipesRepository.FindOne(criteria)

	if errorRecipesFindOne == nil {
		return nil, errorRecipeCreate
	} else if recipeFindOne != nil {
		return nil, errorRecipeExists
	} else {
		recipeDTO.UserId = *userId
		recipeDTO.DateInsert = time.Now().UTC()
		recipeDTO.DateUpdate = time.Now().UTC()

		recipe, errorRecipesInsertOne := recipesRepository.InsertOne(prepareRecipeRepositoryInsert(recipeDTO))

		if errorRecipesInsertOne != nil {
			return nil, errors.Wrapf(errorRecipesInsertOne, "an error occurred while creating a recipe in the database by privided data %s", recipeDTO)
		} else {
			return RecipeInfo(&recipe.Id, &recipe.UserId, nil)
		}
	}
}

func RecipesInfo(userId *uuid.UUID, criteria *persistence.Criteria) ([]*DomainAggregate.Recipe, error) {
	return ApplicationServiceBuilder.BuildRecipesAggregate(nil, userId, criteria)
}

func RecipeInfo(id *uuid.UUID, userId *uuid.UUID, criteria *persistence.Criteria) (*DomainAggregate.Recipe, error) {
	recipesAggregate, errorRecipesAggregate := ApplicationServiceBuilder.BuildRecipesAggregate(id, userId, criteria)

	if errorRecipesAggregate != nil {
		return nil, errors.Wrapf(errorRecipesAggregate, "an error occurred while getting a recipe by privided data id=%s,userId=%s,criteria=%v", id, userId, criteria)
	} else if len(recipesAggregate) == 0 {
		return nil, errorRecipeInfo
	}

	return recipesAggregate[0], nil
}

func RecipeUpdate(id *uuid.UUID, userId *uuid.UUID, recipeDTO *DomainEntity.Recipe) (*DomainAggregate.Recipe, error) {
	recipeRepository := InfrastructureService.GetFactoryRepository().GetRecipeRepository()
	recipeAggregate, errorRecipeAggregate := RecipeInfo(id, userId, nil)

	if errorRecipeAggregate != nil {
		return nil, errors.Wrapf(errorRecipeAggregate, "an error occurred while updating a recipe by privided data id=%s,userId=%s,entityId=%v,criteria=%v", id, userId, nil, nil)
	}

	recipeDTO.Id = *id
	recipeDTO.UserId = *userId
	recipeDTO.DateInsert = recipeAggregate.Entity.DateInsert
	recipeDTO.DateUpdate = time.Now().UTC()

	recipeEntityUpdated, errorRecipeEntityUpdate := service.Update(recipeAggregate.Entity, recipeDTO)

	if errorRecipeEntityUpdate != nil {
		return nil, errors.Wrapf(errorRecipeEntityUpdate, "an error occurred while updating a recipe by privided data %s", recipeDTO)
	}

	restoredRecipeEntityUpdated, okRestoredRecipeEntityUpdated := recipeEntityUpdated.Interface().(*DomainEntity.Recipe)

	if !okRestoredRecipeEntityUpdated {
		return nil, errors.Wrapf(errorUpdateRestored, "an error occurred while restoring updated a recipe process by privided data %s", recipeEntityUpdated)
	}

	updateOne, errorUpdateOne := recipeRepository.UpdateOne(
		recipeRepository.GetCriteria().GetCriteriaById(&restoredRecipeEntityUpdated.Id, nil),
		restoredRecipeEntityUpdated,
	)

	if errorUpdateOne != nil {
		return nil, errors.Wrapf(errorUpdateOne, "an error occurred while updating a recipe entity in the database %s", restoredRecipeEntityUpdated)
	}

	recipeAggregate.Entity = updateOne

	return recipeAggregate, nil
}

func RecipeDelete(id *uuid.UUID, userId *uuid.UUID) (bool, error) {
	recipeRepository := InfrastructureService.GetFactoryRepository().GetRecipeRepository()

	criteria := recipeRepository.GetCriteria().GetCriteriaByUserId(userId, nil)
	criteria = recipeRepository.GetCriteria().GetCriteriaById(id, criteria)

	deleteOneStatus, errorDeleteOne := recipeRepository.DeleteOne(criteria)

	if errorDeleteOne != nil {
		return false, errorDeleteOne
	} else if !deleteOneStatus {
		return deleteOneStatus, errorRecipeDelete
	} else {
		return true, nil
	}
}
