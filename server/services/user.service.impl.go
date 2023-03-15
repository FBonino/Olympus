package services

import (
	"context"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{userCollection, ctx}
}

func (us *UserServiceImpl) FindUserByID(id string) (*models.User, error) {
	var user *models.User

	query := bson.M{"_id": id}

	err := us.userCollection.FindOne(us.ctx, query).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.User{}, err
		}
		return nil, err
	}

	return user, nil
}

func (us *UserServiceImpl) FindUserByIdentifier(identifier string) (*models.User, error) {
	var user *models.User

	query := bson.M{
		"$or": []interface{}{
			bson.M{"email": identifier},
			bson.M{"username": identifier},
		},
	}

	err := us.userCollection.FindOne(us.ctx, query).Decode(&user)

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

	_, err := us.userCollection.UpdateByID(us.ctx, id, update)

	return err
}
