CREATE TABLE `user` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(16) DEFAULT 'default',
    `desc` VARCHAR(258),
    `roler` VARCHAR(16) DEFAULT 'member',
    `phone` VARCHAR(16),
    `msg` VARCHAR(32),
    `sex` INT,
    PRIMARY KEY (`id`),
    INDEX `index_name` (`name`)
);

