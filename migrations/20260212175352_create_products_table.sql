-- +goose Up
-- +goose StatementBegin
CREATE TABLE products (
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
