-- +goose Up
-- +goose StatementBegin
CREATE TABLE `reviews` (
  `comment_id` int NOT NULL AUTO_INCREMENT,
  `product_id` int NOT NULL,
  `user_id` int NOT NULL,
  `comment_text` varchar(500) NOT NULL,
  `created_at` DATE NOT NULL DEFAULT (CURDATE()),
  `grade` int NOT NULL,
  CONSTRAINT grade CHECK (`grade` >= 1 AND `grade` <= 5),
  PRIMARY KEY (`comment_id`),
  FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`),
  FOREIGN KEY (`product_id`) REFERENCES `products` (`product_id`)
  ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS reviews;
-- +goose StatementEnd
