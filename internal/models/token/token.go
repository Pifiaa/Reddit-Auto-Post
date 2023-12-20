package token

import (
	"time"

	"gorm.io/gorm"
)

type AccessToken struct {
	gorm.Model
	Token      string    `json:"token" gorm:"not null"`
	Expiration time.Time `json:"expiration" gorm:"not null"`
}
