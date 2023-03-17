package controllers

import (
	"net/http"
	"server/configs"
	"server/helpers"
	"server/models"
	"server/services"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthController struct {
	authService    services.AuthService
	userService    services.UserService
	sessionService services.SessionService
}

func NewAuthController(authService services.AuthService, userService services.UserService, sessionService services.SessionService) AuthController {
	return AuthController{authService, userService, sessionService}
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

	user, err := ac.userService.FindByIdentifier(input.Identifier)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or password"})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if err = helpers.VerifyPassword(user.Password, input.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}

	session, err := ac.sessionService.Create(user.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Error creating session"})
		return
	}

	config, _ := configs.LoadConfig(".")

	ctx.SetCookie("SID", session.Token, config.TokenMaxAge, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "user": models.UserFilteredResponse(user)})
}

func (ac *AuthController) Logout(ctx *gin.Context) {
	token, _ := ctx.Cookie("SID")

	ac.sessionService.Delete(token)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (ac *AuthController) AutoLogin(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)

	token, _ := ctx.Cookie("SID")

	session, err := ac.sessionService.Update(user.ID, token)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	config, _ := configs.LoadConfig(".")

	ctx.SetCookie("SID", session.Token, config.TokenMaxAge, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "user": models.UserFilteredResponse(user)})
}
