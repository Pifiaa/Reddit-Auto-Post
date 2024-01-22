-- +goose Up
-- +goose StatementBegin
CREATE TABLE Credentials (
    Id int NOT NULL AUTO_INCREMENT,
    Username varchar(255) NOT NULL,
    Password varchar(255) NOT NULL,
    Client_secret varchar(255) NOT NULL,
    Client_id varchar(255) NOT NULL,
    PRIMARY KEY(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Credentials
-- +goose StatementEnd
