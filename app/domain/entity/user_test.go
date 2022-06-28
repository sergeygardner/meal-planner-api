package entity

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/domain/dto"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	tests := []struct {
		name       string
		json       string
		Id         uuid.UUID
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
			name:       "Test case with status register and role admin and the other User properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"username\":\"Username\",\"password\":\"Password\",\"name\":\"Name\",\"surname\":\"Surname\",\"middle_name\":\"MiddleName\",\"birthday\":\"2000-01-01T00:00:00Z\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"status\":\"register\",\"active\":true,\"roles\":[\"admin\"]}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
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
			name:       "Test case with status register and role common and the other User properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"username\":\"Username\",\"password\":\"Password\",\"name\":\"Name\",\"surname\":\"Surname\",\"middle_name\":\"MiddleName\",\"birthday\":\"2000-01-01T00:00:00Z\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"status\":\"register\",\"active\":true,\"roles\":[\"common\"]}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
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
			name:       "Test case with status need confirmation and role common and the other User properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"username\":\"Username\",\"password\":\"Password\",\"name\":\"Name\",\"surname\":\"Surname\",\"middle_name\":\"MiddleName\",\"birthday\":\"2000-01-01T00:00:00Z\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"status\":\"need_confirmation\",\"active\":true,\"roles\":[\"common\"]}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
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
			name:       "Test case with status disabled and role admin and the other User properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"username\":\"Username\",\"password\":\"Password\",\"name\":\"Name\",\"surname\":\"Surname\",\"middle_name\":\"MiddleName\",\"birthday\":\"2000-01-01T00:00:00Z\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"status\":\"disabled\",\"active\":true,\"roles\":[\"admin\"]}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
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
			name:       "Test case with status disabled and role common and the other User properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"username\":\"Username\",\"password\":\"Password\",\"name\":\"Name\",\"surname\":\"Surname\",\"middle_name\":\"MiddleName\",\"birthday\":\"2000-01-01T00:00:00Z\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"status\":\"disabled\",\"active\":true,\"roles\":[\"common\"]}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
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
			name:       "Test case with status register and role admin and active false and the other User properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"username\":\"Username\",\"password\":\"Password\",\"name\":\"Name\",\"surname\":\"Surname\",\"middle_name\":\"MiddleName\",\"birthday\":\"2000-01-01T00:00:00Z\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"status\":\"register\",\"active\":false,\"roles\":[\"admin\"]}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
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
			name:       "Test case with status register and role common and active false and  the other User properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"username\":\"Username\",\"password\":\"Password\",\"name\":\"Name\",\"surname\":\"Surname\",\"middle_name\":\"MiddleName\",\"birthday\":\"2000-01-01T00:00:00Z\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"status\":\"register\",\"active\":false,\"roles\":[\"common\"]}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
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
			name:       "Test case with status need confirmation and role common and active false and  the other User properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"username\":\"Username\",\"password\":\"Password\",\"name\":\"Name\",\"surname\":\"Surname\",\"middle_name\":\"MiddleName\",\"birthday\":\"2000-01-01T00:00:00Z\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"status\":\"need_confirmation\",\"active\":false,\"roles\":[\"common\"]}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
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
			name:       "Test case with status disabled and role admin and active false and  the other User properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"username\":\"Username\",\"password\":\"Password\",\"name\":\"Name\",\"surname\":\"Surname\",\"middle_name\":\"MiddleName\",\"birthday\":\"2000-01-01T00:00:00Z\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"status\":\"disabled\",\"active\":false,\"roles\":[\"admin\"]}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
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
			name:       "Test case with status disabled and role common and active false and  the other User properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"username\":\"Username\",\"password\":\"Password\",\"name\":\"Name\",\"surname\":\"Surname\",\"middle_name\":\"MiddleName\",\"birthday\":\"2000-01-01T00:00:00Z\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"status\":\"disabled\",\"active\":false,\"roles\":[\"common\"]}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
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
				user := User{
					Id: testCase.Id,
					UserDTO: dto.UserDTO{
						UserRegisterDTO: dto.UserRegisterDTO{
							UserCredentialsDTO: dto.UserCredentialsDTO{
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
					},
				}
				assert.Equal(t, testCase.Id, user.Id)
				assert.Equal(t, testCase.Username, user.Username)
				assert.Equal(t, testCase.Password, user.Password)
				assert.Equal(t, testCase.Name, user.Name)
				assert.Equal(t, testCase.Surname, user.Surname)
				assert.Equal(t, testCase.MiddleName, user.MiddleName)
				assert.Equal(t, testCase.Birthday, user.Birthday)
				assert.Equal(t, testCase.DateInsert, user.DateInsert)
				assert.Equal(t, testCase.DateUpdate, user.DateUpdate)
				assert.Equal(t, testCase.Status, user.Status)
				assert.Equal(t, testCase.Active, user.Active)
				assert.Equal(t, testCase.Roles, user.Roles)

				reflectUser := reflect.ValueOf(user)

				for i := 0; i < reflectUser.NumField(); i++ {
					if reflectUser.Field(i).Type().String() != "bool" {
						assert.False(t, reflectUser.Field(i).IsZero())
					}
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(user)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}

func TestUserRole(t *testing.T) {
	tests := []struct {
		name       string
		json       string
		Id         uuid.UUID
		DateInsert time.Time
		DateUpdate time.Time
		Name       string
		Code       kind.UserRole
		Status     kind.UserRoleStatus
	}{
		{
			name:       "Test case with status active and code admin and other UserRole properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"name\":\"Admin\",\"code\":\"admin\",\"status\":\"active\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			DateInsert: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			Name:       "Admin",
			Code:       kind.UserRoleAdmin,
			Status:     kind.UserRoleStatusActive,
		},
		{
			name:       "Test case with status inactive and code admin and other UserRole properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"name\":\"Admin\",\"code\":\"admin\",\"status\":\"inactive\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			DateInsert: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			Name:       "Admin",
			Code:       kind.UserRoleAdmin,
			Status:     kind.UserRoleStatusInActive,
		},
		{
			name:       "Test case with status active and code common and other UserRole properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"name\":\"Common\",\"code\":\"common\",\"status\":\"active\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			DateInsert: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			Name:       "Common",
			Code:       kind.UserRoleCommon,
			Status:     kind.UserRoleStatusActive,
		},
		{
			name:       "Test case with status inactive and code common and other UserRole properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"name\":\"Common\",\"code\":\"common\",\"status\":\"inactive\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			DateInsert: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			Name:       "Common",
			Code:       kind.UserRoleCommon,
			Status:     kind.UserRoleStatusInActive,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				userRole := UserRole{
					Id:         testCase.Id,
					DateInsert: testCase.DateInsert,
					DateUpdate: testCase.DateUpdate,
					Name:       testCase.Name,
					Code:       testCase.Code,
					Status:     testCase.Status,
				}
				assert.Equal(t, testCase.Id, userRole.Id)
				assert.Equal(t, testCase.DateInsert, userRole.DateInsert)
				assert.Equal(t, testCase.DateUpdate, userRole.DateUpdate)
				assert.Equal(t, testCase.Name, userRole.Name)
				assert.Equal(t, testCase.Code, userRole.Code)
				assert.Equal(t, testCase.Status, userRole.Status)

				reflectUser := reflect.ValueOf(userRole)

				for i := 0; i < reflectUser.NumField(); i++ {
					assert.False(t, reflectUser.Field(i).IsZero())
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(userRole)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}

func TestUserToRole(t *testing.T) {
	tests := []struct {
		name       string
		json       string
		Id         uuid.UUID
		UserId     uuid.UUID
		RoleId     uuid.UUID
		DateInsert time.Time
		DateUpdate time.Time
		Rights     []kind.UserRight
		Status     kind.UserToRoleStatus
	}{
		{
			name:       "Test case with rights are write and read and status active and other UserToRole properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"role_id\":\"00000000-0000-0000-0000-000000000003\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"rights\":[\"write\",\"read\"],\"status\":\"active\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			RoleId:     uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			DateInsert: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			Rights:     []kind.UserRight{kind.UserRightWrite, kind.UserRightRead},
			Status:     kind.UserToRoleStatusActive,
		},
		{
			name:       "Test case with rights are write and read and status inactive and other UserToRole properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"role_id\":\"00000000-0000-0000-0000-000000000003\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"rights\":[\"write\",\"read\"],\"status\":\"inactive\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			RoleId:     uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			DateInsert: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			Rights:     []kind.UserRight{kind.UserRightWrite, kind.UserRightRead},
			Status:     kind.UserToRoleStatusInActive,
		},
		{
			name:       "Test case with rights is write and status active and other UserToRole properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"role_id\":\"00000000-0000-0000-0000-000000000003\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"rights\":[\"write\"],\"status\":\"active\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			RoleId:     uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			DateInsert: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			Rights:     []kind.UserRight{kind.UserRightWrite},
			Status:     kind.UserToRoleStatusActive,
		},
		{
			name:       "Test case with rights is write and status inactive and other UserToRole properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"role_id\":\"00000000-0000-0000-0000-000000000003\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"rights\":[\"write\"],\"status\":\"inactive\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			RoleId:     uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			DateInsert: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			Rights:     []kind.UserRight{kind.UserRightWrite},
			Status:     kind.UserToRoleStatusInActive,
		},
		{
			name:       "Test case with rights is read and status active and other UserToRole properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"role_id\":\"00000000-0000-0000-0000-000000000003\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"rights\":[\"read\"],\"status\":\"active\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			RoleId:     uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			DateInsert: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			Rights:     []kind.UserRight{kind.UserRightRead},
			Status:     kind.UserToRoleStatusActive,
		},
		{
			name:       "Test case with rights is read and status inactive and other UserToRole properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"role_id\":\"00000000-0000-0000-0000-000000000003\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"rights\":[\"read\"],\"status\":\"inactive\"}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			RoleId:     uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			DateInsert: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			Rights:     []kind.UserRight{kind.UserRightRead},
			Status:     kind.UserToRoleStatusInActive,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				userToRole := UserToRole{
					Id:         testCase.Id,
					UserId:     testCase.UserId,
					RoleId:     testCase.RoleId,
					DateInsert: testCase.DateInsert,
					DateUpdate: testCase.DateUpdate,
					Rights:     testCase.Rights,
					Status:     testCase.Status,
				}
				assert.Equal(t, testCase.Id, userToRole.Id)
				assert.Equal(t, testCase.UserId, userToRole.UserId)
				assert.Equal(t, testCase.RoleId, userToRole.RoleId)
				assert.Equal(t, testCase.DateInsert, userToRole.DateInsert)
				assert.Equal(t, testCase.DateUpdate, userToRole.DateUpdate)
				assert.Equal(t, testCase.Rights, userToRole.Rights)
				assert.Equal(t, testCase.Status, userToRole.Status)

				reflectUser := reflect.ValueOf(userToRole)

				for i := 0; i < reflectUser.NumField(); i++ {
					if reflectUser.Field(i).Type().String() != "bool" {
						assert.False(t, reflectUser.Field(i).IsZero())
					}
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(userToRole)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}

func TestUserConfirmation(t *testing.T) {
	tests := []struct {
		name       string
		json       string
		Id         uuid.UUID
		DateInsert time.Time
		DateUpdate time.Time
		UserId     uuid.UUID
		Value      string
		Active     bool
	}{
		{
			name:       "Test case with active true and other UserConfirmation properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"value\":\"424242\",\"active\":true}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			DateInsert: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			Value:      "424242",
			Active:     true,
		},
		{
			name:       "Test case with active false and other UserConfirmation properties",
			json:       "{\"id\":\"00000000-0000-0000-0000-000000000001\",\"date_insert\":\"2020-01-01T00:00:00Z\",\"date_update\":\"2020-01-10T00:00:00Z\",\"user_id\":\"00000000-0000-0000-0000-000000000002\",\"value\":\"424242\",\"active\":false}\n",
			Id:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			DateInsert: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			DateUpdate: time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC),
			UserId:     uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			Value:      "424242",
			Active:     false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				userConfirmation := UserConfirmation{
					Id:         testCase.Id,
					DateInsert: testCase.DateInsert,
					DateUpdate: testCase.DateUpdate,
					UserId:     testCase.UserId,
					Value:      testCase.Value,
					Active:     testCase.Active,
				}
				assert.Equal(t, testCase.Id, userConfirmation.Id)
				assert.Equal(t, testCase.DateInsert, userConfirmation.DateInsert)
				assert.Equal(t, testCase.DateUpdate, userConfirmation.DateUpdate)
				assert.Equal(t, testCase.UserId, userConfirmation.UserId)
				assert.Equal(t, testCase.Value, userConfirmation.Value)
				assert.Equal(t, testCase.Active, userConfirmation.Active)

				reflectUser := reflect.ValueOf(userConfirmation)

				for i := 0; i < reflectUser.NumField(); i++ {
					if reflectUser.Field(i).Type().String() != "bool" {
						assert.False(t, reflectUser.Field(i).IsZero())
					}
				}

				var actual strings.Builder
				enc := json.NewEncoder(&actual)
				enc.SetIndent(">", ".")
				enc.SetIndent("", "")

				errorEncode := enc.Encode(userConfirmation)

				assert.Nil(t, errorEncode)
				assert.Equal(t, testCase.json, actual.String())
			},
		)
	}
}
