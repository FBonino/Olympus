package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"server/configs"
	"server/controllers"
	"server/routes"
	"server/services"
	"server/websocket"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	ctx         context.Context
	mongoclient *mongo.Client
	db          *mongo.Database

	authService         services.AuthService
	userService         services.UserService
	sessionService      services.SessionService
	serverService       services.ServerService
	channelService      services.ChannelService
	messageService      services.MessageService
	conversationService services.ConversationService

	AuthController         controllers.AuthController
	UserController         controllers.UserController
	ServerController       controllers.ServerController
	ChannelController      controllers.ChannelController
	ConversationController controllers.ConversationController

	AuthRouteController         routes.AuthRouteController
	UserRouteController         routes.UserRouteController
	ServerRouteController       routes.ServerRouteController
	ChannelRouteController      routes.ChannelRouteController
	ConversationRouteController routes.ConversationRouteController
)

func init() {
	ctx = context.TODO()
	config, err := configs.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	mongoconn := options.Client().ApplyURI(config.DB)
	mongoclient, err = mongo.Connect(ctx, mongoconn)

	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	db = mongoclient.Database(config.DBName)

	// Services
	authService = services.NewAuthService(db, ctx)
	userService = services.NewUserService(db, ctx)
	sessionService = services.NewSessionService(db, ctx)
	serverService = services.NewServerService(db, ctx)
	channelService = services.NewChannelService(db, ctx)
	messageService = services.NewMessageService(db, ctx)
	conversationService = services.NewConversationService(db, ctx)

	// Controllers
	AuthController = controllers.NewAuthController(authService, userService, sessionService)
	UserController = controllers.NewUserController(userService, serverService)
	ServerController = controllers.NewServerController(serverService, channelService)
	ChannelController = controllers.NewChannelController(channelService, messageService)
	ConversationController = controllers.NewConversationController(conversationService, messageService, userService)

	// Routes
	AuthRouteController = routes.NewAuthRouteController(AuthController)
	UserRouteController = routes.NewUserRouteController(UserController)
	ServerRouteController = routes.NewServerRouteController(ServerController)
	ChannelRouteController = routes.NewChannelRouteController(ChannelController)
	ConversationRouteController = routes.NewConversationRouteController(ConversationController)

	server = gin.Default()
}

func main() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	server.Static("/uploads", "./uploads")

	config, err := configs.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load config", err)
	}

	router := server.Group("api")

	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok"})
	})

	wsServer := websocket.NewWebsocketServer()

	go wsServer.Run()

	router.GET("/ws/:user/:channel", func(ctx *gin.Context) {
		userID := ctx.Param("user")
		channelID := ctx.Param("channel")

		websocket.ServeWs(wsServer, ctx, userID, channelID)
	})

	AuthRouteController.AuthRoute(router, userService)
	UserRouteController.UserRoute(router, userService)
	ServerRouteController.ServerRoute(router, userService)
	ChannelRouteController.ChannelRoute(router, userService)
	ConversationRouteController.ConversationRoute(router, userService)

	log.Fatal(server.Run(":" + config.Port))
}
