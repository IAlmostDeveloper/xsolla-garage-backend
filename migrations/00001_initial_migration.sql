-- 00001_initial_migration.sql

-- +goose Up
USE XsollaGarage;
CREATE TABLE IF NOT EXISTS `users` (
                         `id` int PRIMARY KEY auto_increment,
                         `name` varchar(255),
                         `email` varchar(255),
                         `country` varchar(255),
                         `city` varchar(255),
                         `gender` varchar(255),
                         `age` int
);

CREATE TABLE IF NOT EXISTS `tasks` (
                         `id` int PRIMARY KEY auto_increment,
                         `user_id` int,
                         `title` varchar(255),
                         `category` varchar(255),
                         `text_content` varchar(255),
                         `date_create` timestamp,
                         `date_close` timestamp,
                         `date_target` timestamp,
                         `is_completed` bool
);

CREATE TABLE IF NOT EXISTS `subtasks` (
                            `id` int PRIMARY KEY auto_increment,
                            `task_id` int,
                            `depth` int,
                            `index` int,
                            `title` varchar(255),
                            `text_content` varchar(255),
                            `date_create` timestamp,
                            `date_close` timestamp,
                            `date_target` timestamp,
                            `is_completed` bool
);

CREATE TABLE IF NOT EXISTS `notifications` (
                                 `id` int PRIMARY KEY auto_increment,
                                 `task_id` int,
                                 `text_content` varchar(255),
                                 `notify_time` timestamp,
                                 `like` int,
                                 `url` varchar(255),
                                 `partner` varchar(255)
);

CREATE TABLE IF NOT EXISTS `labels` (
                          `id` int PRIMARY KEY auto_increment,
                          `task_id` int,
                          `name` varchar(255)
);

CREATE TABLE IF NOT EXISTS `task_media` (
                              `id` int PRIMARY KEY auto_increment,
                              `task_id` int,
                              `type` varchar(255),
                              `media` blob
);

ALTER TABLE `subtasks` ADD FOREIGN KEY (`task_id`) REFERENCES `tasks` (`id`);

ALTER TABLE `notifications` ADD FOREIGN KEY (`task_id`) REFERENCES `tasks` (`id`);

ALTER TABLE `labels` ADD FOREIGN KEY (`task_id`) REFERENCES `tasks` (`id`);

ALTER TABLE `task_media` ADD FOREIGN KEY (`task_id`) REFERENCES `tasks` (`id`);

-- +goose Down
drop table notifications;
drop table task_media;
drop table labels;
drop table subtasks;
drop table tasks;
drop table users;

