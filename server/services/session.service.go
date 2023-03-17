package services

import (
	"server/models"
)

type SessionService interface {
	Create(string) (*models.Session, error)
	Update(string, string) (*models.Session, error)
	Delete(string) error
}
