package services

import (
	"context"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"nullprogram.com/x/uuid"
)

type MessageServiceImpl struct {
	db  *mongo.Database
	ctx context.Context
}

func NewMessageService(db *mongo.Database, ctx context.Context) MessageService {
	return &MessageServiceImpl{db, ctx}
}

func (ms *MessageServiceImpl) Create(userID string, input *models.CreateMessageInput) (*models.Message, error) {
	now := time.Now()

	message := models.Message{
		ID:        uuid.NewGen().NewV4().String(),
		Author:    userID,
		Content:   input.Content,
		IsDeleted: false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	_, err := ms.db.Collection("messages").InsertOne(ms.ctx, &message)

	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (cs *MessageServiceImpl) FindMessages(messageIDs []string, limit int64) ([]*models.Message, error) {
	var messages []*models.Message

	query := bson.M{
		"_id": bson.M{
			"$in": messageIDs,
		},
		"isDeleted": false,
	}

	options := options.Find().SetSort(bson.M{"createdAt": -1}).SetLimit(limit)

	res, err := cs.db.Collection("messages").Find(cs.ctx, query, options)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []*models.Message{}, err
		}
		return nil, err
	}

	for res.Next(cs.ctx) {
		var message models.Message

		err := res.Decode(&message)

		if err == nil {
			messages = append(messages, &message)
		}
	}

	return messages, nil
}
