package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Showtime struct {
	ID        string    `json:"id" gorm:"primaryKey;column:id"`
	MovieID   string    `json:"movie_id" gorm:"column:movie_id;type:varchar(100);not null;index"`
	StudioID  string    `json:"studio_id" gorm:"column:studio_id;type:varchar(100);not null;index"`
	StartTime time.Time `json:"start_time" gorm:"column:start_time;not null"`
	EndTime   time.Time `json:"end_time" gorm:"column:end_time;not null"`
	Price    float64       `json:"price" gorm:"column:price;type:decimal(10,2);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	// Relasi (Belongs To)
	// Agar saat ambil data jadwal, data Studio & Movie-nya bisa ikut terbawa.
	Studio Studio `json:"studio,omitempty" gorm:"foreignKey:StudioID"`
	Movie  Movie  `json:"movie,omitempty" gorm:"foreignKey:MovieID"`
}

func (showtime *Showtime) TableName() string {
	return "showtimes"
}

func (showtime *Showtime) BeforeCreate(tx *gorm.DB) error {
	if showtime.ID == "" {
		// Generate UUID baru
		uuidObj := uuid.New()
		showtime.ID = "showtime-" + uuidObj.String()
	}
	return nil
}