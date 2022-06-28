package directive

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
)

// Auth /** @see service.CheckTokenFromContext
func Auth(ctx context.Context, _ interface{}, next graphql.Resolver) (interface{}, error) {
	// the validation happens in schema.resolver because we have queries with auth and without it
	// token has obtained with the middleware "jwtauth.Verifier"
	// this is only for config.Directives.Auth as a stub

	return next(ctx)
}
