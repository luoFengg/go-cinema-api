package routes

import (
	authController "go-cinema-api/controllers/auth"
	movieController "go-cinema-api/controllers/movie"
	showtimeController "go-cinema-api/controllers/showtime"
	studioController "go-cinema-api/controllers/studio"
	"go-cinema-api/exceptions"
	"go-cinema-api/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(studioController studioController.StudioController, movieController movieController.MovieController, showtimeController showtimeController.ShowtimeController, authController authController.AuthController) *gin.Engine {
	router := gin.Default()

	// Add error handler middleware
	router.Use(exceptions.ErrorHandler())

	// Studio Routes
	studioGroup := router.Group("v1/studios")
	studioGroup.GET("/:studioID", studioController.GetStudioByID)
	
	// Movie Routes
	movieGroup := router.Group("v1/movies")
	movieGroup.GET("/", movieController.GetAllMovies)

	// Showtime Routes
	showtimeGroup := router.Group("v1/showtimes")
	showtimeGroup.GET("/", showtimeController.GetShowtimeList)

	// Auth Routes
	authGroup := router.Group("v1/auth")
	authGroup.POST("/register", authController.RegisterUser)
	authGroup.POST("/login", authController.LoginUser)

	// --- RUTE DIPROTEKSI (Harus Login / Bawa Token) ---
	// Admin Only (Movie)
	adminMovie := movieGroup.Group("")
	adminMovie.Use(middleware.AuthMiddleware(), middleware.AdminOnlyMiddleware())
	adminMovie.POST("", movieController.CreateMovie)

	// Admin Only (Showtime)
	adminShowtime := showtimeGroup.Group("")
	adminShowtime.Use(middleware.AuthMiddleware(), middleware.AdminOnlyMiddleware())
	adminShowtime.POST("", showtimeController.CreateShowtime)

	// Admin Only (Studio)
	adminStudio := studioGroup.Group("")
	adminStudio.Use(middleware.AuthMiddleware(), middleware.AdminOnlyMiddleware())
	adminStudio.POST("", studioController.CreateStudio)


	return router
}