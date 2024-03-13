-- Create "articles" table
CREATE TABLE `articles` (`id` char(36) NOT NULL, `author_id` char(36) NOT NULL, `slug` varchar(255) NOT NULL, `title` varchar(255) NOT NULL, `description` varchar(255) NOT NULL, `body` varchar(4096) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `slug` (`slug`), UNIQUE INDEX `article_slug` (`slug`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Modify "user_follows" table
ALTER TABLE `user_follows` ADD UNIQUE INDEX `follower_id` (`follower_id`), ADD CONSTRAINT `user_follows_users_follows` FOREIGN KEY (`follower_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION, ADD CONSTRAINT `user_follows_users_followee` FOREIGN KEY (`followee_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION;
