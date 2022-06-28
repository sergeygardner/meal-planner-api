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
	entityIdRecipeProcess  = uuid.New()
	testsRecipeProcessData = testsRecipeProcess{
		{
			name:     "Test case with correct data",
			id:       nil,
			userId:   &testUserId,
			entityId: &entityIdRecipeProcess,
			recipeProcessDTO: &DomainEntity.RecipeProcess{
				Name:   "Test case with correct data",
				Status: kind.RecipeProcessStatusUnPublished,
			},
			toUpdatingRecipeProcessDTO: &DomainEntity.RecipeProcess{
				Status: kind.RecipeProcessStatusPublished,
			},
		},
	}
)

type testsRecipeProcess []struct {
	name                       string
	id                         *uuid.UUID
	userId                     *uuid.UUID
	entityId                   *uuid.UUID
	recipeProcessDTO           *DomainEntity.RecipeProcess
	toUpdatingRecipeProcessDTO *DomainEntity.RecipeProcess
	recipeProcess              *DomainAggregate.RecipeProcess
}

func init() {
	InfrastructureService.PrepareCache()
	InfrastructureService.PreparePersistence()
}

func TestRecipeProcessCreate(t *testing.T) {
	for index, testCase := range testsRecipeProcessData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := RecipeProcessCreate(testCase.userId, testCase.entityId, testCase.recipeProcessDTO)

				assert.Nil(t, errorActual)

				testsRecipeProcessData[index].recipeProcess = actual
				testsRecipeProcessData[index].id = &actual.Entity.Id
				testsRecipeProcessData[index].entityId = &actual.Entity.EntityId

				assert.NotNil(t, actual.Entity.Id)
				assert.Equal(t, testCase.recipeProcessDTO.UserId, actual.Entity.UserId)
				assert.Equal(t, testCase.recipeProcessDTO.EntityId, actual.Entity.EntityId)
				assert.NotNil(t, actual.Entity.DateInsert)
				assert.NotNil(t, actual.Entity.DateUpdate)
				assert.Equal(t, testCase.recipeProcessDTO.Name, actual.Entity.Name)
				assert.Equal(t, testCase.recipeProcessDTO.Status, actual.Entity.Status)
			},
		)
	}
}

func TestRecipeProcessesInfo(t *testing.T) {
	for _, testCase := range testsRecipeProcessData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := RecipeProcessesInfo(testCase.userId, testCase.entityId, nil)

				if testCase.recipeProcess != nil {
					assert.Nil(t, errorActual)

					for _, actualEntity := range actual {
						assert.Equal(t, testCase.recipeProcess.Entity.Id, actualEntity.Entity.Id)
						assert.Equal(t, testCase.recipeProcess.Entity.UserId, actualEntity.Entity.UserId)
						assert.Equal(t, testCase.recipeProcess.Entity.EntityId, actualEntity.Entity.EntityId)
						assert.Equal(t, testCase.recipeProcess.Entity.DateInsert.Format(time.UnixDate), actualEntity.Entity.DateInsert.Format(time.UnixDate))
						assert.Equal(t, testCase.recipeProcess.Entity.DateUpdate.Format(time.UnixDate), actualEntity.Entity.DateUpdate.Format(time.UnixDate))
						assert.Equal(t, testCase.recipeProcess.Entity.Name, actualEntity.Entity.Name)
						assert.Equal(t, testCase.recipeProcess.Entity.Status, actualEntity.Entity.Status)
					}
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestRecipeProcessInfo(t *testing.T) {
	for _, testCase := range testsRecipeProcessData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actualEntity, errorActual := RecipeProcessInfo(testCase.id, testCase.userId, testCase.entityId, nil)

				if testCase.recipeProcess != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.recipeProcess.Entity.Id, actualEntity.Entity.Id)
					assert.Equal(t, testCase.recipeProcess.Entity.UserId, actualEntity.Entity.UserId)
					assert.Equal(t, testCase.recipeProcess.Entity.EntityId, actualEntity.Entity.EntityId)
					assert.Equal(t, testCase.recipeProcess.Entity.DateInsert.Format(time.UnixDate), actualEntity.Entity.DateInsert.Format(time.UnixDate))
					assert.Equal(t, testCase.recipeProcess.Entity.DateUpdate.Format(time.UnixDate), actualEntity.Entity.DateUpdate.Format(time.UnixDate))
					assert.Equal(t, testCase.recipeProcess.Entity.Name, actualEntity.Entity.Name)
					assert.Equal(t, testCase.recipeProcess.Entity.Status, actualEntity.Entity.Status)
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestRecipeProcessUpdate(t *testing.T) {
	for index, testCase := range testsRecipeProcessData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := RecipeProcessUpdate(testCase.id, testCase.userId, testCase.entityId, testCase.toUpdatingRecipeProcessDTO)

				if testCase.recipeProcess != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.toUpdatingRecipeProcessDTO.Status, actual.Entity.Status)
					testsRecipeProcessData[index].recipeProcess.Entity.Status = actual.Entity.Status
					testsRecipeProcessData[index].recipeProcess.Entity.DateUpdate = actual.Entity.DateUpdate
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestGetProcessEntity(t *testing.T) {
	TestRecipeProcessInfo(t)
}

func TestRecipeProcessDelete(t *testing.T) {
	for _, testCase := range testsRecipeProcessData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := RecipeProcessDelete(testCase.id, testCase.userId, testCase.entityId)

				if testCase.recipeProcess != nil {
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
