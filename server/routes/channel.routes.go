package routes

import (
	"server/controllers"
	"server/middlewares"
	"server/services"

	"github.com/gin-gonic/gin"
)

type ChannelRouteController struct {
	channelController controllers.ChannelController
}

func NewChannelRouteController(channelController controllers.ChannelController) ChannelRouteController {
	return ChannelRouteController{channelController}
}

func (rc *ChannelRouteController) ChannelRoute(rg *gin.RouterGroup, userService services.UserService) {
	router := rg.Group("/channel")

	router.GET("/:channel", middlewares.DeserializeUser(userService), rc.channelController.GetChannel)
	router.POST("/:channel/messages", middlewares.DeserializeUser(userService), rc.channelController.NewMessage)
}
