-- +goose Up
-- +goose StatementBegin
CREATE TABLE orderItems (
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orderItems;
-- +goose StatementEnd
