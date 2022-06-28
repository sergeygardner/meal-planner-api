package handler

import (
	"github.com/google/uuid"
	ApplicationService "github.com/sergeygardner/meal-planner-api/application/service"
	"github.com/sergeygardner/meal-planner-api/domain/dto"
	"github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	"time"
)

func prepareUserRepositoryInsert(userRegisterDTO dto.UserRegisterDTO, password string) *entity.User {
	newUUID, _ := uuid.NewUUID()

	return &entity.User{
		Id: newUUID,
		UserDTO: dto.UserDTO{
			DateInsert: time.Now().UTC(),
			DateUpdate: time.Now().UTC(),
			UserRegisterDTO: dto.UserRegisterDTO{
				UserCredentialsDTO: dto.UserCredentialsDTO{
					Username: userRegisterDTO.Username,
					Password: password,
				},
				Name:       userRegisterDTO.Name,
				Surname:    userRegisterDTO.Surname,
				MiddleName: userRegisterDTO.MiddleName,
				Birthday:   userRegisterDTO.Birthday,
			},
			Status: kind.UserStatusNeedConfirmation,
			Active: kind.UserInActive,
			Roles:  kind.UserRoles{kind.UserRoleCommon},
		},
	}
}

func prepareUserConfirmationRepositoryInsert(user entity.User) *entity.UserConfirmation {
	newUUID, _ := uuid.NewUUID()

	return &entity.UserConfirmation{
		Id:         newUUID,
		DateInsert: time.Now().UTC(),
		DateUpdate: time.Now().UTC(),
		UserId:     user.Id,
		Value:      ApplicationService.MathRandomIntAsString(entity.UserConfirmationValueMin, entity.UserConfirmationValueMax),
		Active:     kind.UserConfirmationActive,
	}
}

func prepareRecipeRepositoryInsert(recipe *entity.Recipe) *entity.Recipe {
	newUUID, _ := uuid.NewUUID()
	(*recipe).Id = newUUID

	return recipe
}

func prepareRecipeCategoryRepositoryInsert(recipeCategory *entity.RecipeCategory) *entity.RecipeCategory {
	newUUID, _ := uuid.NewUUID()
	recipeCategory.Id = newUUID

	return recipeCategory
}

func prepareRecipeIngredientRepositoryInsert(recipeIngredient *entity.RecipeIngredient) *entity.RecipeIngredient {
	newUUID, _ := uuid.NewUUID()
	recipeIngredient.Id = newUUID

	return recipeIngredient
}

func prepareRecipeProcessRepositoryInsert(recipeProcess *entity.RecipeProcess) *entity.RecipeProcess {
	newUUID, _ := uuid.NewUUID()
	recipeProcess.Id = newUUID

	return recipeProcess
}

func prepareRecipeMeasureRepositoryInsert(recipeMeasure *entity.RecipeMeasure) *entity.RecipeMeasure {
	newUUID, _ := uuid.NewUUID()
	recipeMeasure.Id = newUUID

	return recipeMeasure
}

func prepareUnitRepositoryInsert(recipeUnit *entity.Unit) *entity.Unit {
	newUUID, _ := uuid.NewUUID()
	recipeUnit.Id = newUUID

	return recipeUnit
}

func prepareCategoryRepositoryInsert(category *entity.Category) *entity.Category {
	newUUID, _ := uuid.NewUUID()
	category.Id = newUUID

	return category
}

func prepareIngredientRepositoryInsert(ingredient *entity.Ingredient) *entity.Ingredient {
	newUUID, _ := uuid.NewUUID()
	ingredient.Id = newUUID

	return ingredient
}

func preparePictureRepositoryInsert(recipePicture *entity.Picture) *entity.Picture {
	newUUID, _ := uuid.NewUUID()
	recipePicture.Id = newUUID

	return recipePicture
}

func prepareAltNameRepositoryInsert(recipeAltName *entity.AltName) *entity.AltName {
	newUUID, _ := uuid.NewUUID()
	recipeAltName.Id = newUUID

	return recipeAltName
}

func preparePlannerRepositoryInsert(planner *entity.Planner) *entity.Planner {
	newUUID, _ := uuid.NewUUID()
	planner.Id = newUUID

	return planner
}

func preparePlannerIntervalRepositoryInsert(plannerInterval *entity.PlannerInterval) *entity.PlannerInterval {
	newUUID, _ := uuid.NewUUID()
	plannerInterval.Id = newUUID

	return plannerInterval
}

func preparePlannerRecipeRepositoryInsert(plannerRecipe *entity.PlannerRecipe) *entity.PlannerRecipe {
	newUUID, _ := uuid.NewUUID()
	plannerRecipe.Id = newUUID

	return plannerRecipe
}
