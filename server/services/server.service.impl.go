package services

import (
	"context"
	"server/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

func (ss *ServerServiceImpl) GetUserServers(userId string) ([]*models.Server, error) {
	var servers []*models.Server

	query := bson.M{"users._id": userId}

	res, err := ss.db.Collection("servers").Find(ss.ctx, query)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []*models.Server{}, err
		}
		return nil, err
	}

	for res.Next(ss.ctx) {
		var server *models.Server

		err := res.Decode(&server)

		if err == nil {
			servers = append(servers, server)
		}
	}

	return servers, nil
}

func (ss *ServerServiceImpl) FindServerByID(serverId string, userId string) (*models.Server, error) {
	var server *models.Server

	query := bson.M{"users._id": userId, "_id": serverId}

	err := ss.db.Collection("servers").FindOne(ss.ctx, query).Decode(&server)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.Server{}, err
		}
		return nil, err
	}

	return server, nil
}
