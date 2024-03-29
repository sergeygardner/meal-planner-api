package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"errors"
	"fmt"
	"github.com/sergeygardner/meal-planner-api/ui/graphql/service"
	RestService "github.com/sergeygardner/meal-planner-api/ui/rest/service"
	"net/http"

	"github.com/sergeygardner/meal-planner-api/application/handler"
	"github.com/sergeygardner/meal-planner-api/domain/dto"
	"github.com/sergeygardner/meal-planner-api/domain/entity"
	"github.com/sergeygardner/meal-planner-api/domain/response"
	"github.com/sergeygardner/meal-planner-api/ui/graphql/model"
)

var (
	errorAuthenticationIsRequired    = errors.New("authentication is required")
	errorAuthenticationIsNotRequired = errors.New("authentication is not required")
)

// Auth is the resolver for the auth field.
func (r *mutationResolver) Auth(_ context.Context) (*model.AuthOps, error) {
	return &model.AuthOps{}, nil
}

// AuthCheck is the resolver for the AuthCheck field.
func (r *queryResolver) AuthCheck(ctx context.Context) (*string, error) {
	if !service.CheckTokenFromContext(ctx) {
		return nil, errorAuthenticationIsRequired
	}

	message := fmt.Sprintf("status is %d", http.StatusOK)

	return &message, nil
}

// AuthCredentials is the resolver for the AuthCredentials field.
func (r *queryResolver) AuthCredentials(ctx context.Context, input dto.UserCredentialsDTO) (*response.AuthConfirmation, error) {
	if service.CheckTokenFromContext(ctx) {
		return nil, errorAuthenticationIsNotRequired
	}

	authConfirmation, _, errorAuthCredentials := handler.AuthCredentials(input)

	if errorAuthCredentials != nil {
		return nil, errorAuthCredentials
	}

	return authConfirmation, nil
}

// AuthConfirmation is the resolver for the AuthConfirmation field.
func (r *queryResolver) AuthConfirmation(ctx context.Context, input dto.AuthConfirmationDTO) (*response.AuthToken, error) {
	if service.CheckTokenFromContext(ctx) {
		return nil, errorAuthenticationIsNotRequired
	}

	authToken, errorAuthConfirmation := handler.AuthConfirmation(input)

	if errorAuthConfirmation != nil {
		return nil, errorAuthConfirmation
	}

	return authToken, nil
}

// AuthRegister is the resolver for the AuthRegister field.
func (r *queryResolver) AuthRegister(ctx context.Context, input dto.UserRegisterDTO) (*entity.User, error) {
	if service.CheckTokenFromContext(ctx) {
		return nil, errorAuthenticationIsNotRequired
	}

	user, errorAuthRegister := handler.AuthRegister(input)

	if errorAuthRegister != nil {
		return nil, errorAuthRegister
	}

	return user, nil
}

// AuthRefresh is the resolver for the AuthRefresh field.
func (r *queryResolver) AuthRefresh(ctx context.Context) (*response.AuthToken, error) {
	if !service.CheckTokenFromContext(ctx) {
		return nil, errorAuthenticationIsRequired
	}

	token, errorExtractClaimsFromContext := RestService.ExtractClaimsFromContext(ctx)

	if errorExtractClaimsFromContext != nil {
		return nil, errorExtractClaimsFromContext
	}

	authToken, errorAuthTokenByUserId := handler.AuthTokenByUserId(&token.UserId)

	if errorAuthTokenByUserId != nil {
		return nil, errorAuthTokenByUserId
	}

	return authToken, nil
}

// ID is the resolver for the id field.
func (r *userResolver) ID(_ context.Context, obj *entity.User) (string, error) {
	return obj.Id.String(), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
