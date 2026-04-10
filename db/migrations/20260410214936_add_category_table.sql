-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS category(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    icon TEXT DEFAULT NULL,
    color INT DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS category;
-- +goose StatementEnd
