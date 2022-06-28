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
	testsCategoryData = testsCategory{
		{
			name:     "Test case with correct data",
			id:       nil,
			userId:   &testUserId,
			entityId: nil,
			categoryDTO: &DomainEntity.Category{
				Name:   "Test case with correct data",
				Status: kind.CategoryStatusUnPublished,
			},
			toUpdatingCategoryDTO: &DomainEntity.Category{
				Status: kind.CategoryStatusPublished,
			},
		},
	}
)

type testsCategory []struct {
	name                  string
	id                    *uuid.UUID
	userId                *uuid.UUID
	entityId              *uuid.UUID
	categoryDTO           *DomainEntity.Category
	toUpdatingCategoryDTO *DomainEntity.Category
	category              *DomainAggregate.Category
}

func init() {
	InfrastructureService.PrepareCache()
	InfrastructureService.PreparePersistence()
}

func TestCategoryCreate(t *testing.T) {
	for index, testCase := range testsCategoryData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := CategoryCreate(testCase.userId, testCase.categoryDTO)

				assert.Nil(t, errorActual)

				testsCategoryData[index].category = actual
				testsCategoryData[index].id = &actual.Entity.Id

				assert.NotNil(t, actual.Entity.Id)
				assert.Equal(t, testCase.categoryDTO.UserId, actual.Entity.UserId)
				assert.NotNil(t, actual.Entity.DateInsert)
				assert.NotNil(t, actual.Entity.DateUpdate)
				assert.Equal(t, testCase.categoryDTO.Name, actual.Entity.Name)
				assert.Equal(t, testCase.categoryDTO.Status, actual.Entity.Status)
			},
		)
	}
}

func TestCategoriesInfo(t *testing.T) {
	for _, testCase := range testsCategoryData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := CategoriesInfo(testCase.userId, nil)

				if testCase.category != nil {
					assert.Nil(t, errorActual)

					for _, actualEntity := range actual {
						assert.Equal(t, testCase.category.Entity.Id, actualEntity.Entity.Id)
						assert.Equal(t, testCase.category.Entity.UserId, actualEntity.Entity.UserId)
						assert.Equal(t, testCase.category.Entity.DateInsert.Format(time.UnixDate), actualEntity.Entity.DateInsert.Format(time.UnixDate))
						assert.Equal(t, testCase.category.Entity.DateUpdate.Format(time.UnixDate), actualEntity.Entity.DateUpdate.Format(time.UnixDate))
						assert.Equal(t, testCase.category.Entity.Name, actualEntity.Entity.Name)
						assert.Equal(t, testCase.category.Entity.Status, actualEntity.Entity.Status)
					}
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestCategoryInfo(t *testing.T) {
	for _, testCase := range testsCategoryData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := CategoryInfo(testCase.id, testCase.userId, nil)

				if testCase.category != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.category.Entity.Id, actual.Entity.Id)
					assert.Equal(t, testCase.category.Entity.UserId, actual.Entity.UserId)
					assert.Equal(t, testCase.category.Entity.DateInsert.Format(time.UnixDate), actual.Entity.DateInsert.Format(time.UnixDate))
					assert.Equal(t, testCase.category.Entity.DateUpdate.Format(time.UnixDate), actual.Entity.DateUpdate.Format(time.UnixDate))
					assert.Equal(t, testCase.category.Entity.Name, actual.Entity.Name)
					assert.Equal(t, testCase.category.Entity.Status, actual.Entity.Status)
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestCategoryUpdate(t *testing.T) {
	for index, testCase := range testsCategoryData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := CategoryUpdate(testCase.id, testCase.userId, testCase.toUpdatingCategoryDTO)

				if testCase.category != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.toUpdatingCategoryDTO.Status, actual.Entity.Status)
					testsCategoryData[index].category.Entity.Status = actual.Entity.Status
					testsCategoryData[index].category.Entity.DateUpdate = actual.Entity.DateUpdate
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestGetCategoryEntity(t *testing.T) {
	TestCategoryInfo(t)
}

func TestCategoryDelete(t *testing.T) {
	for _, testCase := range testsCategoryData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := CategoryDelete(testCase.id, testCase.userId)

				if testCase.category != nil {
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
