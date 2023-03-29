package services

import (
	"context"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
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
