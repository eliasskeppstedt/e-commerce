-- +goose Up
-- +goose StatementBegin
CREATE TABLE `products` (
  `product_id` int NOT NULL AUTO_INCREMENT,
  `product_name` varchar(25) NOT NULL UNIQUE KEY,
  `stock`int NOT NULL DEFAULT 0,
  `price` DECIMAL(10,2) NOT NULL,
  `manufacturer` varchar(100) NOT NULL,
  `description`varchar(255) NOT NULL,
  `category_id` int NOT NULL,
  PRIMARY KEY (`product_id`),
  FOREIGN KEY (`category_id`) REFERENCES `categories` (`category_id`)
  ON DELETE CASCADE
  ON UPDATE CASCADE 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
