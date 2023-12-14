package subreddit

import "gorm.io/gorm"

type subreddit struct {
	gorm.Model
	id   int    `json:"id" gorm:AUTO_INCREMENT;"primaryKey"`
	name string `json:"name" gorm:"not null"`
	url  string `json:"url" gorm:"not null"`
}
