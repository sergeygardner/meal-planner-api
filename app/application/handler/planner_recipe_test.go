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
	testsPlannerRecipeData = testsPlannerRecipe{
		{
			name:     "Test case with correct data",
			id:       nil,
			userId:   &testUserId,
			entityId: &uuidPlanner,
			plannerRecipeDTO: &DomainEntity.PlannerRecipe{
				RecipeId: uuid.Nil,
				Status:   kind.PlannerRecipeStatusInActive,
			},
			toUpdatingPlannerRecipeDTO: &DomainEntity.PlannerRecipe{
				Status: kind.PlannerRecipeStatusActive,
			},
		},
	}
)

type testsPlannerRecipe []struct {
	name                       string
	id                         *uuid.UUID
	userId                     *uuid.UUID
	entityId                   *uuid.UUID
	plannerRecipeDTO           *DomainEntity.PlannerRecipe
	toUpdatingPlannerRecipeDTO *DomainEntity.PlannerRecipe
	plannerRecipe              *DomainAggregate.PlannerRecipe
}

func init() {
	InfrastructureService.PrepareCache()
	InfrastructureService.PreparePersistence()
}

func TestPlannerRecipeCreate(t *testing.T) {
	for index, testCase := range testsPlannerRecipeData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				TestRecipeCreate(t)

				testCase.plannerRecipeDTO.RecipeId = *testsRecipeData[0].id

				actual, errorActual := PlannerRecipeCreate(testCase.userId, testCase.entityId, testCase.plannerRecipeDTO)

				assert.Nil(t, errorActual)

				testsPlannerRecipeData[index].plannerRecipe = actual
				testsPlannerRecipeData[index].id = &actual.Entity.Id

				assert.NotNil(t, actual.Entity.Id)
				assert.Equal(t, testCase.plannerRecipeDTO.UserId, actual.Entity.UserId)
				assert.NotNil(t, actual.Entity.DateInsert)
				assert.NotNil(t, actual.Entity.DateUpdate)
				assert.Equal(t, testCase.plannerRecipeDTO.RecipeId, actual.Entity.RecipeId)
				assert.Equal(t, testCase.plannerRecipeDTO.Status, actual.Entity.Status)
			},
		)
	}
}

func TestPlannerRecipesInfo(t *testing.T) {
	for _, testCase := range testsPlannerRecipeData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := PlannerRecipesInfo(testCase.userId, testCase.entityId, nil)

				if testCase.plannerRecipe != nil {
					assert.Nil(t, errorActual)

					for _, actualEntity := range actual {
						assert.Equal(t, testCase.plannerRecipe.Entity.Id, actualEntity.Entity.Id)
						assert.Equal(t, testCase.plannerRecipe.Entity.UserId, actualEntity.Entity.UserId)
						assert.Equal(t, testCase.plannerRecipe.Entity.DateInsert.Format(time.UnixDate), actualEntity.Entity.DateInsert.Format(time.UnixDate))
						assert.Equal(t, testCase.plannerRecipe.Entity.DateUpdate.Format(time.UnixDate), actualEntity.Entity.DateUpdate.Format(time.UnixDate))
						assert.Equal(t, testCase.plannerRecipe.Entity.RecipeId, actualEntity.Entity.RecipeId)
						assert.Equal(t, testCase.plannerRecipe.Entity.Status, actualEntity.Entity.Status)
					}
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestPlannerRecipeInfo(t *testing.T) {
	for _, testCase := range testsPlannerRecipeData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := PlannerRecipeInfo(testCase.id, testCase.userId, testCase.entityId, nil)

				if testCase.plannerRecipe != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.plannerRecipe.Entity.Id, actual.Entity.Id)
					assert.Equal(t, testCase.plannerRecipe.Entity.UserId, actual.Entity.UserId)
					assert.Equal(t, testCase.plannerRecipe.Entity.DateInsert.Format(time.UnixDate), actual.Entity.DateInsert.Format(time.UnixDate))
					assert.Equal(t, testCase.plannerRecipe.Entity.DateUpdate.Format(time.UnixDate), actual.Entity.DateUpdate.Format(time.UnixDate))
					assert.Equal(t, testCase.plannerRecipe.Entity.RecipeId, actual.Entity.RecipeId)
					assert.Equal(t, testCase.plannerRecipe.Entity.Status, actual.Entity.Status)
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestPlannerRecipeUpdate(t *testing.T) {
	for index, testCase := range testsPlannerRecipeData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := PlannerRecipeUpdate(testCase.id, testCase.userId, testCase.entityId, testCase.toUpdatingPlannerRecipeDTO)

				if testCase.plannerRecipe != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.toUpdatingPlannerRecipeDTO.Status, actual.Entity.Status)
					testsPlannerRecipeData[index].plannerRecipe.Entity.Status = actual.Entity.Status
					testsPlannerRecipeData[index].plannerRecipe.Entity.DateUpdate = actual.Entity.DateUpdate
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestPlannerRecipeDelete(t *testing.T) {
	for _, testCase := range testsPlannerRecipeData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				TestRecipeDelete(t)

				actual, errorActual := PlannerRecipeDelete(testCase.id, testCase.userId, testCase.entityId)

				if testCase.plannerRecipe != nil {
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
