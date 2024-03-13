-- Create "user_favorites" table
CREATE TABLE `user_favorites` (`id` char(36) NOT NULL, `user_id` char(36) NOT NULL, `article_id` char(36) NOT NULL, `created_at` timestamp NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `userfavorite_user_id_article_id` (`user_id`, `article_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
