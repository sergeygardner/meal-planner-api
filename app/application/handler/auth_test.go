package handler

import (
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/domain/dto"
	"github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	"github.com/sergeygardner/meal-planner-api/domain/response"
	"github.com/sergeygardner/meal-planner-api/infrastructure/service/repository"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	testResponseAuthConfirmation *response.AuthConfirmation
	testUserConfirmation         *entity.UserConfirmation
	testResponseAuthToken        *response.AuthToken
	errorAuthCredentials         error
	errorAuthConfirmation        error
)

func TestAuthCredentials(t *testing.T) {
	tests := []struct {
		Name               string
		UserCredentialsDTO dto.UserCredentialsDTO
		MustBePanic        bool
		MustBeFault        bool
	}{
		{
			Name:               "Test case with AuthCredentials",
			UserCredentialsDTO: dto.UserCredentialsDTO{Username: "username", Password: "password"},
			MustBePanic:        false,
			MustBeFault:        false,
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

				testResponseAuthConfirmation, testUserConfirmation, errorAuthCredentials = AuthCredentials(testCase.UserCredentialsDTO)

				if testCase.MustBeFault {
					assert.Nil(t, testResponseAuthConfirmation)
					assert.Nil(t, testUserConfirmation)
					assert.NotNil(t, errorAuthCredentials)
				} else {
					assert.NotNil(t, testResponseAuthConfirmation)
					assert.NotNil(t, testUserConfirmation)
					assert.Nil(t, errorAuthCredentials)
				}
			},
		)
	}
}

func TestAuthConfirmation(t *testing.T) {
	tests := []struct {
		Name                string
		AuthConfirmationDTO dto.AuthConfirmationDTO
		MustBePanic         bool
		MustBeFault         bool
	}{
		{
			Name: "Test case with AuthConfirmationDTO",
			AuthConfirmationDTO: dto.AuthConfirmationDTO{
				UserCredentialsDTO: dto.UserCredentialsDTO{
					Username: "username",
					Password: "password",
				},
				Code: testUserConfirmation.Value,
			},
			MustBeFault: false,
		},
		{
			Name: "Test case with AuthConfirmation",
			AuthConfirmationDTO: dto.AuthConfirmationDTO{
				UserCredentialsDTO: dto.UserCredentialsDTO{
					Username: "username",
					Password: "password",
				},
				Code: "test",
			},
			MustBeFault: true,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				testResponseAuthToken, errorAuthConfirmation = AuthConfirmation(testCase.AuthConfirmationDTO)

				if testCase.MustBeFault {
					assert.Nil(t, testResponseAuthToken)
					assert.NotNil(t, errorAuthConfirmation)
				} else {
					assert.NotNil(t, testResponseAuthToken)
					assert.Nil(t, errorAuthConfirmation)
				}
			},
		)
	}
}

func TestAuthToken(t *testing.T) {
	tests := []struct {
		Name        string
		User        *entity.User
		MustBeFault bool
	}{
		{
			Name: "Test case with AuthToken",
			User: &entity.User{
				Id: uuid.New(),
				UserDTO: dto.UserDTO{
					UserRegisterDTO: dto.UserRegisterDTO{
						UserCredentialsDTO: dto.UserCredentialsDTO{
							Username: "username",
						},
					},
					Roles: kind.UserRoles{kind.UserRoleCommon},
				},
			},
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				authToken, errorAuthToken := AuthToken(testCase.User)

				if testCase.MustBeFault {
					assert.Nil(t, authToken)
					assert.NotNil(t, errorAuthToken)
				} else {
					assert.NotNil(t, authToken)
					assert.Nil(t, errorAuthToken)
				}
			},
		)
	}
}

func TestAuthTokenByUserId(t *testing.T) {
	fakeUserId := uuid.New()
	userRepository := repository.GetFactoryRepository().GetUserRepository()
	user, errorFindOne := userRepository.FindOne(userRepository.GetCriteriaByUsername("username"))

	if errorFindOne != nil {
		panic(errorFindOne.Error())
	}

	tests := []struct {
		Name        string
		UserId      *uuid.UUID
		MustBeFault bool
	}{
		{
			Name:        "Test case with AuthTokenByUserId and correct UserId",
			UserId:      &user.Id,
			MustBeFault: false,
		},
		{
			Name:        "Test case with AuthTokenByUserId and incorrect UserId",
			UserId:      &fakeUserId,
			MustBeFault: true,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				authToken, errorAuthToken := AuthTokenByUserId(testCase.UserId)

				if testCase.MustBeFault {
					assert.Nil(t, authToken)
					assert.NotNil(t, errorAuthToken)
				} else {
					assert.NotNil(t, authToken)
					assert.Nil(t, errorAuthToken)
				}
			},
		)
	}
}

func TestAuthRegister(t *testing.T) {
	tests := []struct {
		Name            string
		UserRegisterDTO dto.UserRegisterDTO
		MustBeFault     bool
	}{
		{
			Name: "Test case with AuthRegister",
			UserRegisterDTO: dto.UserRegisterDTO{
				UserCredentialsDTO: dto.UserCredentialsDTO{
					Username: "usernameTest",
					Password: "passwordTest",
				},
				Name:       "NameTest",
				Surname:    "SurnameTest",
				MiddleName: "MiddleNameTest",
				Birthday:   time.Now().UTC(),
			},
			MustBeFault: false,
		},
	}

	userRepository := repository.GetFactoryRepository().GetUserRepository()

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				user, errorAuthRegister := AuthRegister(testCase.UserRegisterDTO)

				if testCase.MustBeFault {
					assert.Nil(t, user)
					assert.NotNil(t, errorAuthRegister)
				} else {
					assert.NotNil(t, user)
					assert.Nil(t, errorAuthRegister)

					statusDeleteOne, errorDeleteOne := userRepository.DeleteOne(userRepository.GetCriteriaByUserId(&user.Id))

					assert.True(t, statusDeleteOne)
					assert.Nil(t, errorDeleteOne)
				}
			},
		)
	}
}

func TestGetNewTokenSignedString(t *testing.T) {
	tests := []struct {
		Name        string
		User        *entity.User
		ExpiresAt   time.Duration
		MustBeFault bool
	}{
		{
			Name: "Test case with GetNewTokenSignedString",
			User: &entity.User{
				Id: uuid.New(),
				UserDTO: dto.UserDTO{
					UserRegisterDTO: dto.UserRegisterDTO{
						UserCredentialsDTO: dto.UserCredentialsDTO{
							Username: "username",
						},
					},
					Roles: kind.UserRoles{kind.UserRoleCommon},
				},
			},
			ExpiresAt:   time.Duration(1),
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				tokenSignedString, errorTokenSignedString := getNewTokenSignedString(testCase.User, testCase.ExpiresAt)

				if testCase.MustBeFault {
					assert.Nil(t, tokenSignedString)
					assert.NotNil(t, errorTokenSignedString)
				} else {
					assert.NotNil(t, tokenSignedString)
					assert.Nil(t, errorTokenSignedString)
				}
			},
		)
	}
}

func TestMakeTokens(t *testing.T) {
	tests := []struct {
		Name        string
		User        *entity.User
		MustBeFault bool
	}{
		{
			Name: "Test case with MakeTokens",
			User: &entity.User{
				Id: uuid.New(),
				UserDTO: dto.UserDTO{
					UserRegisterDTO: dto.UserRegisterDTO{
						UserCredentialsDTO: dto.UserCredentialsDTO{
							Username: "username",
						},
					},
					Roles: kind.UserRoles{kind.UserRoleCommon},
				},
			},
			MustBeFault: false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.Name,
			func(t *testing.T) {
				responseAuthToken, errorResponseAuthToken := makeTokens(testCase.User)

				if testCase.MustBeFault {
					assert.Nil(t, responseAuthToken)
					assert.NotNil(t, errorResponseAuthToken)
				} else {
					assert.NotNil(t, responseAuthToken)
					assert.NotNil(t, responseAuthToken.AccessToken)
					assert.NotNil(t, responseAuthToken.RefreshToken)
					assert.Nil(t, errorResponseAuthToken)
				}
			},
		)
	}
}
