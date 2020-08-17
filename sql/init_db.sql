DROP TABLE IF EXISTS `demos`;

CREATE TABLE `demos` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255),
  `description` VARCHAR(255),
  `create_time` DATETIME,
  PRIMARY KEY (`id`)
) character set = utf8;
