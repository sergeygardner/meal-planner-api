package handler

import (
	"github.com/google/uuid"
	DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	InfrastructureService "github.com/sergeygardner/meal-planner-api/infrastructure/service"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	testsIngredientData = testsIngredient{
		{
			name:     "Test case with correct data",
			id:       nil,
			userId:   &testUserId,
			entityId: nil,
			ingredientDTO: &DomainEntity.Ingredient{
				Name:   "Test case with correct data",
				Status: kind.IngredientStatusUnPublished,
			},
			toUpdatingIngredientDTO: &DomainEntity.Ingredient{
				Status: kind.IngredientStatusPublished,
			},
		},
	}
)

type testsIngredient []struct {
	name                    string
	id                      *uuid.UUID
	userId                  *uuid.UUID
	entityId                *uuid.UUID
	ingredientDTO           *DomainEntity.Ingredient
	toUpdatingIngredientDTO *DomainEntity.Ingredient
	ingredient              *DomainEntity.Ingredient
}

func init() {
	InfrastructureService.PrepareCache()
	InfrastructureService.PreparePersistence()
}

func TestIngredientCreate(t *testing.T) {
	for index, testCase := range testsIngredientData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := IngredientCreate(testCase.userId, testCase.ingredientDTO)

				assert.Nil(t, errorActual)

				testsIngredientData[index].ingredient = actual
				testsIngredientData[index].id = &actual.Id

				assert.NotNil(t, actual.Id)
				assert.Equal(t, testCase.ingredientDTO.UserId, actual.UserId)
				assert.NotNil(t, actual.DateInsert)
				assert.NotNil(t, actual.DateUpdate)
				assert.Equal(t, testCase.ingredientDTO.Name, actual.Name)
				assert.Equal(t, testCase.ingredientDTO.Status, actual.Status)
			},
		)
	}
}

func TestIngredientsInfo(t *testing.T) {
	for _, testCase := range testsIngredientData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := IngredientsInfo(testCase.userId, nil)

				if testCase.ingredient != nil {
					assert.Nil(t, errorActual)

					for _, actualEntity := range actual {
						assert.Equal(t, testCase.ingredient.Id, actualEntity.Id)
						assert.Equal(t, testCase.ingredient.UserId, actualEntity.UserId)
						assert.Equal(t, testCase.ingredient.DateInsert.Format(time.UnixDate), actualEntity.DateInsert.Format(time.UnixDate))
						assert.Equal(t, testCase.ingredient.DateUpdate.Format(time.UnixDate), actualEntity.DateUpdate.Format(time.UnixDate))
						assert.Equal(t, testCase.ingredient.Name, actualEntity.Name)
						assert.Equal(t, testCase.ingredient.Status, actualEntity.Status)
					}
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestIngredientInfo(t *testing.T) {
	for _, testCase := range testsIngredientData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := IngredientInfo(testCase.id, testCase.userId, nil)

				if testCase.ingredient != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.ingredient.Id, actual.Id)
					assert.Equal(t, testCase.ingredient.UserId, actual.UserId)
					assert.Equal(t, testCase.ingredient.DateInsert.Format(time.UnixDate), actual.DateInsert.Format(time.UnixDate))
					assert.Equal(t, testCase.ingredient.DateUpdate.Format(time.UnixDate), actual.DateUpdate.Format(time.UnixDate))
					assert.Equal(t, testCase.ingredient.Name, actual.Name)
					assert.Equal(t, testCase.ingredient.Status, actual.Status)
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestIngredientUpdate(t *testing.T) {
	for index, testCase := range testsIngredientData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := IngredientUpdate(testCase.id, testCase.userId, testCase.toUpdatingIngredientDTO)

				if testCase.ingredient != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.toUpdatingIngredientDTO.Status, actual.Status)
					testsIngredientData[index].ingredient.Status = actual.Status
					testsIngredientData[index].ingredient.DateUpdate = actual.DateUpdate
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestGetIngredientEntity(t *testing.T) {
	TestIngredientInfo(t)
}

func TestIngredientDelete(t *testing.T) {
	for _, testCase := range testsIngredientData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := IngredientDelete(testCase.id, testCase.userId)

				if testCase.ingredient != nil {
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
