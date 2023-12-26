package token

import (
	"time"
)

type AccessToken struct {
	Token      string    `json:"token" gorm:"not null"`
	Expiration time.Time `json:"expiration" gorm:"not null"`
}
