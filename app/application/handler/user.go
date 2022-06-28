package handler

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	ApplicationServicePassword "github.com/sergeygardner/meal-planner-api/application/service/password"
	"github.com/sergeygardner/meal-planner-api/domain/dto"
	"github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/domain/service"
	InfrastructureService "github.com/sergeygardner/meal-planner-api/infrastructure/service/repository"
	"time"
)

var (
	errorUpdateRestored = errors.New("error has gotten while reflection datum is being restored")
)

func UserInfo(userId *uuid.UUID) (*entity.User, error) {
	userRepository := InfrastructureService.GetFactoryRepository().GetUserRepository()

	return userRepository.FindOne(userRepository.GetCriteriaByUserId(userId))
}

func UserUpdate(userId *uuid.UUID, userDTO *dto.UserDTO) (*entity.User, error) {
	userRepository := InfrastructureService.GetFactoryRepository().GetUserRepository()
	user, errorUser := UserInfo(userId)

	if errorUser != nil {
		return nil, errors.Wrapf(errorUser, "an error occurred while getting a user by provided data userId=%s", userId)
	}

	castPassword, errorCastPassword := ApplicationServicePassword.CastPassword(userDTO.Password)

	if errorCastPassword != nil {
		return nil, errors.Wrapf(errorCastPassword, "an error occurred while casting a password for a user by provided data password=%s", userDTO.Password)
	} else {
		userDTO.Password = string(castPassword)
	}

	userDTO.DateInsert = user.DateInsert
	userDTO.DateUpdate = time.Now().UTC()
	userUpdated, errorUserUpdate := service.Update(user, userDTO)

	if errorUserUpdate != nil {
		return nil, errors.Wrapf(errorUserUpdate, "an error occurred while updating a user by privided data %v", userDTO)
	}

	restoredUserUpdated, okRestoredUserUpdated := userUpdated.Interface().(*entity.User)

	if !okRestoredUserUpdated {
		return nil, errors.Wrapf(errorUpdateRestored, "an error occurred while restoring updated a user by privided data %s", userUpdated)
	}

	updateOne, errorUpdateOne := InfrastructureService.GetFactoryRepository().GetUserRepository().UpdateOne(
		userRepository.GetCriteriaByUserId(&restoredUserUpdated.Id),
		restoredUserUpdated,
	)

	if errorUpdateOne != nil {
		return nil, errors.Wrapf(errorUpdateOne, "an error occurred while updating a user entity in the database %v", restoredUserUpdated)
	}

	return updateOne, nil
}

func UserDelete(userId *uuid.UUID) (bool, error) {
	userRepository := InfrastructureService.GetFactoryRepository().GetUserRepository()

	return userRepository.DeleteOne(userRepository.GetCriteriaByUserId(userId))
}
