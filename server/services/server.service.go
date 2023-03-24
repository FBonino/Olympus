package services

import "server/models"

type ServerService interface {
	CreateServer(string, *models.CreateServerInput) (*models.Server, error)
	GetUserServers(string) ([]*models.Server, error)
	FindServerByID(string, string) (*models.Server, []*models.User, error)
	FindChannelByID(string, string, string) (*models.Channel, error)
}
