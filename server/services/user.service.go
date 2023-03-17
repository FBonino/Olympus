package services

import (
	"server/models"
)

type UserService interface {
	FindByID(string) (*models.User, error)
	FindByIdentifier(string) (*models.User, error)
	UpdateAvatar(string, string) error
}
