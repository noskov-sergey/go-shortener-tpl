-- +goose Up
-- +goose StatementBegin
ALTER TABLE shortener ADD COLUMN username VARCHAR;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE shortener DROP COLUMN username;
-- +goose StatementEnd
