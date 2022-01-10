CREATE DATABASE IF NOT EXISTS test default charset utf8 COLLATE utf8_general_ci;

use test;

DROP TABLE IF EXISTS `person`;

CREATE TABLE `person`
(
    `user_id`  int(11) NOT NULL AUTO_INCREMENT,
    `username` varchar(260) DEFAULT NULL,
    `sex`      varchar(260) DEFAULT NULL,
    `email`    varchar(260) DEFAULT NULL,
    PRIMARY KEY (`user_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 6 DEFAULT CHARSET = utf8;

INSERT INTO `test`.`person` (`user_id`, `username`, `sex`, `email`) VALUES (5, 'test11', 'man', 'test1@qq.com');