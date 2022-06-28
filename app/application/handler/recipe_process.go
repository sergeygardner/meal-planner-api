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
	errorRecipeProcessCreate = errors.New("recipe process has not created by provided data")
	errorRecipeProcessExists = errors.New("recipe process has not created by provided data")
	errorRecipeProcessInfo   = errors.New("recipe process cannot be showed by provided data")
)

func RecipeProcessCreate(userId *uuid.UUID, entityId *uuid.UUID, recipeProcessDTO *DomainEntity.RecipeProcess) (*DomainAggregate.RecipeProcess, error) {
	recipeProcessRepository := InfrastructureService.GetFactoryRepository().GetRecipeProcessRepository()
	criteria := recipeProcessRepository.GetCriteria().GetCriteriaByName(&recipeProcessDTO.Name, nil)
	criteria = recipeProcessRepository.GetCriteria().GetCriteriaByUserId(userId, criteria)
	criteria = recipeProcessRepository.GetCriteria().GetCriteriaByEntityId(entityId, criteria)
	recipeProcessFindOne, errorRecipeProcessFindOne := recipeProcessRepository.FindOne(criteria)

	if errorRecipeProcessFindOne == nil {
		return nil, errorRecipeProcessCreate
	} else if recipeProcessFindOne != nil {
		return nil, errorRecipeProcessExists
	} else {
		recipeProcessDTO.UserId = *userId
		recipeProcessDTO.EntityId = *entityId
		recipeProcessDTO.DateInsert = time.Now().UTC()
		recipeProcessDTO.DateUpdate = time.Now().UTC()

		recipeProcess, errorRecipeProcessesInsertOne := recipeProcessRepository.InsertOne(prepareRecipeProcessRepositoryInsert(recipeProcessDTO))

		if errorRecipeProcessesInsertOne != nil {
			return nil, errors.Wrapf(errorRecipeProcessesInsertOne, "an error occurred while creating a recipe process in the database by privided data %s", recipeProcessDTO)
		} else {
			return getProcessAggregate(&recipeProcess.Id, &recipeProcess.UserId, &recipeProcess.EntityId, nil)
		}
	}
}

func RecipeProcessesInfo(userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) ([]*DomainAggregate.RecipeProcess, error) {
	return ApplicationService.BuildRecipeProcessesAggregate(nil, userId, entityId, criteria)
}

func RecipeProcessInfo(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) (*DomainAggregate.RecipeProcess, error) {
	return getProcessAggregate(id, userId, entityId, criteria)
}

func RecipeProcessUpdate(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, recipeProcessDTO *DomainEntity.RecipeProcess) (*DomainAggregate.RecipeProcess, error) {
	recipeProcessRepository := InfrastructureService.GetFactoryRepository().GetRecipeProcessRepository()
	recipeProcessAggregate, errorRecipeProcessAggregate := getProcessAggregate(id, userId, entityId, nil)

	if errorRecipeProcessAggregate != nil {
		return nil, errors.Wrapf(errorRecipeProcessAggregate, "an error occurred while updating a recipe process by privided data id=%s,userId=%s,entityId=%v,criteria=%v", id, userId, nil, nil)
	}

	recipeProcessDTO.Id = *id
	recipeProcessDTO.UserId = *userId
	recipeProcessDTO.EntityId = *entityId
	recipeProcessDTO.DateInsert = recipeProcessAggregate.Entity.DateInsert
	recipeProcessDTO.DateUpdate = time.Now().UTC()

	recipeProcessEntityUpdated, errorRecipeProcessEntityUpdate := service.Update(recipeProcessAggregate.Entity, recipeProcessDTO)

	if errorRecipeProcessEntityUpdate != nil {
		return nil, errors.Wrapf(errorRecipeProcessEntityUpdate, "an error occurred while updating a recipe process by privided data %s", recipeProcessDTO)
	}

	restoredRecipeProcessEntityUpdated, okRestoredRecipeProcessEntityUpdated := recipeProcessEntityUpdated.Interface().(*DomainEntity.RecipeProcess)

	if !okRestoredRecipeProcessEntityUpdated {
		return nil, errors.Wrapf(errorUpdateRestored, "an error occurred while restoring updated a recipe process by privided data %s", recipeProcessEntityUpdated)
	}

	updateOne, errorUpdateOne := recipeProcessRepository.UpdateOne(
		recipeProcessRepository.GetCriteria().GetCriteriaById(&restoredRecipeProcessEntityUpdated.Id, nil),
		restoredRecipeProcessEntityUpdated,
	)

	if errorUpdateOne != nil {
		return nil, errors.Wrapf(errorUpdateOne, "an error occurred while updating a recipe process entity in the database %s", restoredRecipeProcessEntityUpdated)
	}

	recipeProcessAggregate.Entity = updateOne

	return recipeProcessAggregate, nil
}

func RecipeProcessDelete(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID) (bool, error) {
	recipeProcessRepository := InfrastructureService.GetFactoryRepository().GetRecipeProcessRepository()

	criteria := recipeProcessRepository.GetCriteria().GetCriteriaById(id, nil)
	criteria = recipeProcessRepository.GetCriteria().GetCriteriaByUserId(userId, criteria)
	criteria = recipeProcessRepository.GetCriteria().GetCriteriaByEntityId(entityId, criteria)

	return recipeProcessRepository.DeleteOne(criteria)
}

func getProcessAggregate(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) (*DomainAggregate.RecipeProcess, error) {
	recipeProcessesAggregate, errorRecipeProcessesAggregate := ApplicationService.BuildRecipeProcessesAggregate(id, userId, entityId, criteria)
	if errorRecipeProcessesAggregate != nil {
		return nil, errors.Wrapf(errorRecipeProcessesAggregate, "an error occurred while getting a recipe process by privided data id=%s,userId=%s,entityId=%s,criteria=%v", id, userId, entityId, criteria)
	} else if len(recipeProcessesAggregate) == 0 {
		return nil, errorRecipeProcessInfo
	}
	return recipeProcessesAggregate[0], nil
}
