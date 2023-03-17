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

	authService    services.AuthService
	userService    services.UserService
	sessionService services.SessionService

	AuthController controllers.AuthController
	UserController controllers.UserController

	AuthRouteController routes.AuthRouteController
	UserRouteController routes.UserRouteController
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

	// Controllers
	AuthController = controllers.NewAuthController(authService, userService, sessionService)
	UserController = controllers.NewUserController(userService)

	// Routes
	AuthRouteController = routes.NewAuthRouteController(AuthController)
	UserRouteController = routes.NewUserRouteController(UserController)

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

	AuthRouteController.AuthRoute(router, userService)
	UserRouteController.UserRoute(router, userService)

	log.Fatal(server.Run(":" + config.Port))
}
