package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Movie struct {
	ID          string `json:"id" gorm:"primaryKey;column:id"`
	Title       string `json:"title" gorm:"column:title;type:varchar(255);not null"`
	Description string `json:"description" gorm:"column:description;type:text"`
	DurationMin int    `json:"duration_min" gorm:"column:duration_min;type:int;not null"` // Durasi dalam menit
	ReleaseDate time.Time `json:"release_date" gorm:"column:release_date;type:date"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (movie *Movie) TableName() string {
	return "movies"
}

func (movie *Movie) BeforeCreate(tx *gorm.DB) error {
	if movie.ID == "" {
		// Generate UUID baru
		uuidObj := uuid.New()
		movie.ID = "movie-" + uuidObj.String()
	}
	return nil
}
