package update

import (
	"github.com/pkg/errors"
	"github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	InfrastructureService "github.com/sergeygardner/meal-planner-api/infrastructure/service/repository"
)

func SetUserConfirmationInActive(userConfirmation *entity.UserConfirmation) error {
	userConfirmationRepository := InfrastructureService.GetFactoryRepository().GetUserConfirmationRepository()
	(*userConfirmation).Active = kind.UserConfirmationInActive

	_, errorUpdateOne := userConfirmationRepository.UpdateOne(
		userConfirmationRepository.GetCriteriaById(&userConfirmation.Id),
		userConfirmation,
	)
	if errorUpdateOne != nil {
		return errors.Wrapf(errorUpdateOne, "an error occurred while updating a conformation for a user by provided data %p", userConfirmation)
	}

	return nil
}
