package dtos

import (
	"server/models"
)

type ChannelBasicDTO struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Type string `json:"type" bson:"type"`
}

type ChannelDTO struct {
	ID       string           `json:"id" bson:"_id"`
	Name     string           `json:"name" bson:"name"`
	Type     string           `json:"type" bson:"type"`
	Topic    string           `json:"topic" bson:"topic"`
	Messages []models.Message `json:"messages" bson:"messages"`
}

func MapChannelBasicDTO(channel models.Channel) ChannelBasicDTO {
	return ChannelBasicDTO{
		ID:   channel.ID,
		Name: channel.Name,
		Type: channel.Type,
	}
}

func MapChannelsBasicDTO(channels []models.Channel) []ChannelBasicDTO {
	var channelsDTO []ChannelBasicDTO

	for _, channel := range channels {
		channelDTO := MapChannelBasicDTO(channel)
		channelsDTO = append(channelsDTO, channelDTO)
	}

	return channelsDTO
}

func MapChannelDTO(channel models.Channel, messages []models.Message) ChannelDTO {
	return ChannelDTO{
		ID:       channel.ID,
		Name:     channel.Name,
		Type:     channel.Type,
		Topic:    channel.Topic,
		Messages: messages,
	}
}
