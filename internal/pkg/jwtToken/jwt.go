package jwtToken

import (
	"context"
	"errors"
	"fmt"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

// GetJwtSecret gets the JWT secret from the environment variable `JWT_SECRET`.
func GetJwtSecret() string {
	utils.LoadEnv()
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return ""
	}
	return secretKey
}

// EncodeToken encodes data into a JSON Web Token (JWT) with an expiration time.
func EncodeToken(ctx context.Context, data map[string]interface{}, expiryTime time.Time) (string, error) {
	secretKey := GetJwtSecret()
	currentTime := time.Now()
	if expiryTime.Before(currentTime) {
		return "", errors.New("expiration time should be in the future")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(data))
	token.Claims.(jwt.MapClaims)["exp"] = expiryTime.Unix()

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to encode token: %w", err)
	}
	return signedToken, nil
}

// DecodeToken decodes a JSON Web Token (JWT) and retrieves the data, validating the expiration time.
func DecodeToken(ctx context.Context, token string) (map[string]interface{}, error) {
	secretKey := GetJwtSecret()
	validationKey := func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	}

	parsedToken, err := jwt.Parse(token, validationKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode token: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	// Validate the expiration time
	expiryTime, ok := claims["exp"].(int64)
	if !ok {
		return nil, fmt.Errorf("token does not contain an expiration time")
	}

	if time.Now().Unix() > expiryTime {
		return nil, fmt.Errorf("token has expired")
	}

	return claims, nil
}
