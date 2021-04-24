CREATE TABLE IF NOT EXISTS `even`(
    `id` INT UNSIGNED AUTO_INCREMENT,
    `qq_num` VARCHAR(1000),
    `even_content_more_info` VARCHAR(1000),
    `even_content` VARCHAR(1000),
    `verify_status` INT,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;