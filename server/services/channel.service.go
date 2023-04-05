package services

import "server/models"

type ChannelService interface {
	Create(models.CreateChannelInput) (*models.Channel, error)
	FindByID(string) (*models.Channel, error)
	FindManyByID([]string) ([]*models.Channel, error)
	AddMessage(string, string) error
}
