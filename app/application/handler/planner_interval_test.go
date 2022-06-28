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
	testsPlannerIntervalData = testsPlannerInterval{
		{
			name:     "Test case with correct data",
			id:       nil,
			userId:   &testUserId,
			entityId: &uuidPlanner,
			plannerIntervalDTO: &DomainEntity.PlannerInterval{
				Name:   "PlannerInterval",
				Status: kind.PlannerIntervalStatusInActive,
			},
			toUpdatingPlannerIntervalDTO: &DomainEntity.PlannerInterval{
				Status: kind.PlannerIntervalStatusActive,
			},
		},
	}
)

type testsPlannerInterval []struct {
	name                         string
	id                           *uuid.UUID
	userId                       *uuid.UUID
	entityId                     *uuid.UUID
	plannerIntervalDTO           *DomainEntity.PlannerInterval
	toUpdatingPlannerIntervalDTO *DomainEntity.PlannerInterval
	plannerInterval              *DomainAggregate.PlannerInterval
}

func init() {
	InfrastructureService.PrepareCache()
	InfrastructureService.PreparePersistence()
}

func TestPlannerIntervalCreate(t *testing.T) {
	for index, testCase := range testsPlannerIntervalData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := PlannerIntervalCreate(testCase.userId, testCase.entityId, testCase.plannerIntervalDTO)

				assert.Nil(t, errorActual)

				testsPlannerIntervalData[index].plannerInterval = actual
				testsPlannerIntervalData[index].id = &actual.Entity.Id

				assert.NotNil(t, actual.Entity.Id)
				assert.Equal(t, testCase.plannerIntervalDTO.UserId, actual.Entity.UserId)
				assert.NotNil(t, actual.Entity.DateInsert)
				assert.NotNil(t, actual.Entity.DateUpdate)
				assert.Equal(t, testCase.plannerIntervalDTO.Name, actual.Entity.Name)
				assert.Equal(t, testCase.plannerIntervalDTO.Status, actual.Entity.Status)
			},
		)
	}
}

func TestPlannerIntervalsInfo(t *testing.T) {
	for _, testCase := range testsPlannerIntervalData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := PlannerIntervalsInfo(testCase.userId, testCase.entityId, nil)

				if testCase.plannerInterval != nil {
					assert.Nil(t, errorActual)

					for _, actualEntity := range actual {
						assert.Equal(t, testCase.plannerInterval.Entity.Id, actualEntity.Entity.Id)
						assert.Equal(t, testCase.plannerInterval.Entity.UserId, actualEntity.Entity.UserId)
						assert.Equal(t, testCase.plannerInterval.Entity.DateInsert.Format(time.UnixDate), actualEntity.Entity.DateInsert.Format(time.UnixDate))
						assert.Equal(t, testCase.plannerInterval.Entity.DateUpdate.Format(time.UnixDate), actualEntity.Entity.DateUpdate.Format(time.UnixDate))
						assert.Equal(t, testCase.plannerInterval.Entity.Name, actualEntity.Entity.Name)
						assert.Equal(t, testCase.plannerInterval.Entity.Status, actualEntity.Entity.Status)
					}
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestPlannerIntervalInfo(t *testing.T) {
	for _, testCase := range testsPlannerIntervalData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := PlannerIntervalInfo(testCase.id, testCase.userId, testCase.entityId, nil)

				if testCase.plannerInterval != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.plannerInterval.Entity.Id, actual.Entity.Id)
					assert.Equal(t, testCase.plannerInterval.Entity.UserId, actual.Entity.UserId)
					assert.Equal(t, testCase.plannerInterval.Entity.DateInsert.Format(time.UnixDate), actual.Entity.DateInsert.Format(time.UnixDate))
					assert.Equal(t, testCase.plannerInterval.Entity.DateUpdate.Format(time.UnixDate), actual.Entity.DateUpdate.Format(time.UnixDate))
					assert.Equal(t, testCase.plannerInterval.Entity.Name, actual.Entity.Name)
					assert.Equal(t, testCase.plannerInterval.Entity.Status, actual.Entity.Status)
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestPlannerIntervalUpdate(t *testing.T) {
	for index, testCase := range testsPlannerIntervalData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := PlannerIntervalUpdate(testCase.id, testCase.userId, testCase.entityId, testCase.toUpdatingPlannerIntervalDTO)

				if testCase.plannerInterval != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.toUpdatingPlannerIntervalDTO.Status, actual.Entity.Status)
					testsPlannerIntervalData[index].plannerInterval.Entity.Status = actual.Entity.Status
					testsPlannerIntervalData[index].plannerInterval.Entity.DateUpdate = actual.Entity.DateUpdate
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestPlannerIntervalDelete(t *testing.T) {
	for _, testCase := range testsPlannerIntervalData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := PlannerIntervalDelete(testCase.id, testCase.userId, testCase.entityId)

				if testCase.plannerInterval != nil {
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
