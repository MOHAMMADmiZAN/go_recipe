package jwtToken

import (
	"context"
	"errors"
	"fmt"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"log"
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
// DecodeToken decodes a JSON Web Token (JWT) and returns the claims if valid.
func DecodeToken(ctx context.Context, tokenString string) (map[string]interface{}, error) {
	// Retrieve the secret key for validating the token
	secretKey := GetJwtSecret()
	if secretKey == "" {
		log.Println("JWT_SECRET is not set")
		return nil, errors.New("invalid token: JWT_SECRET is not set")
	}

	// Custom key function for token validation
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key for validation
		return []byte(secretKey), nil
	}

	// Parse the token using the custom key function
	parsedToken, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// Validate token
	if !parsedToken.Valid {
		return nil, errors.New("invalid token: token is not valid")
	}

	// Extract and validate claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token: unable to extract claims")
	}

	// Validate expiration time
	expiryTime, ok := claims["exp"].(float64)
	if !ok {
		return nil, errors.New("invalid token: expiration time not found in claims")
	}
	if int64(expiryTime) < time.Now().Unix() {
		return nil, errors.New("invalid token: token has expired")
	}
	// Token is valid, return the claims
	return claims, nil
}
