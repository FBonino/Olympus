package controllers

import (
	"net/http"
	"server/models"
	"server/services"

	"github.com/gin-gonic/gin"
	"nullprogram.com/x/uuid"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{userService}
}

func (uc *UserController) UpdateAvatar(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)

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

	err = uc.userService.UpdateAvatar(user.ID, filename)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "avatar": filename})
}
