package services

import (
	"context"
	"server/configs"
	"server/helpers"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"nullprogram.com/x/uuid"
)

type SessionServiceImpl struct {
	db  *mongo.Database
	ctx context.Context
}

func NewSessionService(db *mongo.Database, ctx context.Context) SessionService {
	return &SessionServiceImpl{db, ctx}
}

func (ss *SessionServiceImpl) Create(userId string) (*models.Session, error) {
	config, _ := configs.LoadConfig(".")

	token, err := helpers.CreateToken(config.TokenExpiration, userId, config.TokenPrivateKey)

	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	var session models.Session = models.Session{
		ID:        uuid.NewGen().NewV4().String(),
		UserID:    userId,
		Token:     token,
		CreatedAt: now,
		ExpireAt:  now.Add(config.TokenExpiration),
	}

	sessionsCollection := ss.db.Collection("sessions")

	_, err = sessionsCollection.InsertOne(ss.ctx, &session)

	if err != nil {
		return nil, err
	}

	_, err = sessionsCollection.Indexes().CreateMany(ss.ctx, []mongo.IndexModel{
		{
			Keys: bson.M{"token": 1}, Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.M{"expireAt": 1}, Options: options.Index().SetExpireAfterSeconds(0),
		},
	})

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (ss *SessionServiceImpl) Update(userId string, token string) (*models.Session, error) {
	var session *models.Session

	config, _ := configs.LoadConfig(".")

	newToken, _ := helpers.CreateToken(config.TokenExpiration, userId, config.TokenPrivateKey)

	query := bson.M{"token": token}

	update := bson.M{
		"$set": bson.M{
			"expireAt": time.Now().UTC().Add(config.TokenExpiration),
			"token":    newToken,
		},
	}

	sessionCollection := ss.db.Collection("sessions")

	_, err := sessionCollection.UpdateOne(ss.ctx, query, update)

	if err != nil {
		return nil, err
	}

	query = bson.M{"token": newToken}

	sessionCollection.FindOne(ss.ctx, query).Decode(&session)

	return session, nil
}

func (ss *SessionServiceImpl) Delete(token string) error {
	query := bson.M{"token": token}

	_, err := ss.db.Collection("sessions").DeleteOne(ss.ctx, query)

	return err
}
