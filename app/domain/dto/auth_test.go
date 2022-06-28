package dto

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestAuthConfirmationDTO(t *testing.T) {
	tests := []struct {
		name     string
		Username string
		Password string
		Code     string
	}{
		{
			name:     "Test case with AuthConfirmationDTO properties",
			Username: "Username",
			Password: "Password",
			Code:     "123456",
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				authConfirmationDTO := AuthConfirmationDTO{
					UserCredentialsDTO: UserCredentialsDTO{
						Username: testCase.Username,
						Password: testCase.Password,
					},
					Code: testCase.Code,
				}
				assert.Equal(t, testCase.Username, authConfirmationDTO.Username)
				assert.Equal(t, testCase.Password, authConfirmationDTO.Password)
				assert.Equal(t, testCase.Code, authConfirmationDTO.Code)

				reflectAuthConfirmationDTO := reflect.ValueOf(authConfirmationDTO)

				for i := 0; i < reflectAuthConfirmationDTO.NumField(); i++ {
					assert.False(t, reflectAuthConfirmationDTO.Field(i).IsZero())
				}
			},
		)
	}
}
