-- +goose Up
CREATE TABLE RedditCredentials (
    id int NOT NULL,
    username varchar,
    password varchar,

    PRIMARY KEY(id),
)
-- +goose StatementBegin
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
