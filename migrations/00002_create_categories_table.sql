-- +goose Up
-- +goose StatementBegin
CREATE TABLE `categories` (
  `category_id` int NOT NULL AUTO_INCREMENT,
  `category_name` varchar(25) NOT NULL UNIQUE KEY,
  PRIMARY KEY (`category_id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS categories;
-- +goose StatementEnd
