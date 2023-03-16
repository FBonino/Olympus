package middlewares

import (
	"fmt"
	"net/http"
	"server/services"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func DeserializeSession(userService services.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		userID := fmt.Sprintf("%v", session.Get("id"))

		user, err := userService.FindUserByID(userID)

		if user == nil || err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "This session is no longer avaialable"})
			return
		}

		ctx.Set("user", user)

		ctx.Next()
	}
}
