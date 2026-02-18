-- +goose Up
-- +goose StatementBegin
CREATE TABLE `orderItems` (
  `product_id` int NOT NULL,
  `quantity` int NOT NULL,
  `order_id` int NOT NULL AUTO_INCREMENT,
  KEY `product_id` (`product_id`),
  KEY `order_id_idx` (`order_id`),
  CONSTRAINT `order_id` FOREIGN KEY (`order_id`) REFERENCES `orders` (`order_id`),
  CONSTRAINT `orderItems_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `products` (`product_id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orderItems;
-- +goose StatementEnd
