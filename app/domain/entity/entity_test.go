package entity

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	"github.com/sergeygardner/meal-planner-api/domain/service/validator"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestUnit(t *testing.T) {
	tests := []struct {
		name        string
		json        string
		Id          uuid.UUID
		DateInsert  time.Time
		DateUpdate  time.Time
		Name        string
		Status      kind.UnitStatus
		MustBeFault bool
	}{
		{
			name:        "Test case with published unit properties",
			json:        "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Unit\",\"status\":\"published\"}\n",
			Id:          uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			DateInsert:  time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate:  time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Name:        "Unit",
			Status:      kind.UnitStatusPublished,
			MustBeFault: false,
		},
		{
			name:        "Test case with unpublished unit properties",
			json:        "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Unit\",\"status\":\"unpublished\"}\n",
			Id:          uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			DateInsert:  time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate:  time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Name:        "Unit",
			Status:      kind.UnitStatusUnPublished,
			MustBeFault: false,
		},
		{
			name:        "Test case with unpublished unit properties with incorrect name",
			json:        "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Test case with unpublished unit propertiesTest case with unpublished unit propertiesTest case with unpublished unit propertiesTest case with unpublished unit propertiesTest case with unpublished unit propertiesTest case with unpublished unit properties unit properties\",\"status\":\"unpublished\"}\n",
			Id:          uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			DateInsert:  time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate:  time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Name:        "Test case with unpublished unit propertiesTest case with unpublished unit propertiesTest case with unpublished unit propertiesTest case with unpublished unit propertiesTest case with unpublished unit propertiesTest case with unpublished unit properties unit properties",
			Status:      kind.UnitStatusUnPublished,
			MustBeFault: true,
		},
		{
			name:        "Test case with unpublished unit properties with incorrect name",
			json:        "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"T\",\"status\":\"unpublished\"}\n",
			Id:          uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			DateInsert:  time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate:  time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Name:        "T",
			Status:      kind.UnitStatusUnPublished,
			MustBeFault: true,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				unit := Unit{
					Id:         testCase.Id,
					DateInsert: testCase.DateInsert,
					DateUpdate: testCase.DateUpdate,
					Name:       testCase.Name,
					Status:     testCase.Status,
				}
				assert.Equal(t, testCase.Id, unit.Id)
				assert.Equal(t, testCase.DateInsert, unit.DateInsert)
				assert.Equal(t, testCase.DateUpdate, unit.DateUpdate)
				assert.Equal(t, testCase.Name, unit.Name)
				assert.Equal(t, testCase.Status, unit.Status)

				reflectUnit := reflect.ValueOf(unit)

				for i := 0; i < reflectUnit.NumField(); i++ {
					assert.False(t, reflectUnit.Field(i).IsZero())
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(unit)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())

				errorValidate := validator.Validate(unit)

				if testCase.MustBeFault {
					assert.NotNil(t, errorValidate)
				} else {
					assert.Nil(t, errorValidate)
				}
			},
		)
	}
}

func TestAltName(t *testing.T) {
	tests := []struct {
		name       string
		json       string
		Id         uuid.UUID
		UserId     uuid.UUID
		EntityId   uuid.UUID
		DateInsert time.Time
		DateUpdate time.Time
		Name       string
		Status     kind.AltNameStatus
	}{
		{
			name:       "Test case with published alt name properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000003\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"published\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Name:       "AltName",
			Status:     kind.AltNameStatusPublished,
		},
		{
			name:       "Test case with unpublished alt name properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000003\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"AltName\",\"status\":\"unpublished\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Name:       "AltName",
			Status:     kind.AltNameStatusUnPublished,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				altName := AltName{
					Id:         testCase.Id,
					UserId:     testCase.UserId,
					EntityId:   testCase.EntityId,
					DateInsert: testCase.DateInsert,
					DateUpdate: testCase.DateUpdate,
					Name:       testCase.Name,
					Status:     testCase.Status,
				}
				assert.Equal(t, testCase.Id, altName.Id)
				assert.Equal(t, testCase.UserId, altName.UserId)
				assert.Equal(t, testCase.EntityId, altName.EntityId)
				assert.Equal(t, testCase.DateInsert, altName.DateInsert)
				assert.Equal(t, testCase.DateUpdate, altName.DateUpdate)
				assert.Equal(t, testCase.Name, altName.Name)
				assert.Equal(t, testCase.Status, altName.Status)

				reflectAltName := reflect.ValueOf(altName)

				for i := 0; i < reflectAltName.NumField(); i++ {
					assert.False(t, reflectAltName.Field(i).IsZero())
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(altName)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}

func TestPicture(t *testing.T) {
	tests := []struct {
		name       string
		json       string
		Id         uuid.UUID
		UserId     uuid.UUID
		EntityId   uuid.UUID
		DateInsert time.Time
		DateUpdate time.Time
		Name       string
		URL        string
		Width      int64
		Height     int64
		Size       int64
		Type       string
		Status     kind.PictureStatus
	}{
		{
			name:       "Test case with published picture properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000003\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"published\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Name:       "Picture",
			URL:        "https://google.com/doodle.png",
			Width:      512,
			Height:     512,
			Size:       1024,
			Type:       "image/png",
			Status:     kind.PictureStatusPublished,
		},
		{
			name:       "Test case with unpublished picture properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000003\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Picture\",\"url\":\"https://google.com/doodle.png\",\"width\":512,\"height\":512,\"size\":1024,\"type\":\"image/png\",\"status\":\"unpublished\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Name:       "Picture",
			URL:        "https://google.com/doodle.png",
			Width:      512,
			Height:     512,
			Size:       1024,
			Type:       "image/png",
			Status:     kind.PictureStatusUnPublished,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				picture := Picture{
					Id:         testCase.Id,
					UserId:     testCase.UserId,
					EntityId:   testCase.EntityId,
					DateInsert: testCase.DateInsert,
					DateUpdate: testCase.DateUpdate,
					Name:       testCase.Name,
					URL:        testCase.URL,
					Width:      testCase.Width,
					Height:     testCase.Height,
					Size:       testCase.Size,
					Type:       testCase.Type,
					Status:     testCase.Status,
				}
				assert.Equal(t, testCase.Id, picture.Id)
				assert.Equal(t, testCase.UserId, picture.UserId)
				assert.Equal(t, testCase.EntityId, picture.EntityId)
				assert.Equal(t, testCase.DateInsert, picture.DateInsert)
				assert.Equal(t, testCase.DateUpdate, picture.DateUpdate)
				assert.Equal(t, testCase.Name, picture.Name)
				assert.Equal(t, testCase.URL, picture.URL)
				assert.Equal(t, testCase.Width, picture.Width)
				assert.Equal(t, testCase.Height, picture.Height)
				assert.Equal(t, testCase.Size, picture.Size)
				assert.Equal(t, testCase.Status, picture.Status)

				reflectPicture := reflect.ValueOf(picture)

				for i := 0; i < reflectPicture.NumField(); i++ {
					assert.False(t, reflectPicture.Field(i).IsZero())
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(picture)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}

func TestCategory(t *testing.T) {
	tests := []struct {
		name       string
		json       string
		Id         uuid.UUID
		UserId     uuid.UUID
		DateInsert time.Time
		DateUpdate time.Time
		Name       string
		Status     kind.CategoryStatus
	}{
		{
			name:       "Test case with published category properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Category\",\"status\":\"published\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Name:       "Category",
			Status:     kind.CategoryStatusPublished,
		},
		{
			name:       "Test case with unpublished category properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Category\",\"status\":\"unpublished\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Name:       "Category",
			Status:     kind.CategoryStatusUnPublished,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				category := Category{
					Id:         testCase.Id,
					UserId:     testCase.UserId,
					DateInsert: testCase.DateInsert,
					DateUpdate: testCase.DateUpdate,
					Name:       testCase.Name,
					Status:     testCase.Status,
				}
				assert.Equal(t, testCase.Id, category.Id)
				assert.Equal(t, testCase.UserId, category.UserId)
				assert.Equal(t, testCase.DateInsert, category.DateInsert)
				assert.Equal(t, testCase.DateUpdate, category.DateUpdate)
				assert.Equal(t, testCase.Name, category.Name)
				assert.Equal(t, testCase.Status, category.Status)

				reflectCategory := reflect.ValueOf(category)

				for i := 0; i < reflectCategory.NumField(); i++ {
					assert.False(t, reflectCategory.Field(i).IsZero())
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(category)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}

func TestIngredient(t *testing.T) {
	tests := []struct {
		name       string
		json       string
		Id         uuid.UUID
		UserId     uuid.UUID
		DateInsert time.Time
		DateUpdate time.Time
		Name       string
		Status     kind.IngredientStatus
	}{
		{
			name:       "Test case with published ingredient properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Ingredient\",\"status\":\"published\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Name:       "Ingredient",
			Status:     kind.IngredientStatusPublished,
		},
		{
			name:       "Test case with unpublished ingredient properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"name\":\"Ingredient\",\"status\":\"unpublished\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Name:       "Ingredient",
			Status:     kind.IngredientStatusUnPublished,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				ingredient := Ingredient{
					Id:         testCase.Id,
					UserId:     testCase.UserId,
					DateInsert: testCase.DateInsert,
					DateUpdate: testCase.DateUpdate,
					Name:       testCase.Name,
					Status:     testCase.Status,
				}
				assert.Equal(t, testCase.Id, ingredient.Id)
				assert.Equal(t, testCase.UserId, ingredient.UserId)
				assert.Equal(t, testCase.DateInsert, ingredient.DateInsert)
				assert.Equal(t, testCase.DateUpdate, ingredient.DateUpdate)
				assert.Equal(t, testCase.Name, ingredient.Name)
				assert.Equal(t, testCase.Status, ingredient.Status)

				reflectIngredient := reflect.ValueOf(ingredient)

				for i := 0; i < reflectIngredient.NumField(); i++ {
					assert.False(t, reflectIngredient.Field(i).IsZero())
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(ingredient)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}
