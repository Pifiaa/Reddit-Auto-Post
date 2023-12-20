-- +goose Up
CREATE TABLE Token (
    Token Text NOT NULL,
    Expiration Datetime NOT NULL
)
-- +goose StatementBegin
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS Token
-- +goose StatementBegin
-- +goose StatementEnd
