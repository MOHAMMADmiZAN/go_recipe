package model

import "github.com/kamva/mgm/v3"

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string   `json:"name" bson:"name"`
	Email            string   `json:"email" bson:"email"`
	Roles            []string `json:"roles" bson:"roles"`
	Password         string   `json:"password" bson:"password"`
}

func UserModel(name string, email string, roles []string, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Roles:    roles,
		Password: password,
	}
}

func (u *User) CollectionName() string {
	return "users"
}
