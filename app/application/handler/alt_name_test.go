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
	entityIdAltName  = uuid.New()
	testsAltNameData = testsAltName{
		{
			name:     "Test case with correct data",
			id:       nil,
			userId:   &testUserId,
			entityId: &entityIdAltName,
			altNameDTO: &DomainEntity.AltName{
				Name:   "Test case with correct data",
				Status: kind.AltNameStatusUnPublished,
			},
			toUpdatingAltNameDTO: &DomainEntity.AltName{
				Status: kind.AltNameStatusPublished,
			},
		},
	}
)

type testsAltName []struct {
	name                 string
	id                   *uuid.UUID
	userId               *uuid.UUID
	entityId             *uuid.UUID
	altNameDTO           *DomainEntity.AltName
	toUpdatingAltNameDTO *DomainEntity.AltName
	altName              *DomainEntity.AltName
}

func init() {
	InfrastructureService.PrepareCache()
	InfrastructureService.PreparePersistence()
}

func TestAltNameCreate(t *testing.T) {
	for index, testCase := range testsAltNameData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := AltNameCreate(testCase.userId, testCase.entityId, testCase.altNameDTO)

				assert.Nil(t, errorActual)

				testsAltNameData[index].altName = actual
				testsAltNameData[index].id = &actual.Id
				testsAltNameData[index].entityId = &actual.EntityId

				assert.NotNil(t, actual.Id)
				assert.Equal(t, testCase.altNameDTO.UserId, actual.UserId)
				assert.Equal(t, testCase.altNameDTO.EntityId, actual.EntityId)
				assert.NotNil(t, actual.DateInsert)
				assert.NotNil(t, actual.DateUpdate)
				assert.Equal(t, testCase.altNameDTO.Name, actual.Name)
				assert.Equal(t, testCase.altNameDTO.Status, actual.Status)
			},
		)
	}
}

func TestAltNamesInfo(t *testing.T) {
	for _, testCase := range testsAltNameData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := AltNamesInfo(testCase.userId, testCase.entityId, nil)

				if testCase.altName != nil {
					assert.Nil(t, errorActual)

					for _, actualEntity := range actual {
						assert.Equal(t, testCase.altName.Id, actualEntity.Id)
						assert.Equal(t, testCase.altName.UserId, actualEntity.UserId)
						assert.Equal(t, testCase.altName.EntityId, actualEntity.EntityId)
						assert.Equal(t, testCase.altName.DateInsert.Format(time.UnixDate), actualEntity.DateInsert.Format(time.UnixDate))
						assert.Equal(t, testCase.altName.DateUpdate.Format(time.UnixDate), actualEntity.DateUpdate.Format(time.UnixDate))
						assert.Equal(t, testCase.altName.Name, actualEntity.Name)
						assert.Equal(t, testCase.altName.Status, actualEntity.Status)
					}
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestAltNameInfo(t *testing.T) {
	for _, testCase := range testsAltNameData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := AltNameInfo(testCase.id, testCase.userId, testCase.entityId, nil)

				if testCase.altName != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.altName.Id, actual.Id)
					assert.Equal(t, testCase.altName.UserId, actual.UserId)
					assert.Equal(t, testCase.altName.Id, actual.Id)
					assert.Equal(t, testCase.altName.DateInsert.Format(time.UnixDate), actual.DateInsert.Format(time.UnixDate))
					assert.Equal(t, testCase.altName.DateUpdate.Format(time.UnixDate), actual.DateUpdate.Format(time.UnixDate))
					assert.Equal(t, testCase.altName.Name, actual.Name)
					assert.Equal(t, testCase.altName.Status, actual.Status)
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestAltNameUpdate(t *testing.T) {
	for index, testCase := range testsAltNameData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := AltNameUpdate(testCase.id, testCase.userId, testCase.entityId, testCase.toUpdatingAltNameDTO)

				if testCase.altName != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.toUpdatingAltNameDTO.Status, actual.Status)
					testsAltNameData[index].altName.Status = actual.Status
					testsAltNameData[index].altName.DateUpdate = actual.DateUpdate
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestGetAltNameEntity(t *testing.T) {
	TestAltNameInfo(t)
}

func TestAltNameDelete(t *testing.T) {
	for _, testCase := range testsAltNameData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := AltNameDelete(testCase.id, testCase.userId, testCase.entityId)

				if testCase.altName != nil {
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
