package controllers

import "github.com/gin-gonic/gin"

type BookingController interface {
	CreateBooking(ctx *gin.Context)
	GetBookingHistory(ctx *gin.Context)
}