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
	entityIdRecipeCategory  = uuid.New()
	testsRecipeCategoryData = testsRecipeCategory{
		{
			name:     "Test case with correct data",
			id:       nil,
			userId:   &testUserId,
			entityId: &entityIdRecipeCategory,
			deriveId: nil,
			categoryDTO: &DomainEntity.Category{
				Name:   "Test case with correct data",
				Status: kind.CategoryStatusUnPublished,
			},
			recipeCategoryDTO: &DomainEntity.RecipeCategory{
				Status: kind.RecipeCategoryStatusUnPublished,
			},
			toUpdatingRecipeCategoryDTO: &DomainEntity.RecipeCategory{
				Status: kind.RecipeCategoryStatusPublished,
			},
		},
	}
)

type testsRecipeCategory []struct {
	name                        string
	id                          *uuid.UUID
	userId                      *uuid.UUID
	entityId                    *uuid.UUID
	deriveId                    *uuid.UUID
	categoryDTO                 *DomainEntity.Category
	category                    *DomainAggregate.Category
	recipeCategoryDTO           *DomainEntity.RecipeCategory
	toUpdatingRecipeCategoryDTO *DomainEntity.RecipeCategory
	recipeCategory              *DomainAggregate.RecipeCategory
}

func init() {
	InfrastructureService.PrepareCache()
	InfrastructureService.PreparePersistence()
}

func TestRecipeCategoryCreate(t *testing.T) {
	for index, testCase := range testsRecipeCategoryData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actualCategory, errorCategoryCreateActual := CategoryCreate(testCase.userId, testCase.categoryDTO)

				assert.Nil(t, errorCategoryCreateActual)

				testsRecipeCategoryData[index].category = actualCategory

				testCase.recipeCategoryDTO.DeriveId = actualCategory.Entity.Id

				actual, errorActual := RecipeCategoryCreate(testCase.userId, testCase.entityId, testCase.recipeCategoryDTO)

				assert.Nil(t, errorActual)

				testsRecipeCategoryData[index].recipeCategory = actual
				testsRecipeCategoryData[index].id = &actual.Entity.Id
				testsRecipeCategoryData[index].entityId = &actual.Entity.EntityId

				assert.NotNil(t, actual.Entity.Id)
				assert.Equal(t, testCase.recipeCategoryDTO.UserId, actual.Entity.UserId)
				assert.Equal(t, testCase.recipeCategoryDTO.EntityId, actual.Entity.EntityId)
				assert.Equal(t, testCase.recipeCategoryDTO.DeriveId, actual.Entity.DeriveId)
				assert.NotNil(t, actual.Entity.DateInsert)
				assert.NotNil(t, actual.Entity.DateUpdate)
				assert.Equal(t, testCase.recipeCategoryDTO.Status, actual.Entity.Status)
			},
		)
	}
}

func TestRecipeCategoriesInfo(t *testing.T) {
	for _, testCase := range testsRecipeCategoryData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := RecipeCategoriesInfo(testCase.userId, testCase.entityId, nil)

				if testCase.recipeCategory != nil {
					assert.Nil(t, errorActual)

					for _, actualEntity := range actual {
						assert.Equal(t, testCase.recipeCategory.Entity.Id, actualEntity.Entity.Id)
						assert.Equal(t, testCase.recipeCategory.Entity.UserId, actualEntity.Entity.UserId)
						assert.Equal(t, testCase.recipeCategory.Entity.EntityId, actualEntity.Entity.EntityId)
						assert.Equal(t, testCase.recipeCategory.Entity.DeriveId, actualEntity.Entity.DeriveId)
						assert.Equal(t, testCase.recipeCategory.Entity.DateInsert.Format(time.UnixDate), actualEntity.Entity.DateInsert.Format(time.UnixDate))
						assert.Equal(t, testCase.recipeCategory.Entity.DateUpdate.Format(time.UnixDate), actualEntity.Entity.DateUpdate.Format(time.UnixDate))
						assert.Equal(t, testCase.recipeCategory.Entity.Status, actualEntity.Entity.Status)
					}
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestRecipeCategoryInfo(t *testing.T) {
	for _, testCase := range testsRecipeCategoryData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actualEntity, errorActual := RecipeCategoryInfo(testCase.id, testCase.userId, testCase.entityId, nil)

				if testCase.recipeCategory != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.recipeCategory.Entity.Id, actualEntity.Entity.Id)
					assert.Equal(t, testCase.recipeCategory.Entity.UserId, actualEntity.Entity.UserId)
					assert.Equal(t, testCase.recipeCategory.Entity.EntityId, actualEntity.Entity.EntityId)
					assert.Equal(t, testCase.recipeCategory.Entity.DeriveId, actualEntity.Entity.DeriveId)
					assert.Equal(t, testCase.recipeCategory.Entity.DateInsert.Format(time.UnixDate), actualEntity.Entity.DateInsert.Format(time.UnixDate))
					assert.Equal(t, testCase.recipeCategory.Entity.DateUpdate.Format(time.UnixDate), actualEntity.Entity.DateUpdate.Format(time.UnixDate))
					assert.Equal(t, testCase.recipeCategory.Entity.Status, actualEntity.Entity.Status)
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestRecipeCategoryUpdate(t *testing.T) {
	for index, testCase := range testsRecipeCategoryData {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				actual, errorActual := RecipeCategoryUpdate(testCase.id, testCase.userId, testCase.entityId, testCase.toUpdatingRecipeCategoryDTO)

				if testCase.recipeCategory != nil {
					assert.Nil(t, errorActual)
					assert.Equal(t, testCase.toUpdatingRecipeCategoryDTO.Status, actual.Entity.Status)
					testsRecipeCategoryData[index].recipeCategory.Entity.Status = actual.Entity.Status
					testsRecipeCategoryData[index].recipeCategory.Entity.DateUpdate = actual.Entity.DateUpdate
				} else {
					assert.NotNil(t, errorActual)
				}
			},
		)
	}
}

func TestGetRecipeCategoryEntity(t *testing.T) {
	TestRecipeCategoryInfo(t)
}

func TestRecipeCategoryDelete(t *testing.T) {
	for index, testCase := range testsRecipeCategoryData {
		t.Run(
			testCase.name,
			func(t *testing.T) {

				actualCategory, errorCategoryDeleteActual := CategoryDelete(&testsRecipeCategoryData[index].category.Entity.Id, testCase.userId)
				actualRecipeCategory, errorRecipeCategoryDeleteActual := RecipeCategoryDelete(testCase.id, testCase.userId, testCase.entityId)

				if testCase.recipeCategory != nil {
					assert.Nil(t, errorCategoryDeleteActual)
					assert.Nil(t, errorRecipeCategoryDeleteActual)
					assert.True(t, actualCategory)
					assert.True(t, actualRecipeCategory)
				} else {
					assert.NotNil(t, errorCategoryDeleteActual)
					assert.NotNil(t, errorRecipeCategoryDeleteActual)
					assert.False(t, actualCategory)
					assert.False(t, actualRecipeCategory)
				}
			},
		)
	}
}
