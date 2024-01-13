-- +goose Up
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
-- +goose StatementBegin
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS Posts
-- +goose StatementBegin
-- +goose StatementEnd

