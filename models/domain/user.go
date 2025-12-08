package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `json:"id" gorm:"primaryKey;column:id"`
	
	Username  string    `json:"username" gorm:"column:username;type:varchar(30);unique;not null;"`
	Email     string    `json:"email" gorm:"column:email;type:varchar(255);unique;not null"`
	FirstName string    `json:"first_name" gorm:"column:first_name;type:varchar(100)"`
	LastName  string    `json:"last_name" gorm:"column:last_name;type:varchar(100)"`
	Password  string    `json:"-" gorm:"column:password;type:varchar(255);not null"`
	Role      string    `json:"role" gorm:"column:role;type:varchar(50);not null;default:'customer'"`
	
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;not null;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;not null;autoCreateTime;autoUpdateTime"`
}

func (user *User) TableName() string {
	return "users"
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	if user.ID == "" {
		// Generate UUID baru
		uuidObj := uuid.New()
		user.ID = "user-" + uuidObj.String()
	}
	return nil
}