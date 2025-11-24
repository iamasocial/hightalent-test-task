-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS questions (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    text TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE questions;
-- +goose StatementEnd
