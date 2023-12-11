-- +goose Up
CREATE TABLE Posts {
    id int NOT NULL,
    title varchar,
    message varchar,
    image varchar,
    nsfw boolean,
    subreddit_id int,
    create_at timestamp,
    PRIMARY KEY(id),
    FOREIGN KEY (subreddit_id) REFERENCES Subreddits(id)
}
-- +goose StatementBegin
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS Posts
-- +goose StatementBegin
-- +goose StatementEnd

