package credentials

import "gorm.io/gorm"

type credentials struct {
	gorm.Model
	id            int    `json:"id" gorm:AUTO_INCREMENT;"primaryKey"`
	usename       int    `json:"usename" gorm:"not null"`
	password      int    `json:"password" gorm:"not null"`
	client_secret string `json:"client_secret" gorm:"not null"`
	client_id     string `json:"client_id" gorm:"not null"`
}
