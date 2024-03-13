-- Create "tags" table
CREATE TABLE `tags` (`id` char(36) NOT NULL, `description` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `description` (`description`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
