package services

import (
	"context"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/exp/slices"
	"nullprogram.com/x/uuid"
)

type ConversationServiceImpl struct {
	db  *mongo.Database
	ctx context.Context
}

func NewConversationService(db *mongo.Database, ctx context.Context) ConversationService {
	return &ConversationServiceImpl{db, ctx}
}

func (cs *ConversationServiceImpl) FindUserConversations(id string) ([]*models.Conversation, []string, error) {
	var usersIDs []string = []string{id}
	var conversations []*models.Conversation

	query := bson.M{"users": id}

	res, err := cs.db.Collection("conversations").Find(cs.ctx, query)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []*models.Conversation{}, []string{}, err
		}
		return nil, nil, err
	}

	for res.Next(cs.ctx) {
		var conversation *models.Conversation

		err = res.Decode(&conversation)

		if err == nil {
			conversations = append(conversations, conversation)
		}

		for _, userID := range conversation.Users {
			if userID != id && !slices.Contains(usersIDs, userID) {
				usersIDs = append(usersIDs, userID)
			}
		}
	}

	return conversations, usersIDs, nil
}

func (cs *ConversationServiceImpl) FindOrCreate(input *models.CreateConversationInput) (*models.Conversation, error) {
	var conversation models.Conversation

	query := bson.M{
		"users": bson.M{
			"$size": len(input.Users),
			"$all":  input.Users,
		},
	}

	err := cs.db.Collection("conversations").FindOne(cs.ctx, query).Decode(&conversation)

	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
	} else {
		return &conversation, nil
	}

	now := time.Now()

	conversation = models.Conversation{
		ID:        uuid.NewGen().NewV4().String(),
		Avatar:    "default-group-avatar.png",
		Users:     input.Users,
		Owner:     input.Owner,
		Messages:  []string{},
		CreatedAt: now,
		UpdatedAt: now,
	}

	_, err = cs.db.Collection("conversations").InsertOne(cs.ctx, &conversation)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.Conversation{}, err
		}
		return nil, err
	}

	return &conversation, nil
}

func (cs *ConversationServiceImpl) FindByID(conversationID string) (*models.Conversation, error) {
	var conversation *models.Conversation

	query := bson.M{"_id": conversationID}

	err := cs.db.Collection("conversations").FindOne(cs.ctx, query).Decode(&conversation)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.Conversation{}, err
		}
		return nil, err
	}

	return conversation, nil
}

func (cs *ConversationServiceImpl) AddMessage(conversationID string, messageID string) error {
	update := bson.M{
		"$push": bson.M{
			"messages": messageID,
		},
	}

	_, err := cs.db.Collection("conversations").UpdateByID(cs.ctx, conversationID, update)

	return err
}
