package services

import "server/models"

type MessageService interface {
	Create(string, *models.CreateMessageInput) (*models.Message, error)
}
