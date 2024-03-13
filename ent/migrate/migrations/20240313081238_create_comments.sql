-- Create "comments" table
CREATE TABLE `comments` (`id` char(36) NOT NULL, `author_id` char(36) NOT NULL, `article_id` char(36) NOT NULL, `body` varchar(4096) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
