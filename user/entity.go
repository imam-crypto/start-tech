package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             int            `json:"id" gorm:"primaryKey"`
	First_name     string         `json:"first_name"`
	Email          string         `gorm:"not null;email"`
	Password       string         `gorm:"not null"`
	AvatarFileName string         `gorm:"null"`
	CreatedAt      time.Time      `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
}
