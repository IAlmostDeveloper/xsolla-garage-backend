-- 00001_initial_migration.sql

-- +goose Up
USE XsollaGarage;


CREATE TABLE IF NOT EXISTS `tasks` (
                         `id` int PRIMARY KEY AUTO_INCREMENT,
                         `user_id` INT,
                         `title` VARCHAR(255),
                         `text_content` VARCHAR(255),
                         `date_create` TIMESTAMP,
                         `date_close` TIMESTAMP,
                         `date_target` TIMESTAMP,
                         `is_completed` BOOL
);


-- +goose Down
DROP TABLE `tasks`;
