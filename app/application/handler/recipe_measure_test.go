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
	entityIdRecipeMeasure  = uuid.New()
	testsRecipeMeasureData = testsRecipeMeasure{
		{
			name:     "Test case with correct data",
			id:       nil,
			userId:   &testUserId,
			entityId: &entityIdRecipeMeasure,
			recipeMeasureDTO: &DomainEntity.RecipeMeasure{
				Status: kind.RecipeMeasureStatusUnPublished,
			},
			toUpdatingRecipeMeasureDTO: &DomainEntity.RecipeMeasure{
				Status: kind.RecipeMeasureStatusPublished,
			},
		},
	}
)

type testsRecipeMeasure []struct {
	name                       string
	id                         *uuid.UUID
	userId                     *uuid.UUID
	entityId                   *uuid.UUID
	recipeMeasureDTO           *DomainEntity.RecipeMeasure
	toUpdatingRecipeMeasureDTO *DomainEntity.RecipeMeasure
	recipeMeasure              *DomainAggregate.RecipeMeasure
}

func init() {
	InfrastructureService.PrepareCache()
	InfrastructureService.PreparePersistence()
}

func TestRecipeMeasureCreate(t *testing.T) {
	for index, testCase := range testsRecipeMeasureData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := RecipeMeasureCreate(testCase.userId, testCase.entityId, testCase.recipeMeasureDTO)

				assert.Nil(t, errorActual)

				testsRecipeMeasureData[index].recipeMeasure = actual
				testsRecipeMeasureData[index].id = &actual.Entity.Id
				testsRecipeMeasureData[index].entityId = &actual.Entity.EntityId

				assert.NotNil(t, actual.Entity.Id)
				assert.Equal(t, testCase.recipeMeasureDTO.UserId, actual.Entity.UserId)
				assert.Equal(t, testCase.recipeMeasureDTO.EntityId, actual.Entity.EntityId)
				assert.NotNil(t, actual.Entity.DateInsert)
				assert.NotNil(t, actual.Entity.DateUpdate)
				assert.Equal(t, testCase.recipeMeasureDTO.Status, actual.Entity.Status)
			},
		)
	}
}

func TestRecipeMeasuresInfo(t *testing.T) {
	for _, testCase := range testsRecipeMeasureData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := RecipeMeasuresInfo(testCase.userId, testCase.entityId, nil)

				if testCase.recipeMeasure != nil {
					assert.Nil(t, errorActual)

					for _, actualEntity := range actual {
						assert.Equal(t, testCase.recipeMeasure.Entity.Id, actualEntity.Entity.Id)
						assert.Equal(t, testCase.recipeMeasure.Entity.UserId, actualEntity.Entity.UserId)
						assert.Equal(t, testCase.recipeMeasure.Entity.EntityId, actualEntity.Entity.EntityId)
						assert.Equal(t, testCase.recipeMeasure.Entity.DateInsert.Format(time.UnixDate), actualEntity.Entity.DateInsert.Format(time.UnixDate))
						assert.Equal(t, testCase.recipeMeasure.Entity.DateUpdate.Format(time.UnixDate), actualEntity.Entity.DateUpdate.Format(time.UnixDate))
						assert.Equal(t, testCase.recipeMeasure.Entity.Status, actualEntity.Entity.Status)
					}
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestRecipeMeasureInfo(t *testing.T) {
	for _, testCase := range testsRecipeMeasureData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := RecipeMeasureInfo(testCase.id, testCase.userId, testCase.entityId, nil)

				if testCase.recipeMeasure != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.recipeMeasure.Entity.Id, actual.Entity.Id)
					assert.Equal(t, testCase.recipeMeasure.Entity.UserId, actual.Entity.UserId)
					assert.Equal(t, testCase.recipeMeasure.Entity.EntityId, actual.Entity.EntityId)
					assert.Equal(t, testCase.recipeMeasure.Entity.DateInsert.Format(time.UnixDate), actual.Entity.DateInsert.Format(time.UnixDate))
					assert.Equal(t, testCase.recipeMeasure.Entity.DateUpdate.Format(time.UnixDate), actual.Entity.DateUpdate.Format(time.UnixDate))
					assert.Equal(t, testCase.recipeMeasure.Entity.Status, actual.Entity.Status)
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestRecipeMeasureUpdate(t *testing.T) {
	for index, testCase := range testsRecipeMeasureData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := RecipeMeasureUpdate(testCase.id, testCase.userId, testCase.entityId, testCase.toUpdatingRecipeMeasureDTO)

				if testCase.recipeMeasure != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.toUpdatingRecipeMeasureDTO.Status, actual.Entity.Status)
					testsRecipeMeasureData[index].recipeMeasure.Entity.Status = actual.Entity.Status
					testsRecipeMeasureData[index].recipeMeasure.Entity.DateUpdate = actual.Entity.DateUpdate
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestGetMeasureEntity(t *testing.T) {
	TestRecipeMeasureInfo(t)
}

func TestRecipeMeasureDelete(t *testing.T) {
	for _, testCase := range testsRecipeMeasureData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := RecipeMeasureDelete(testCase.id, testCase.userId, testCase.entityId)

				if testCase.recipeMeasure != nil {
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
