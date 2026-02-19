-- +goose Up
-- +goose StatementBegin
CREATE TABLE `carts` (
  `cart_id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  PRIMARY KEY (`cart_id`),
  FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS carts;
-- +goose StatementEnd
