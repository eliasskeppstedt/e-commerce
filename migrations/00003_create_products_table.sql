-- +goose Up
-- +goose StatementBegin
CREATE TABLE `products` (
  `product_id` int NOT NULL AUTO_INCREMENT,
  `product_name` varchar(25) NOT NULL UNIQUE KEY,
  `price` float NOT NULL,
  `category_name` varchar(25) NOT NULL UNIQUE KEY,
  PRIMARY KEY (`product_id`),
  FOREIGN KEY (`category_name`) REFERENCES `categories` (`category_name`) 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
