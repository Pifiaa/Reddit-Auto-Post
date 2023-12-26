package token

import (
	"time"
)

type Tokens struct {
	Token      string    `json:"token" gorm:"not null"`
	Expiration time.Time `json:"expiration" gorm:"not null"`
}
