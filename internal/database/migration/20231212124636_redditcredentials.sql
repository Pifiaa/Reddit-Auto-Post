-- +goose Up
CREATE TABLE RedditCredentials (
    id int NOT NULL,
    username varchar(255),
    password varchar(255),
    client_secret varchar(255),
    client_id varchar(255),
    PRIMARY KEY(id)
)
-- +goose StatementBegin
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
