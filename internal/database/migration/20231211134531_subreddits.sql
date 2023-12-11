-- +goose Up
CREATE TABLE Subreddits {
    id int NOT NULL,
    name varchar,
    url varchar,
    PRIMARY KEY(id)
} 
-- +goose StatementBegin
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS Subreddits
-- +goose StatementBegin
-- +goose StatementEnd
