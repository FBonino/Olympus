package dtos

import "server/models"

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
	ID             string       `json:"id" bson:"_id"`
	Name           string       `json:"name" bson:"name"`
	Avatar         string       `json:"avatar" bson:"avatar"`
	Channels       []ChannelDTO `json:"channels" bson:"channels"`
	DefaultChannel string       `json:"defaultChannel" bson:"defaultChannel"`
	Users          []UserDTO    `json:"users" bson:"users"`
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
	return ServerExtendedDTO{
		ID:             server.ID,
		Name:           server.Name,
		Avatar:         server.Avatar,
		Channels:       MapChannelsDTO(server.Channels),
		DefaultChannel: server.DefaultChannel,
		Users:          MapUsersDTO(users),
	}
}
