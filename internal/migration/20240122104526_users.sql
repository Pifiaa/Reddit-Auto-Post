-- +goose Up
-- +goose StatementBegin
CREATE TABLE Users (
    Id int NOT NULL AUTO_INCREMENT,
    Username varchar(255) NOT NULL,
    Password varchar(255) NOT NULL,
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Users
-- +goose StatementEnd
