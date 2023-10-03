package appMiddleware

import (
	"context"
	userService "github.com/MOHAMMADmiZAN/go_recipe/internal/app/lib/user"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/appResponse"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/jwtToken"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
)

type User struct {
	ID    primitive.ObjectID `json:"id"`
	Name  string             `json:"name"`
	Email string             `json:"email"`
	Roles []string           `json:"roles"`
}

// AuthMiddleware is a middleware that authenticates users using JWT.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the request header.
		token := r.Header.Get("Authorization")
		if token == "" {
			appResponse.ResponseMessage(w, http.StatusUnauthorized, "Missing token")
			return
		}
		// Split the token into two parts.
		tokenParts := strings.Split(token, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			appResponse.ResponseMessage(w, http.StatusUnauthorized, "Invalid token")
			return
		}
		// Get the token value.
		tokenValue := tokenParts[1]
		// Decode the token.
		claims, err := jwtToken.DecodeToken(r.Context(), tokenValue)
		if err != nil {
			appResponse.ResponseMessage(w, http.StatusUnauthorized, err.Error())
			return
		}
		// Get the user ID from the token.
		userID := utils.HexToObjectId(claims["id"].(string))

		// Get the user from the database.
		user := userService.GetUserById(userID)
		if err != nil {
			appResponse.ResponseMessage(w, http.StatusUnauthorized, "failed to get user")
			return
		}
		payload := &User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Roles: user.Roles,
		}
		// Attach the user to the request context.
		ctx := context.WithValue(r.Context(), "user", payload)
		r = r.WithContext(ctx)

		// Call the next middleware.
		next.ServeHTTP(w, r)
	})
}
