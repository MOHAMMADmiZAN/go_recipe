package encrypt

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// SaltRounds is the number of rounds to use when generating bcrypt salts.
const SaltRounds = 10

// HashPassword hashes a password using bcrypt.
func HashPassword(ctx context.Context, password string) (string, error) {
	salt, err := bcrypt.GenerateFromPassword([]byte(password), SaltRounds)
	if err != nil {
		return "", fmt.Errorf("failed to generate bcrypt salt: %w", err)
	}

	return string(salt), nil

}

// ComparePassword compares a plain text password to a hashed password.
func ComparePassword(ctx context.Context, password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			// The passwords do not match.
			return fmt.Errorf("passwords do not match: %w", err)
		}

		// An unexpected error occurred.
		return fmt.Errorf("failed to compare bcrypt passwords: %w", err)
	}

	return nil
}
