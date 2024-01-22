-- +goose Up
-- +goose StatementBegin
CREATE TABLE Posts (
    Id int NOT NULL AUTO_INCREMENT,
    Title varchar(255) NOT NULL,
    Message varchar(255) NOT NULL,
    Image varchar(255),
    Nsfw boolean,
    Subreddit_id int NOT NULL,
    Create_at timestamp,
    PRIMARY KEY(id),
    FOREIGN KEY (subreddit_id) REFERENCES Subreddits(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Posts
-- +goose StatementEnd

