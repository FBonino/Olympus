package services

import "server/models"

type ServerService interface {
	CreateServer(string, *models.CreateServerInput) (*models.Server, error)
	GetUserServers(string) ([]*models.Server, error)
}
