-- Create "article_tags" table
CREATE TABLE `article_tags` (`id` char(36) NOT NULL, `article_id` char(36) NOT NULL, `tag_id` char(36) NOT NULL, `created_at` timestamp NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `articletag_article_id_tag_id` (`article_id`, `tag_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
