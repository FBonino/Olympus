package middlewares

import (
	"fmt"
	"net/http"
	"server/configs"
	"server/helpers"
	"server/services"

	"github.com/gin-gonic/gin"
)

func DeserializeUser(userService services.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("SID")

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		config, _ := configs.LoadConfig(".")

		sub, err := helpers.ValidateToken(token, config.TokenPublicKey)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		user, err := userService.FindByID(fmt.Sprint(sub))

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "The user belonging to this token no logger exists"})
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
