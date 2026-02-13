-- +goose Up
-- +goose StatementBegin
CREATE TABLE `orderItems` (
  `productID` int NOT NULL,
  `quantity` int NOT NULL,
  `orderID` int NOT NULL AUTO_INCREMENT,
  KEY `productID` (`productID`),
  KEY `orderID_idx` (`orderID`),
  CONSTRAINT `orderID` FOREIGN KEY (`orderID`) REFERENCES `orders` (`orderID`),
  CONSTRAINT `orderItems_ibfk_1` FOREIGN KEY (`productID`) REFERENCES `products` (`productID`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orderItems;
-- +goose StatementEnd
