package model

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User models object
type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username,omitempty"`
	Password string             `json:"-" bson:"password,omitempty"`
	Email    string             `json:"email" bson:"email,omitempty"`
	Profile  Profile            `json:"profile" bson:"-"`
}

// FromJSON load data from json
func (u *User) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(u)
}
