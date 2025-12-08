package main

import (
	"fmt"
	"go-cinema-api/apps/databases"
	"go-cinema-api/apps/routes"
	"go-cinema-api/utils"
	"log"
	"os"

	studioController "go-cinema-api/controllers/studio"
	studioRepository "go-cinema-api/repositories/studio"
	studioService "go-cinema-api/services/studio"

	movieController "go-cinema-api/controllers/movie"
	movieRepository "go-cinema-api/repositories/movie"
	movieService "go-cinema-api/services/movie"

	showtimeController "go-cinema-api/controllers/showtime"
	showtimeRepository "go-cinema-api/repositories/showtime"
	showtimeService "go-cinema-api/services/showtime"

	authController "go-cinema-api/controllers/auth"
	authRepository "go-cinema-api/repositories/user"
	authService "go-cinema-api/services/auth"

	bookingController "go-cinema-api/controllers/booking"
	bookingRepository "go-cinema-api/repositories/booking"
	bookingService "go-cinema-api/services/booking"

	paymentController "go-cinema-api/controllers/payment"
	paymentService "go-cinema-api/services/payment"

	"github.com/joho/godotenv"
)

func main() {
	// 1. Load configuraion from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Note: .env file not found, using system environment variables")
	}

	// 2. Connect to the database
	db, errDB := databases.ConnectDatabase()
	if errDB != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 3. Initialize Midtrans
	utils.InitMidtrans()

	// 4. Initialize repository, service, and controller

	// Studio
	studioRepository := studioRepository.NewStudioRepository(db)
	studioService := studioService.NewStudioService(studioRepository)
	studioController := studioController.NewStudioController(studioService)
	
	// Movie
	movieRepository := movieRepository.NewMovieRepository(db)
	movieService := movieService.NewMovieService(movieRepository)
	movieController := movieController.NewMovieController(movieService)

	// Showtime
	showtimeRepository := showtimeRepository.NewShowtimeRepository(db)
	// Booking repository is needed by showtime service, initialize it first
	bookingRepository := bookingRepository.NewBookingRepository(db)
	showtimeService := showtimeService.NewShowtimeService(showtimeRepository, movieRepository, studioRepository, bookingRepository)
	showtimeController := showtimeController.NewShowtimeController(showtimeService)

	// Booking
	bookingService := bookingService.NewBookingService(bookingRepository, showtimeRepository, db)
	bookingController := bookingController.NewBookingController(bookingService)

	// Auth
	authRepository := authRepository.NewUserRepository(db)
	authService := authService.NewAuthService(authRepository)
	authController := authController.NewAuthController(authService)

	// Payment
	paymentService := paymentService.NewPaymentService(bookingRepository, db)
	paymentController := paymentController.NewPaymentController(paymentService)

	// 5. Set up the router
	router := routes.NewRouter(studioController, movieController, showtimeController, authController, bookingController, paymentController)

	// 6. Start the server
	address := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))

	err = router.Run(address)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	
}