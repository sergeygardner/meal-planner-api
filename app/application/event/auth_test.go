package event

import (
	"github.com/google/uuid"
	DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestConstants(t *testing.T) {
	tests := []struct {
		Name     string
		Const    string
		Expected string
	}{
		{
			Name:     "Test case with AuthConfirmationTopicName",
			Const:    AuthConfirmationTopicName,
			Expected: "auth:confirmation",
		},
		{
			Name:     "Test case with AuthConfirmationTopicName",
			Const:    UserConfirmationEventName,
			Expected: "UserConfirmationEvent",
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				assert.Equal(t, testCase.Expected, testCase.Const)
			},
		)
	}
}

func TestUserConfirmationEvent(t *testing.T) {
	tests := []struct {
		Name             string
		UserConfirmation *DomainEntity.UserConfirmation
		MustBePanic      bool
	}{
		{
			Name: "Test case with UserConfirmation with active true",
			UserConfirmation: &DomainEntity.UserConfirmation{
				Id:         uuid.New(),
				DateInsert: time.Now().UTC(),
				DateUpdate: time.Now().UTC(),
				UserId:     uuid.New(),
				Value:      "test",
				Active:     true,
			},
			MustBePanic: false,
		},
		{
			Name: "Test case with UserConfirmation with active false",
			UserConfirmation: &DomainEntity.UserConfirmation{
				Id:         uuid.New(),
				DateInsert: time.Now().UTC(),
				DateUpdate: time.Now().UTC(),
				UserId:     uuid.New(),
				Value:      "test",
				Active:     false,
			},
			MustBePanic: true,
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

				UserConfirmationEvent(testCase.UserConfirmation)
			},
		)
	}
}
