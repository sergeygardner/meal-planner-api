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
	errorRecipeMeasureCreate = errors.New("recipe measure has not created by provided data")
	errorRecipeMeasureExists = errors.New("recipe measure has not created by provided data")
	errorRecipeMeasureInfo   = errors.New("recipe measure cannot be showed by provided data")
)

func RecipeMeasureCreate(userId *uuid.UUID, entityId *uuid.UUID, recipeMeasureDTO *DomainEntity.RecipeMeasure) (*DomainAggregate.RecipeMeasure, error) {
	recipeMeasureRepository := InfrastructureService.GetFactoryRepository().GetRecipeMeasureRepository()
	criteria := recipeMeasureRepository.GetCriteria().GetCriteriaByUserId(userId, nil)
	criteria = recipeMeasureRepository.GetCriteria().GetCriteriaByEntityId(entityId, criteria)
	criteria = recipeMeasureRepository.GetCriteria().GetCriteriaByUnitId(&recipeMeasureDTO.UnitId, criteria)
	recipeMeasureFindOne, errorRecipeMeasureFindOne := recipeMeasureRepository.FindOne(criteria)

	if errorRecipeMeasureFindOne == nil {
		return nil, errorRecipeMeasureCreate
	} else if recipeMeasureFindOne != nil {
		return nil, errorRecipeMeasureExists
	} else {
		recipeMeasureDTO.UserId = *userId
		recipeMeasureDTO.EntityId = *entityId
		recipeMeasureDTO.DateInsert = time.Now().UTC()
		recipeMeasureDTO.DateUpdate = time.Now().UTC()

		recipeMeasure, errorRecipeMeasuresInsertOne := recipeMeasureRepository.InsertOne(prepareRecipeMeasureRepositoryInsert(recipeMeasureDTO))

		if errorRecipeMeasuresInsertOne != nil {
			return nil, errors.Wrapf(errorRecipeMeasuresInsertOne, "an error occurred while creating a recipe measure in the database by privided data %v", recipeMeasureDTO)
		} else {
			return getMeasureAggregate(&recipeMeasure.Id, &recipeMeasure.UserId, &recipeMeasure.EntityId, nil)
		}
	}
}

func RecipeMeasuresInfo(userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) ([]*DomainAggregate.RecipeMeasure, error) {
	return ApplicationService.BuildRecipeMeasuresAggregate(nil, userId, entityId, criteria)
}

func RecipeMeasureInfo(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) (*DomainAggregate.RecipeMeasure, error) {
	return getMeasureAggregate(id, userId, entityId, criteria)
}

func RecipeMeasureUpdate(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, recipeMeasureDTO *DomainEntity.RecipeMeasure) (*DomainAggregate.RecipeMeasure, error) {
	recipeMeasureAggregate, errorRecipeMeasureAggregate := getMeasureAggregate(id, userId, entityId, nil)
	recipeMeasureRepository := InfrastructureService.GetFactoryRepository().GetRecipeMeasureRepository()

	if errorRecipeMeasureAggregate != nil {
		return nil, errors.Wrapf(errorRecipeMeasureAggregate, "an error occurred while updating a recipe measure by privided data id=%s,userId=%s,entityId=%v,criteria=%v", id, userId, nil, nil)
	}

	recipeMeasureDTO.Id = *id
	recipeMeasureDTO.UserId = *userId
	recipeMeasureDTO.EntityId = *entityId
	recipeMeasureDTO.DateInsert = recipeMeasureAggregate.Entity.DateInsert
	recipeMeasureDTO.DateUpdate = time.Now().UTC()

	recipeMeasureEntityUpdated, errorRecipeMeasureEntityUpdate := service.Update(recipeMeasureAggregate.Entity, recipeMeasureDTO)

	if errorRecipeMeasureEntityUpdate != nil {
		return nil, errors.Wrapf(errorRecipeMeasureEntityUpdate, "an error occurred while updating a recipe measure by privided data %v", recipeMeasureDTO)
	}

	restoredRecipeMeasureEntityUpdated, okRestoredRecipeMeasureEntityUpdated := recipeMeasureEntityUpdated.Interface().(*DomainEntity.RecipeMeasure)

	if !okRestoredRecipeMeasureEntityUpdated {
		return nil, errors.Wrapf(errorUpdateRestored, "an error occurred while restoring updated a recipe measure by privided data %v", restoredRecipeMeasureEntityUpdated)
	}

	updateOne, errorUpdateOne := recipeMeasureRepository.UpdateOne(
		recipeMeasureRepository.GetCriteria().GetCriteriaById(&restoredRecipeMeasureEntityUpdated.Id, nil),
		restoredRecipeMeasureEntityUpdated,
	)

	if errorUpdateOne != nil {
		return nil, errors.Wrapf(errorUpdateOne, "an error occurred while updating a recipe measure entity in the database %v", restoredRecipeMeasureEntityUpdated)
	}

	recipeMeasureAggregate.Entity = updateOne

	return recipeMeasureAggregate, nil
}

func RecipeMeasureDelete(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID) (bool, error) {
	recipeMeasureRepository := InfrastructureService.GetFactoryRepository().GetRecipeMeasureRepository()

	criteria := recipeMeasureRepository.GetCriteria().GetCriteriaById(id, nil)
	criteria = recipeMeasureRepository.GetCriteria().GetCriteriaByUserId(userId, criteria)
	criteria = recipeMeasureRepository.GetCriteria().GetCriteriaByEntityId(entityId, criteria)

	return recipeMeasureRepository.DeleteOne(criteria)
}

func getMeasureAggregate(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) (*DomainAggregate.RecipeMeasure, error) {
	recipeMeasuresAggregate, errorRecipeMeasuresAggregate := ApplicationService.BuildRecipeMeasuresAggregate(id, userId, entityId, criteria)
	if errorRecipeMeasuresAggregate != nil {
		return nil, errors.Wrapf(errorRecipeMeasuresAggregate, "an error occurred while getting a recipe measure by privided data id=%s,userId=%s,entityId=%s,criteria=%v", id, userId, entityId, criteria)
	} else if len(recipeMeasuresAggregate) == 0 {
		return nil, errorRecipeMeasureInfo
	}
	return recipeMeasuresAggregate[0], nil
}
