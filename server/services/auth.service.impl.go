package services

import (
	"context"
	"server/helpers"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"nullprogram.com/x/uuid"
)

type AuthServiceImpl struct {
	db  *mongo.Database
	ctx context.Context
}

func NewAuthService(db *mongo.Database, ctx context.Context) AuthService {
	return &AuthServiceImpl{db, ctx}
}

func (as *AuthServiceImpl) Signup(signupInput *models.SignupInput) (*models.User, error) {
	now := time.Now()

	hashedPassword, _ := helpers.HashPassword(signupInput.Password)

	var user models.User = models.User{
		ID:           uuid.NewGen().NewV4().String(),
		Username:     signupInput.Username,
		Email:        signupInput.Email,
		Avatar:       "default-avatar.png",
		Password:     hashedPassword,
		Status:       0,
		CustomStatus: "",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	usersCollection := as.db.Collection("users")

	_, err := usersCollection.InsertOne(as.ctx, &user)

	if err != nil {
		return nil, err
	}

	_, err = usersCollection.Indexes().CreateMany(as.ctx, []mongo.IndexModel{
		{
			Keys: bson.M{"username": 1}, Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.M{"email": 1}, Options: options.Index().SetUnique(true),
		},
	})

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (as *AuthServiceImpl) Login(loginInput *models.LoginInput) (*models.User, error) {
	return nil, nil
}
