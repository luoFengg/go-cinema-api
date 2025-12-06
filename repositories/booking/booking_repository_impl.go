package repositories

import (
	"context"
	"go-cinema-api/models/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookingRepositoryImpl struct {
	DB *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &BookingRepositoryImpl{
		DB: db,
	}
}

func (repo *BookingRepositoryImpl) CreateBooking(ctx context.Context, tx *gorm.DB, booking *domain.Booking) (*domain.Booking, error) {
	dbExecutor := repo.DB
	if tx != nil {
		dbExecutor = tx
	}
	err := dbExecutor.WithContext(ctx).Create(booking).Error
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (repo *BookingRepositoryImpl) GetBookingsByUserID(ctx context.Context, userID string) ([]domain.Booking, error) {
	var bookings []domain.Booking
	err := repo.DB.WithContext(ctx).Where("user_id = ?", userID).
	Preload("Showtime").
	Preload("User").
	Preload("Showtime.Studio").
	Preload("Showtime.Movie").
	Preload("BookingSeats").
	Preload("BookingSeats.Seat").
	Order("created_at DESC").
	Find(&bookings).Error

	return bookings, err
}

func (repo *BookingRepositoryImpl) CheckSeatsAvailability(ctx context.Context, tx *gorm.DB, showtimeID string, seatIDs []string) (bool, error) {
	var count int64

	err := tx.WithContext(ctx).Table("booking_seats").
	Joins("JOIN bookings ON bookings.id = booking_seats.booking_id").
	Where("bookings.showtime_id = ?", showtimeID).
	Where("booking_seats.seat_id IN ?", seatIDs).
	Where("bookings.status IN ?", []string{"pending", "confirmed"}).
	Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repo *BookingRepositoryImpl) LockSeats(tx *gorm.DB, seatIDs []string) error {
	var seats []domain.Seat

	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id IN ?", seatIDs).Find(&seats).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *BookingRepositoryImpl) UpdateSeatsAvailability(ctx context.Context, tx *gorm.DB, seatIDs []string, isAvailable bool) error {
	var dbExecutor *gorm.DB
	if tx != nil {
		dbExecutor = tx
	}

	err := dbExecutor.WithContext(ctx).Model(&domain.Seat{}).Where("id IN ?", seatIDs).
	Update("is_available", isAvailable).Error
	return err
}