-- +goose Up
CREATE TABLE Posts (
    id int NOT NULL,
    title varchar(255),
    message varchar(255),
    image varchar(255),
    nsfw boolean,
    subreddit_id int,
    create_at timestamp,
    PRIMARY KEY(id),
    FOREIGN KEY (subreddit_id) REFERENCES Subreddits(id)
)
-- +goose StatementBegin
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS Posts
-- +goose StatementBegin
-- +goose StatementEnd

