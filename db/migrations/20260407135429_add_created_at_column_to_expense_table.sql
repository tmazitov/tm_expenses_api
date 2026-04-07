-- +goose Up
-- +goose StatementBegin
ALTER TABLE expense ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT now();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE expense DROP COLUMN created_at;
-- +goose StatementEnd
