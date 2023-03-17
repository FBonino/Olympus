package routes

import (
	"server/controllers"
	"server/middlewares"
	"server/services"

	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup, userService services.UserService) {
	router := rg.Group("/auth")

	router.POST("/login", rc.authController.Login)
	router.POST("/logout", rc.authController.Logout)
	router.POST("/signup", rc.authController.Signup)
	router.GET("/auto-login", middlewares.DeserializeUser(userService), rc.authController.AutoLogin)
}
