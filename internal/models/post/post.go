package post

import (
	"time"

	"gorm.io/gorm"
)

type post struct {
	gorm.Model
	id        int       `json:"id" gorm:AUTO_INCREMENT;"primaryKey"`
	message   string    `json:"message" gorm:"not null"`
	image     string    `json:"image" gorm:"not null"`
	nsfw      string    `json:"nsfw" `
	subreddit int       `json:"subreddit" gorm:"not null"`
	create_at time.Time `json:"create_at"`
}
