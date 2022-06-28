package entity

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestRecipe(t *testing.T) {
	tests := []struct {
		name        string
		json        string
		Id          uuid.UUID
		UserId      uuid.UUID
		DateInsert  time.Time
		DateUpdate  time.Time
		Name        string
		Description string
		Notes       string
		Status      kind.RecipeStatus
	}{
		{
			name:        "Test case with published recipe properties",
			json:        "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Recipe\",\"description\":\"Description\",\"notes\":\"Notes\",\"status\":\"published\"}\n",
			Id:          uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:      uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			DateInsert:  time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate:  time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Name:        "Recipe",
			Description: "Description",
			Notes:       "Notes",
			Status:      kind.RecipeStatusPublished,
		},
		{
			name:        "Test case with unpublished recipe properties",
			json:        "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Recipe\",\"description\":\"Description\",\"notes\":\"Notes\",\"status\":\"unpublished\"}\n",
			Id:          uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:      uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			DateInsert:  time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate:  time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Name:        "Recipe",
			Description: "Description",
			Notes:       "Notes",
			Status:      kind.RecipeStatusUnPublished,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				recipe := Recipe{
					Id:          testCase.Id,
					UserId:      testCase.UserId,
					DateInsert:  testCase.DateInsert,
					DateUpdate:  testCase.DateUpdate,
					Name:        testCase.Name,
					Description: testCase.Description,
					Notes:       testCase.Notes,
					Status:      testCase.Status,
				}
				assert.Equal(t, testCase.Id, recipe.Id)
				assert.Equal(t, testCase.UserId, recipe.UserId)
				assert.Equal(t, testCase.DateInsert, recipe.DateInsert)
				assert.Equal(t, testCase.DateUpdate, recipe.DateUpdate)
				assert.Equal(t, testCase.Name, recipe.Name)
				assert.Equal(t, testCase.Description, recipe.Description)
				assert.Equal(t, testCase.Notes, recipe.Notes)
				assert.Equal(t, testCase.Status, recipe.Status)

				reflectRecipe := reflect.ValueOf(recipe)

				for i := 0; i < reflectRecipe.NumField(); i++ {
					assert.False(t, reflectRecipe.Field(i).IsZero())
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(recipe)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}

func TestRecipeCategory(t *testing.T) {
	tests := []struct {
		name       string
		json       string
		Id         uuid.UUID
		UserId     uuid.UUID
		EntityId   uuid.UUID
		DeriveId   uuid.UUID
		DateInsert time.Time
		DateUpdate time.Time
		Status     kind.RecipeCategoryStatus
	}{
		{
			name:       "Test case with published recipe category properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000003\",\"derive_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"status\":\"published\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			DeriveId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
			DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Status:     kind.RecipeCategoryStatusPublished,
		},
		{
			name:       "Test case with unpublished recipe category properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000003\",\"derive_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"status\":\"unpublished\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			DeriveId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
			DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Status:     kind.RecipeCategoryStatusUnPublished,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				recipeCategory := RecipeCategory{
					Id:         testCase.Id,
					UserId:     testCase.UserId,
					EntityId:   testCase.EntityId,
					DeriveId:   testCase.DeriveId,
					DateInsert: testCase.DateInsert,
					DateUpdate: testCase.DateUpdate,
					Status:     testCase.Status,
				}
				assert.Equal(t, testCase.Id, recipeCategory.Id)
				assert.Equal(t, testCase.UserId, recipeCategory.UserId)
				assert.Equal(t, testCase.EntityId, recipeCategory.EntityId)
				assert.Equal(t, testCase.DeriveId, recipeCategory.DeriveId)
				assert.Equal(t, testCase.DateInsert, recipeCategory.DateInsert)
				assert.Equal(t, testCase.DateUpdate, recipeCategory.DateUpdate)
				assert.Equal(t, testCase.Status, recipeCategory.Status)

				reflectRecipeCategory := reflect.ValueOf(recipeCategory)

				for i := 0; i < reflectRecipeCategory.NumField(); i++ {
					assert.False(t, reflectRecipeCategory.Field(i).IsZero())
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(recipeCategory)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}

func TestRecipeIngredient(t *testing.T) {
	tests := []struct {
		name       string
		json       string
		Id         uuid.UUID
		UserId     uuid.UUID
		EntityId   uuid.UUID
		DeriveId   uuid.UUID
		DateInsert time.Time
		DateUpdate time.Time
		Name       string
		Status     kind.RecipeIngredientStatus
	}{
		{
			name:       "Test case with published recipe ingredient properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000003\",\"derive_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"RecipeIngredient\",\"status\":\"published\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			DeriveId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
			DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Name:       "RecipeIngredient",
			Status:     kind.RecipeIngredientStatusPublished,
		},
		{
			name:       "Test case with unpublished recipe ingredient properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000003\",\"derive_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"RecipeIngredient\",\"status\":\"unpublished\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			DeriveId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
			DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Name:       "RecipeIngredient",
			Status:     kind.RecipeIngredientStatusUnPublished,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				recipeIngredient := RecipeIngredient{
					Id:         testCase.Id,
					UserId:     testCase.UserId,
					EntityId:   testCase.EntityId,
					DeriveId:   testCase.DeriveId,
					DateInsert: testCase.DateInsert,
					DateUpdate: testCase.DateUpdate,
					Name:       testCase.Name,
					Status:     testCase.Status,
				}
				assert.Equal(t, testCase.Id, recipeIngredient.Id)
				assert.Equal(t, testCase.UserId, recipeIngredient.UserId)
				assert.Equal(t, testCase.EntityId, recipeIngredient.EntityId)
				assert.Equal(t, testCase.DeriveId, recipeIngredient.DeriveId)
				assert.Equal(t, testCase.DateInsert, recipeIngredient.DateInsert)
				assert.Equal(t, testCase.DateUpdate, recipeIngredient.DateUpdate)
				assert.Equal(t, testCase.Name, recipeIngredient.Name)
				assert.Equal(t, testCase.Status, recipeIngredient.Status)

				reflectRecipeIngredient := reflect.ValueOf(recipeIngredient)

				for i := 0; i < reflectRecipeIngredient.NumField(); i++ {
					assert.False(t, reflectRecipeIngredient.Field(i).IsZero())
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(recipeIngredient)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}

func TestRecipeMeasure(t *testing.T) {
	tests := []struct {
		name       string
		json       string
		Id         uuid.UUID
		UserId     uuid.UUID
		EntityId   uuid.UUID
		UnitId     uuid.UUID
		DateInsert time.Time
		DateUpdate time.Time
		Value      int64
		Status     kind.RecipeMeasureStatus
	}{
		{
			name:       "Test case with published recipe measure properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000003\",\"unit_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"value\":42,\"status\":\"published\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			UnitId:     uuid.MustParse("00000000-0000-0000-0000-000000000004"),
			DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Value:      42,
			Status:     kind.RecipeMeasureStatusPublished,
		},
		{
			name:       "Test case with unpublished recipe measure properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000003\",\"unit_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"value\":42,\"status\":\"unpublished\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			UnitId:     uuid.MustParse("00000000-0000-0000-0000-000000000004"),
			DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Value:      42,
			Status:     kind.RecipeMeasureStatusUnPublished,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				recipeMeasure := RecipeMeasure{
					Id:         testCase.Id,
					UserId:     testCase.UserId,
					EntityId:   testCase.EntityId,
					UnitId:     testCase.UnitId,
					DateInsert: testCase.DateInsert,
					DateUpdate: testCase.DateUpdate,
					Value:      testCase.Value,
					Status:     testCase.Status,
				}
				assert.Equal(t, testCase.Id, recipeMeasure.Id)
				assert.Equal(t, testCase.UserId, recipeMeasure.UserId)
				assert.Equal(t, testCase.EntityId, recipeMeasure.EntityId)
				assert.Equal(t, testCase.UnitId, recipeMeasure.UnitId)
				assert.Equal(t, testCase.DateInsert, recipeMeasure.DateInsert)
				assert.Equal(t, testCase.DateUpdate, recipeMeasure.DateUpdate)
				assert.Equal(t, testCase.Value, recipeMeasure.Value)
				assert.Equal(t, testCase.Status, recipeMeasure.Status)

				reflectRecipeMeasure := reflect.ValueOf(recipeMeasure)

				for i := 0; i < reflectRecipeMeasure.NumField(); i++ {
					assert.False(t, reflectRecipeMeasure.Field(i).IsZero())
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(recipeMeasure)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}

func TestRecipeProcess(t *testing.T) {
	tests := []struct {
		name        string
		json        string
		Id          uuid.UUID
		UserId      uuid.UUID
		EntityId    uuid.UUID
		DateInsert  time.Time
		DateUpdate  time.Time
		Name        string
		Description string
		Notes       string
		Status      kind.RecipeProcessStatus
	}{
		{
			name:        "Test case with published recipe process properties",
			json:        "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000003\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"RecipeProcess\",\"description\":\"Description\",\"notes\":\"Notes\",\"status\":\"published\"}\n",
			Id:          uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:      uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			EntityId:    uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			DateInsert:  time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate:  time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Name:        "RecipeProcess",
			Description: "Description",
			Notes:       "Notes",
			Status:      kind.RecipeProcessStatusPublished,
		},
		{
			name:        "Test case with unpublished recipe process properties",
			json:        "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000003\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"RecipeProcess\",\"description\":\"Description\",\"notes\":\"Notes\",\"status\":\"unpublished\"}\n",
			Id:          uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:      uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			EntityId:    uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			DateInsert:  time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate:  time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Name:        "RecipeProcess",
			Description: "Description",
			Notes:       "Notes",
			Status:      kind.RecipeProcessStatusUnPublished,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				recipeProcess := RecipeProcess{
					Id:          testCase.Id,
					UserId:      testCase.UserId,
					EntityId:    testCase.EntityId,
					DateInsert:  testCase.DateInsert,
					DateUpdate:  testCase.DateUpdate,
					Name:        testCase.Name,
					Description: testCase.Description,
					Notes:       testCase.Notes,
					Status:      testCase.Status,
				}
				assert.Equal(t, testCase.Id, recipeProcess.Id)
				assert.Equal(t, testCase.UserId, recipeProcess.UserId)
				assert.Equal(t, testCase.EntityId, recipeProcess.EntityId)
				assert.Equal(t, testCase.DateInsert, recipeProcess.DateInsert)
				assert.Equal(t, testCase.DateUpdate, recipeProcess.DateUpdate)
				assert.Equal(t, testCase.Name, recipeProcess.Name)
				assert.Equal(t, testCase.Description, recipeProcess.Description)
				assert.Equal(t, testCase.Notes, recipeProcess.Notes)
				assert.Equal(t, testCase.Status, recipeProcess.Status)

				reflectRecipeProcess := reflect.ValueOf(recipeProcess)

				for i := 0; i < reflectRecipeProcess.NumField(); i++ {
					assert.False(t, reflectRecipeProcess.Field(i).IsZero())
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(recipeProcess)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}
