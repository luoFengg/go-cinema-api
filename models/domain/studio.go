package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Studio struct {
	ID       string `json:"id" gorm:"primaryKey;column:id"`
	Name     string `json:"name" gorm:"column:name;type:varchar(100);not null"`
	Capacity int    `json:"capacity" gorm:"column:capacity;type:int;not null"`

	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`

	// Relasi: Satu Studio punya BANYAK Kursi
	// 'constraint:OnDelete:CASCADE' artinya: Kalau Studio dihapus, kursinya ikut terhapus otomatis.
	Seats []Seat `json:"seats,omitempty" gorm:"foreignKey:StudioID;constraint:OnDelete:CASCADE"`
}

func (studio *Studio) TableName() string {
	return "studios"
}

func (studio *Studio) BeforeCreate(tx *gorm.DB) error {
	if studio.ID == "" {
		// Generate UUID baru
		uuidObj := uuid.New()
		studio.ID = "studio-" + uuidObj.String()
	}
	return nil
}