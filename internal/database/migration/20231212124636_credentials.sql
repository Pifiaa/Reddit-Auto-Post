-- +goose Up
CREATE TABLE Credentials (
    id int NOT NULL AUTO_INCREMENT,
    username varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    client_secret varchar(255) NOT NULL,
    client_id varchar(255) NOT NULL,
    PRIMARY KEY(id)
)
-- +goose StatementBegin
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS Credentials
-- +goose StatementBegin
-- +goose StatementEnd
