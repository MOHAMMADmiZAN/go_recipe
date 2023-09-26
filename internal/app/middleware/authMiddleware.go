package middleware

//
//import (
//	"context"
//	"fmt"
//	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/jwtToken"
//	"net/http"
//	"strings"
//)
//
//// AuthMiddleware is a middleware that authenticates users using JWT.
//func AuthMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//
//		// Get the token from the request header.
//		token := r.Header.Get("Authorization")
//		if token == "" {
//			http.Error(w, "Unauthorized", http.StatusUnauthorized)
//			return
//		}
//
//		// Split the token into two parts.
//		tokenParts := strings.Split(token, " ")
//		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
//			http.Error(w, "Invalid token format", http.StatusUnauthorized)
//			return
//		}
//
//		// Get the token value.
//		tokenValue := tokenParts[1]
//
//		// Decode the token.
//		claims, err := jwtToken.DecodeToken(r.Context(), tokenValue)
//		if err != nil {
//			http.Error(w, fmt.Sprintf("Failed to decode token: %v", err), http.StatusUnauthorized)
//			return
//		}
//
//		// Get the user ID from the token.
//		userID := claims["id"].(int)
//
//		// Get the user from the database.
//		user, err := GetUserById(r.Context(), userID)
//		if err != nil {
//			http.Error(w, fmt.Sprintf("Failed to get user: %v", err), http.StatusUnauthorized)
//			return
//		}
//
//		// Attach the user to the request context.
//		ctx := context.WithValue(r.Context(), "user", user)
//		r = r.WithContext(ctx)
//
//		// Call the next middleware.
//		next.ServeHTTP(w, r)
//	})
//}
