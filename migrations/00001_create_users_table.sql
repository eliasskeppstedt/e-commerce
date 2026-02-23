-- +goose Up
-- +goose StatementBegin
CREATE TABLE `users` (
  `user_id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL UNIQUE KEY,
  `password` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL UNIQUE KEY,
  `first_name` varchar(25) DEFAULT NULL,
  `last_name` varchar(25) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `zip_code` varchar(6) DEFAULT NULL,
  `phone_number` varchar(25) DEFAULT NULL,
  PRIMARY KEY (`user_id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
