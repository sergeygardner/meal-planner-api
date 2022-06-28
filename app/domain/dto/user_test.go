package dto

import (
	"encoding/json"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestUserCredentialsDTO(t *testing.T) {
	tests := []struct {
		name     string
		json     string
		Username string
		Password string
	}{
		{
			name:     "Test case with UserCredentialsDTO properties",
			json:     "{\"username\":\"Username\",\"password\":\"Password\"}\n",
			Username: "Username",
			Password: "Password",
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				userCredentialsDTO := UserCredentialsDTO{
					Username: testCase.Username,
					Password: testCase.Password,
				}
				assert.Equal(t, testCase.Username, userCredentialsDTO.Username)
				assert.Equal(t, testCase.Password, userCredentialsDTO.Password)

				reflectUserCredentialsDTO := reflect.ValueOf(userCredentialsDTO)

				for i := 0; i < reflectUserCredentialsDTO.NumField(); i++ {
					assert.False(t, reflectUserCredentialsDTO.Field(i).IsZero())
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(userCredentialsDTO)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}

func TestUserRegisterDTO(t *testing.T) {
	tests := []struct {
		name       string
		json       string
		Username   string
		Password   string
		Name       string
		Surname    string
		MiddleName string
		Birthday   time.Time
	}{
		{
			name:       "Test case with UserRegisterDTO properties",
			json:       "{\"username\":\"Username\",\"password\":\"Password\",\"name\":\"Name\",\"surname\":\"Surname\",\"middle_name\":\"MiddleName\",\"birthday\":\"2000-01-01T00:00:00Z\"}\n",
			Username:   "Username",
			Password:   "Password",
			Name:       "Name",
			Surname:    "Surname",
			MiddleName: "MiddleName",
			Birthday:   time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				userRegisterDTO := UserRegisterDTO{
					UserCredentialsDTO: UserCredentialsDTO{
						Username: testCase.Username,
						Password: testCase.Password,
					},
					Name:       testCase.Name,
					Surname:    testCase.Surname,
					MiddleName: testCase.MiddleName,
					Birthday:   testCase.Birthday,
				}
				assert.Equal(t, testCase.Username, userRegisterDTO.Username)
				assert.Equal(t, testCase.Password, userRegisterDTO.Password)
				assert.Equal(t, testCase.Name, userRegisterDTO.Name)
				assert.Equal(t, testCase.Surname, userRegisterDTO.Surname)
				assert.Equal(t, testCase.MiddleName, userRegisterDTO.MiddleName)
				assert.Equal(t, testCase.Birthday, userRegisterDTO.Birthday)

				reflectUserRegisterDTO := reflect.ValueOf(userRegisterDTO)

				for i := 0; i < reflectUserRegisterDTO.NumField(); i++ {
					assert.False(t, reflectUserRegisterDTO.Field(i).IsZero())
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(userRegisterDTO)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}

func TestUserDTO(t *testing.T) {
	tests := []struct {
		name       string
		json       string
		Username   string
		Password   string
		Name       string
		Surname    string
		MiddleName string
		Birthday   time.Time
		DateInsert time.Time
		DateUpdate time.Time
		Status     kind.UserStatus
		Active     bool
		Roles      kind.UserRoles
	}{
		{
			name:       "Test case with status register and role admin and the other UserDTO properties",
			json:       "{\"username\":\"Username\",\"password\":\"Password\",\"name\":\"Name\",\"surname\":\"Surname\",\"middle_name\":\"MiddleName\",\"birthday\":\"2000-01-01T00:00:00Z\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"status\":\"register\",\"active\":true,\"roles\":[\"admin\"]}\n",
			Username:   "Username",
			Password:   "Password",
			Name:       "Name",
			Surname:    "Surname",
			MiddleName: "MiddleName",
			Birthday:   time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateInsert: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			Status:     kind.UserStatusRegister,
			Active:     true,
			Roles:      kind.UserRoles{kind.UserRoleAdmin},
		},
		{
			name:       "Test case with status register and role common and the other UserDTO properties",
			json:       "{\"username\":\"Username\",\"password\":\"Password\",\"name\":\"Name\",\"surname\":\"Surname\",\"middle_name\":\"MiddleName\",\"birthday\":\"2000-01-01T00:00:00Z\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"status\":\"register\",\"active\":true,\"roles\":[\"common\"]}\n",
			Username:   "Username",
			Password:   "Password",
			Name:       "Name",
			Surname:    "Surname",
			MiddleName: "MiddleName",
			Birthday:   time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateInsert: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			Status:     kind.UserStatusRegister,
			Active:     true,
			Roles:      kind.UserRoles{kind.UserRoleCommon},
		},
		{
			name:       "Test case with status need confirmation and role common and the other UserDTO properties",
			json:       "{\"username\":\"Username\",\"password\":\"Password\",\"name\":\"Name\",\"surname\":\"Surname\",\"middle_name\":\"MiddleName\",\"birthday\":\"2000-01-01T00:00:00Z\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"status\":\"need_confirmation\",\"active\":true,\"roles\":[\"common\"]}\n",
			Username:   "Username",
			Password:   "Password",
			Name:       "Name",
			Surname:    "Surname",
			MiddleName: "MiddleName",
			Birthday:   time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateInsert: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			Status:     kind.UserStatusNeedConfirmation,
			Active:     true,
			Roles:      kind.UserRoles{kind.UserRoleCommon},
		},
		{
			name:       "Test case with status disabled and role admin and the other UserDTO properties",
			json:       "{\"username\":\"Username\",\"password\":\"Password\",\"name\":\"Name\",\"surname\":\"Surname\",\"middle_name\":\"MiddleName\",\"birthday\":\"2000-01-01T00:00:00Z\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"status\":\"disabled\",\"active\":true,\"roles\":[\"admin\"]}\n",
			Username:   "Username",
			Password:   "Password",
			Name:       "Name",
			Surname:    "Surname",
			MiddleName: "MiddleName",
			Birthday:   time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateInsert: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			Status:     kind.UserStatusDisabled,
			Active:     true,
			Roles:      kind.UserRoles{kind.UserRoleAdmin},
		},
		{
			name:       "Test case with status disabled and role common and the other UserDTO properties",
			json:       "{\"username\":\"Username\",\"password\":\"Password\",\"name\":\"Name\",\"surname\":\"Surname\",\"middle_name\":\"MiddleName\",\"birthday\":\"2000-01-01T00:00:00Z\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"status\":\"disabled\",\"active\":true,\"roles\":[\"common\"]}\n",
			Username:   "Username",
			Password:   "Password",
			Name:       "Name",
			Surname:    "Surname",
			MiddleName: "MiddleName",
			Birthday:   time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateInsert: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			Status:     kind.UserStatusDisabled,
			Active:     true,
			Roles:      kind.UserRoles{kind.UserRoleCommon},
		},
		{
			name:       "Test case with status register and role admin and active false and the other UserDTO properties",
			json:       "{\"username\":\"Username\",\"password\":\"Password\",\"name\":\"Name\",\"surname\":\"Surname\",\"middle_name\":\"MiddleName\",\"birthday\":\"2000-01-01T00:00:00Z\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"status\":\"register\",\"active\":false,\"roles\":[\"admin\"]}\n",
			Username:   "Username",
			Password:   "Password",
			Name:       "Name",
			Surname:    "Surname",
			MiddleName: "MiddleName",
			Birthday:   time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateInsert: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			Status:     kind.UserStatusRegister,
			Active:     false,
			Roles:      kind.UserRoles{kind.UserRoleAdmin},
		},
		{
			name:       "Test case with status register and role common and active false and  the other UserDTO properties",
			json:       "{\"username\":\"Username\",\"password\":\"Password\",\"name\":\"Name\",\"surname\":\"Surname\",\"middle_name\":\"MiddleName\",\"birthday\":\"2000-01-01T00:00:00Z\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"status\":\"register\",\"active\":false,\"roles\":[\"common\"]}\n",
			Username:   "Username",
			Password:   "Password",
			Name:       "Name",
			Surname:    "Surname",
			MiddleName: "MiddleName",
			Birthday:   time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateInsert: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			Status:     kind.UserStatusRegister,
			Active:     false,
			Roles:      kind.UserRoles{kind.UserRoleCommon},
		},
		{
			name:       "Test case with status need confirmation and role common and active false and  the other UserDTO properties",
			json:       "{\"username\":\"Username\",\"password\":\"Password\",\"name\":\"Name\",\"surname\":\"Surname\",\"middle_name\":\"MiddleName\",\"birthday\":\"2000-01-01T00:00:00Z\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"status\":\"need_confirmation\",\"active\":false,\"roles\":[\"common\"]}\n",
			Username:   "Username",
			Password:   "Password",
			Name:       "Name",
			Surname:    "Surname",
			MiddleName: "MiddleName",
			Birthday:   time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateInsert: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			Status:     kind.UserStatusNeedConfirmation,
			Active:     false,
			Roles:      kind.UserRoles{kind.UserRoleCommon},
		},
		{
			name:       "Test case with status disabled and role admin and active false and  the other UserDTO properties",
			json:       "{\"username\":\"Username\",\"password\":\"Password\",\"name\":\"Name\",\"surname\":\"Surname\",\"middle_name\":\"MiddleName\",\"birthday\":\"2000-01-01T00:00:00Z\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"status\":\"disabled\",\"active\":false,\"roles\":[\"admin\"]}\n",
			Username:   "Username",
			Password:   "Password",
			Name:       "Name",
			Surname:    "Surname",
			MiddleName: "MiddleName",
			Birthday:   time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateInsert: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			Status:     kind.UserStatusDisabled,
			Active:     false,
			Roles:      kind.UserRoles{kind.UserRoleAdmin},
		},
		{
			name:       "Test case with status disabled and role common and active false and  the other UserDTO properties",
			json:       "{\"username\":\"Username\",\"password\":\"Password\",\"name\":\"Name\",\"surname\":\"Surname\",\"middle_name\":\"MiddleName\",\"birthday\":\"2000-01-01T00:00:00Z\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"status\":\"disabled\",\"active\":false,\"roles\":[\"common\"]}\n",
			Username:   "Username",
			Password:   "Password",
			Name:       "Name",
			Surname:    "Surname",
			MiddleName: "MiddleName",
			Birthday:   time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateInsert: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			Status:     kind.UserStatusDisabled,
			Active:     false,
			Roles:      kind.UserRoles{kind.UserRoleCommon},
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				userDTO := UserDTO{
					UserRegisterDTO: UserRegisterDTO{
						UserCredentialsDTO: UserCredentialsDTO{
							Username: testCase.Username,
							Password: testCase.Password,
						},
						Name:       testCase.Name,
						Surname:    testCase.Surname,
						MiddleName: testCase.MiddleName,
						Birthday:   testCase.Birthday,
					},
					DateInsert: testCase.DateInsert,
					DateUpdate: testCase.DateUpdate,
					Status:     testCase.Status,
					Active:     testCase.Active,
					Roles:      testCase.Roles,
				}
				assert.Equal(t, testCase.Username, userDTO.Username)
				assert.Equal(t, testCase.Password, userDTO.Password)
				assert.Equal(t, testCase.Name, userDTO.Name)
				assert.Equal(t, testCase.Surname, userDTO.Surname)
				assert.Equal(t, testCase.MiddleName, userDTO.MiddleName)
				assert.Equal(t, testCase.Birthday, userDTO.Birthday)
				assert.Equal(t, testCase.DateInsert, userDTO.DateInsert)
				assert.Equal(t, testCase.DateUpdate, userDTO.DateUpdate)
				assert.Equal(t, testCase.Status, userDTO.Status)
				assert.Equal(t, testCase.Active, userDTO.Active)
				assert.Equal(t, testCase.Roles, userDTO.Roles)

				reflectUserDTO := reflect.ValueOf(userDTO)

				for i := 0; i < reflectUserDTO.NumField(); i++ {
					if reflectUserDTO.Field(i).Type().String() != "bool" {
						assert.False(t, reflectUserDTO.Field(i).IsZero())
					}
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(userDTO)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}
