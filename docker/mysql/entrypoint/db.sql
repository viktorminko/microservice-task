CREATE TABLE `Client` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `Name` varchar(200) DEFAULT NULL,
  `Email` varchar(1024) DEFAULT NULL,
  `Mobile` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `ID_UNIQUE` (`ID`),
  UNIQUE KEY `EMAIL_UNIQUE` (`Email`)
);