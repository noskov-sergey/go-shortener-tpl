-- +goose Up
-- +goose StatementBegin
ALTER TABLE shortener ADD COLUMN is_deleted BOOLEAN DEFAULT FALSE NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE shortener DROP COLUMN is_deleted;
-- +goose StatementEnd
