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

func TestPlanner(t *testing.T) {
	tests := []struct {
		name       string
		json       string
		Id         uuid.UUID
		UserId     uuid.UUID
		DateInsert time.Time
		DateUpdate time.Time
		StartTime  time.Time
		EndTime    time.Time
		Name       string
		Status     kind.PlannerStatus
	}{
		{
			name:       "Test case with active planner properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"start_time\":\"2000-01-11T00:00:00Z\",\"end_time\":\"2000-01-17T23:59:59Z\",\"name\":\"Planner\",\"status\":\"active\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			StartTime:  time.Date(2000, time.January, 11, 0, 0, 0, 0, time.UTC),
			EndTime:    time.Date(2000, time.January, 17, 23, 59, 59, 0, time.UTC),
			Name:       "Planner",
			Status:     kind.PlannerStatusActive,
		},
		{
			name:       "Test case with inactive planner properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"start_time\":\"2000-01-11T00:00:00Z\",\"end_time\":\"2000-01-17T23:59:59Z\",\"name\":\"Planner\",\"status\":\"inactive\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			StartTime:  time.Date(2000, time.January, 11, 0, 0, 0, 0, time.UTC),
			EndTime:    time.Date(2000, time.January, 17, 23, 59, 59, 0, time.UTC),
			Name:       "Planner",
			Status:     kind.PlannerStatusInActive,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				planner := Planner{
					Id:         testCase.Id,
					UserId:     testCase.UserId,
					DateInsert: testCase.DateInsert,
					DateUpdate: testCase.DateUpdate,
					StartTime:  testCase.StartTime,
					EndTime:    testCase.EndTime,
					Name:       testCase.Name,
					Status:     testCase.Status,
				}
				assert.Equal(t, testCase.Id, planner.Id)
				assert.Equal(t, testCase.UserId, planner.UserId)
				assert.Equal(t, testCase.DateInsert, planner.DateInsert)
				assert.Equal(t, testCase.DateUpdate, planner.DateUpdate)
				assert.Equal(t, testCase.StartTime, planner.StartTime)
				assert.Equal(t, testCase.EndTime, planner.EndTime)
				assert.Equal(t, testCase.Name, planner.Name)
				assert.Equal(t, testCase.Status, planner.Status)

				reflectPlanner := reflect.ValueOf(planner)

				for i := 0; i < reflectPlanner.NumField(); i++ {
					assert.False(t, reflectPlanner.Field(i).IsZero())
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(planner)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}

func TestPlannerInterval(t *testing.T) {
	tests := []struct {
		name       string
		json       string
		Id         uuid.UUID
		UserId     uuid.UUID
		EntityId   uuid.UUID
		DateInsert time.Time
		DateUpdate time.Time
		StartTime  time.Time
		EndTime    time.Time
		Name       string
		Status     kind.PlannerIntervalStatus
	}{
		{
			name:       "Test case with active planner interval properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000003\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"start_time\":\"2000-01-11T00:00:00Z\",\"end_time\":\"2000-01-11T23:59:59Z\",\"name\":\"PlannerInterval\",\"status\":\"active\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			StartTime:  time.Date(2000, time.January, 11, 0, 0, 0, 0, time.UTC),
			EndTime:    time.Date(2000, time.January, 11, 23, 59, 59, 0, time.UTC),
			Name:       "PlannerInterval",
			Status:     kind.PlannerIntervalStatusActive,
		},
		{
			name:       "Test case with inactive planner interval properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000003\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"start_time\":\"2000-01-11T00:00:00Z\",\"end_time\":\"2000-01-11T23:59:59Z\",\"name\":\"PlannerInterval\",\"status\":\"inactive\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			StartTime:  time.Date(2000, time.January, 11, 0, 0, 0, 0, time.UTC),
			EndTime:    time.Date(2000, time.January, 11, 23, 59, 59, 0, time.UTC),
			Name:       "PlannerInterval",
			Status:     kind.PlannerIntervalStatusInActive,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				plannerInterval := PlannerInterval{
					Id:         testCase.Id,
					UserId:     testCase.UserId,
					EntityId:   testCase.EntityId,
					DateInsert: testCase.DateInsert,
					DateUpdate: testCase.DateUpdate,
					StartTime:  testCase.StartTime,
					EndTime:    testCase.EndTime,
					Name:       testCase.Name,
					Status:     testCase.Status,
				}
				assert.Equal(t, testCase.Id, plannerInterval.Id)
				assert.Equal(t, testCase.UserId, plannerInterval.UserId)
				assert.Equal(t, testCase.EntityId, plannerInterval.EntityId)
				assert.Equal(t, testCase.DateInsert, plannerInterval.DateInsert)
				assert.Equal(t, testCase.DateUpdate, plannerInterval.DateUpdate)
				assert.Equal(t, testCase.StartTime, plannerInterval.StartTime)
				assert.Equal(t, testCase.EndTime, plannerInterval.EndTime)
				assert.Equal(t, testCase.Name, plannerInterval.Name)
				assert.Equal(t, testCase.Status, plannerInterval.Status)

				reflectPlannerInterval := reflect.ValueOf(plannerInterval)

				for i := 0; i < reflectPlannerInterval.NumField(); i++ {
					assert.False(t, reflectPlannerInterval.Field(i).IsZero())
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(plannerInterval)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}

func TestPlannerRecipe(t *testing.T) {
	tests := []struct {
		name       string
		json       string
		Id         uuid.UUID
		UserId     uuid.UUID
		EntityId   uuid.UUID
		RecipeId   uuid.UUID
		DateInsert time.Time
		DateUpdate time.Time
		Status     kind.PlannerRecipeStatus
	}{
		{
			name:       "Test case with active planner recipe properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000003\",\"recipe_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"status\":\"active\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			RecipeId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
			DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Status:     kind.PlannerRecipeStatusActive,
		},
		{
			name:       "Test case with inactive planner recipe properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"entity_id\":\"00000000-0000-0000-0000-000000000003\",\"recipe_id\":\"00000000-0000-0000-0000-000000000004\",\"date_insert\":\"2000-01-01T00:00:00Z\",\"date_update\":\"2000-01-10T00:00:00Z\",\"status\":\"inactive\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			EntityId:   uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			RecipeId:   uuid.MustParse("00000000-0000-0000-0000-000000000004"),
			DateInsert: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2000, time.January, 10, 0, 0, 0, 0, time.UTC),
			Status:     kind.PlannerRecipeStatusInActive,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				plannerRecipe := PlannerRecipe{
					Id:         testCase.Id,
					UserId:     testCase.UserId,
					EntityId:   testCase.EntityId,
					RecipeId:   testCase.RecipeId,
					DateInsert: testCase.DateInsert,
					DateUpdate: testCase.DateUpdate,
					Status:     testCase.Status,
				}
				assert.Equal(t, testCase.Id, plannerRecipe.Id)
				assert.Equal(t, testCase.UserId, plannerRecipe.UserId)
				assert.Equal(t, testCase.EntityId, plannerRecipe.EntityId)
				assert.Equal(t, testCase.DateInsert, plannerRecipe.DateInsert)
				assert.Equal(t, testCase.DateUpdate, plannerRecipe.DateUpdate)
				assert.Equal(t, testCase.Status, plannerRecipe.Status)

				reflectPlannerRecipe := reflect.ValueOf(plannerRecipe)

				for i := 0; i < reflectPlannerRecipe.NumField(); i++ {
					assert.False(t, reflectPlannerRecipe.Field(i).IsZero())
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(plannerRecipe)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}
