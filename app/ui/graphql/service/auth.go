package service

import (
	"context"
	"github.com/go-chi/jwtauth/v5"
)

func CheckTokenFromContext(ctx context.Context) bool {
	return ctx.Value(jwtauth.TokenCtxKey) != nil
}
