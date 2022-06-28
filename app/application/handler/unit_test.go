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
	testsUnitData = testsUnit{
		{
			name: "Test case with correct data",
			id:   nil,
			unitDTO: &DomainEntity.Unit{
				Name:   "Test case with correct data",
				Status: kind.UnitStatusUnPublished,
			},
			toUpdatingUnitDTO: &DomainEntity.Unit{
				Status: kind.UnitStatusPublished,
			},
		},
	}
)

type testsUnit []struct {
	name              string
	id                *uuid.UUID
	unitDTO           *DomainEntity.Unit
	toUpdatingUnitDTO *DomainEntity.Unit
	unit              *DomainEntity.Unit
}

func init() {
	InfrastructureService.PrepareCache()
	InfrastructureService.PreparePersistence()
}

func TestUnitCreate(t *testing.T) {
	for index, testCase := range testsUnitData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := UnitCreate(testCase.unitDTO)

				assert.Nil(t, errorActual)

				testsUnitData[index].unit = actual
				testsUnitData[index].id = &actual.Id

				assert.NotNil(t, actual.Id)
				assert.NotNil(t, actual.DateInsert)
				assert.NotNil(t, actual.DateUpdate)
				assert.Equal(t, testCase.unitDTO.Name, actual.Name)
				assert.Equal(t, testCase.unitDTO.Status, actual.Status)
			},
		)
	}
}

func TestUnitsInfo(t *testing.T) {
	for _, testCase := range testsUnitData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := UnitsInfo(nil)

				if testCase.unit != nil {
					assert.Nil(t, errorActual)

					for _, actualEntity := range actual {
						if testCase.unit.Id == actualEntity.Id {
							//assert.Equal(t, testCase.unit.Id, actualEntity.Id)
							assert.Equal(t, testCase.unit.DateInsert.Format(time.UnixDate), actualEntity.DateInsert.Format(time.UnixDate))
							assert.Equal(t, testCase.unit.DateUpdate.Format(time.UnixDate), actualEntity.DateUpdate.Format(time.UnixDate))
							assert.Equal(t, testCase.unit.Name, actualEntity.Name)
							assert.Equal(t, testCase.unit.Status, actualEntity.Status)
						}
					}
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestUnitInfo(t *testing.T) {
	for _, testCase := range testsUnitData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := UnitInfo(testCase.id, nil)

				if testCase.unit != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.unit.Id, actual.Id)
					assert.Equal(t, testCase.unit.DateInsert.Format(time.UnixDate), actual.DateInsert.Format(time.UnixDate))
					assert.Equal(t, testCase.unit.DateUpdate.Format(time.UnixDate), actual.DateUpdate.Format(time.UnixDate))
					assert.Equal(t, testCase.unit.Name, actual.Name)
					assert.Equal(t, testCase.unit.Status, actual.Status)
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestUnitUpdate(t *testing.T) {
	for index, testCase := range testsUnitData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := UnitUpdate(testCase.id, testCase.toUpdatingUnitDTO)

				if testCase.unit != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.toUpdatingUnitDTO.Status, actual.Status)
					testsUnitData[index].unit.Status = actual.Status
					testsUnitData[index].unit.DateUpdate = actual.DateUpdate
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestGetUnitAggregate(t *testing.T) {
	TestUnitInfo(t)
}

func TestUnitDelete(t *testing.T) {
	for _, testCase := range testsUnitData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := UnitDelete(testCase.id)

				if testCase.unit != nil {
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
