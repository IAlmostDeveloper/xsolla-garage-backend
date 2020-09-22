-- 00001_initial_migration.sql

-- +goose Up
USE XsollaGarage;

ALTER TABLE `task_tags` DROP FOREIGN KEY task_tags_ibfk_1;
ALTER TABLE `task_tags` ADD CONSTRAINT `task_id_fk` FOREIGN KEY (`task_id`) REFERENCES `tasks` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE;


-- +goose Down
ALTER TABLE `task_tags` DROP CONSTRAINT `task_id_fk`;
ALTER TABLE `task_tags` ADD CONSTRAINT `task_tags_ibfk_1` FOREIGN KEY (`task_id`) REFERENCES `tasks` (`id`);



