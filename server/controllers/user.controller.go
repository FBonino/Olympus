package controllers

import (
	"net/http"
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
	userId := ctx.Param("id")

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

	err = uc.userService.UpdateAvatar(userId, filename)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "avatar": filename})
}
