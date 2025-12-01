package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Seat struct {
	ID       string `json:"id" gorm:"primaryKey;column:id"`
	StudioID string `json:"studio_id" gorm:"column:studio_id;type:varchar(50);not null;index"`
	Row      string `json:"row" gorm:"column:row;type:varchar(5);not null"`
	Number   int    `json:"number" gorm:"column:number;type:int;not null"`

	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`

	// Status kursi (Available/Maintenance).
	// Default true artinya kursi bisa dipakai.
	IsAvailable bool `gorm:"default:true" json:"is_available"`
}

func (seat *Seat) TableName() string {
	return "seats"
}

func (seat *Seat) BeforeCreate(tx *gorm.DB) error {
	if seat.ID == "" {
		// Generate UUID baru
		uuidObj := uuid.New()
		seat.ID = "seat-" + uuidObj.String()
	}
	return nil
}