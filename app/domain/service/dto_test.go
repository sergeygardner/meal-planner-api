package service

import (
	"bytes"
	"github.com/sergeygardner/meal-planner-api/domain/dto"
	"github.com/stretchr/testify/assert"
	"testing"
	"testing/iotest"
	"time"
)

func TestCreateDTOFromAuthCredentials(t *testing.T) {
	tests := []struct {
		name     string
		JSON     string
		Expected dto.UserCredentialsDTO
	}{
		{
			name:     "Test case for CreateDTOFromAuthCredentials",
			JSON:     "{\"username\":\"username\",\"password\":\"password\"}",
			Expected: dto.UserCredentialsDTO{Username: "username", Password: "password"},
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				buffer := new(bytes.Buffer)
				buffer.WriteString(testCase.JSON)
				oneByteReader := iotest.OneByteReader(buffer)
				authCredentials, errorCreateDTOFromAuthCredentials := CreateDTOFromAuthCredentials(oneByteReader)

				assert.Equal(t, testCase.Expected, authCredentials)
				assert.Nil(t, errorCreateDTOFromAuthCredentials)
			},
		)
	}
}

func TestCreateDTOFromAuthConfirmation(t *testing.T) {
	tests := []struct {
		name     string
		JSON     string
		Expected dto.AuthConfirmationDTO
	}{
		{
			name: "Test case for CreateDTOFromAuthConfirmation",
			JSON: "{\"username\":\"username\",\"password\":\"password\",\"code\":\"123456\"}",
			Expected: dto.AuthConfirmationDTO{
				UserCredentialsDTO: dto.UserCredentialsDTO{
					Username: "username",
					Password: "password",
				},
				Code: "123456",
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
				authConfirmation, errorCreateDTOFromAuthConfirmation := CreateDTOFromAuthConfirmation(oneByteReader)

				assert.Equal(t, testCase.Expected, authConfirmation)
				assert.Nil(t, errorCreateDTOFromAuthConfirmation)
			},
		)
	}
}

func TestCreateDTOFromUserRegister(t *testing.T) {
	tests := []struct {
		name     string
		JSON     string
		Expected dto.UserRegisterDTO
	}{
		{
			name: "Test case for CreateDTOFromUserRegister",
			JSON: "{\"username\":\"username\",\"password\":\"password\",\"name\":\"name\",\"surname\":\"surname\",\"middle_name\":\"middle name\",\"birthday\":\"2000-01-01T00:00:00Z\"}",
			Expected: dto.UserRegisterDTO{
				UserCredentialsDTO: dto.UserCredentialsDTO{
					Username: "username",
					Password: "password",
				},
				Name:       "name",
				Surname:    "surname",
				MiddleName: "middle name",
				Birthday:   time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
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
				userRegister, errorCreateDTOFromUserRegister := CreateDTOFromUserRegister(oneByteReader)

				assert.Equal(t, testCase.Expected, userRegister)
				assert.Nil(t, errorCreateDTOFromUserRegister)
			},
		)
	}
}

func TestCreateDTOFromUserUpdate(t *testing.T) {
	tests := []struct {
		name     string
		JSON     string
		Expected dto.UserDTO
	}{
		{
			name: "Test case for CreateDTOFromUserRegister",
			JSON: "{\"username\":\"username\",\"password\":\"password\",\"name\":\"name\",\"surname\":\"surname\",\"middle_name\":\"middle name\",\"birthday\":\"2000-01-01T00:00:00Z\"}",
			Expected: dto.UserDTO{
				UserRegisterDTO: dto.UserRegisterDTO{
					UserCredentialsDTO: dto.UserCredentialsDTO{
						Username: "username",
						Password: "password",
					},
					Name:       "name",
					Surname:    "surname",
					MiddleName: "middle name",
					Birthday:   time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
				},
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
				userUpdate, errorCreateDTOFromUserUpdate := CreateDTOFromUserUpdate(oneByteReader)

				assert.Equal(t, testCase.Expected, userUpdate)
				assert.Nil(t, errorCreateDTOFromUserUpdate)
			},
		)
	}
}
