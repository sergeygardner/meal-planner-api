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
	entityIdRecipeIngredient  = uuid.New()
	testsRecipeIngredientData = testsRecipeIngredient{
		{
			name:     "Test case with correct data",
			id:       nil,
			userId:   &testUserId,
			entityId: &entityIdRecipeIngredient,
			recipeIngredientDTO: &DomainEntity.RecipeIngredient{
				Name:   "Test case with correct data",
				Status: kind.RecipeIngredientStatusUnPublished,
			},
			toUpdatingRecipeIngredientDTO: &DomainEntity.RecipeIngredient{
				Status: kind.RecipeIngredientStatusPublished,
			},
		},
	}
)

type testsRecipeIngredient []struct {
	name                          string
	id                            *uuid.UUID
	userId                        *uuid.UUID
	entityId                      *uuid.UUID
	recipeIngredientDTO           *DomainEntity.RecipeIngredient
	toUpdatingRecipeIngredientDTO *DomainEntity.RecipeIngredient
	recipeIngredient              *DomainAggregate.RecipeIngredient
}

func init() {
	InfrastructureService.PrepareCache()
	InfrastructureService.PreparePersistence()
}

func TestRecipeIngredientCreate(t *testing.T) {
	for index, testCase := range testsRecipeIngredientData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := RecipeIngredientCreate(testCase.userId, testCase.entityId, testCase.recipeIngredientDTO)

				assert.Nil(t, errorActual)

				testsRecipeIngredientData[index].recipeIngredient = actual
				testsRecipeIngredientData[index].id = &actual.Entity.Id
				testsRecipeIngredientData[index].entityId = &actual.Entity.EntityId

				assert.NotNil(t, actual.Entity.Id)
				assert.Equal(t, testCase.recipeIngredientDTO.UserId, actual.Entity.UserId)
				assert.Equal(t, testCase.recipeIngredientDTO.EntityId, actual.Entity.EntityId)
				assert.NotNil(t, actual.Entity.DateInsert)
				assert.NotNil(t, actual.Entity.DateUpdate)
				assert.Equal(t, testCase.recipeIngredientDTO.Name, actual.Entity.Name)
				assert.Equal(t, testCase.recipeIngredientDTO.Status, actual.Entity.Status)
			},
		)
	}
}

func TestRecipeIngredientsInfo(t *testing.T) {
	for _, testCase := range testsRecipeIngredientData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := RecipeIngredientsInfo(testCase.userId, testCase.entityId, nil)

				if testCase.recipeIngredient != nil {
					assert.Nil(t, errorActual)

					for _, actualEntity := range actual {
						assert.Equal(t, testCase.recipeIngredient.Entity.Id, actualEntity.Entity.Id)
						assert.Equal(t, testCase.recipeIngredient.Entity.UserId, actualEntity.Entity.UserId)
						assert.Equal(t, testCase.recipeIngredient.Entity.EntityId, actualEntity.Entity.EntityId)
						assert.Equal(t, testCase.recipeIngredient.Entity.DateInsert.Format(time.UnixDate), actualEntity.Entity.DateInsert.Format(time.UnixDate))
						assert.Equal(t, testCase.recipeIngredient.Entity.DateUpdate.Format(time.UnixDate), actualEntity.Entity.DateUpdate.Format(time.UnixDate))
						assert.Equal(t, testCase.recipeIngredient.Entity.Name, actualEntity.Entity.Name)
						assert.Equal(t, testCase.recipeIngredient.Entity.Status, actualEntity.Entity.Status)
					}
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestRecipeIngredientInfo(t *testing.T) {
	for _, testCase := range testsRecipeIngredientData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := RecipeIngredientInfo(testCase.id, testCase.userId, testCase.entityId, nil)

				if testCase.recipeIngredient != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.recipeIngredient.Entity.Id, actual.Entity.Id)
					assert.Equal(t, testCase.recipeIngredient.Entity.UserId, actual.Entity.UserId)
					assert.Equal(t, testCase.recipeIngredient.Entity.EntityId, actual.Entity.EntityId)
					assert.Equal(t, testCase.recipeIngredient.Entity.DateInsert.Format(time.UnixDate), actual.Entity.DateInsert.Format(time.UnixDate))
					assert.Equal(t, testCase.recipeIngredient.Entity.DateUpdate.Format(time.UnixDate), actual.Entity.DateUpdate.Format(time.UnixDate))
					assert.Equal(t, testCase.recipeIngredient.Entity.Name, actual.Entity.Name)
					assert.Equal(t, testCase.recipeIngredient.Entity.Status, actual.Entity.Status)
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestRecipeIngredientUpdate(t *testing.T) {
	for index, testCase := range testsRecipeIngredientData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := RecipeIngredientUpdate(testCase.id, testCase.userId, testCase.entityId, testCase.toUpdatingRecipeIngredientDTO)

				if testCase.recipeIngredient != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.toUpdatingRecipeIngredientDTO.Status, actual.Entity.Status)
					testsRecipeIngredientData[index].recipeIngredient.Entity.Status = actual.Entity.Status
					testsRecipeIngredientData[index].recipeIngredient.Entity.DateUpdate = actual.Entity.DateUpdate
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestGetRecipeIngredientEntity(t *testing.T) {
	TestRecipeIngredientInfo(t)
}

func TestRecipeIngredientDelete(t *testing.T) {
	for _, testCase := range testsRecipeIngredientData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := RecipeIngredientDelete(testCase.id, testCase.userId, testCase.entityId)

				if testCase.recipeIngredient != nil {
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
