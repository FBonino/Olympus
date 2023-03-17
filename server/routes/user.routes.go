package routes

import (
	"server/controllers"
	"server/middlewares"
	"server/services"

	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewUserRouteController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (rc *UserRouteController) UserRoute(rg *gin.RouterGroup, userService services.UserService) {
	router := rg.Group("/user")

	router.PUT("/avatar", middlewares.DeserializeUser(userService), rc.userController.UpdateAvatar)
}
