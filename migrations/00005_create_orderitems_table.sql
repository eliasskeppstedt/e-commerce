-- +goose Up
-- +goose StatementBegin
CREATE TABLE order_items (
  `order_item_id` INT NOT NULL AUTO_INCREMENT,
  `order_id` INT NOT NULL,
  `product_id` INT NOT NULL,
  `quantity` INT NOT NULL,
  `price_at_purchase` DECIMAL(10,2) NOT NULL,
  `product_name` VARCHAR(25),
  PRIMARY KEY (`order_item_id`),
  FOREIGN KEY (`order_id`) REFERENCES orders(order_id) ON DELETE CASCADE,
  UNIQUE (`order_id`, `product_id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS order_items;
-- +goose StatementEnd
