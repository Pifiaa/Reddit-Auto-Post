-- +goose Up
CREATE TABLE Credentials (
    Id int NOT NULL AUTO_INCREMENT,
    Username varchar(255) NOT NULL,
    Password varchar(255) NOT NULL,
    Client_secret varchar(255) NOT NULL,
    Client_id varchar(255) NOT NULL,
    PRIMARY KEY(id)
)
-- +goose StatementBegin
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS Credentials
-- +goose StatementBegin
-- +goose StatementEnd
