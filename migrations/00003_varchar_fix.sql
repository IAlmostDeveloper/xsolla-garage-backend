-- 00001_initial_migration.sql

-- +goose Up
USE XsollaGarage;

ALTER TABLE `tasks` MODIFY `title` TEXT;
ALTER TABLE `tasks` MODIFY `text_content` TEXT;

-- +goose Down
ALTER TABLE `tasks` MODIFY `title` VARCHAR(255);
ALTER TABLE `tasks` MODIFY `text_content` VARCHAR(255);



