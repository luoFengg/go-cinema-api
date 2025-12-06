package services

import (
	"context"
	"go-cinema-api/models/domain"
	"go-cinema-api/models/web"
)

type BookingService interface {
	CreateBooking(ctx context.Context, userID string, request web.BookingCreateRequest) (*domain.Booking, error)
	GetBookingsByUserID(ctx context.Context, userID string) ([]domain.Booking, error)
}