package services

import "server/models"

type ServerService interface {
	Create(string, *models.CreateServerInput) (*models.Server, error)
	GetUserServers(string) ([]*models.Server, error)
	FindByID(string, string) (*models.Server, []*models.User, error)
}
