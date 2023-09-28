-- +goose Up
-- +goose StatementBegin
ALTER TABLE snippets ADD COLUMN expires_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE snippets DROP COLUMN expires_at;
-- +goose StatementEnd
