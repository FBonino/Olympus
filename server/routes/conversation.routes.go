package routes

import (
	"server/controllers"
	"server/middlewares"
	"server/services"

	"github.com/gin-gonic/gin"
)

type ConversationRouteController struct {
	conversationController controllers.ConversationController
}

func NewConversationRouteController(conversationController controllers.ConversationController) ConversationRouteController {
	return ConversationRouteController{conversationController}
}

func (rc *ConversationRouteController) ConversationRoute(rg *gin.RouterGroup, userService services.UserService) {
	router := rg.Group("/conversation")

	router.GET("/me", middlewares.DeserializeUser(userService), rc.conversationController.GetUserConversations)
	router.POST("", middlewares.DeserializeUser(userService), rc.conversationController.Create)
	router.GET("/:conversation", middlewares.DeserializeUser(userService), rc.conversationController.GetConversation)
	router.POST("/:conversation/messages", middlewares.DeserializeUser(userService), rc.conversationController.NewMessage)
}
