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

type ChannelServiceImpl struct {
	db  *mongo.Database
	ctx context.Context
}

func NewChannelService(db *mongo.Database, ctx context.Context) ChannelService {
	return &ChannelServiceImpl{db, ctx}
}

func (cs *ChannelServiceImpl) Create(input models.CreateChannelInput) (*models.Channel, error) {
	now := time.Now()

	channel := models.Channel{
		ID:        uuid.NewGen().NewV4().String(),
		Name:      input.Name,
		Type:      input.Type,
		Topic:     "",
		Messages:  []string{},
		CreatedAt: now,
		UpdatedAt: now,
	}

	_, err := cs.db.Collection("channels").InsertOne(cs.ctx, &channel)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.Channel{}, err
		}
		return nil, err
	}

	return &channel, nil
}

func (cs *ChannelServiceImpl) FindByID(channelID string) (*models.Channel, error) {
	var channel *models.Channel

	query := bson.M{"_id": channelID}

	err := cs.db.Collection("channels").FindOne(cs.ctx, query).Decode(&channel)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.Channel{}, err
		}
		return nil, err
	}

	return channel, nil
}

func (cs *ChannelServiceImpl) FindManyByID(channelIDs []string) ([]*models.Channel, error) {
	var channels []*models.Channel

	query := bson.M{
		"_id": bson.M{
			"$in": channelIDs,
		},
	}

	res, err := cs.db.Collection("channels").Find(cs.ctx, query)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []*models.Channel{}, err
		}
		return nil, err
	}

	for res.Next(cs.ctx) {
		var channel *models.Channel

		err := res.Decode(&channel)

		if err == nil {
			channels = append(channels, channel)
		}
	}

	return channels, nil
}

func (cs *ChannelServiceImpl) AddMessage(channelID string, messageID string) error {
	update := bson.M{
		"$push": bson.M{
			"messages": messageID,
		},
	}

	_, err := cs.db.Collection("channels").UpdateByID(cs.ctx, channelID, update)

	return err
}

func (cs *ChannelServiceImpl) FindMessages(messageIDs []string, limit int64) ([]*models.Message, error) {
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
