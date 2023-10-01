package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	// Register custom validation
	err := validate.RegisterValidation("roleEnum", validateRoles)
	if err != nil {
		panic(err)
	}
}

// ValidateStruct validates a struct and returns validation errors.
func ValidateStruct(input interface{}) map[string]string {
	err := validate.Struct(input)
	if err != nil {
		validationErrors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			var message string

			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s is required", field)
			case "email":
				message = "Invalid email address"
			case "min":
				message = fmt.Sprintf("%s must be at least %s characters long", field, err.Param())
			case "roleEnum":
				message = fmt.Sprintf("%s must be either 'admin' or 'user'", field)
			default:
				message = fmt.Sprintf("Validation failed for %s", field)
			}

			validationErrors[field] = message
		}
		return validationErrors
	}
	return nil
}

// validateRoles validates the roles field.
func validateRoles(fl validator.FieldLevel) bool {
	// Allowed roles
	allowedRoles := map[string]bool{"admin": true, "user": true}
	// Get the roles from the field
	roles := fl.Field().Interface().([]string)

	// Iterate through roles and check if they are allowed
	for _, role := range roles {
		if !allowedRoles[role] {
			// If role is not allowed, return a validation error
			return false
		}
	}
	// All roles are allowed
	return true
}
