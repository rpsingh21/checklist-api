package models

// Profile model object
type Profile struct {
	ID             string `json:"id"`
	UserID         string `json:"-"`
	ProfilePicture string `json:"profilePicture"`
}
