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
	errorPictureCreate = errors.New("picture has not created by provided data")
	errorPictureExists = errors.New("picture has not created by provided data")
	errorPictureInfo   = errors.New("picture cannot be showed by provided data")
)

func PictureCreate(userId *uuid.UUID, entityId *uuid.UUID, pictureDTO *DomainEntity.Picture) (*DomainAggregate.Picture, error) {
	pictureRepository := InfrastructureService.GetFactoryRepository().GetPictureRepository()
	criteria := pictureRepository.GetCriteria().GetCriteriaByName(&pictureDTO.Name, nil)
	criteria = pictureRepository.GetCriteria().GetCriteriaByUserId(userId, criteria)
	criteria = pictureRepository.GetCriteria().GetCriteriaByEntityId(entityId, criteria)
	pictureFindOne, errorPictureFindOne := pictureRepository.FindOne(criteria)

	if errorPictureFindOne == nil {
		return nil, errorPictureCreate
	} else if pictureFindOne != nil {
		return nil, errorPictureExists
	} else {
		pictureDTO.UserId = *userId
		pictureDTO.EntityId = *entityId
		pictureDTO.DateInsert = time.Now().UTC()
		pictureDTO.DateUpdate = time.Now().UTC()

		picture, errorPicturesInsertOne := pictureRepository.InsertOne(preparePictureRepositoryInsert(pictureDTO))

		if errorPicturesInsertOne != nil {
			return nil, errors.Wrapf(errorPicturesInsertOne, "an error occurred while creating a picture in the database by privided data %v", pictureDTO)
		} else {
			return getPictureAggregate(&picture.Id, &picture.UserId, &picture.EntityId, nil)
		}
	}
}

func PicturesInfo(userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) ([]*DomainAggregate.Picture, error) {
	return ApplicationService.BuildPicturesAggregate(nil, userId, entityId, criteria)
}

func PictureInfo(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) (*DomainAggregate.Picture, error) {
	return getPictureAggregate(id, userId, entityId, criteria)
}

func PictureUpdate(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, pictureDTO *DomainEntity.Picture) (*DomainAggregate.Picture, error) {
	pictureRepository := InfrastructureService.GetFactoryRepository().GetPictureRepository()
	pictureAggregate, errorPictureAggregate := getPictureAggregate(id, userId, entityId, nil)

	if errorPictureAggregate != nil {
		return nil, errors.Wrapf(errorPictureAggregate, "an error occurred while updating a picture by privided data id=%s,userId=%s,entityId=%v,criteria=%v", id, userId, nil, nil)
	}

	pictureDTO.Id = *id
	pictureDTO.UserId = *userId
	pictureDTO.EntityId = *entityId
	pictureDTO.DateInsert = pictureAggregate.Entity.DateInsert
	pictureDTO.DateUpdate = time.Now().UTC()

	pictureEntityUpdated, errorPictureEntityUpdate := service.Update(pictureAggregate.Entity, pictureDTO)

	if errorPictureEntityUpdate != nil {
		return nil, errors.Wrapf(errorPictureEntityUpdate, "an error occurred while updating a picture by privided data %v", pictureDTO)
	}

	restoredPictureEntityUpdated, okRestoredPictureEntityUpdated := pictureEntityUpdated.Interface().(*DomainEntity.Picture)

	if !okRestoredPictureEntityUpdated {
		return nil, errors.Wrapf(errorUpdateRestored, "an error occurred while restoring updated a picture by privided data %s", pictureEntityUpdated)
	}

	updateOne, errorUpdateOne := pictureRepository.UpdateOne(
		pictureRepository.GetCriteria().GetCriteriaById(&restoredPictureEntityUpdated.Id, nil),
		restoredPictureEntityUpdated,
	)

	if errorUpdateOne != nil {
		return nil, errors.Wrapf(errorUpdateOne, "an error occurred while updating a picture entity in the database %v", restoredPictureEntityUpdated)
	}

	pictureAggregate.Entity = updateOne

	return pictureAggregate, nil
}

func PictureDelete(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID) (bool, error) {
	pictureRepository := InfrastructureService.GetFactoryRepository().GetPictureRepository()

	criteria := pictureRepository.GetCriteria().GetCriteriaById(id, nil)
	criteria = pictureRepository.GetCriteria().GetCriteriaByUserId(userId, criteria)
	criteria = pictureRepository.GetCriteria().GetCriteriaByEntityId(entityId, criteria)

	return pictureRepository.DeleteOne(criteria)
}

func getPictureAggregate(id *uuid.UUID, userId *uuid.UUID, entityId *uuid.UUID, criteria *persistence.Criteria) (*DomainAggregate.Picture, error) {
	picturesAggregate, errorPicturesAggregate := ApplicationService.BuildPicturesAggregate(id, userId, entityId, criteria)
	if errorPicturesAggregate != nil {
		return nil, errors.Wrapf(errorPicturesAggregate, "an error occurred while getting a picture by privided data id=%s,userId=%s,entityId=%s,criteria=%v", id, userId, entityId, criteria)
	} else if len(picturesAggregate) == 0 {
		return nil, errorPictureInfo
	}
	return picturesAggregate[0], nil
}
