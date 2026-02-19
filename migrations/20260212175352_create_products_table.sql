  -- +goose Up
  -- +goose StatementBegin
  CREATE TABLE `products` (
    `product_id` int NOT NULL AUTO_INCREMENT,
    `product_name` varchar(100) NOT NULL,
    `stock` int NOT NULL,
    `price` float NOT NULL,
    `manufacturer` varchar(100) NOT NULL,
    `category_name` varchar(255) NOT NULL,
    PRIMARY KEY (`product_id`),
    UNIQUE KEY `product_id` (`product_id`),
    KEY `category_name` (`category_name`),
    CONSTRAINT `category_name_fk` FOREIGN KEY (`category_name`) REFERENCES `categories` (`category_name`)
  );
  -- +goose StatementEnd

  -- +goose Down
  -- +goose StatementBegin
  DROP TABLE IF EXISTS products;
  -- +goose StatementEnd
