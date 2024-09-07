
CREATE DATABASE `testdb`  DEFAULT CHARACTER SET utf8mb4;
USE `testdb`;

START TRANSACTION;

CREATE TABLE `user` (
  `itemId` bigint(20)NOT NULL AUTO_INCREMENT,
  `title` varchar(64) NOT NULL,
  `description` text NOT NULL,
  PRIMARY KEY (`itemId`),
  KEY `title` (`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `user` (`itemId`, `title`, `description`) VALUES (NULL, 'Foo D', 'Foo Description');

COMMIT;