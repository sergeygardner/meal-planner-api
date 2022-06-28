package model

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	tests := []struct {
		name           string
		UserId         uuid.UUID
		Username       string
		UserRoles      kind.UserRoles
		UserRoleWanted kind.UserRole
		MustBePassed   bool
	}{
		{
			name:           "Test case with a user role exists",
			UserId:         uuid.New(),
			Username:       "Username",
			UserRoles:      kind.UserRoles{kind.UserRoleAdmin},
			UserRoleWanted: kind.UserRoleAdmin,
			MustBePassed:   true,
		},
		{
			name:           "Test case with a user role doesn't exist",
			UserId:         uuid.New(),
			Username:       "Username1",
			UserRoles:      kind.UserRoles{kind.UserRoleCommon},
			UserRoleWanted: kind.UserRoleAdmin,
			MustBePassed:   false,
		},
	}

	for _, testCase := range tests {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				token := Token{
					UserId:    testCase.UserId,
					Username:  testCase.Username,
					UserRoles: testCase.UserRoles,
					RegisteredClaims: jwt.RegisteredClaims{
						Issuer:    testCase.name,
						Subject:   "testing",
						Audience:  jwt.ClaimStrings{"testing"},
						ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Duration(10)).UTC()},
						NotBefore: &jwt.NumericDate{Time: time.Now().Add(time.Duration(1)).UTC()},
						IssuedAt:  &jwt.NumericDate{Time: time.Now().Add(time.Duration(-10)).UTC()},
						ID:        uuid.New().String(),
					},
				}
				assert.Equal(t, testCase.UserId, token.UserId)
				assert.Equal(t, testCase.Username, token.Username)
				assert.Equal(t, testCase.UserRoles, token.UserRoles)

				ensureRoleExists, errorEnsureRoleExists := token.EnsureRoleExists(testCase.UserRoleWanted)

				if testCase.MustBePassed {
					assert.True(t, ensureRoleExists)
					assert.Nil(t, errorEnsureRoleExists)
				} else {
					assert.False(t, ensureRoleExists)
					assert.NotNil(t, errorEnsureRoleExists)
				}

				reflectToken := reflect.ValueOf(token)

				for i := 0; i < reflectToken.NumField(); i++ {
					assert.False(t, reflectToken.Field(i).IsZero())
				}
			},
		)
	}
}
