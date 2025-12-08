package routes

import (
	authController "go-cinema-api/controllers/auth"
	bookingController "go-cinema-api/controllers/booking"
	movieController "go-cinema-api/controllers/movie"
	paymentController "go-cinema-api/controllers/payment"
	showtimeController "go-cinema-api/controllers/showtime"
	studioController "go-cinema-api/controllers/studio"
	"go-cinema-api/exceptions"
	"go-cinema-api/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(studioController studioController.StudioController, movieController movieController.MovieController, showtimeController showtimeController.ShowtimeController, authController authController.AuthController, bookingController bookingController.BookingController, paymentController paymentController.PaymentController) *gin.Engine {
	router := gin.Default()

	// Add error handler middleware
	router.Use(exceptions.ErrorHandler())

	// Studio Routes
	studioGroup := router.Group("v1/studios")
	studioGroup.GET("/:studioID", studioController.GetStudioByID)

	// Webhook route for payment gateway
	webhookGroup := router.Group("v1/webhooks")
	webhookGroup.POST("/payments", paymentController.HandleWebhook)

	
	// Movie Routes
	movieGroup := router.Group("v1/movies")
	movieGroup.GET("/", movieController.GetAllMovies)

	// Showtime Routes
	showtimeGroup := router.Group("v1/showtimes")
	showtimeGroup.GET("/", showtimeController.GetShowtimeList)
	showtimeGroup.GET("/:showtimeID/seat-map", showtimeController.GetSeatMap)

	// Auth Routes
	authGroup := router.Group("v1/auth")
	authGroup.POST("/register", authController.RegisterUser)
	authGroup.POST("/login", authController.LoginUser)

	// --- PROTECTED ROUTES ---
	// Booking Routes
	bookingGroup := router.Group("v1/bookings")
	bookingGroup.Use(middleware.AuthMiddleware())
	bookingGroup.POST("/", bookingController.CreateBooking)
	bookingGroup.GET("/", bookingController.GetBookingHistory)

	// --- ADMIN ROUTE  ---
	// Admin Only (Movie)
	adminMovie := movieGroup.Group("/")
	adminMovie.Use(middleware.AuthMiddleware(), middleware.AdminOnlyMiddleware())
	adminMovie.POST("/", movieController.CreateMovie)

	// Admin Only (Showtime)
	adminShowtime := showtimeGroup.Group("")
	adminShowtime.Use(middleware.AuthMiddleware(), middleware.AdminOnlyMiddleware())
	adminShowtime.POST("/", showtimeController.CreateShowtime)

	// Admin Only (Studio)
	adminStudio := studioGroup.Group("/")
	adminStudio.Use(middleware.AuthMiddleware(), middleware.AdminOnlyMiddleware())
	adminStudio.POST("/", studioController.CreateStudio)


	return router
}