package controllers

import (
	"go-cinema-api/models/web"
	services "go-cinema-api/services/booking"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingControllerImpl struct {
	bookingService services.BookingService
}

func NewBookingController(bookingService services.BookingService) BookingController {
	return &BookingControllerImpl{
		bookingService: bookingService,
	}
}

func (controller *BookingControllerImpl) CreateBooking(ctx *gin.Context) {
	// 1. Get userID from context
	userIDInterface, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, web.WebResponse{
			Success: false,
			Message: "Unauthorized",
			Data:    nil,
		})
		return
	}
	
	userID := userIDInterface.(string)

	// 2. Take booking request from body
	var request web.BookingCreateRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(400, web.WebResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	// 3. Call service to create booking
	booking, errBooking  := controller.bookingService.CreateBooking(ctx, userID, request)
	if errBooking != nil {
		ctx.Error(errBooking)
		return
	}

	// 4. Return response
	ctx.JSON(http.StatusCreated, web.WebResponse{
		Success: true,
		Message: "Booking created successfully",
		Data:    booking,
	})

}
func (controller *BookingControllerImpl) GetBookingHistory(ctx *gin.Context) {
	// 1. Get userID from context
	userIDInterface, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, web.WebResponse{
			Success: false,
			Message: "Unauthorized",
			Data:    nil,
		})
		return
	}
	// Convert userID to string
	userID := userIDInterface.(string)


	// 2. Call service to get booking history
	bookings, err := controller.bookingService.GetBookingsByUserID(ctx.Request.Context(), userID)
	if err != nil {
		ctx.Error(err)
		return
	}
	
	// 3. Return response
	ctx.JSON(http.StatusOK, web.WebResponse{
		Success: true,
		Message: "Booking history retrieved successfully",
		Data:    bookings,
	})
}