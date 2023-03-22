package controllers

import (
	"net/http"
	"server/dtos"
	"server/models"
	"server/services"

	"github.com/gin-gonic/gin"
)

type ServerController struct {
	serverService services.ServerService
}

func NewServerController(serverService services.ServerService) ServerController {
	return ServerController{serverService}
}

func (sc *ServerController) CreateServer(ctx *gin.Context) {
	var input *models.CreateServerInput

	user := ctx.MustGet("user").(*models.User)

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	server, err := sc.serverService.CreateServer(user.ID, input)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "server": dtos.MapServerDTO(server)})
}

func (sc *ServerController) GetServer(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)

	serverId := ctx.Param("id")

	server, err := sc.serverService.FindServerByID(serverId, user.ID)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "server": server})
}
