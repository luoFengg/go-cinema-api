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
	GetBookedSeatIDsForShowtime(ctx context.Context, showtimeID string) ([]string, error)
	UpdatePaymentInfo(ctx context.Context, tx *gorm.DB, bookingID string, paymentURL string, paymentToken string) error
	UpdateBookingStatus(ctx context.Context, tx *gorm.DB, bookingID string, status string) error
}