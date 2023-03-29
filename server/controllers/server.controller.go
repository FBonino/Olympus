package controllers

import (
	"net/http"
	"server/dtos"
	"server/models"
	"server/services"

	"github.com/gin-gonic/gin"
)

type ServerController struct {
	serverService  services.ServerService
	channelService services.ChannelService
}

func NewServerController(serverService services.ServerService, channelService services.ChannelService) ServerController {
	return ServerController{serverService, channelService}
}

func (sc *ServerController) Create(ctx *gin.Context) {
	var input *models.CreateServerInput

	user := ctx.MustGet("user").(*models.User)

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	defaultTextInput := models.CreateChannelInput{Name: "general", Type: "text"}

	defaultText, _ := sc.channelService.Create(defaultTextInput)

	defaultVoiceInput := models.CreateChannelInput{Name: "general", Type: "voice"}

	defaultVoice, _ := sc.channelService.Create(defaultVoiceInput)

	input.Channels = []string{defaultText.ID, defaultVoice.ID}

	server, err := sc.serverService.Create(user.ID, input)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "server": dtos.MapServerBasicDTO(server)})
}

func (sc *ServerController) GetServer(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)

	serverID := ctx.Param("id")

	server, users, err := sc.serverService.FindByID(serverID, user.ID)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	channels, err := sc.channelService.FindManyByID(server.Channels)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "server": dtos.MapServerDTO(server, users, channels)})
}
