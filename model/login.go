package model

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

// Login request model
type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"Password" validate:"required"`
}

// Token responce
type Token struct {
	Role        []string `json:"role"`
	Username    string   `json:"username"`
	TokenString string   `json:"accessToken"`
}

// FromJSON load data from json
func (l *Login) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(l)
}

// Validate json input body
func (l *Login) Validate() error {
	validator := validator.New()
	return validator.Struct(l)
}
