-- +goose Up
-- +goose StatementBegin
CREATE TABLE `cartItems` (
  `cart_id` int NOT NULL,
  `product_id` int NOT NULL,
  `quantity` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`cart_id`, `product_id`),
  FOREIGN KEY (`cart_id`) REFERENCES `carts` (`cart_id`),
  FOREIGN KEY (`product_id`) REFERENCES `products` (`product_id`)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cartItems;
-- +goose StatementEnd
