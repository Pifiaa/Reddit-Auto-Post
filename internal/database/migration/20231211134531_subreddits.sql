-- +goose Up
CREATE TABLE Subreddits (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    url varchar(255) NOT NULL,
    PRIMARY KEY(id)
) 
-- +goose StatementBegin
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS Subreddits
-- +goose StatementBegin
-- +goose StatementEnd
