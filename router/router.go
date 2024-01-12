package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// InitRoutes sets up the routes for the application
func InitRoutes(r *gin.Engine) {
	// Static file serving
	r.Static("/static", "./static")

	// Route for the "Table Example" page
	// r.GET("/table-example", tableExampleHandler)

	r.GET("/albums", getAlbums)
	r.POST("/albums", postAlbums)

	// Route for the "About" page
	r.GET("/about", aboutHandler)

	// Create a Prometheus histogram
	httpRequestDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests.",
			Buckets: prometheus.DefBuckets, // You can define custom buckets
		},
		[]string{"path"},
	)

	// Register the histogram with Prometheus
	prometheus.MustRegister(httpRequestDuration)

	// Use a Gin middleware to measure request duration
	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next() // Process request
		duration := time.Since(start)
		httpRequestDuration.WithLabelValues(c.Request.URL.Path).Observe(duration.Seconds())
	})
	// Expose the Prometheus metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbums(c *gin.Context) {
	// Handle the "Table Example" page
	c.IndentedJSON(http.StatusOK, albums)
	// c.HTML(http.StatusOK, "views/table-example.html", albums)
}

func aboutHandler(c *gin.Context) {
	// Handle the "About" page
	c.HTML(http.StatusOK, "views/about.html", nil)
}
