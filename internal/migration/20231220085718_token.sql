-- +goose Up
-- +goose StatementBegin
CREATE TABLE Tokens (
    id int NOT NULL AUTO_INCREMENT,
    Token Text NOT NULL,
    Credential_id int NOT NULL, 
    user_id int NOT NULL, 
    Expiration Datetime NOT NULL,
    PRIMARY KEY (id)
    FOREIGN KEY (Credential_id) REFERENCES Credentials(Credential_id)
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Tokens
-- +goose StatementEnd
