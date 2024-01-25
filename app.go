package main

import (
	"log"
	"reacttables/config"
	"reacttables/router"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router       *gin.Engine
	AlbumService *services.AlbumService
	// ... other services
}

func NewApp() *App {
	// Load env vars from dotenv
	cfg := config.LoadConfig()

	// Initialize MongoDB connection
	mongoDB, err := mongodb.NewMongoDB(cfg.MongoDbUri, "albums")
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Initialize services
	albumService := services.NewAlbumService(mongoDB)

	// Setup Gin Router
	router := gin.Default()

	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	// - AllowOriginFunc allows us to create custom logic to allow specific hosts.
	//     Whereas AllowOrigins is a static list of allowed hosts, AllowOriginFunc can be dymanic
	//     Situations on where this is useful:
	//    -Environment use case: "localhost" ok for DEVELOPMENT but not for PRODUCTION
	//    -Pattern matching: "*.example.com"
	router.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: time.Duration(cfg.CorsMaxAge) * time.Hour,
	}))

	router.LoadHTMLGlob("views/*")

	app := &App{
		Router:       router,
		AlbumService: albumService,
		// ... initialize other services
	}

	return app
}

func (app *App) InitializeRoutes() {
	router.InitRoutes(app.Router, app.AlbumService)
	// ... initialize other routes
}

func (app *App) Run() {
	app.Router.Run(":8080")
}
