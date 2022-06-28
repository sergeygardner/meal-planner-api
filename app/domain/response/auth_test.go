package response

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestAuthToken(t *testing.T) {
	tests := []struct {
		name         string
		AccessToken  string
		RefreshToken string
	}{
		{
			name:         "Test case with AuthToken properties",
			AccessToken:  "AccessToken",
			RefreshToken: "RefreshToken",
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				authToken := AuthToken{
					AccessToken:  testCase.AccessToken,
					RefreshToken: testCase.RefreshToken,
				}
				assert.Equal(t, testCase.AccessToken, authToken.AccessToken)
				assert.Equal(t, testCase.RefreshToken, authToken.RefreshToken)

				reflectAuthToken := reflect.ValueOf(authToken)

				for i := 0; i < reflectAuthToken.NumField(); i++ {
					assert.False(t, reflectAuthToken.Field(i).IsZero())
				}
			},
		)
	}
}

func TestAuthConfirmation(t *testing.T) {
	tests := []struct {
		name    string
		Message string
		Status  int
	}{
		{
			name:    "Test case with AuthConfirmation properties",
			Message: "Message",
			Status:  42,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				authConfirmation := AuthConfirmation{
					Message: testCase.Message,
					Status:  testCase.Status,
				}
				assert.Equal(t, testCase.Message, authConfirmation.Message)
				assert.Equal(t, testCase.Status, authConfirmation.Status)

				reflectAuthConfirmation := reflect.ValueOf(authConfirmation)

				for i := 0; i < reflectAuthConfirmation.NumField(); i++ {
					assert.False(t, reflectAuthConfirmation.Field(i).IsZero())
				}
			},
		)
	}
}
