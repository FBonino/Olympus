package routes

import (
	"server/controllers"
	"server/middlewares"
	"server/services"

	"github.com/gin-gonic/gin"
)

type ServerRouteController struct {
	serverController controllers.ServerController
}

func NewServerRouteController(serverController controllers.ServerController) ServerRouteController {
	return ServerRouteController{serverController}
}

func (rc *ServerRouteController) ServerRoute(rg *gin.RouterGroup, userService services.UserService) {
	router := rg.Group("/server")

	router.POST("", middlewares.DeserializeUser(userService), rc.serverController.CreateServer)
}
