package post

import "time"

type post struct {
	id        int       `json:"id"`
	message   string    `json:"message"`
	image     string    `json:"image"`
	nsfw      string    `json:"nsfw"`
	subreddit int       `json:"subreddit"`
	create_at time.Time `json:"create_at"`
}
