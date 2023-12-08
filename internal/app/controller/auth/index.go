package auth

import (
	"context"
	"encoding/json"
	userService "github.com/MOHAMMADmiZAN/go_recipe/internal/app/lib/user"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/appResponse"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/encrypt"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/jwtToken"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/validator"
	"net/http"
	"time"
)

type createRequestUser struct {
	Name     string   `json:"name" validate:"required"`
	Email    string   `json:"email" validate:"required,email"`
	Password string   `json:"password" validate:"required,min=6"`
	Roles    []string `json:"roles,omitempty" validate:"roleEnum"`
}

type loginRequestedUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type authUser struct {
	ID    string   `json:"id"`
	Email string   `json:"email"`
	Roles []string `json:"roles"`
}

type LoginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID    string   `json:"id"`
		Email string   `json:"email"`
		Roles []string `json:"roles"`
		Token string   `json:"token"`
	} `json:"data"`
	Links struct {
		Logout  Link `json:"logout"`
		Profile Link `json:"profile"`
	} `json:"links"`
}

type Link struct {
	Rel    string `json:"rel"`
	Href   string `json:"href"`
	Method string `json:"method"`
}

type createdUserResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Links   struct {
		SignIn struct {
			Rel    string `json:"rel"`
			Href   string `json:"href"`
			Method string `json:"method"`
		} `json:"sign_in"`
	} `json:"links"`
}

func SignUP(w http.ResponseWriter, r *http.Request) {
	var newUser createRequestUser
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		appResponse.ResponseMessage(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if validationErrors := validateUserData(newUser); validationErrors != nil {
		appResponse.ResponseMessage(w, http.StatusBadRequest, validationErrors)
		return
	}

	if userService.DuplicateUser(newUser.Email) {
		appResponse.ResponseMessage(w, http.StatusBadRequest, "User Already Exists")
		return
	}

	hashedPassword, err := encryptPassword(newUser.Password)
	if err != nil {
		appResponse.ResponseMessage(w, http.StatusBadRequest, "Password Hash Failed")
		return
	}

	err = userService.CreateUser(newUser.Name, newUser.Email, newUser.Roles, hashedPassword)
	if err != nil {
		appResponse.ResponseMessage(w, http.StatusBadRequest, "User Create Failed")
		return
	}

	// Create and send the response
	response := createResponse("/auth/signin")
	appResponse.RawResponse(w, http.StatusCreated, response)
}

func SignIN(w http.ResponseWriter, r *http.Request) {
	var authUserRequest loginRequestedUser
	if err := json.NewDecoder(r.Body).Decode(&authUserRequest); err != nil {
		appResponse.ResponseMessage(w, http.StatusUnauthorized, "Invalid Credential")
		return
	}
	if err := validateLoginData(authUserRequest); err != nil {
		appResponse.ResponseMessage(w, http.StatusBadRequest, err)
		return
	}
	// find user exit or not
	existUserPassword := userService.ExistsUserPassword(w, authUserRequest.Email)
	// password compare
	if err := encrypt.ComparePassword(context.Background(), authUserRequest.Password, existUserPassword); err != nil {
		appResponse.ResponseMessage(w, http.StatusUnauthorized, "Invalid Credential")
		return
	}
	// get User
	user := userService.GetUser(w, authUserRequest.Email)
	// token payload map
	payload := map[string]interface{}{
		"id":    user.ID.Hex(),
		"email": user.Email,
		"roles": user.Roles,
	}

	// Set the expiration time for the token (e.g., 60 minutes)
	expirationTime := time.Now().Add(120 * time.Minute)
	// Encode the payload into a JWT token
	token, err := jwtToken.EncodeToken(context.Background(), payload, expirationTime)
	if err != nil {
		appResponse.ResponseMessage(w, http.StatusBadRequest, err)
		return
	}

	response := LoginUserResponse(authUser{
		ID:    user.ID.Hex(),
		Email: user.Email,
		Roles: user.Roles,
	}, token)
	appResponse.RawResponse(w, 200, response)

}

// validate requested user Data
func validateUserData(user createRequestUser) map[string]string {
	return validator.ValidateStruct(user)
}

// login user Data Validation
func validateLoginData(user loginRequestedUser) map[string]string {
	return validator.ValidateStruct(user)

}

// password encrypt
func encryptPassword(password string) (string, error) {
	ctx := context.Background()
	return encrypt.HashPassword(ctx, password)
}

// create Response
func createResponse(signInURL string) createdUserResponse {
	return createdUserResponse{
		Code:    http.StatusCreated,
		Message: "User created successfully",
		Links: struct {
			SignIn struct {
				Rel    string `json:"rel"`
				Href   string `json:"href"`
				Method string `json:"method"`
			} `json:"sign_in"`
		}{
			SignIn: struct {
				Rel    string `json:"rel"`
				Href   string `json:"href"`
				Method string `json:"method"`
			}{
				Rel:    "sign_in",
				Href:   signInURL,
				Method: "POST",
			},
		},
	}
}

// LoginUserResponse login Response
func LoginUserResponse(user authUser, token string) LoginResponse {
	// Create the response object
	response := LoginResponse{
		Code:    http.StatusOK,
		Message: "User login successfully",
		Data: struct {
			ID    string   `json:"id"`
			Email string   `json:"email"`
			Roles []string `json:"roles"`
			Token string   `json:"token"`
		}{
			ID:    user.ID,
			Email: user.Email,
			Roles: user.Roles,
			Token: token,
		},
		Links: struct {
			Logout  Link `json:"logout"`
			Profile Link `json:"profile"`
		}{
			Logout: Link{
				Rel:    "logout",
				Href:   "/auth/logout",
				Method: "POST",
			},
			Profile: Link{
				Rel:    "profile",
				Href:   "/users/" + user.ID,
				Method: "GET",
			},
		},
	}

	return response
}
