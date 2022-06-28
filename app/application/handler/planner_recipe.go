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
	errorPlannerRecipeCreate = errors.New("planner recipe has not created by provided data")
	errorPlannerRecipeExists = errors.New("planner recipe has not created by provided data")
	errorPlannerRecipeInfo   = errors.New("planner recipe cannot be showed by provided data")
)

func PlannerRecipeCreate(userId *uuid.UUID, entityId *uuid.UUID, plannerRecipeDTO *DomainEntity.PlannerRecipe) (*DomainAggregate.PlannerRecipe, error) {
	plannerRecipeRepository := InfrastructureService.GetFactoryRepository().GetPlannerRecipeRepository()
	criteria := plannerRecipeRepository.GetCriteria().GetCriteriaByRecipeId(&plannerRecipeDTO.RecipeId, nil)
	criteria = plannerRecipeRepository.GetCriteria().GetCriteriaByUserId(userId, criteria)
	criteria = plannerRecipeRepository.GetCriteria().GetCriteriaByEntityId(entityId, criteria)
	plannerRecipeFindOne, errorPlannerRecipeFindOne := plannerRecipeRepository.FindOne(criteria)

	if errorPlannerRecipeFindOne == nil {
		return nil, errorPlannerRecipeCreate
	} else if plannerRecipeFindOne != nil {
		return nil, errorPlannerRecipeExists
	} else {
		plannerRecipeDTO.UserId = *userId
		plannerRecipeDTO.EntityId = *entityId
		plannerRecipeDTO.DateInsert = time.Now().UTC()
		plannerRecipeDTO.DateUpdate = time.Now().UTC()

		plannerRecipe, errorPlannerRecipeInsertOne := plannerRecipeRepository.InsertOne(preparePlannerRecipeRepositoryInsert(plannerRecipeDTO))

		if errorPlannerRecipeInsertOne != nil {
			return nil, errors.Wrapf(errorPlannerRecipeInsertOne, "an error occurred while creating a planner recipe in the database by privided data %v", plannerRecipeDTO)
		} else {
			return getPlannerRecipeAggregate(&plannerRecipe.Id, &plannerRecipe.UserId, &plannerRecipe.EntityId, nil)
		}
	}
}

func PlannerRecipesInfo(userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) ([]*DomainAggregate.PlannerRecipe, error) {
	return ApplicationService.BuildPlannerRecipeAggregates(nil, userId, entityId, criteria)
}

func PlannerRecipeInfo(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) (*DomainAggregate.PlannerRecipe, error) {
	return getPlannerRecipeAggregate(id, userId, entityId, criteria)
}

func PlannerRecipeUpdate(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, plannerRecipeDTO *DomainEntity.PlannerRecipe) (*DomainAggregate.PlannerRecipe, error) {
	plannerRecipeRepository := InfrastructureService.GetFactoryRepository().GetPlannerRecipeRepository()
	plannerRecipe, errorPlannerRecipe := getPlannerRecipeAggregate(id, userId, entityId, nil)

	if errorPlannerRecipe != nil {
		return nil, errors.Wrapf(errorPlannerRecipe, "an error occurred while updating a planner recipe by privided data id=%s,userId=%s,entityId=%v,criteria=%v", id, userId, entityId, nil)
	}

	plannerRecipeDTO.Id = *id
	plannerRecipeDTO.UserId = *userId
	plannerRecipeDTO.EntityId = *entityId
	plannerRecipeDTO.DateInsert = plannerRecipe.Entity.DateInsert
	plannerRecipeDTO.DateUpdate = time.Now().UTC()

	plannerRecipeUpdated, errorPlannerRecipeUpdated := service.Update(plannerRecipe.Entity, plannerRecipeDTO)

	if errorPlannerRecipeUpdated != nil {
		return nil, errors.Wrapf(errorPlannerRecipeUpdated, "an error occurred while updating a planner recipe by privided data %v", plannerRecipeDTO)
	}

	restoredPlannerRecipeUpdated, okRestoredPlannerRecipeUpdated := plannerRecipeUpdated.Interface().(*DomainEntity.PlannerRecipe)

	if !okRestoredPlannerRecipeUpdated {
		return nil, errors.Wrapf(errorUpdateRestored, "an error occurred while restoring updated a planner recipe by privided data %s", plannerRecipeUpdated)
	}

	updateOne, errorUpdateOne := plannerRecipeRepository.UpdateOne(
		plannerRecipeRepository.GetCriteria().GetCriteriaById(&restoredPlannerRecipeUpdated.Id, nil),
		restoredPlannerRecipeUpdated,
	)

	if errorUpdateOne != nil {
		return nil, errors.Wrapf(errorUpdateOne, "an error occurred while updating a planner recipe entity in the database %v", restoredPlannerRecipeUpdated)
	}

	return getPlannerRecipeAggregate(&updateOne.Id, &updateOne.UserId, &updateOne.EntityId, nil)
}

func PlannerRecipeDelete(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID) (bool, error) {
	plannerRecipeRepository := InfrastructureService.GetFactoryRepository().GetPlannerRecipeRepository()

	criteria := plannerRecipeRepository.GetCriteria().GetCriteriaById(id, nil)
	criteria = plannerRecipeRepository.GetCriteria().GetCriteriaByUserId(userId, criteria)
	criteria = plannerRecipeRepository.GetCriteria().GetCriteriaByEntityId(entityId, criteria)

	return plannerRecipeRepository.DeleteOne(criteria)
}

func getPlannerRecipeAggregate(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) (*DomainAggregate.PlannerRecipe, error) {
	plannerRecipeAggregates, errorPlannerRecipeAggregates := ApplicationService.BuildPlannerRecipeAggregates(id, userId, entityId, criteria)
	if errorPlannerRecipeAggregates != nil {
		return nil, errors.Wrapf(errorPlannerRecipeAggregates, "an error occurred while getting a planner recipe by privided data id=%s,userId=%s,entityId=%s,criteria=%v", id, userId, entityId, criteria)
	} else if len(plannerRecipeAggregates) == 0 {
		return nil, errorPlannerRecipeInfo
	}
	return plannerRecipeAggregates[0], nil
}
