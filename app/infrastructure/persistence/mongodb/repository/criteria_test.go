package repository

import (
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/infrastructure/persistence"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testId = uuid.New()

func TestGetCriteriaById(t *testing.T) {
	tests := []struct {
		Name        string
		Id          uuid.UUID
		Criteria    *persistence.Criteria
		Expected    *persistence.Criteria
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:     "Test case with GetCriteriaById with empty criteria",
			Id:       testId,
			Criteria: nil,
			Expected: &persistence.Criteria{
				Where: map[string]interface{}{"id": &testId},
			},
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name: "Test case with GetCriteriaById with not empty criteria",
			Id:   testId,
			Criteria: &persistence.Criteria{
				Limit:  100,
				Offset: 100,
				Order:  map[string]interface{}{"sort": "ASC"},
			},
			Expected: &persistence.Criteria{
				Limit:  100,
				Offset: 100,
				Order:  map[string]interface{}{"sort": "ASC"},
				Where:  map[string]interface{}{"id": &testId},
			},
			MustBePanic: false,
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				defer func() {
					if testCase.MustBePanic {
						assert.NotNil(t, recover())
					} else {
						assert.Nil(t, recover())
					}
				}()

				actualCriteria := CriteriaRepository.GetCriteriaById(&testCase.Id, testCase.Criteria)

				if testCase.MustBeFault {
					assert.NotEqual(t, *testCase.Expected, *actualCriteria)
				} else {
					assert.Equal(t, *testCase.Expected, *actualCriteria)
				}
			},
		)
	}
}

func TestGetCriteriaByIds(t *testing.T) {
	tests := []struct {
		Name        string
		Id          []*uuid.UUID
		Criteria    *persistence.Criteria
		Expected    *persistence.Criteria
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:     "Test case with GetCriteriaByIds with empty criteria",
			Id:       []*uuid.UUID{&testId},
			Criteria: nil,
			Expected: &persistence.Criteria{
				Where: map[string]interface{}{"id": []*uuid.UUID{&testId}},
			},
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name: "Test case with GetCriteriaByIds with not empty criteria",
			Id:   []*uuid.UUID{&testId},
			Criteria: &persistence.Criteria{
				Limit:  100,
				Offset: 100,
				Order:  map[string]interface{}{"sort": "ASC"},
			},
			Expected: &persistence.Criteria{
				Limit:  100,
				Offset: 100,
				Order:  map[string]interface{}{"sort": "ASC"},
				Where:  map[string]interface{}{"id": []*uuid.UUID{&testId}},
			},
			MustBePanic: false,
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				defer func() {
					if testCase.MustBePanic {
						assert.NotNil(t, recover())
					} else {
						assert.Nil(t, recover())
					}
				}()

				actualCriteria := CriteriaRepository.GetCriteriaByIds(testCase.Id, testCase.Criteria)

				if testCase.MustBeFault {
					assert.NotEqual(t, *testCase.Expected, *actualCriteria)
				} else {
					assert.Equal(t, *testCase.Expected, *actualCriteria)
				}
			},
		)
	}
}

func TestGetCriteriaByEntityId(t *testing.T) {
	tests := []struct {
		Name        string
		Id          uuid.UUID
		Criteria    *persistence.Criteria
		Expected    *persistence.Criteria
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:     "Test case with GetCriteriaByEntityId with empty criteria",
			Id:       testId,
			Criteria: nil,
			Expected: &persistence.Criteria{
				Where: map[string]interface{}{"entity_id": &testId},
			},
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name: "Test case with GetCriteriaByEntityId with not empty criteria",
			Id:   testId,
			Criteria: &persistence.Criteria{
				Limit:  100,
				Offset: 100,
				Order:  map[string]interface{}{"sort": "ASC"},
			},
			Expected: &persistence.Criteria{
				Limit:  100,
				Offset: 100,
				Order:  map[string]interface{}{"sort": "ASC"},
				Where:  map[string]interface{}{"entity_id": &testId},
			},
			MustBePanic: false,
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				defer func() {
					if testCase.MustBePanic {
						assert.NotNil(t, recover())
					} else {
						assert.Nil(t, recover())
					}
				}()

				actualCriteria := CriteriaRepository.GetCriteriaByEntityId(&testCase.Id, testCase.Criteria)

				if testCase.MustBeFault {
					assert.NotEqual(t, *testCase.Expected, *actualCriteria)
				} else {
					assert.Equal(t, *testCase.Expected, *actualCriteria)
				}
			},
		)
	}
}

func TestGetCriteriaByDeriveId(t *testing.T) {
	tests := []struct {
		Name        string
		Id          uuid.UUID
		Criteria    *persistence.Criteria
		Expected    *persistence.Criteria
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:     "Test case with GetCriteriaByDeriveId with empty criteria",
			Id:       testId,
			Criteria: nil,
			Expected: &persistence.Criteria{
				Where: map[string]interface{}{"derive_id": &testId},
			},
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name: "Test case with GetCriteriaByDeriveId with not empty criteria",
			Id:   testId,
			Criteria: &persistence.Criteria{
				Limit:  100,
				Offset: 100,
				Order:  map[string]interface{}{"sort": "ASC"},
			},
			Expected: &persistence.Criteria{
				Limit:  100,
				Offset: 100,
				Order:  map[string]interface{}{"sort": "ASC"},
				Where:  map[string]interface{}{"derive_id": &testId},
			},
			MustBePanic: false,
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				defer func() {
					if testCase.MustBePanic {
						assert.NotNil(t, recover())
					} else {
						assert.Nil(t, recover())
					}
				}()

				actualCriteria := CriteriaRepository.GetCriteriaByDeriveId(&testCase.Id, testCase.Criteria)

				if testCase.MustBeFault {
					assert.NotEqual(t, *testCase.Expected, *actualCriteria)
				} else {
					assert.Equal(t, *testCase.Expected, *actualCriteria)
				}
			},
		)
	}
}

func TestGetCriteriaByUnitId(t *testing.T) {
	tests := []struct {
		Name        string
		Id          uuid.UUID
		Criteria    *persistence.Criteria
		Expected    *persistence.Criteria
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:     "Test case with GetCriteriaByUnitId with empty criteria",
			Id:       testId,
			Criteria: nil,
			Expected: &persistence.Criteria{
				Where: map[string]interface{}{"unit_id": &testId},
			},
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name: "Test case with GetCriteriaByUnitId with not empty criteria",
			Id:   testId,
			Criteria: &persistence.Criteria{
				Limit:  100,
				Offset: 100,
				Order:  map[string]interface{}{"sort": "ASC"},
			},
			Expected: &persistence.Criteria{
				Limit:  100,
				Offset: 100,
				Order:  map[string]interface{}{"sort": "ASC"},
				Where:  map[string]interface{}{"unit_id": &testId},
			},
			MustBePanic: false,
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				defer func() {
					if testCase.MustBePanic {
						assert.NotNil(t, recover())
					} else {
						assert.Nil(t, recover())
					}
				}()

				actualCriteria := CriteriaRepository.GetCriteriaByUnitId(&testCase.Id, testCase.Criteria)

				if testCase.MustBeFault {
					assert.NotEqual(t, *testCase.Expected, *actualCriteria)
				} else {
					assert.Equal(t, *testCase.Expected, *actualCriteria)
				}
			},
		)
	}
}

func TestGetCriteriaByRecipeId(t *testing.T) {
	tests := []struct {
		Name        string
		Id          uuid.UUID
		Criteria    *persistence.Criteria
		Expected    *persistence.Criteria
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:     "Test case with GetCriteriaByRecipeId with empty criteria",
			Id:       testId,
			Criteria: nil,
			Expected: &persistence.Criteria{
				Where: map[string]interface{}{"recipe_id": &testId},
			},
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name: "Test case with GetCriteriaByRecipeId with not empty criteria",
			Id:   testId,
			Criteria: &persistence.Criteria{
				Limit:  100,
				Offset: 100,
				Order:  map[string]interface{}{"sort": "ASC"},
			},
			Expected: &persistence.Criteria{
				Limit:  100,
				Offset: 100,
				Order:  map[string]interface{}{"sort": "ASC"},
				Where:  map[string]interface{}{"recipe_id": &testId},
			},
			MustBePanic: false,
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				defer func() {
					if testCase.MustBePanic {
						assert.NotNil(t, recover())
					} else {
						assert.Nil(t, recover())
					}
				}()

				actualCriteria := CriteriaRepository.GetCriteriaByRecipeId(&testCase.Id, testCase.Criteria)

				if testCase.MustBeFault {
					assert.NotEqual(t, *testCase.Expected, *actualCriteria)
				} else {
					assert.Equal(t, *testCase.Expected, *actualCriteria)
				}
			},
		)
	}
}

func TestGetCriteriaByUserId(t *testing.T) {
	tests := []struct {
		Name        string
		Id          uuid.UUID
		Criteria    *persistence.Criteria
		Expected    *persistence.Criteria
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:     "Test case with GetCriteriaByUserId with empty criteria",
			Id:       testId,
			Criteria: nil,
			Expected: &persistence.Criteria{
				Where: map[string]interface{}{"user_id": &testId},
			},
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name: "Test case with GetCriteriaByUserId with not empty criteria",
			Id:   testId,
			Criteria: &persistence.Criteria{
				Limit:  100,
				Offset: 100,
				Order:  map[string]interface{}{"sort": "ASC"},
			},
			Expected: &persistence.Criteria{
				Limit:  100,
				Offset: 100,
				Order:  map[string]interface{}{"sort": "ASC"},
				Where:  map[string]interface{}{"user_id": &testId},
			},
			MustBePanic: false,
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				defer func() {
					if testCase.MustBePanic {
						assert.NotNil(t, recover())
					} else {
						assert.Nil(t, recover())
					}
				}()

				actualCriteria := CriteriaRepository.GetCriteriaByUserId(&testCase.Id, testCase.Criteria)

				if testCase.MustBeFault {
					assert.NotEqual(t, *testCase.Expected, *actualCriteria)
				} else {
					assert.Equal(t, *testCase.Expected, *actualCriteria)
				}
			},
		)
	}
}

func TestGetCriteriaByName(t *testing.T) {
	name := "name"
	tests := []struct {
		Name        string
		Criteria    *persistence.Criteria
		Expected    *persistence.Criteria
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:     "Test case with GetCriteriaByUserId with empty criteria",
			Criteria: nil,
			Expected: &persistence.Criteria{
				Where: map[string]interface{}{"name": &name},
			},
			MustBePanic: false,
			MustBeFault: false,
		},
		{
			Name: "Test case with GetCriteriaByUserId with not empty criteria",
			Criteria: &persistence.Criteria{
				Limit:  100,
				Offset: 100,
				Order:  map[string]interface{}{"sort": "ASC"},
			},
			Expected: &persistence.Criteria{
				Limit:  100,
				Offset: 100,
				Order:  map[string]interface{}{"sort": "ASC"},
				Where:  map[string]interface{}{"name": &name},
			},
			MustBePanic: false,
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				defer func() {
					if testCase.MustBePanic {
						assert.NotNil(t, recover())
					} else {
						assert.Nil(t, recover())
					}
				}()

				actualCriteria := CriteriaRepository.GetCriteriaByName(&name, testCase.Criteria)

				if testCase.MustBeFault {
					assert.NotEqual(t, *testCase.Expected, *actualCriteria)
				} else {
					assert.Equal(t, *testCase.Expected, *actualCriteria)
				}
			},
		)
	}
}
