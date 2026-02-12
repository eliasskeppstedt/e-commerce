-- +goose Up
-- +goose StatementBegin
CREATE TABLE categories (
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS categories;
-- +goose StatementEnd
