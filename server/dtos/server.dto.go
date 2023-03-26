package dtos

import (
	"server/models"
)

type ServerUserDTO struct {
	ID           string              `json:"id" bson:"_id"`
	Username     string              `json:"username" bson:"username"`
	Email        string              `json:"email" bson:"email"`
	Avatar       string              `json:"avatar" bson:"avatar"`
	Status       string              `json:"status" bson:"status"`
	CustomStatus string              `json:"customStatus" bson:"customStatus"`
	Roles        []models.ServerRole `json:"roles" bson:"roles"`
}

type ChannelDTO struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Type string `json:"type" bson:"type"`
}

type ServerDTO struct {
	ID             string       `json:"id" bson:"_id"`
	Name           string       `json:"name" bson:"name"`
	Avatar         string       `json:"avatar" bson:"avatar"`
	Channels       []ChannelDTO `json:"channels" bson:"channels"`
	DefaultChannel string       `json:"defaultChannel" bson:"defaultChannel"`
}

type ServerExtendedDTO struct {
	ID             string          `json:"id" bson:"_id"`
	Name           string          `json:"name" bson:"name"`
	Avatar         string          `json:"avatar" bson:"avatar"`
	Channels       []ChannelDTO    `json:"channels" bson:"channels"`
	DefaultChannel string          `json:"defaultChannel" bson:"defaultChannel"`
	Users          []ServerUserDTO `json:"users" bson:"users"`
}

func MapChannelDTO(channel models.Channel) ChannelDTO {
	return ChannelDTO{
		ID:   channel.ID,
		Name: channel.Name,
		Type: channel.Type,
	}
}

func MapChannelsDTO(channels []models.Channel) []ChannelDTO {
	var channelsDTO []ChannelDTO

	for _, channel := range channels {
		channelDTO := MapChannelDTO(channel)
		channelsDTO = append(channelsDTO, channelDTO)
	}

	return channelsDTO
}

func MapServerDTO(server *models.Server) ServerDTO {
	return ServerDTO{
		ID:             server.ID,
		Name:           server.Name,
		Avatar:         server.Avatar,
		Channels:       MapChannelsDTO(server.Channels),
		DefaultChannel: server.DefaultChannel,
	}
}

func MapServersDTO(servers []*models.Server) []ServerDTO {
	var serversDTO []ServerDTO

	for _, server := range servers {
		serverDTO := MapServerDTO(server)
		serversDTO = append(serversDTO, serverDTO)
	}

	return serversDTO
}

func MapServerExtendedDTO(server *models.Server, users []*models.User) ServerExtendedDTO {
	var serverUsersDTO []ServerUserDTO

	usersDTO := MapUsersDTO(users)

	for i, user := range server.Users {
		var serverUser ServerUserDTO = ServerUserDTO{
			ID:           usersDTO[i].ID,
			Username:     usersDTO[i].Username,
			Email:        usersDTO[i].Email,
			Avatar:       usersDTO[i].Avatar,
			Status:       usersDTO[i].Status,
			CustomStatus: usersDTO[i].CustomStatus,
			Roles:        user.Roles,
		}

		serverUsersDTO = append(serverUsersDTO, serverUser)
	}

	return ServerExtendedDTO{
		ID:             server.ID,
		Name:           server.Name,
		Avatar:         server.Avatar,
		Channels:       MapChannelsDTO(server.Channels),
		DefaultChannel: server.DefaultChannel,
		Users:          serverUsersDTO,
	}
}
