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
	uuidPlanner      = uuid.New()
	testsPlannerData = testsPlanner{
		{
			name:   "Test case with correct data",
			id:     nil,
			userId: &testUserId,
			plannerDTO: &DomainEntity.Planner{
				Name:   "Planner",
				Status: kind.PlannerStatusInActive,
			},
			toUpdatingPlannerDTO: &DomainEntity.Planner{
				Status: kind.PlannerStatusActive,
			},
		},
	}
)

type testsPlanner []struct {
	name                 string
	id                   *uuid.UUID
	userId               *uuid.UUID
	plannerDTO           *DomainEntity.Planner
	toUpdatingPlannerDTO *DomainEntity.Planner
	planner              *DomainAggregate.Planner
}

func init() {
	InfrastructureService.PrepareCache()
	InfrastructureService.PreparePersistence()
}

func TestPlannerCreate(t *testing.T) {
	for index, testCase := range testsPlannerData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := PlannerCreate(testCase.userId, testCase.plannerDTO)

				assert.Nil(t, errorActual)

				testsPlannerData[index].planner = actual
				testsPlannerData[index].id = &actual.Entity.Id

				assert.NotNil(t, actual.Entity.Id)
				assert.Equal(t, testCase.plannerDTO.UserId, actual.Entity.UserId)
				assert.NotNil(t, actual.Entity.DateInsert)
				assert.NotNil(t, actual.Entity.DateUpdate)
				assert.Equal(t, testCase.plannerDTO.Name, actual.Entity.Name)
				assert.Equal(t, testCase.plannerDTO.Status, actual.Entity.Status)
			},
		)
	}
}

func TestPlannersInfo(t *testing.T) {
	for _, testCase := range testsPlannerData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := PlannersInfo(testCase.userId, nil)

				if testCase.planner != nil {
					assert.Nil(t, errorActual)

					for _, actualEntity := range actual {
						assert.Equal(t, testCase.planner.Entity.Id, actualEntity.Entity.Id)
						assert.Equal(t, testCase.planner.Entity.UserId, actualEntity.Entity.UserId)
						assert.Equal(t, testCase.planner.Entity.DateInsert.Format(time.UnixDate), actualEntity.Entity.DateInsert.Format(time.UnixDate))
						assert.Equal(t, testCase.planner.Entity.DateUpdate.Format(time.UnixDate), actualEntity.Entity.DateUpdate.Format(time.UnixDate))
						assert.Equal(t, testCase.planner.Entity.Name, actualEntity.Entity.Name)
						assert.Equal(t, testCase.planner.Entity.Status, actualEntity.Entity.Status)
					}
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestPlannerInfo(t *testing.T) {
	for _, testCase := range testsPlannerData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := PlannerInfo(testCase.id, testCase.userId, nil)

				if testCase.planner != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.planner.Entity.Id, actual.Entity.Id)
					assert.Equal(t, testCase.planner.Entity.UserId, actual.Entity.UserId)
					assert.Equal(t, testCase.planner.Entity.DateInsert.Format(time.UnixDate), actual.Entity.DateInsert.Format(time.UnixDate))
					assert.Equal(t, testCase.planner.Entity.DateUpdate.Format(time.UnixDate), actual.Entity.DateUpdate.Format(time.UnixDate))
					assert.Equal(t, testCase.planner.Entity.Name, actual.Entity.Name)
					assert.Equal(t, testCase.planner.Entity.Status, actual.Entity.Status)
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestPlannerUpdate(t *testing.T) {
	for index, testCase := range testsPlannerData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := PlannerUpdate(testCase.id, testCase.userId, testCase.toUpdatingPlannerDTO)

				if testCase.planner != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.toUpdatingPlannerDTO.Status, actual.Entity.Status)
					testsPlannerData[index].planner.Entity.Status = actual.Entity.Status
					testsPlannerData[index].planner.Entity.DateUpdate = actual.Entity.DateUpdate
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestPlannerDelete(t *testing.T) {
	for _, testCase := range testsPlannerData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := PlannerDelete(testCase.id, testCase.userId)

				if testCase.planner != nil {
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
