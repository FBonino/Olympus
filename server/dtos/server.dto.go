package dtos

import (
	"server/models"
)

type ServerUserDTO struct {
	ID           string   `json:"id" bson:"_id"`
	Username     string   `json:"username" bson:"username"`
	Email        string   `json:"email" bson:"email"`
	Avatar       string   `json:"avatar" bson:"avatar"`
	Status       string   `json:"status" bson:"status"`
	CustomStatus string   `json:"customStatus" bson:"customStatus"`
	Roles        []string `json:"roles" bson:"roles"`
}

type ServerBasicDTO struct {
	ID             string `json:"id" bson:"_id"`
	Name           string `json:"name" bson:"name"`
	Avatar         string `json:"avatar" bson:"avatar"`
	DefaultChannel string `json:"defaultChannel" bson:"defaultChannel"`
}

type ServerDTO struct {
	ID             string              `json:"id" bson:"_id"`
	Name           string              `json:"name" bson:"name"`
	Avatar         string              `json:"avatar" bson:"avatar"`
	Channels       []ChannelBasicDTO   `json:"channels" bson:"channels"`
	DefaultChannel string              `json:"defaultChannel" bson:"defaultChannel"`
	Users          []ServerUserDTO     `json:"users" bson:"users"`
	Roles          []models.ServerRole `json:"roles" bson:"roles"`
}

func MapServerBasicDTO(server *models.Server) ServerBasicDTO {
	return ServerBasicDTO{
		ID:             server.ID,
		Name:           server.Name,
		Avatar:         server.Avatar,
		DefaultChannel: server.DefaultChannel,
	}
}

func MapServersBasicDTO(servers []*models.Server) []ServerBasicDTO {
	var serversDTO []ServerBasicDTO

	for _, server := range servers {
		serverDTO := MapServerBasicDTO(server)
		serversDTO = append(serversDTO, serverDTO)
	}

	return serversDTO
}

func MapServerDTO(server *models.Server, users []*models.User, channels []*models.Channel) ServerDTO {
	serverUsersDTO := []ServerUserDTO{}

	for _, svUser := range server.Users {
		var userDTO UserDTO

		for _, user := range users {
			if svUser.ID == user.ID {
				userDTO = MapUserDTO(user)
				break
			}
		}

		var serverUser ServerUserDTO = ServerUserDTO{
			ID:           userDTO.ID,
			Username:     userDTO.Username,
			Email:        userDTO.Email,
			Avatar:       userDTO.Avatar,
			Status:       userDTO.Status,
			CustomStatus: userDTO.CustomStatus,
			Roles:        svUser.Roles,
		}

		serverUsersDTO = append(serverUsersDTO, serverUser)
	}

	return ServerDTO{
		ID:             server.ID,
		Name:           server.Name,
		Avatar:         server.Avatar,
		Channels:       MapChannelsBasicDTO(channels),
		DefaultChannel: server.DefaultChannel,
		Users:          serverUsersDTO,
		Roles:          server.Roles,
	}
}
