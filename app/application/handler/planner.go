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
	errorPlannerCreate = errors.New("planner has not created by provided data")
	errorPlannerExists = errors.New("planner has not created by provided data")
	errorPlannerInfo   = errors.New("planner cannot be showed by provided data")
)

func PlannerCreate(userId *uuid.UUID, plannerDTO *DomainEntity.Planner) (*DomainAggregate.Planner, error) {
	plannerRepository := InfrastructureService.GetFactoryRepository().GetPlannerRepository()
	criteria := plannerRepository.GetCriteria().GetCriteriaByName(&plannerDTO.Name, nil)
	criteria = plannerRepository.GetCriteria().GetCriteriaByUserId(userId, criteria)
	plannerFindOne, errorPlannerFindOne := plannerRepository.FindOne(criteria)

	if errorPlannerFindOne == nil {
		return nil, errorPlannerCreate
	} else if plannerFindOne != nil {
		return nil, errorPlannerExists
	} else {
		plannerDTO.UserId = *userId
		plannerDTO.DateInsert = time.Now().UTC()
		plannerDTO.DateUpdate = time.Now().UTC()

		planner, errorPlannerInsertOne := plannerRepository.InsertOne(preparePlannerRepositoryInsert(plannerDTO))

		if errorPlannerInsertOne != nil {
			return nil, errors.Wrapf(errorPlannerInsertOne, "an error occurred while creating a planner in the database by privided data %v", plannerDTO)
		} else {
			return getPlannerAggregate(&planner.Id, &planner.UserId, nil)
		}
	}
}

func PlannersInfo(userId *uuid.UUID, criteria *persistence.Criteria) ([]*DomainAggregate.Planner, error) {
	return ApplicationService.BuildPlannersAggregate(nil, userId, criteria)
}

func PlannerInfo(id *uuid.UUID, userId *uuid.UUID, criteria *persistence.Criteria) (*DomainAggregate.Planner, error) {
	return getPlannerAggregate(id, userId, criteria)
}

func PlannerUpdate(id *uuid.UUID, userId *uuid.UUID, plannerDTO *DomainEntity.Planner) (*DomainAggregate.Planner, error) {
	plannerRepository := InfrastructureService.GetFactoryRepository().GetPlannerRepository()
	planner, errorPlanner := getPlannerAggregate(id, userId, nil)

	if errorPlanner != nil {
		return nil, errors.Wrapf(errorPlanner, "an error occurred while updating a planner by privided data id=%s,userId=%s,criteria=%v", id, userId, nil)
	}

	plannerDTO.Id = *id
	plannerDTO.UserId = *userId
	plannerDTO.DateInsert = planner.Entity.DateInsert
	plannerDTO.DateUpdate = time.Now().UTC()

	plannerUpdated, errorPlannerUpdated := service.Update(planner.Entity, plannerDTO)

	if errorPlannerUpdated != nil {
		return nil, errors.Wrapf(errorPlannerUpdated, "an error occurred while updating a planner by privided data %v", plannerDTO)
	}

	restoredPlannerUpdated, okRestoredPlannerUpdated := plannerUpdated.Interface().(*DomainEntity.Planner)

	if !okRestoredPlannerUpdated {
		return nil, errors.Wrapf(errorUpdateRestored, "an error occurred while restoring updated a planner by privided data %s", plannerUpdated)
	}

	updateOne, errorUpdateOne := plannerRepository.UpdateOne(
		plannerRepository.GetCriteria().GetCriteriaById(&restoredPlannerUpdated.Id, nil),
		restoredPlannerUpdated,
	)

	if errorUpdateOne != nil {
		return nil, errors.Wrapf(errorUpdateOne, "an error occurred while updating a planner entity in the database %v", restoredPlannerUpdated)
	}

	return getPlannerAggregate(&updateOne.Id, &updateOne.UserId, nil)
}

func PlannerDelete(id *uuid.UUID, userId *uuid.UUID) (bool, error) {
	plannerRepository := InfrastructureService.GetFactoryRepository().GetPlannerRepository()

	criteria := plannerRepository.GetCriteria().GetCriteriaByUserId(userId, nil)
	criteria = plannerRepository.GetCriteria().GetCriteriaById(id, criteria)

	return plannerRepository.DeleteOne(criteria)
}

func PlannerCalculate(id *uuid.UUID, userId *uuid.UUID) ([]*DomainAggregate.PlannerCalculation, error) {
	var plannerCalculations []*DomainAggregate.PlannerCalculation

	mapPlannerCalculations := map[string]*DomainAggregate.PlannerCalculation{}
	planner, errorPlanner := getPlannerAggregate(id, userId, nil)

	if errorPlanner != nil {
		return nil, errors.Wrapf(errorPlanner, "an error occurred while calculating the planner with id=%s", id)
	}

	for _, interval := range planner.Intervals {
		for _, recipe := range interval.Recipes {
			for _, ingredient := range recipe.Recipe.Ingredients {
				for _, measure := range ingredient.Measures {
					mapKey := ingredient.Derive.Id.String() + measure.Unit.Id.String()
					_, ok := mapPlannerCalculations[mapKey]

					if !ok {
						mapPlannerCalculations[mapKey] = &DomainAggregate.PlannerCalculation{
							Ingredient: ingredient.Derive,
							Unit:       measure.Unit,
							Amount:     0,
						}
						plannerCalculations = append(plannerCalculations, mapPlannerCalculations[mapKey])
					}

					mapPlannerCalculations[mapKey].Amount += measure.Entity.Value
				}
			}
		}
	}

	return plannerCalculations, nil
}

func getPlannerAggregate(id *uuid.UUID, userId *uuid.UUID, criteria *persistence.Criteria) (*DomainAggregate.Planner, error) {
	plannerEntities, errorPlannerEntities := ApplicationService.BuildPlannersAggregate(id, userId, criteria)
	if errorPlannerEntities != nil {
		return nil, errors.Wrapf(errorPlannerEntities, "an error occurred while getting a planner by privided data id=%s,userId=%s,criteria=%v", id, userId, criteria)
	} else if len(plannerEntities) == 0 {
		return nil, errorPlannerInfo
	}
	return plannerEntities[0], nil
}
