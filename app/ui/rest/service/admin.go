package service

import (
	"github.com/sergeygardner/meal-planner-api/domain/kind"
	"net/http"
)

func EnsureAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, errorExtractRoleFromClaims := ExtractClaimsFromContext(r.Context())

		if errorExtractRoleFromClaims != nil {
			http.Error(w, errorExtractRoleFromClaims.Error(), 401)

			return
		}

		roleExists, errorRoleExists := token.EnsureRoleExists(kind.UserRoleAdmin)

		if errorRoleExists != nil {
			http.Error(w, errorRoleExists.Error(), 401)

			return
		} else if !roleExists {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		next.ServeHTTP(w, r)
	})
}
