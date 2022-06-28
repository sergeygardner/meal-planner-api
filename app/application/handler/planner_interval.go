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
	errorPlannerIntervalCreate = errors.New("planner interval has not created by provided data")
	errorPlannerIntervalExists = errors.New("planner interval has not created by provided data")
	errorPlannerIntervalInfo   = errors.New("planner interval cannot be showed by provided data")
)

func PlannerIntervalCreate(userId *uuid.UUID, entityId *uuid.UUID, plannerIntervalDTO *DomainEntity.PlannerInterval) (*DomainAggregate.PlannerInterval, error) {
	plannerIntervalRepository := InfrastructureService.GetFactoryRepository().GetPlannerIntervalRepository()
	criteria := plannerIntervalRepository.GetCriteria().GetCriteriaByName(&plannerIntervalDTO.Name, nil)
	criteria = plannerIntervalRepository.GetCriteria().GetCriteriaByUserId(userId, criteria)
	criteria = plannerIntervalRepository.GetCriteria().GetCriteriaByEntityId(entityId, criteria)
	plannerIntervalFindOne, errorPlannerIntervalFindOne := plannerIntervalRepository.FindOne(criteria)

	if errorPlannerIntervalFindOne == nil {
		return nil, errorPlannerIntervalCreate
	} else if plannerIntervalFindOne != nil {
		return nil, errorPlannerIntervalExists
	} else {
		plannerIntervalDTO.UserId = *userId
		plannerIntervalDTO.EntityId = *entityId
		plannerIntervalDTO.DateInsert = time.Now().UTC()
		plannerIntervalDTO.DateUpdate = time.Now().UTC()

		plannerInterval, errorPlannerIntervalInsertOne := plannerIntervalRepository.InsertOne(preparePlannerIntervalRepositoryInsert(plannerIntervalDTO))

		if errorPlannerIntervalInsertOne != nil {
			return nil, errors.Wrapf(errorPlannerIntervalInsertOne, "an error occurred while creating a planner interval in the database by privided data %v", plannerIntervalDTO)
		} else {
			return getPlannerIntervalAggregate(&plannerInterval.Id, &plannerInterval.UserId, &plannerInterval.EntityId, nil)
		}
	}
}

func PlannerIntervalsInfo(userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) ([]*DomainAggregate.PlannerInterval, error) {
	return ApplicationService.BuildPlannerIntervalAggregates(nil, userId, entityId, criteria)
}

func PlannerIntervalInfo(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) (*DomainAggregate.PlannerInterval, error) {
	return getPlannerIntervalAggregate(id, userId, entityId, criteria)
}

func PlannerIntervalUpdate(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, plannerIntervalDTO *DomainEntity.PlannerInterval) (*DomainAggregate.PlannerInterval, error) {
	plannerIntervalRepository := InfrastructureService.GetFactoryRepository().GetPlannerIntervalRepository()
	plannerInterval, errorPlannerInterval := getPlannerIntervalAggregate(id, userId, entityId, nil)

	if errorPlannerInterval != nil {
		return nil, errors.Wrapf(errorPlannerInterval, "an error occurred while updating a planner interval by privided data id=%s,userId=%s,entityId=%v,criteria=%v", id, userId, entityId, nil)
	}

	plannerIntervalDTO.Id = *id
	plannerIntervalDTO.UserId = *userId
	plannerIntervalDTO.EntityId = *entityId
	plannerIntervalDTO.DateInsert = plannerInterval.Entity.DateInsert
	plannerIntervalDTO.DateUpdate = time.Now().UTC()

	plannerIntervalUpdated, errorPlannerIntervalUpdated := service.Update(plannerInterval.Entity, plannerIntervalDTO)

	if errorPlannerIntervalUpdated != nil {
		return nil, errors.Wrapf(errorPlannerIntervalUpdated, "an error occurred while updating a planner interval by privided data %v", plannerIntervalDTO)
	}

	restoredPlannerIntervalUpdated, okRestoredPlannerIntervalUpdated := plannerIntervalUpdated.Interface().(*DomainEntity.PlannerInterval)

	if !okRestoredPlannerIntervalUpdated {
		return nil, errors.Wrapf(errorUpdateRestored, "an error occurred while restoring updated a planner interval by privided data %s", plannerIntervalUpdated)
	}

	updateOne, errorUpdateOne := plannerIntervalRepository.UpdateOne(
		plannerIntervalRepository.GetCriteria().GetCriteriaById(&restoredPlannerIntervalUpdated.Id, nil),
		restoredPlannerIntervalUpdated,
	)

	if errorUpdateOne != nil {
		return nil, errors.Wrapf(errorUpdateOne, "an error occurred while updating a planner interval entity in the database %v", restoredPlannerIntervalUpdated)
	}

	return getPlannerIntervalAggregate(&updateOne.Id, &updateOne.UserId, &updateOne.EntityId, nil)
}

func PlannerIntervalDelete(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID) (bool, error) {
	plannerIntervalRepository := InfrastructureService.GetFactoryRepository().GetPlannerIntervalRepository()

	criteria := plannerIntervalRepository.GetCriteria().GetCriteriaById(id, nil)
	criteria = plannerIntervalRepository.GetCriteria().GetCriteriaByUserId(userId, criteria)
	criteria = plannerIntervalRepository.GetCriteria().GetCriteriaByEntityId(entityId, criteria)

	return plannerIntervalRepository.DeleteOne(criteria)
}

func getPlannerIntervalAggregate(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) (*DomainAggregate.PlannerInterval, error) {
	plannerIntervalEntities, errorPlannerIntervalEntities := ApplicationService.BuildPlannerIntervalAggregates(id, userId, entityId, criteria)
	if errorPlannerIntervalEntities != nil {
		return nil, errors.Wrapf(errorPlannerIntervalEntities, "an error occurred while getting a planner interval by privided data id=%s,userId=%s,entityId=%s,criteria=%v", id, userId, entityId, criteria)
	} else if len(plannerIntervalEntities) == 0 {
		return nil, errorPlannerIntervalInfo
	}
	return plannerIntervalEntities[0], nil
}
