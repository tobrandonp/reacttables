package main

import (
	"reacttables/config"
	"reacttables/router" // Import your router package
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))

	r.LoadHTMLGlob("views/*")

	// Initialize routes
	router.InitRoutes(r)

	// Start the server
	r.Run(":8080")
}
