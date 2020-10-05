-- 00001_initial_migration.sql

-- +goose Up
USE XsollaGarage;

ALTER TABLE `tasks` MODIFY COLUMN `user_id` varchar(50);

CREATE TABLE IF NOT EXISTS `users`
(
    `id`          VARCHAR(50)  NOT NULL PRIMARY KEY ,
    `email`       VARCHAR(255) NOT NULL,
    `name`        VARCHAR(510) NOT NULL,
    `given_name`  VARCHAR(510) NOT NULL,
    `family_name` VARCHAR(510) NOT NULL,
    `locale`      VARCHAR(10) NOT NULL,
    `picture`     VARCHAR(510) NOT NULL
);



ALTER TABLE `tasks` ADD CONSTRAINT `user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE;

-- +goose Down
ALTER TABLE `tasks` DROP CONSTRAINT `user_id_fk`;

DROP TABLE `users`;

ALTER TABLE `tasks` MODIFY COLUMN `user_id` INT;
