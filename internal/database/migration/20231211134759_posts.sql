-- +goose Up
CREATE TABLE Posts (
    id int NOT NULL AUTO_INCREMENT,
    title varchar(255) NOT NULL,
    message varchar(255) NOT NULL,
    image varchar(255),
    nsfw boolean,
    subreddit_id int NOT NULL,
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

