package update

import (
	"github.com/google/uuid"
	DomainEntity "github.com/sergeygardner/meal-planner-api/domain/entity"
	InfrastructureService "github.com/sergeygardner/meal-planner-api/infrastructure/service"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func init() {
	InfrastructureService.PrepareCache()
	InfrastructureService.PreparePersistence()
}

func TestSetUserConfirmationInActive(t *testing.T) {
	tests := []struct {
		Name             string
		UserConfirmation *DomainEntity.UserConfirmation
		MustBePanic      bool
		MustBeFault      bool
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

				errorSetUserConfirmationInActive := SetUserConfirmationInActive(testCase.UserConfirmation)

				if testCase.MustBeFault {
					assert.NotNil(t, errorSetUserConfirmationInActive)
				} else {
					assert.Nil(t, errorSetUserConfirmationInActive)
				}
			},
		)
	}
}
