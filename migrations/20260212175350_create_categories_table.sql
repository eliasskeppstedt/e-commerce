-- +goose Up
-- +goose StatementBegin
CREATE TABLE `categories` (
  `categoryID` int NOT NULL AUTO_INCREMENT,
  `categoryName` varchar(25) NOT NULL,
  UNIQUE KEY `categoryID` (`categoryID`),
  UNIQUE KEY `categoryName` (`categoryName`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS categories;
-- +goose StatementEnd
