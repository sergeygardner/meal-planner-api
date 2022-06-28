package service

import (
	"context"
	"github.com/go-chi/jwtauth/v5"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	"github.com/sergeygardner/meal-planner-api/domain/model"
)

var errorParentId = errors.New("the picture is not found by criteria")

func ExtractClaimsFromContext(ctx context.Context) (*model.Token, error) {
	token, claims, errorFromContext := jwtauth.FromContext(ctx)

	if errorFromContext != nil {
		return nil, errorFromContext
	}

	userId, okUserId := claims["user_id"]
	username, okUsername := claims["username"]
	userRoles, okUserRoles := claims["user_roles"]

	if !okUserId || !okUsername || !okUserRoles {
		return nil, errors.New("error get claims from context")
	}

	uuidParsed, errorUuidParsed := uuid.Parse(userId.(string))
	usernameParsed, okUsernameParsed := username.(string)
	rolesInterfaces := userRoles.([]interface{})
	userRolesParsed := make(kind.UserRoles, len(rolesInterfaces))

	for i, value := range rolesInterfaces {
		okUserRoleString, okUserRolesParsed := value.(string)

		if !okUserRolesParsed {
			return nil, errors.New("error get claims from context")
		}

		userRolesParsed[i] = kind.UserRole(okUserRoleString)
	}

	if errorUuidParsed != nil || !okUsernameParsed {
		return nil, errors.New("error get claims from context")
	}

	return &model.Token{
		UserId:    uuidParsed,
		Username:  usernameParsed,
		UserRoles: userRolesParsed,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(token.Expiration()),
		}}, nil
}

func GetParentId(keys []string, values []string, exclude string) (*uuid.UUID, error) {
	keyIDs := []string{"picture_id", "category_id", "ingredient_id", "process_id", "recipe_id"}

	for _, keyId := range keyIDs {
		for key, value := range keys {
			if exclude != keyId && value == keyId {
				id, errorId := uuid.Parse(values[key])

				if errorId != nil {
					return nil, errorId
				}

				return &id, nil
			}
		}
	}

	return nil, errorParentId
}
