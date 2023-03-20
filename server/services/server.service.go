package services

import "server/models"

type ServerService interface {
	CreateServer(string, *models.CreateServerInput) (*models.Server, error)
}
