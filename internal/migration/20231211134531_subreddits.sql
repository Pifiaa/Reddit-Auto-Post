-- +goose Up
-- +goose StatementBegin
CREATE TABLE Subreddits (
    Id int NOT NULL AUTO_INCREMENT,
    Name varchar(255) NOT NULL,
    Url varchar(255) NOT NULL,
    PRIMARY KEY(id)
) 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Subreddits
-- +goose StatementEnd
