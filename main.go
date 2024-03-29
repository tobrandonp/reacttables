package main

import (
	"fmt"
	"log"
	"reacttables/config"
	"reacttables/router"
	"reacttables/utils"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load env vars from dotenv
	cfg := config.LoadConfig()

	// Setup MongoDB connection
	// mongoDB, _ := mongodb.NewMongoDB(cfg.MongoDbUri, "albums") // First attempt
	mongoDB, err := utils.NewMongoDB(cfg.MongoDbUri, "albums")
	if err != nil {
		log.Fatalf("Error initializing MongoDB connection: %v", err)
	}
	defer mongoDB.CloseConnection()

	r := gin.Default()

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
	fmt.Printf("Allowed origins: %v", cfg.AllowedOrigins)
	r.Use(cors.New(cors.Config{
		// AllowOrigins:     cfg.AllowedOrigins,
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type"}, // Include Content-Type if you're sending data
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: time.Duration(cfg.CorsMaxAge) * time.Hour,
	}))

	// r.LoadHTMLGlob("views/*")

	// Initialize routes
	router.InitRoutes(r)

	// Start the server
	r.Run(":8080")
}
