package dtos

import "server/models"

type ServerDTO struct {
	ID     string `json:"id" bson:"_id"`
	Name   string `json:"name" bson:"name"`
	Avatar string `json:"avatar" bson:"avatar"`
}

func MapServerDTO(server *models.Server) ServerDTO {
	return ServerDTO{
		ID:     server.ID,
		Name:   server.Name,
		Avatar: server.Avatar,
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
