package subreddit

type Subreddit struct {
	ID   int    `json:"id" gorm:"AUTO_INCREMENT;primaryKey"`
	Name string `json:"name" gorm:"not null"`
	URL  string `json:"url" gorm:"not null"`
}
