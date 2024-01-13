-- +goose Up
CREATE TABLE Tokens (
    id int NOT NULL AUTO_INCREMENT,
    Token Text NOT NULL,
    Credential_id int NOT NULL, 
    Expiration Datetime NOT NULL,
    PRIMARY KEY (id)
    FOREIGN KEY (Credential_id) REFERENCES Credentials(Credential_id)
)
-- +goose StatementBegin
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS Tokens
-- +goose StatementBegin
-- +goose StatementEnd
