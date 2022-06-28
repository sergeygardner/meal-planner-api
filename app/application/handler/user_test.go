package handler

import (
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/domain/dto"
	"github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	testUser   *entity.User
	testUserId = uuid.New()
)

func TestUserInfo(t *testing.T) {
	tests := []struct {
		Name            string
		UserRegisterDTO dto.UserRegisterDTO
		MustBePanic     bool
		MustBeFault     bool
	}{
		{
			Name: "Test case with UserInfo",
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

				user, errorAuthRegister := AuthRegister(testCase.UserRegisterDTO)

				if testCase.MustBeFault {
					assert.Nil(t, user)
					assert.NotNil(t, errorAuthRegister)
				} else {
					assert.NotNil(t, user)
					assert.Nil(t, errorAuthRegister)

					testUser = user
				}
			},
		)
	}
}

func TestUserUpdate(t *testing.T) {
	tests := []struct {
		Name        string
		User        *entity.User
		UserDTO     *dto.UserDTO
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with UserUpdate",
			User:        testUser,
			UserDTO:     &dto.UserDTO{},
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

				userUpdated, errorUserUpdate := UserUpdate(&testCase.User.Id, testCase.UserDTO)

				if testCase.MustBeFault {
					assert.Nil(t, userUpdated)
					assert.NotNil(t, errorUserUpdate)
				} else {
					assert.NotNil(t, userUpdated)
					assert.Nil(t, errorUserUpdate)

					testUser = userUpdated
				}
			},
		)
	}
}

func TestUserDelete(t *testing.T) {
	tests := []struct {
		Name        string
		User        *entity.User
		MustBePanic bool
		MustBeFault bool
	}{
		{
			Name:        "Test case with UserDelete",
			User:        testUser,
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

				userStatusDeleted, errorUserDelete := UserDelete(&testCase.User.Id)

				testUser = nil

				if testCase.MustBeFault {
					assert.False(t, userStatusDeleted)
					assert.NotNil(t, errorUserDelete)
				} else {
					assert.True(t, userStatusDeleted)
					assert.Nil(t, errorUserDelete)
				}
			},
		)
	}
}
