-- +goose Up
-- +goose StatementBegin
CREATE TABLE `orders` (
  `orderID` int NOT NULL AUTO_INCREMENT,
  `totalPrice` float NOT NULL,
  `orderStatus` enum('Proccesing','In Transit/Transport','Ready for pickup') NOT NULL,
  `userID` int DEFAULT NULL,
  PRIMARY KEY (`orderID`),
  UNIQUE KEY `orderID` (`orderID`),
  KEY `userID` (`userID`),
  CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`userID`) REFERENCES `users` (`userID`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
