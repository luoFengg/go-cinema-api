package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Booking struct {
	ID         string  `gorm:"primaryKey;type:varchar(100)" json:"id"`
	
	UserID     string  `gorm:"type:varchar(100);not null" json:"user_id"`
	ShowtimeID string  `gorm:"type:varchar(100);not null" json:"showtime_id"`
	TotalPrice float64 `gorm:"type:decimal(10,2);not null" json:"total_price"`
	Status     string  `gorm:"type:varchar(50);not null;default:'pending'" json:"status"`
	PaymentURL string `gorm:"type:text" json:"payment_url,omitempty"`
	PaymentToken string `gorm:"type:varchar(255)" json:"payment_token,omitempty"`
	
	CreatedAt  time.Time  `gorm:"type:timestamptz;not null;autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"type:timestamptz;not null;autoCreateTime;autoUpdateTime" json:"updated_at"`

	// Relations
	// 1. Relasi ke User (Yang pesan siapa?)
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`

	// 2. Relasi ke Showtime (Nonton apa?)
	Showtime Showtime `gorm:"foreignKey:ShowtimeID" json:"showtime,omitempty"`

	// 3. Relasi ke Detail Tiket (Has Many)
	// Satu Booking punya BANYAK BookingSeats
	BookingSeats []BookingSeat `gorm:"foreignKey:BookingID;constraint:OnDelete:CASCADE" json:"booking_seats,omitempty"`
}

type BookingSeat struct {
	ID        string  `gorm:"primaryKey;type:varchar(100)" json:"id"`
	BookingID string  `gorm:"type:varchar(100);not null;index" json:"booking_id"`
	SeatID    string  `gorm:"type:varchar(100);not null;index" json:"seat_id"`
	Price     float64 `gorm:"type:decimal(10,2);not null" json:"price"`

	// Relations
	// Relasi balik ke kursi fisik (untuk tahu ini kursi nomor berapa, row apa)
	Seat Seat `gorm:"foreignKey:SeatID" json:"seat,omitempty"`
}

func (booking *Booking) TableName() string {
	return "bookings"
}

func (bookingSeat *BookingSeat) TableName() string {
	return "booking_seats"
}

func (booking *Booking) BeforeCreate(tx *gorm.DB) error {
	if booking.ID == "" {
		// Generate UUID baru
		uuidObj := uuid.New()
		booking.ID = "booking-" + uuidObj.String()
	}
	return nil
}

func (bookingSeat *BookingSeat) BeforeCreate(tx *gorm.DB) error {
	if bookingSeat.ID == "" {
		// Generate UUID baru
		uuidObj := uuid.New()
		bookingSeat.ID = "bookingseat-" + uuidObj.String()
	}
	return nil
}