-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
