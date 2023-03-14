package services

import (
	"context"
	"server/models"
	"server/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"nullprogram.com/x/uuid"
)

type AuthServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewAuthService(userCollection *mongo.Collection, ctx context.Context) AuthService {
	return &AuthServiceImpl{userCollection, ctx}
}

func (as *AuthServiceImpl) Signup(signupInput *models.SignupInput) (*models.User, error) {
	var newUser *models.User

	now := time.Now()
	hashedPassword, _ := utils.HashPassword(signupInput.Password)

	var user models.User = models.User{
		ID:           uuid.NewGen().NewV4().String(),
		Username:     signupInput.Username,
		Email:        signupInput.Email,
		Avatar:       "",
		Password:     hashedPassword,
		Status:       0,
		CustomStatus: "",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	res, err := as.userCollection.InsertOne(as.ctx, &user)

	if err != nil {
		return nil, err
	}

	_, err = as.userCollection.Indexes().CreateMany(as.ctx, []mongo.IndexModel{
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

	if err = as.userCollection.FindOne(as.ctx, bson.M{"_id": res.InsertedID}).Decode(&newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (as *AuthServiceImpl) Login(loginInput *models.LoginInput) (*models.User, error) {
	return nil, nil
}
