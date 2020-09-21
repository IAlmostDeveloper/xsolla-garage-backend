-- 00001_initial_migration.sql

-- +goose Up
USE XsollaGarage;

CREATE TABLE `tags` (
    `id` int PRIMARY KEY AUTO_INCREMENT,
    `name` varchar(255) UNIQUE NOT NULL
);

CREATE TABLE `task_tags` (
     `id` int PRIMARY KEY AUTO_INCREMENT,
     `task_id` int NOT NULL,
     `tag_id` int NOT NULL
);

ALTER TABLE `task_tags` ADD FOREIGN KEY (`task_id`) REFERENCES `tasks` (`id`);

ALTER TABLE `task_tags` ADD FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`);

ALTER TABLE `task_tags` ADD UNIQUE INDEX `ix_task_tag_unique` (`task_id`, tag_id);

-- +goose Down
DROP TABLE `task_tags`;
DROP TABLE `tags`;

