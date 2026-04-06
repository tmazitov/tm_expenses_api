-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS expense (
    id   UUID PRIMARY KEY,
    name TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS expense;
-- +goose StatementEnd
