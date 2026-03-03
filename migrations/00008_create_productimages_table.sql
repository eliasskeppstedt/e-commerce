-- +goose Up
-- +goose StatementBegin
CREATE TABLE `product_images` (
  `product_image_id` int NOT NULL AUTO_INCREMENT,
  `product_id` int NOT NULL,
  `product_image_uri` float NOT NULL,
  PRIMARY KEY (`product_image_id`),
  FOREIGN KEY (`product_id`) REFERENCES `products` (`product_id`) 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS product_images;
-- +goose StatementEnd