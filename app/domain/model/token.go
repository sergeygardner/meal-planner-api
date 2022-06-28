package model

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	"golang.org/x/exp/slices"
	"time"
)

const (
	JWTAccessTokenExpire  = 15 * time.Minute
	JWTRefreshTokenExpire = 168 * time.Hour
)

type Token struct {
	UserId    uuid.UUID      `json:"user_id"`
	Username  string         `json:"username"`
	UserRoles kind.UserRoles `json:"user_roles"`
	jwt.RegisteredClaims
}

func (t *Token) EnsureRoleExists(role kind.UserRole) (bool, error) {
	index := slices.Index(t.UserRoles, role)

	if index < 0 {
		return false, errors.New("role is not found in UserRoles")
	}

	return true, nil
}
