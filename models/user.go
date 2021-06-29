package models

// User models object
type User struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"-"`
	Email    string  `json:"email"`
	Profile  Profile `json:"profile"`
}
