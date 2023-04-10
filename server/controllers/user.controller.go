package controllers

import (
	"net/http"
	"server/dtos"
	"server/models"
	"server/services"

	"github.com/gin-gonic/gin"
	"nullprogram.com/x/uuid"
)

type UserController struct {
	userService   services.UserService
	serverService services.ServerService
}

func NewUserController(userService services.UserService, serverService services.ServerService) UserController {
	return UserController{userService, serverService}
}

func (uc *UserController) GetMyUser(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)

	servers, err := uc.serverService.GetUserServers(user.ID)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	friends, err := uc.userService.GetUserFriends(user.Friends)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "user": dtos.MapMyUserDTO(user, friends), "servers": dtos.MapServersBasicDTO(servers)})
}

func (uc *UserController) UploadAvatar(ctx *gin.Context) {
	file, err := ctx.FormFile("file")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	filename := uuid.NewGen().NewV4().String() + file.Filename

	err = ctx.SaveUploadedFile(file, "uploads/"+filename)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "avatar": filename})
}
