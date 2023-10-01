package userService

import (
	"github.com/MOHAMMADmiZAN/go_recipe/internal/app/model"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/appResponse"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// CreateUser Create User

func CreateUser(name string, email string, roles []string, password string) error {
	user := model.UserModel(name, email, roles, password)
	err := mgm.Coll(user).Create(user)
	return err
}

// DuplicateUser Duplicate User Find
func DuplicateUser(email string) bool {
	if ExistsUser(email) {
		return true
	}
	return false
}

// ExistsUserPassword Exits User Password
func ExistsUserPassword(w http.ResponseWriter, email string) string {
	user := &model.User{}
	err := mgm.Coll(user).First(bson.M{"email": email}, user)
	if err != nil {
		appResponse.ResponseMessage(w, http.StatusNotFound, "User Not Exists")
		return ""
	}

	return user.Password
}

// UserId  find by email
func GetUserId(w http.ResponseWriter, email string) string {
	user := &model.User{}
	err := mgm.Coll(user).First(bson.M{"email": email}, user)
	if err != nil {
		appResponse.ResponseMessage(w, http.StatusNotFound, "User Not Exists")
	}
	return user.ID.Hex()
}

// GetUserRoles find by email
func GetUserRoles(w http.ResponseWriter, email string) []string {
	user := &model.User{}
	err := mgm.Coll(user).First(bson.M{"email": email}, user)
	if err != nil {
		appResponse.ResponseMessage(w, http.StatusNotFound, "User Not Exists")
	}
	return user.Roles
}

// GetUser find by email
func GetUser(w http.ResponseWriter, email string) *model.User {
	user := &model.User{}
	err := mgm.Coll(user).First(bson.M{"email": email}, user)
	if err != nil {
		appResponse.ResponseMessage(w, http.StatusNotFound, "User Not Exists")
	}
	return user
}

// ExistsUser Exists User
func ExistsUser(email string) bool {
	user := &model.User{}
	err := mgm.Coll(user).First(bson.M{"email": email}, user)
	if err != nil {
		return false
	}
	return true
}

// HexToObjectId hex to ObjectId
func HexToObjectId(hex string) primitive.ObjectID {
	var w http.ResponseWriter
	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		appResponse.ResponseMessage(w, http.StatusBadRequest, "ObjectId Create Failed")
	}
	return id
}
