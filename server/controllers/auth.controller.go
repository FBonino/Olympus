package controllers

import (
	"net/http"
	"server/models"
	"server/services"
	"server/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthController struct {
	authService services.AuthService
	userService services.UserService
}

func NewAuthController(authService services.AuthService, userService services.UserService) AuthController {
	return AuthController{authService, userService}
}

func (ac *AuthController) Signup(ctx *gin.Context) {
	var input *models.SignupInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if input.Password != input.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
		return
	}

	_, err := ac.authService.Signup(input)

	if err != nil {
		if strings.Contains(err.Error(), "already exist") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success"})
}

func (ac *AuthController) Login(ctx *gin.Context) {
	var input *models.LoginInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	user, err := ac.userService.FindUserByIdentifier(input.Identifier)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or password"})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if err = utils.VerifyPassword(user.Password, input.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "user": models.UserFilteredResponse(user)})
}
