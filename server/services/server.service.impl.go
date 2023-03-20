package services

import (
	"context"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"nullprogram.com/x/uuid"
)

type ServerServiceImpl struct {
	db  *mongo.Database
	ctx context.Context
}

func NewServerService(db *mongo.Database, ctx context.Context) ServerService {
	return &ServerServiceImpl{db, ctx}
}

func (ss *ServerServiceImpl) CreateServer(userId string, input *models.CreateServerInput) (*models.Server, error) {
	now := time.Now()

	var avatar string = input.Avatar

	if avatar == "" {
		avatar = "default-server-avatar.png"
	}

	var defaultRole models.ServerRole = models.ServerRole{
		ID:    uuid.NewGen().NewV4().String(),
		Name:  "Owner",
		Color: "#D30000",
	}

	var owner models.ServerUser = models.ServerUser{
		ID:    userId,
		Roles: []models.ServerRole{defaultRole},
	}

	var defaultTextGeneral models.Channel = models.Channel{
		ID:   uuid.NewGen().NewV4().String(),
		Name: "general",
		Type: "text",
	}

	var defaultVoiceGeneral models.Channel = models.Channel{
		ID:   uuid.NewGen().NewV4().String(),
		Name: "general",
		Type: "voice",
	}

	var server models.Server = models.Server{
		ID:        uuid.NewGen().NewV4().String(),
		Name:      input.Name,
		Avatar:    avatar,
		Roles:     []models.ServerRole{defaultRole},
		Users:     []models.ServerUser{owner},
		Channels:  []models.Channel{defaultTextGeneral, defaultVoiceGeneral},
		CreatedAt: now,
		UpdatedAt: now,
	}

	serversCollection := ss.db.Collection("servers")

	_, err := serversCollection.InsertOne(ss.ctx, &server)

	if err != nil {
		return nil, err
	}

	return &server, nil
}
