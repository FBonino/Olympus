package controllers

import (
	"net/http"
	"server/dtos"
	"server/models"
	"server/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChannelController struct {
	channelService services.ChannelService
	messageService services.MessageService
}

func NewChannelController(channelService services.ChannelService, messageService services.MessageService) ChannelController {
	return ChannelController{channelService, messageService}
}

func (cc *ChannelController) GetChannel(ctx *gin.Context) {
	channelID := ctx.Param("channel")

	var queryLimit int

	limit := ctx.Query("limit")

	queryLimit, err := strconv.Atoi(limit)

	if err != nil {
		queryLimit = 50
	}

	channel, err := cc.channelService.FindByID(channelID)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	messages, err := cc.channelService.FindMessages(channel.Messages, int64(queryLimit))

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "channel": dtos.MapChannelDTO(channel, messages)})
}

func (cc *ChannelController) NewMessage(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)

	channelID := ctx.Param("channel")

	var input *models.CreateMessageInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	message, err := cc.messageService.Create(user.ID, input)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	err = cc.channelService.AddMessage(channelID, message.ID)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
}
