package model

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User models object
type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username,omitempty" validate:"required"`
	Password string             `json:"password" bson:"password,omitempty" validate:"required,gte=8"`
	Email    string             `json:"email" bson:"email,omitempty" validate:"required"`
	Profile  Profile            `json:"profile" bson:"-"`
}

// FromJSON load data from json
func (u *User) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(u)
}

// Validate json input body
func (u *User) Validate() error {
	validator := validator.New()
	return validator.Struct(u)
}

// SetPassword for update password value
func (u *User) SetPassword(hs string) {
	u.Password = hs
}
