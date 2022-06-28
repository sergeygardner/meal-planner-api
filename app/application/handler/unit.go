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
	errorUnitCreate = errors.New("unit has not created by provided data")
	errorUnitExists = errors.New("unit has not created by provided data")
	errorUnitInfo   = errors.New("unit cannot be showed by provided data")
)

func UnitCreate(unitDTO *DomainEntity.Unit) (*DomainEntity.Unit, error) {
	unitRepository := InfrastructureService.GetFactoryRepository().GetUnitRepository()
	criteria := unitRepository.GetCriteria().GetCriteriaByName(&unitDTO.Name, nil)
	unitFindOne, errorUnitFindOne := unitRepository.FindOne(criteria)

	if errorUnitFindOne == nil {
		return nil, errorUnitCreate
	} else if unitFindOne != nil {
		return nil, errorUnitExists
	} else {
		unitDTO.DateInsert = time.Now().UTC()
		unitDTO.DateUpdate = time.Now().UTC()

		unit, errorUnitInsertOne := unitRepository.InsertOne(prepareUnitRepositoryInsert(unitDTO))

		if errorUnitInsertOne != nil {
			return nil, errors.Wrapf(errorUnitInsertOne, "an error occurred while creating an unit in the database by privided data %s", unitDTO)
		} else {
			return unit, nil
		}
	}
}

func UnitsInfo(criteria *persistence.Criteria) ([]*DomainEntity.Unit, error) {
	return ApplicationService.BuildUnitEntities(nil, criteria)
}

func UnitInfo(id *uuid.UUID, criteria *persistence.Criteria) (*DomainEntity.Unit, error) {
	return getUnitEntity(id, criteria)
}

func UnitUpdate(id *uuid.UUID, unitDTO *DomainEntity.Unit) (*DomainEntity.Unit, error) {
	unitRepository := InfrastructureService.GetFactoryRepository().GetUnitRepository()
	unit, errorUnit := getUnitEntity(id, nil)

	if errorUnit != nil {
		return nil, errors.Wrapf(errorUnit, "an error occurred while updating an unit by privided data id=%s", id)
	}

	unitDTO.Id = *id
	unitDTO.DateInsert = unit.DateInsert
	unitDTO.DateUpdate = time.Now().UTC()

	unitUpdated, errorUnitUpdated := service.Update(unit, unitDTO)

	if errorUnitUpdated != nil {
		return nil, errors.Wrapf(errorUnitUpdated, "an error occurred while updating an unit by privided data %s", unitDTO)
	}

	restoredUnitUpdated, okRestoredUnitUpdated := unitUpdated.Interface().(*DomainEntity.Unit)

	if !okRestoredUnitUpdated {
		return nil, errors.Wrapf(errorUpdateRestored, "an error occurred while restoring updated an unit by privided data %s", unitUpdated)
	}

	updateOne, errorUpdateOne := unitRepository.UpdateOne(
		unitRepository.GetCriteria().GetCriteriaById(&restoredUnitUpdated.Id, nil),
		restoredUnitUpdated,
	)

	if errorUpdateOne != nil {
		return nil, errors.Wrapf(errorUpdateOne, "an error occurred while updating an unit entity in the database %s", restoredUnitUpdated)
	}

	return updateOne, nil
}

func UnitDelete(id *uuid.UUID) (bool, error) {
	unitRepository := InfrastructureService.GetFactoryRepository().GetUnitRepository()

	criteria := unitRepository.GetCriteria().GetCriteriaById(id, nil)

	return unitRepository.DeleteOne(criteria)
}

func getUnitEntity(id *uuid.UUID, criteria *persistence.Criteria) (*DomainEntity.Unit, error) {
	unitEntities, errorUnitEntities := ApplicationService.BuildUnitEntities(id, criteria)
	if errorUnitEntities != nil {
		return nil, errors.Wrapf(errorUnitEntities, "an error occurred while getting an unit by privided data id=%s,criteria=%v", id, criteria)
	} else if len(unitEntities) == 0 {
		return nil, errorUnitInfo
	}
	return unitEntities[0], nil
}
