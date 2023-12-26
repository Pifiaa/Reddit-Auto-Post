package post

import (
	"time"
)

type Post struct {
	ID        int       `json:"id" gorm:"AUTO_INCREMENT;primaryKey"`
	Message   string    `json:"message" gorm:"not null"`
	Image     string    `json:"image" gorm:"not null"`
	NSFW      string    `json:"nsfw"`
	Subreddit int       `json:"subreddit" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
}
