package dtos

import (
	"server/helpers"
	"server/models"
)

type UserDTO struct {
	ID           string `json:"id" bson:"_id"`
	Username     string `json:"username" bson:"username"`
	Email        string `json:"email" bson:"email"`
	Avatar       string `json:"avatar" bson:"avatar"`
	Status       string `json:"status" bson:"status"`
	CustomStatus string `json:"customStatus" bson:"customStatus"`
}

func MapUserDTO(user *models.User) UserDTO {
	return UserDTO{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		Avatar:       user.Avatar,
		Status:       helpers.TransformStatus(user.Status),
		CustomStatus: user.CustomStatus,
	}
}

func MapUsersDTO(users []*models.User) []UserDTO {
	var usersDTO []UserDTO

	for _, user := range users {
		userDTO := MapUserDTO(user)
		usersDTO = append(usersDTO, userDTO)
	}

	return usersDTO
}
