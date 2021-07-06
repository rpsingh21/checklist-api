package repository

import (
	"context"

	"github.com/rpsingh21/checklist-api/config"
	"github.com/rpsingh21/checklist-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository object
type UserRepository struct {
	db         *mongo.Database
	Collection *mongo.Collection
}

// NewUserRepository new object
func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		db:         db,
		Collection: db.Collection(config.UserCollection),
	}
}

// Get method for return all object
func (ur *UserRepository) Get() []model.User {
	return nil
}

// Create method to create new entry
func (ur *UserRepository) Create(u *model.User) error {
	result, err := ur.Collection.InsertOne(context.TODO(), u)
	if err != nil {
		return err
	}
	u.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

// Update method to update entry
func (ur *UserRepository) Update(id string, u *model.User) (model.User, error) {
	return *u, nil
}

// Delete methd to update entry
func (ur *UserRepository) Delete(id string) (bool, error) {
	return false, nil
}

// FindByUserName find user object ny username
func (ur *UserRepository) FindByUserName(username string) (*model.User, error) {
	user := &model.User{}
	if err := ur.Collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}
