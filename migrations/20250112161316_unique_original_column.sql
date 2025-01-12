-- +goose Up
-- +goose StatementBegin
--ALTER TABLE shortener ADD CONSTRAINT original_unique UNIQUE(original_url, short_url);
CREATE UNIQUE INDEX original_unique ON shortener (original_url);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE shortener DROP CONSTRAINT IF EXISTS original_unique;
-- +goose StatementEnd
