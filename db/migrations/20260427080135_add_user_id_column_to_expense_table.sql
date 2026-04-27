-- +goose Up
-- +goose StatementBegin
ALTER TABLE expense
    ADD COLUMN user_id INT NOT NULL,
    ADD CONSTRAINT fk_expense_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE expense
    DROP CONSTRAINT fk_expense_user,
    DROP COLUMN user_id;
-- +goose StatementEnd
