package repositories

import (
	"context"
	"go-cinema-api/models/domain"

	"gorm.io/gorm"
)

type BookingRepository interface {
	CreateBooking(ctx context.Context, tx *gorm.DB, booking *domain.Booking) (*domain.Booking, error)
	GetBookingsByUserID(ctx context.Context, userID string) ([]domain.Booking, error)
	CheckSeatsAvailability(ctx context.Context,tx *gorm.DB, showtimeID string, seatIDs []string) (bool, error)
	LockSeats(tx *gorm.DB, seatIDs []string) error
	UpdateSeatsAvailability(ctx context.Context, tx *gorm.DB, seatIDs []string, isAvailable bool) error
}