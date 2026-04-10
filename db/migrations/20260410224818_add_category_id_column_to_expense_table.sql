-- +goose Up
-- +goose StatementBegin
ALTER TABLE expense
    ADD COLUMN category_id UUID NULL REFERENCES category(id) ON DELETE SET NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE expense DROP COLUMN category_id;
-- +goose StatementEnd
