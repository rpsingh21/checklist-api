package repository

import (
	"github.com/rpsingh21/checklist-api/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// UserRepository object
type UserRepository struct {
	logger *zap.SugaredLogger
	db     *mongo.Database
}

// NewUserRepository new object
func NewUserRepository(logger *zap.SugaredLogger, db *mongo.Database) *UserRepository {
	return &UserRepository{logger: logger, db: db}
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
