package controllers

import (
	"net/http"
	"server/dtos"
	"server/models"
	"server/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ConversationController struct {
	conversationService services.ConversationService
	messageService      services.MessageService
	userService         services.UserService
}

func NewConversationController(conversationService services.ConversationService, messageService services.MessageService, userService services.UserService) ConversationController {
	return ConversationController{conversationService, messageService, userService}
}

func (cc *ConversationController) GetUserConversations(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)

	conversations, usersIDs, err := cc.conversationService.FindUserConversations(user.ID)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	users, _ := cc.userService.FindManyByID(usersIDs)

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "conversations": dtos.MapConversationsBasicDTO(conversations, user.ID, users)})
}

func (cc *ConversationController) FindOrCreate(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)

	var input *models.CreateConversationInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	input.Users = append(input.Users, user.ID)

	input.Owner = user.ID

	conversation, err := cc.conversationService.FindOrCreate(input)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	users, err := cc.userService.FindManyByID(conversation.Users)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "conversation": dtos.MapConversationBasicDTO(conversation, user.ID, users)})
}

func (cc *ConversationController) GetConversation(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)

	conversationID := ctx.Param("conversation")

	var queryLimit int

	limit := ctx.Query("limit")

	queryLimit, err := strconv.Atoi(limit)

	if err != nil {
		queryLimit = 50
	}

	conversation, err := cc.conversationService.FindByID(conversationID)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	users, err := cc.userService.FindManyByID(conversation.Users)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	messages, err := cc.messageService.FindMessages(conversation.Messages, int64(queryLimit))

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "conversation": dtos.MapConversationDTO(conversation, user.ID, users, messages)})
}

func (cc *ConversationController) NewMessage(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)

	conversationID := ctx.Param("conversation")

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

	err = cc.conversationService.AddMessage(conversationID, message.ID)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "message": dtos.MapMessageDTO(message)})
}
