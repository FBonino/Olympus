package routes

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewUserRouteController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (rc *UserRouteController) UserRoute(rg *gin.RouterGroup) {
	router := rg.Group("/user")

	router.PUT("/:id/avatar", rc.userController.UpdateAvatar)
}
