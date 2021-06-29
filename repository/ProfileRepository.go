package repository

import (
	"github.com/rpsingh21/checklist-api/model"
	"go.uber.org/zap"
)

// ProfileRepository object
type ProfileRepository struct {
	logger *zap.Logger
}

// Get method for return all object
func (pr *ProfileRepository) Get() []model.User {
	return nil
}

// Create method to create new entry
func (pr *ProfileRepository) Create(u *model.User) (model.User, error) {
	return *u, nil
}

// Update method to update entry
func (pr *ProfileRepository) Update(id string, u *model.User) (model.User, error) {
	return *u, nil
}

// Delete methd to update entry
func (pr *ProfileRepository) Delete(id string) (bool, error) {
	return false, nil
}
