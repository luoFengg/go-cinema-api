package routes

import (
	movieController "go-cinema-api/controllers/movie"
	showtimeController "go-cinema-api/controllers/showtime"
	studioController "go-cinema-api/controllers/studio"
	"go-cinema-api/exceptions"

	"github.com/gin-gonic/gin"
)

func NewRouter(studioController studioController.StudioController, movieController movieController.MovieController, showtimeController showtimeController.ShowtimeController) *gin.Engine {
	router := gin.Default()

	// Add error handler middleware
	router.Use(exceptions.ErrorHandler())

	// Studio Routes
	studioGroup := router.Group("v1/studios")
	studioGroup.POST("/", studioController.CreateStudio)
	studioGroup.GET("/:studioID", studioController.GetStudioByID)
	
	// Movie Routes
	movieGroup := router.Group("v1/movies")
	movieGroup.POST("/", movieController.CreateMovie)
	movieGroup.GET("/", movieController.GetAllMovies)

	// Showtime Routes
	showtimeGroup := router.Group("v1/showtimes")
	showtimeGroup.POST("/", showtimeController.CreateShowtime)
	showtimeGroup.GET("/", showtimeController.GetShowtimeList)
	return router
}