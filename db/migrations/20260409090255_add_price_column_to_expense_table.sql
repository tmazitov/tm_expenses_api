-- +goose Up
-- +goose StatementBegin
ALTER TABLE expense ADD COLUMN price NUMERIC NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE expense DROP COLUMN price; 
-- +goose StatementEnd
