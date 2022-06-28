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
	errorAltNameCreate = errors.New("alt name has not created by provided data")
	errorAltNameExists = errors.New("alt name has not created by provided data")
	errorAltNameInfo   = errors.New("alt name cannot be showed by provided data")
)

func AltNameCreate(userId *uuid.UUID, entityId *uuid.UUID, altNameDTO *DomainEntity.AltName) (*DomainEntity.AltName, error) {
	altNameRepository := InfrastructureService.GetFactoryRepository().GetAltNameRepository()
	criteria := altNameRepository.GetCriteria().GetCriteriaByName(&altNameDTO.Name, nil)
	criteria = altNameRepository.GetCriteria().GetCriteriaByUserId(userId, criteria)
	criteria = altNameRepository.GetCriteria().GetCriteriaByEntityId(entityId, criteria)
	altNameFindOne, errorAltNameFindOne := altNameRepository.FindOne(criteria)

	if errorAltNameFindOne == nil {
		return nil, errorAltNameCreate
	} else if altNameFindOne != nil {
		return nil, errorAltNameExists
	} else {
		altNameDTO.UserId = *userId
		altNameDTO.EntityId = *entityId
		altNameDTO.DateInsert = time.Now().UTC()
		altNameDTO.DateUpdate = time.Now().UTC()

		altName, errorAltNameInsertOne := altNameRepository.InsertOne(prepareAltNameRepositoryInsert(altNameDTO))

		if errorAltNameInsertOne != nil {
			return nil, errors.Wrapf(errorAltNameInsertOne, "an error occurred while creating an alt name in the database by privided data %s", altNameDTO)
		} else {
			return altName, nil
		}
	}
}

func AltNamesInfo(userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) ([]*DomainEntity.AltName, error) {
	return ApplicationService.BuildAltNameEntities(nil, userId, entityId, criteria)
}

func AltNameInfo(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) (*DomainEntity.AltName, error) {
	return getAltNameEntity(id, userId, entityId, criteria)
}

func AltNameUpdate(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, altNameDTO *DomainEntity.AltName) (*DomainEntity.AltName, error) {
	altNameRepository := InfrastructureService.GetFactoryRepository().GetAltNameRepository()
	altName, errorAltName := getAltNameEntity(id, userId, entityId, nil)

	if errorAltName != nil {
		return nil, errors.Wrapf(errorAltName, "an error occurred while updating an alt name by privided data id=%s,userId=%s,entityId=%v,criteria=%v", id, userId, nil, nil)
	}

	altNameDTO.Id = *id
	altNameDTO.UserId = *userId
	altNameDTO.EntityId = *entityId
	altNameDTO.DateInsert = altName.DateInsert
	altNameDTO.DateUpdate = time.Now().UTC()

	altNameUpdated, errorAltNameUpdated := service.Update(altName, altNameDTO)

	if errorAltNameUpdated != nil {
		return nil, errors.Wrapf(errorAltNameUpdated, "an error occurred while updating an alt name by privided data %s", altNameDTO)
	}

	restoredAltNameUpdated, okRestoredAltNameUpdated := altNameUpdated.Interface().(*DomainEntity.AltName)

	if !okRestoredAltNameUpdated {
		return nil, errors.Wrapf(errorUpdateRestored, "an error occurred while restoring updated an alt name by privided data %s", altNameUpdated)
	}

	updateOne, errorUpdateOne := altNameRepository.UpdateOne(
		altNameRepository.GetCriteria().GetCriteriaById(&restoredAltNameUpdated.Id, nil),
		restoredAltNameUpdated,
	)

	if errorUpdateOne != nil {
		return nil, errors.Wrapf(errorUpdateOne, "an error occurred while updating an alt name entity in the database %s", restoredAltNameUpdated)
	}

	return updateOne, nil
}

func AltNameDelete(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID) (bool, error) {
	altNameRepository := InfrastructureService.GetFactoryRepository().GetAltNameRepository()

	criteria := altNameRepository.GetCriteria().GetCriteriaById(id, nil)
	criteria = altNameRepository.GetCriteria().GetCriteriaByUserId(userId, criteria)
	criteria = altNameRepository.GetCriteria().GetCriteriaByEntityId(entityId, criteria)

	return altNameRepository.DeleteOne(criteria)
}

func getAltNameEntity(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) (*DomainEntity.AltName, error) {
	altNameEntities, errorAltNameEntities := ApplicationService.BuildAltNameEntities(id, userId, entityId, criteria)
	if errorAltNameEntities != nil {
		return nil, errors.Wrapf(errorAltNameEntities, "an error occurred while getting an alt name by privided data id=%s,userId=%s,entityId=%s,criteria=%v", id, userId, entityId, criteria)
	} else if len(altNameEntities) == 0 {
		return nil, errorAltNameInfo
	}
	return altNameEntities[0], nil
}
