-- +goose Up
-- +goose StatementBegin
CREATE TABLE cart_items (
  `cart_item_id` INT NOT NULL AUTO_INCREMENT,
  `cart_id` INT NOT NULL,
  `product_id` INT NOT NULL,
  `quantity` INT NOT NULL,
  PRIMARY KEY (cart_item_id),
  FOREIGN KEY (cart_id) REFERENCES carts(cart_id),
  FOREIGN KEY (product_id) REFERENCES products(product_id),
  UNIQUE (cart_id, product_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cart_items;
-- +goose StatementEnd
