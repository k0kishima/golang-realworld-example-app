-- Create "user_follows" table
CREATE TABLE `user_follows` (`id` char(36) NOT NULL, `follower_id` char(36) NOT NULL, `followee_id` char(36) NOT NULL, `created_at` timestamp NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `userfollow_follower_id_followee_id` (`follower_id`, `followee_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
