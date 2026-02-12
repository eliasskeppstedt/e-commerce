-- +goose Up
-- +goose StatementBegin
CREATE TABLE `products` (
  `productID` int NOT NULL AUTO_INCREMENT,
  `productName` varchar(100) NOT NULL,
  `stock` int NOT NULL,
  `price` float NOT NULL,
  `manufacturer` varchar(100) NOT NULL,
  `categoryID` int NOT NULL,
  PRIMARY KEY (`productID`),
  UNIQUE KEY `productID` (`productID`),
  KEY `CategoryID_idx` (`categoryID`),
  CONSTRAINT `CategoryID` FOREIGN KEY (`categoryID`) REFERENCES `categories` (`categoryID`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
