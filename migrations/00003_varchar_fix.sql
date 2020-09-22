-- 00001_initial_migration.sql

-- +goose Up
USE XsollaGarage;

ALTER TABLE `tasks` MODIFY `title` TEXT;

-- +goose Down
ALTER TABLE `tasks` MODIFY `title` VARCHAR(255);



