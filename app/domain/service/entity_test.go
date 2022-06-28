package service

import (
	"bytes"
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	"github.com/stretchr/testify/assert"
	"testing"
	"testing/iotest"
	"time"
)

func TestCreateEntityFromRecipeUpdate(t *testing.T) {
	tests := []struct {
		name     string
		JSON     string
		Expected entity.Recipe
	}{
		{
			name: "Test case for CreateEntityFromRecipeUpdate with status published",
			JSON: "{\"name\":\"name\",\"description\":\"description\",\"notes\":\"notes\",\"status\":\"published\"}",
			Expected: entity.Recipe{
				Name:        "name",
				Description: "description",
				Notes:       "notes",
				Status:      kind.RecipeStatusPublished,
			},
		},
		{
			name: "Test case for CreateEntityFromRecipeUpdate with status unpublished",
			JSON: "{\"name\":\"name\",\"description\":\"description\",\"notes\":\"notes\",\"status\":\"unpublished\"}",
			Expected: entity.Recipe{
				Name:        "name",
				Description: "description",
				Notes:       "notes",
				Status:      kind.RecipeStatusUnPublished,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				buffer := new(bytes.Buffer)
				buffer.WriteString(testCase.JSON)
				oneByteReader := iotest.OneByteReader(buffer)
				recipeUpdate, errorCreateDTOFromRecipeUpdate := CreateEntityFromRecipeUpdate(oneByteReader)

				assert.Equal(t, testCase.Expected, recipeUpdate)
				assert.Nil(t, errorCreateDTOFromRecipeUpdate)
			},
		)
	}
}

func TestCreateEntityFromRecipeCategoryUpdate(t *testing.T) {
	tests := []struct {
		name     string
		JSON     string
		Expected entity.RecipeCategory
	}{
		{
			name: "Test case for CreateEntityFromRecipeCategoryUpdate with status published",
			JSON: "{\"derive_id\":\"00000000-0000-0000-0000-000000000001\",\"status\":\"published\"}",
			Expected: entity.RecipeCategory{
				DeriveId: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Status:   kind.RecipeCategoryStatusPublished,
			},
		},
		{
			name: "Test case for CreateEntityFromRecipeCategoryUpdate with status unpublished",
			JSON: "{\"derive_id\":\"00000000-0000-0000-0000-000000000001\",\"status\":\"unpublished\"}",
			Expected: entity.RecipeCategory{
				DeriveId: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Status:   kind.RecipeCategoryStatusUnPublished,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				buffer := new(bytes.Buffer)
				buffer.WriteString(testCase.JSON)
				oneByteReader := iotest.OneByteReader(buffer)
				recipeCategoryUpdate, errorCreateDTOFromRecipeCategoryUpdate := CreateEntityFromRecipeCategoryUpdate(oneByteReader)

				assert.Equal(t, testCase.Expected, recipeCategoryUpdate)
				assert.Nil(t, errorCreateDTOFromRecipeCategoryUpdate)
			},
		)
	}
}

func TestCreateEntityFromRecipeIngredientUpdate(t *testing.T) {
	tests := []struct {
		name     string
		JSON     string
		Expected entity.RecipeIngredient
	}{
		{
			name: "Test case for CreateEntityFromRecipeIngredientUpdate with status published",
			JSON: "{\"derive_id\":\"00000000-0000-0000-0000-000000000001\",\"name\":\"name\",\"status\":\"published\"}",
			Expected: entity.RecipeIngredient{
				DeriveId: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Name:     "name",
				Status:   kind.RecipeIngredientStatusPublished,
			},
		},
		{
			name: "Test case for CreateEntityFromRecipeIngredientUpdate with status unpublished",
			JSON: "{\"derive_id\":\"00000000-0000-0000-0000-000000000001\",\"name\":\"name\",\"status\":\"unpublished\"}",
			Expected: entity.RecipeIngredient{
				DeriveId: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Name:     "name",
				Status:   kind.RecipeIngredientStatusUnPublished,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				buffer := new(bytes.Buffer)
				buffer.WriteString(testCase.JSON)
				oneByteReader := iotest.OneByteReader(buffer)
				recipeIngredientUpdate, errorCreateDTOFromRecipeIngredientUpdate := CreateEntityFromRecipeIngredientUpdate(oneByteReader)

				assert.Equal(t, testCase.Expected, recipeIngredientUpdate)
				assert.Nil(t, errorCreateDTOFromRecipeIngredientUpdate)
			},
		)
	}
}

func TestCreateEntityFromRecipeProcessUpdate(t *testing.T) {
	tests := []struct {
		name     string
		JSON     string
		Expected entity.RecipeProcess
	}{
		{
			name: "Test case for CreateEntityFromRecipeProcessUpdate with status published",
			JSON: "{\"name\":\"name\",\"description\":\"description\",\"notes\":\"notes\",\"status\":\"published\"}",
			Expected: entity.RecipeProcess{
				Name:        "name",
				Description: "description",
				Notes:       "notes",
				Status:      kind.RecipeProcessStatusPublished,
			},
		},
		{
			name: "Test case for CreateEntityFromRecipeProcessUpdate with status unpublished",
			JSON: "{\"name\":\"name\",\"description\":\"description\",\"notes\":\"notes\",\"status\":\"unpublished\"}",
			Expected: entity.RecipeProcess{
				Name:        "name",
				Description: "description",
				Notes:       "notes",
				Status:      kind.RecipeProcessStatusUnPublished,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				buffer := new(bytes.Buffer)
				buffer.WriteString(testCase.JSON)
				oneByteReader := iotest.OneByteReader(buffer)
				recipeProcessUpdate, errorCreateDTOFromRecipeProcessUpdate := CreateEntityFromRecipeProcessUpdate(oneByteReader)

				assert.Equal(t, testCase.Expected, recipeProcessUpdate)
				assert.Nil(t, errorCreateDTOFromRecipeProcessUpdate)
			},
		)
	}
}

func TestCreateEntityFromPictureUpdate(t *testing.T) {
	tests := []struct {
		name     string
		JSON     string
		Expected entity.Picture
	}{
		{
			name: "Test case for CreateEntityFromPictureUpdate with status published",
			JSON: "{\"name\":\"name\",\"url\":\"https://www.google.com/images/branding/googlelogo/2x/googlelogo_light_color_272x92dp.png\",\"Width\":42,\"Height\":42,\"Size\":1764,\"Type\":\"image/png\",\"status\":\"published\"}",
			Expected: entity.Picture{
				Name:   "name",
				URL:    "https://www.google.com/images/branding/googlelogo/2x/googlelogo_light_color_272x92dp.png",
				Width:  42,
				Height: 42,
				Size:   1764,
				Type:   "image/png",
				Status: kind.PictureStatusPublished,
			},
		},
		{
			name: "Test case for CreateEntityFromPictureUpdate with status unpublished",
			JSON: "{\"name\":\"name\",\"url\":\"https://www.google.com/images/branding/googlelogo/2x/googlelogo_light_color_272x92dp.png\",\"Width\":42,\"Height\":42,\"Size\":1764,\"Type\":\"image/png\",\"status\":\"unpublished\"}",
			Expected: entity.Picture{
				Name:   "name",
				URL:    "https://www.google.com/images/branding/googlelogo/2x/googlelogo_light_color_272x92dp.png",
				Width:  42,
				Height: 42,
				Size:   1764,
				Type:   "image/png",
				Status: kind.PictureStatusUnPublished,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				buffer := new(bytes.Buffer)
				buffer.WriteString(testCase.JSON)
				oneByteReader := iotest.OneByteReader(buffer)
				pictureUpdate, errorCreateDTOFromPictureUpdate := CreateEntityFromPictureUpdate(oneByteReader)

				assert.Equal(t, testCase.Expected, pictureUpdate)
				assert.Nil(t, errorCreateDTOFromPictureUpdate)
			},
		)
	}
}

func TestCreateEntityFromRecipeMeasureUpdate(t *testing.T) {
	tests := []struct {
		name     string
		JSON     string
		Expected entity.RecipeMeasure
	}{
		{
			name: "Test case for CreateEntityFromRecipeMeasureUpdate with status published",
			JSON: "{\"unit_id\":\"00000000-0000-0000-0000-000000000001\",\"value\":42,\"status\":\"published\"}",
			Expected: entity.RecipeMeasure{
				UnitId: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Value:  42,
				Status: kind.RecipeMeasureStatusPublished,
			},
		},
		{
			name: "Test case for CreateEntityFromRecipeMeasureUpdate with status unpublished",
			JSON: "{\"unit_id\":\"00000000-0000-0000-0000-000000000001\",\"value\":42,\"status\":\"unpublished\"}",
			Expected: entity.RecipeMeasure{
				UnitId: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Value:  42,
				Status: kind.RecipeMeasureStatusUnPublished,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				buffer := new(bytes.Buffer)
				buffer.WriteString(testCase.JSON)
				oneByteReader := iotest.OneByteReader(buffer)
				recipeMeasureUpdate, errorCreateDTOFromRecipeMeasureUpdate := CreateEntityFromRecipeMeasureUpdate(oneByteReader)

				assert.Equal(t, testCase.Expected, recipeMeasureUpdate)
				assert.Nil(t, errorCreateDTOFromRecipeMeasureUpdate)
			},
		)
	}
}

func TestCreateEntityFromAltNameUpdate(t *testing.T) {
	tests := []struct {
		name     string
		JSON     string
		Expected entity.AltName
	}{
		{
			name: "Test case for CreateEntityFromAltNameUpdate with status published",
			JSON: "{\"name\":\"name\",\"status\":\"published\"}",
			Expected: entity.AltName{
				Name:   "name",
				Status: kind.AltNameStatusPublished,
			},
		},
		{
			name: "Test case for CreateEntityFromAltNameUpdate with status unpublished",
			JSON: "{\"name\":\"name\",\"status\":\"unpublished\"}",
			Expected: entity.AltName{
				Name:   "name",
				Status: kind.AltNameStatusUnPublished,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				buffer := new(bytes.Buffer)
				buffer.WriteString(testCase.JSON)
				oneByteReader := iotest.OneByteReader(buffer)
				altNameUpdate, errorCreateDTOFromAltNameUpdate := CreateEntityFromAltNameUpdate(oneByteReader)

				assert.Equal(t, testCase.Expected, altNameUpdate)
				assert.Nil(t, errorCreateDTOFromAltNameUpdate)
			},
		)
	}
}

func TestCreateEntityFromUnitUpdate(t *testing.T) {
	tests := []struct {
		name     string
		JSON     string
		Expected entity.Unit
	}{
		{
			name: "Test case for CreateEntityFromUnitUpdate with status published",
			JSON: "{\"name\":\"name\",\"status\":\"published\"}",
			Expected: entity.Unit{
				Name:   "name",
				Status: kind.UnitStatusPublished,
			},
		},
		{
			name: "Test case for CreateEntityFromUnitUpdate with status unpublished",
			JSON: "{\"name\":\"name\",\"status\":\"unpublished\"}",
			Expected: entity.Unit{
				Name:   "name",
				Status: kind.UnitStatusUnPublished,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				buffer := new(bytes.Buffer)
				buffer.WriteString(testCase.JSON)
				oneByteReader := iotest.OneByteReader(buffer)
				unitUpdate, errorCreateDTOFromUnitUpdate := CreateEntityFromUnitUpdate(oneByteReader)

				assert.Equal(t, testCase.Expected, unitUpdate)
				assert.Nil(t, errorCreateDTOFromUnitUpdate)
			},
		)
	}
}

func TestCreateEntityFromCategoryUpdate(t *testing.T) {
	tests := []struct {
		name     string
		JSON     string
		Expected entity.Category
	}{
		{
			name: "Test case for CreateEntityFromCategoryUpdate with status published",
			JSON: "{\"name\":\"name\",\"status\":\"published\"}",
			Expected: entity.Category{
				Name:   "name",
				Status: kind.CategoryStatusPublished,
			},
		},
		{
			name: "Test case for CreateEntityFromCategoryUpdate with status unpublished",
			JSON: "{\"name\":\"name\",\"status\":\"unpublished\"}",
			Expected: entity.Category{
				Name:   "name",
				Status: kind.CategoryStatusUnPublished,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				buffer := new(bytes.Buffer)
				buffer.WriteString(testCase.JSON)
				oneByteReader := iotest.OneByteReader(buffer)
				categoryUpdate, errorCreateDTOFromCategoryUpdate := CreateEntityFromCategoryUpdate(oneByteReader)

				assert.Equal(t, testCase.Expected, categoryUpdate)
				assert.Nil(t, errorCreateDTOFromCategoryUpdate)
			},
		)
	}
}

func TestCreateEntityFromIngredientUpdate(t *testing.T) {
	tests := []struct {
		name     string
		JSON     string
		Expected entity.Ingredient
	}{
		{
			name: "Test case for CreateEntityFromIngredientUpdate with status published",
			JSON: "{\"name\":\"name\",\"status\":\"published\"}",
			Expected: entity.Ingredient{
				Name:   "name",
				Status: kind.IngredientStatusPublished,
			},
		},
		{
			name: "Test case for CreateEntityFromIngredientUpdate with status unpublished",
			JSON: "{\"name\":\"name\",\"status\":\"unpublished\"}",
			Expected: entity.Ingredient{
				Name:   "name",
				Status: kind.IngredientStatusUnPublished,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				buffer := new(bytes.Buffer)
				buffer.WriteString(testCase.JSON)
				oneByteReader := iotest.OneByteReader(buffer)
				ingredientUpdate, errorCreateDTOFromIngredientUpdate := CreateEntityFromIngredientUpdate(oneByteReader)

				assert.Equal(t, testCase.Expected, ingredientUpdate)
				assert.Nil(t, errorCreateDTOFromIngredientUpdate)
			},
		)
	}
}

func TestCreateEntityFromPlannerUpdate(t *testing.T) {
	tests := []struct {
		name     string
		JSON     string
		Expected entity.Planner
	}{
		{
			name: "Test case for CreateEntityFromPlannerUpdate with status published",
			JSON: "{\"start_time\":\"2000-01-01T00:00:00Z\",\"end_time\":\"2000-01-07T23:59:59Z\",\"name\":\"name\",\"status\":\"active\"}",
			Expected: entity.Planner{
				StartTime: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2000, time.January, 7, 23, 59, 59, 0, time.UTC),
				Name:      "name",
				Status:    kind.PlannerStatusActive,
			},
		},
		{
			name: "Test case for CreateEntityFromPlannerUpdate with status unpublished",
			JSON: "{\"start_time\":\"2000-01-01T00:00:00Z\",\"end_time\":\"2000-01-07T23:59:59Z\",\"name\":\"name\",\"status\":\"inactive\"}",
			Expected: entity.Planner{
				StartTime: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2000, time.January, 7, 23, 59, 59, 0, time.UTC),
				Name:      "name",
				Status:    kind.PlannerStatusInActive,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				buffer := new(bytes.Buffer)
				buffer.WriteString(testCase.JSON)
				oneByteReader := iotest.OneByteReader(buffer)
				plannerUpdate, errorCreateDTOFromPlannerUpdate := CreateEntityFromPlannerUpdate(oneByteReader)

				assert.Equal(t, testCase.Expected, plannerUpdate)
				assert.Nil(t, errorCreateDTOFromPlannerUpdate)
			},
		)
	}
}

func TestCreateEntityFromPlannerIntervalUpdate(t *testing.T) {
	tests := []struct {
		name     string
		JSON     string
		Expected entity.PlannerInterval
	}{
		{
			name: "Test case for CreateEntityFromPlannerIntervalUpdate with status published",
			JSON: "{\"start_time\":\"2000-01-01T00:00:00Z\",\"end_time\":\"2000-01-07T23:59:59Z\",\"name\":\"name\",\"status\":\"active\"}",
			Expected: entity.PlannerInterval{
				StartTime: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2000, time.January, 7, 23, 59, 59, 0, time.UTC),
				Name:      "name",
				Status:    kind.PlannerIntervalStatusActive,
			},
		},
		{
			name: "Test case for CreateEntityFromPlannerIntervalUpdate with status unpublished",
			JSON: "{\"start_time\":\"2000-01-01T00:00:00Z\",\"end_time\":\"2000-01-07T23:59:59Z\",\"name\":\"name\",\"status\":\"inactive\"}",
			Expected: entity.PlannerInterval{
				StartTime: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2000, time.January, 7, 23, 59, 59, 0, time.UTC),
				Name:      "name",
				Status:    kind.PlannerIntervalStatusInActive,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				buffer := new(bytes.Buffer)
				buffer.WriteString(testCase.JSON)
				oneByteReader := iotest.OneByteReader(buffer)
				plannerIntervalUpdate, errorCreateDTOFromPlannerIntervalUpdate := CreateEntityFromPlannerIntervalUpdate(oneByteReader)

				assert.Equal(t, testCase.Expected, plannerIntervalUpdate)
				assert.Nil(t, errorCreateDTOFromPlannerIntervalUpdate)
			},
		)
	}
}

func TestCreateEntityFromPlannerRecipeUpdate(t *testing.T) {
	tests := []struct {
		name     string
		JSON     string
		Expected entity.PlannerRecipe
	}{
		{
			name: "Test case for CreateEntityFromPlannerRecipeUpdate with status published",
			JSON: "{\"recipe_id\":\"00000000-0000-0000-0000-000000000001\",\"status\":\"active\"}",
			Expected: entity.PlannerRecipe{
				RecipeId: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Status:   kind.PlannerRecipeStatusActive,
			},
		},
		{
			name: "Test case for CreateEntityFromPlannerRecipeUpdate with status unpublished",
			JSON: "{\"recipe_id\":\"00000000-0000-0000-0000-000000000001\",\"status\":\"inactive\"}",
			Expected: entity.PlannerRecipe{
				RecipeId: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Status:   kind.PlannerRecipeStatusInActive,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				buffer := new(bytes.Buffer)
				buffer.WriteString(testCase.JSON)
				oneByteReader := iotest.OneByteReader(buffer)
				plannerRecipeUpdate, errorCreateDTOFromPlannerRecipeUpdate := CreateEntityFromPlannerRecipeUpdate(oneByteReader)

				assert.Equal(t, testCase.Expected, plannerRecipeUpdate)
				assert.Nil(t, errorCreateDTOFromPlannerRecipeUpdate)
			},
		)
	}
}
