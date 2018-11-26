CREATE TABLE `Client` (
  `id` bigint(20) NOT NULL,
  `Name` varchar(200) DEFAULT NULL,
  `Email` varchar(1024) DEFAULT NULL,
  `Mobile` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`)
);