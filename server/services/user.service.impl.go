package services

import (
	"context"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	db  *mongo.Database
	ctx context.Context
}

func NewUserService(db *mongo.Database, ctx context.Context) UserService {
	return &UserServiceImpl{db, ctx}
}

func (us *UserServiceImpl) FindByID(id string) (*models.User, error) {
	var user *models.User

	query := bson.M{"_id": id}

	err := us.db.Collection("users").FindOne(us.ctx, query).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.User{}, err
		}
		return nil, err
	}

	return user, nil
}

func (us *UserServiceImpl) FindByIdentifier(identifier string) (*models.User, error) {
	var user *models.User

	query := bson.M{
		"$or": []interface{}{
			bson.M{"email": identifier},
			bson.M{"username": identifier},
		},
	}

	err := us.db.Collection("users").FindOne(us.ctx, query).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.User{}, err
		}
		return nil, err
	}

	return user, nil
}

func (us *UserServiceImpl) UpdateAvatar(id string, avatar string) error {
	now := time.Now()

	update := bson.M{
		"$set": bson.M{
			"avatar":    avatar,
			"updatedAt": now,
		},
	}

	_, err := us.db.Collection("users").UpdateByID(us.ctx, id, update)

	return err
}
