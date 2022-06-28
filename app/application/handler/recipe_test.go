package handler

import (
	"github.com/google/uuid"
	DomainAggregate "github.com/sergeygardner/meal-planner-api/domain/aggregate"
	DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	InfrastructureService "github.com/sergeygardner/meal-planner-api/infrastructure/service"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	testsRecipeData = testsRecipe{
		{
			name:   "Test case with correct data",
			id:     nil,
			userId: &testUserId,
			recipeDTO: &DomainEntity.Recipe{
				Name:   "Test case with correct data",
				Status: kind.RecipeStatusUnPublished,
			},
			toUpdatingRecipeDTO: &DomainEntity.Recipe{
				Status: kind.RecipeStatusPublished,
			},
		},
	}
)

type testsRecipe []struct {
	name                string
	id                  *uuid.UUID
	userId              *uuid.UUID
	recipeDTO           *DomainEntity.Recipe
	toUpdatingRecipeDTO *DomainEntity.Recipe
	recipe              *DomainAggregate.Recipe
}

func init() {
	InfrastructureService.PrepareCache()
	InfrastructureService.PreparePersistence()
}

func TestRecipeCreate(t *testing.T) {
	for index, testCase := range testsRecipeData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := RecipeCreate(testCase.userId, testCase.recipeDTO)

				assert.Nil(t, errorActual)

				testsRecipeData[index].recipe = actual
				testsRecipeData[index].id = &actual.Entity.Id

				assert.NotNil(t, actual.Entity.Id)
				assert.Equal(t, testCase.recipeDTO.UserId, actual.Entity.UserId)
				assert.NotNil(t, actual.Entity.DateInsert)
				assert.NotNil(t, actual.Entity.DateUpdate)
				assert.Equal(t, testCase.recipeDTO.Name, actual.Entity.Name)
				assert.Equal(t, testCase.recipeDTO.Status, actual.Entity.Status)
			},
		)
	}
}

func TestRecipesInfo(t *testing.T) {
	for _, testCase := range testsRecipeData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := RecipesInfo(testCase.userId, nil)

				if testCase.recipe != nil {
					assert.Nil(t, errorActual)

					for _, actualEntity := range actual {
						assert.Equal(t, testCase.recipe.Entity.Id, actualEntity.Entity.Id)
						assert.Equal(t, testCase.recipe.Entity.UserId, actualEntity.Entity.UserId)
						assert.Equal(t, testCase.recipe.Entity.DateInsert.Format(time.UnixDate), actualEntity.Entity.DateInsert.Format(time.UnixDate))
						assert.Equal(t, testCase.recipe.Entity.DateUpdate.Format(time.UnixDate), actualEntity.Entity.DateUpdate.Format(time.UnixDate))
						assert.Equal(t, testCase.recipe.Entity.Name, actualEntity.Entity.Name)
						assert.Equal(t, testCase.recipe.Entity.Status, actualEntity.Entity.Status)
					}
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestRecipeInfo(t *testing.T) {
	for _, testCase := range testsRecipeData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := RecipeInfo(testCase.id, testCase.userId, nil)

				if testCase.recipe != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.recipe.Entity.Id, actual.Entity.Id)
					assert.Equal(t, testCase.recipe.Entity.UserId, actual.Entity.UserId)
					assert.Equal(t, testCase.recipe.Entity.DateInsert.Format(time.UnixDate), actual.Entity.DateInsert.Format(time.UnixDate))
					assert.Equal(t, testCase.recipe.Entity.DateUpdate.Format(time.UnixDate), actual.Entity.DateUpdate.Format(time.UnixDate))
					assert.Equal(t, testCase.recipe.Entity.Name, actual.Entity.Name)
					assert.Equal(t, testCase.recipe.Entity.Status, actual.Entity.Status)
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestRecipeUpdate(t *testing.T) {
	for index, testCase := range testsRecipeData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := RecipeUpdate(testCase.id, testCase.userId, testCase.toUpdatingRecipeDTO)

				if testCase.recipe != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.toUpdatingRecipeDTO.Status, actual.Entity.Status)
					testsRecipeData[index].recipe.Entity.Status = actual.Entity.Status
					testsRecipeData[index].recipe.Entity.DateUpdate = actual.Entity.DateUpdate
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestGetEntity(t *testing.T) {
	TestRecipeInfo(t)
}

func TestRecipeDelete(t *testing.T) {
	for _, testCase := range testsRecipeData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := RecipeDelete(testCase.id, testCase.userId)

				if testCase.recipe != nil {
					assert.Nil(t, errorActual)
					assert.True(t, actual)
				} else {
					assert.NotNil(t, errorActual)
					assert.False(t, actual)
				}
			},
		)
	}
}
