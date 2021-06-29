package repository

import (
	"github.com/rpsingh21/checklist-api/model"
	"go.uber.org/zap"
)

// UserRepository object
type UserRepository struct {
	logger *zap.Logger
}

// Get method for return all object
func (ur *UserRepository) Get() []model.User {
	return nil
}

// Create method to create new entry
func (ur *UserRepository) Create(u *model.User) (model.User, error) {
	return *u, nil
}

// Update method to update entry
func (ur *UserRepository) Update(id string, u *model.User) (model.User, error) {
	return *u, nil
}

// Delete methd to update entry
func (ur *UserRepository) Delete(id string) (bool, error) {
	return false, nil
}
