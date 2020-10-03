-- 00001_initial_migration.sql

-- +goose Up
USE XsollaGarage;

CREATE TABLE IF NOT EXISTS feedback (
    `id` int PRIMARY KEY AUTO_INCREMENT,
    `date_create` TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
    `content` TEXT
);

-- +goose Down
DROP TABLE feedback;