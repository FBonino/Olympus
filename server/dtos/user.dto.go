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

type FriendDTO struct {
	User     UserDTO `json:"user" bson:"user"`
	Relation string  `json:"relation" bson:"relation"`
}

type MyUserDTO struct {
	ID           string      `json:"id" bson:"_id"`
	Username     string      `json:"username" bson:"username"`
	Email        string      `json:"email" bson:"email"`
	Avatar       string      `json:"avatar" bson:"avatar"`
	Status       string      `json:"status" bson:"status"`
	CustomStatus string      `json:"customStatus" bson:"customStatus"`
	Friends      []FriendDTO `json:"friends" bson:"friends"`
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
	usersDTO := []UserDTO{}

	for _, user := range users {
		userDTO := MapUserDTO(user)
		usersDTO = append(usersDTO, userDTO)
	}

	return usersDTO
}

func MapMyUserDTO(user *models.User, friends []*models.User) MyUserDTO {
	friendsDTO := []FriendDTO{}

	for _, friend := range friends {
		for _, userFriend := range user.Friends {
			if friend.ID == userFriend.ID {
				friendDTO := FriendDTO{
					User:     MapUserDTO(friend),
					Relation: helpers.TransformRelation(userFriend.Relation),
				}
				friendsDTO = append(friendsDTO, friendDTO)
			}
		}
	}

	return MyUserDTO{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		Avatar:       user.Avatar,
		Status:       helpers.TransformStatus(user.Status),
		CustomStatus: user.CustomStatus,
		Friends:      friendsDTO,
	}
}
