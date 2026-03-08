-- +goose Up
-- +goose StatementBegin
CREATE TABLE `orders` (
  `order_id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT NOT NULL,
`status` ENUM('Processing','In Transit/Transport','Ready for pickup') NOT NULL DEFAULT 'Processing',
  `date` DATE DEFAULT (CURDATE()),
  `subtotal_at_purchase` DECIMAL(10,2) NOT NULL,
  PRIMARY KEY (`order_id`),
  FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
