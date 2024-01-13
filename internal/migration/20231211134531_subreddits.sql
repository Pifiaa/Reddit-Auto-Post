-- +goose Up
CREATE TABLE Subreddits (
    Id int NOT NULL AUTO_INCREMENT,
    Name varchar(255) NOT NULL,
    Url varchar(255) NOT NULL,
    PRIMARY KEY(id)
) 
-- +goose StatementBegin
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS Subreddits
-- +goose StatementBegin
-- +goose StatementEnd
