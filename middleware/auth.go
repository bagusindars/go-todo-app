package middleware

import (
	"context"
	"net/http"
	"simple-todo-app/internal/helpers"
	"strings"
)

func AuthJWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			helpers.ApiResponse(w, 401, "Authorization header is required", nil)
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		if token == "" {
			helpers.ApiResponse(w, 401, "Unauthorized", nil)
			return
		}

		claims, err := helpers.ValidateToken(token)

		if err != nil {
			helpers.ApiResponse(w, 401, "Unauthorized", nil)
			return
		}

		ctx := context.WithValue(context.Background(), "userInfo", claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
