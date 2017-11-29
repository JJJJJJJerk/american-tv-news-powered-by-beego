# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: www.cnfluffy.com (MySQL 5.5.51-MariaDB)
# Database: trytv
# Generation Time: 2017-11-29 17:34:00 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table article_episode
# ------------------------------------------------------------

CREATE TABLE `article_episode` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(10) unsigned NOT NULL,
  `episode_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `article_episode_article_id_index` (`article_id`),
  KEY `article_episode_episode_id_index` (`episode_id`),
  CONSTRAINT `article_episode_article_id_foreign` FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`) ON DELETE CASCADE,
  CONSTRAINT `article_episode_episode_id_foreign` FOREIGN KEY (`episode_id`) REFERENCES `episodes` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table article_movie
# ------------------------------------------------------------

CREATE TABLE `article_movie` (
  `article_id` int(10) unsigned NOT NULL,
  `movie_id` int(10) unsigned NOT NULL,
  KEY `article_movie_article_id_index` (`article_id`),
  KEY `article_movie_movie_id_index` (`movie_id`),
  CONSTRAINT `article_movie_article_id_foreign` FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`) ON DELETE CASCADE,
  CONSTRAINT `article_movie_movie_id_foreign` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table article_show
# ------------------------------------------------------------

CREATE TABLE `article_show` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `show_id` int(10) unsigned NOT NULL,
  `article_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `article_show_show_id_index` (`show_id`),
  KEY `article_show_article_id_index` (`article_id`),
  CONSTRAINT `article_show_article_id_foreign` FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`) ON DELETE CASCADE,
  CONSTRAINT `article_show_show_id_foreign` FOREIGN KEY (`show_id`) REFERENCES `shows` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table article_tag
# ------------------------------------------------------------

CREATE TABLE `article_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `article_id` int(10) unsigned NOT NULL,
  `tag_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `article_tag_article_id_index` (`article_id`),
  KEY `article_tag_tag_id_index` (`tag_id`),
  CONSTRAINT `article_tag_article_id_foreign` FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`) ON DELETE CASCADE,
  CONSTRAINT `article_tag_tag_id_foreign` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table articles
# ------------------------------------------------------------

CREATE TABLE `articles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `body` longtext COLLATE utf8_unicode_ci NOT NULL,
  `url_provider` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `key_word` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `description` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `raw_title` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '爬去的标题',
  `raw_content` longtext COLLATE utf8_unicode_ci COMMENT '爬去的内容',
  `video_title` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT '视频标题',
  `url_video` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '视频地址',
  `coverage_uri` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '封面图片的uri/key',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `video_code` text COLLATE utf8_unicode_ci COMMENT '视频播放代码(准备放弃)',
  `is_show` tinyint(1) unsigned NOT NULL COMMENT '废弃和deleted_at重复',
  PRIMARY KEY (`id`),
  KEY `url_title_unique` (`raw_title`,`url_provider`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table episodes
# ------------------------------------------------------------

CREATE TABLE `episodes` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '重新计算的结果',
  `raw_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `provider` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `url_provider` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `season` smallint(6) NOT NULL,
  `episode` smallint(6) NOT NULL,
  `show_id` int(10) unsigned NOT NULL,
  `size` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `quality` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `is_auth` tinyint(1) NOT NULL,
  `url_torrent` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `url_magnet` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `url_ed2k` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `url_baidupan` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `url_other` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `url_video` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `like_count` mediumint(8) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `movie_id` int(10) unsigned NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `episodes_raw_name_unique` (`raw_name`),
  KEY `episodes_show_id_foreign` (`show_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table ht_user
# ------------------------------------------------------------

CREATE TABLE `ht_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `phone` varchar(12) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `email_index` (`email`) USING BTREE,
  KEY `phone_index` (`phone`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;



# Dump of table image_movie
# ------------------------------------------------------------

CREATE TABLE `image_movie` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `image_id` int(10) unsigned NOT NULL,
  `movie_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `image_movie_image_id_index` (`image_id`),
  KEY `image_movie_movie_id_index` (`movie_id`),
  CONSTRAINT `image_movie_image_id_foreign` FOREIGN KEY (`image_id`) REFERENCES `images` (`id`),
  CONSTRAINT `image_movie_movie_id_foreign` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table image_show
# ------------------------------------------------------------

CREATE TABLE `image_show` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `show_id` int(10) unsigned NOT NULL,
  `image_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `image_show_show_id_index` (`show_id`),
  KEY `image_show_image_id_index` (`image_id`),
  CONSTRAINT `image_show_image_id_foreign` FOREIGN KEY (`image_id`) REFERENCES `images` (`id`),
  CONSTRAINT `image_show_show_id_foreign` FOREIGN KEY (`show_id`) REFERENCES `shows` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table images
# ------------------------------------------------------------

CREATE TABLE `images` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `key` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `article_id` int(10) unsigned NOT NULL,
  `description` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `bucket` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `fname` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `fsize` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `width` int(10) unsigned NOT NULL,
  `height` int(10) unsigned NOT NULL,
  `format` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `imgur_id` int(10) unsigned DEFAULT NULL COMMENT '管理imgur模型',
  PRIMARY KEY (`id`),
  UNIQUE KEY `images_key_unique` (`key`),
  KEY `imgur_id` (`imgur_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table imgurs
# ------------------------------------------------------------

CREATE TABLE `imgurs` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `imgur_id` varchar(11) COLLATE utf8_unicode_ci DEFAULT NULL,
  `title` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `title_translation` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `keywords` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `description_translation` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `is_published` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table jobs
# ------------------------------------------------------------

CREATE TABLE `jobs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `queue` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `payload` longtext COLLATE utf8_unicode_ci NOT NULL,
  `attempts` tinyint(3) unsigned NOT NULL,
  `reserved` tinyint(3) unsigned NOT NULL,
  `reserved_at` int(10) unsigned DEFAULT NULL,
  `available_at` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `jobs_queue_reserved_reserved_at_index` (`queue`,`reserved`,`reserved_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table migrations
# ------------------------------------------------------------

CREATE TABLE `migrations` (
  `migration` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `batch` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table movies
# ------------------------------------------------------------

CREATE TABLE `movies` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `name_en` text COLLATE utf8_unicode_ci NOT NULL,
  `name_zh` text COLLATE utf8_unicode_ci NOT NULL,
  `playwright` text COLLATE utf8_unicode_ci NOT NULL,
  `publish_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `grade` text COLLATE utf8_unicode_ci NOT NULL,
  `length` text COLLATE utf8_unicode_ci NOT NULL,
  `official_url` text COLLATE utf8_unicode_ci NOT NULL,
  `detail` longtext COLLATE utf8_unicode_ci NOT NULL,
  `key_word` text COLLATE utf8_unicode_ci NOT NULL,
  `description` text COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table password_resets
# ------------------------------------------------------------

CREATE TABLE `password_resets` (
  `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `token` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  KEY `password_resets_email_index` (`email`),
  KEY `password_resets_token_index` (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table quotes
# ------------------------------------------------------------

CREATE TABLE `quotes` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `english` mediumtext COLLATE utf8_unicode_ci NOT NULL,
  `chinese` mediumtext COLLATE utf8_unicode_ci NOT NULL,
  `writer` text COLLATE utf8_unicode_ci NOT NULL,
  `image_uri` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table sessions
# ------------------------------------------------------------

CREATE TABLE `sessions` (
  `id` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `ip_address` varchar(45) COLLATE utf8_unicode_ci DEFAULT NULL,
  `user_agent` text COLLATE utf8_unicode_ci,
  `payload` text COLLATE utf8_unicode_ci NOT NULL,
  `last_activity` int(11) NOT NULL,
  UNIQUE KEY `sessions_id_unique` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table shows
# ------------------------------------------------------------

CREATE TABLE `shows` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name_en` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `name_zh` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `detail` longtext COLLATE utf8_unicode_ci NOT NULL,
  `key_word` text COLLATE utf8_unicode_ci NOT NULL,
  `description` text COLLATE utf8_unicode_ci NOT NULL,
  `type` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `director` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `play_writer` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `actor` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `alias` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '美剧其他中文名称,逗号分隔',
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table subtitles
# ------------------------------------------------------------

CREATE TABLE `subtitles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `source_url` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '抓取地址',
  `name_zh` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '中文名称',
  `name_en` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '英文名称',
  `version` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '视频格式版本',
  `file_name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '字幕文件名',
  `lang` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '语言',
  `format` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '字幕格式',
  `url` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '字幕原来地址',
  `uri` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'oss地址',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `source_url` (`source_url`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table tags
# ------------------------------------------------------------

CREATE TABLE `tags` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `name_en` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `image_id` int(11) NOT NULL,
  `key_word` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `description` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `deleted_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table users
# ------------------------------------------------------------

CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `password` varchar(60) COLLATE utf8_unicode_ci NOT NULL,
  `honor_point` mediumint(9) NOT NULL,
  `nick_name` char(20) COLLATE utf8_unicode_ci NOT NULL,
  `slogan` char(40) COLLATE utf8_unicode_ci NOT NULL,
  `avatar_image` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `remember_token` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `github_id` int(11) NOT NULL,
  `github_token` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `github_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `github_nickname` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `github_email` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `github_avatar` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `github_json` text COLLATE utf8_unicode_ci NOT NULL,
  `weibo_id` int(11) unsigned NOT NULL,
  `weibo_token` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `weibo_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `weibo_nickname` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `weibo_email` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `weibo_avatar` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `weibo_json` text COLLATE utf8_unicode_ci NOT NULL,
  `qq_id` int(11) NOT NULL,
  `qq_token` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `qq_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `qq_nickname` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `qq_email` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `qq_avatar` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `qq_json` text COLLATE utf8_unicode_ci NOT NULL,
  `weixin_id` int(11) NOT NULL,
  `weixin_token` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `weixin_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `weixin_nickname` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `weixin_email` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `weixin_avatar` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `weixin_json` text COLLATE utf8_unicode_ci NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_email_unique` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table votes
# ------------------------------------------------------------

CREATE TABLE `votes` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(10) unsigned DEFAULT NULL,
  `movie_id` int(10) unsigned DEFAULT NULL,
  `show_id` int(10) unsigned DEFAULT NULL,
  `visit` int(10) unsigned NOT NULL,
  `favorate_count` int(10) unsigned DEFAULT NULL,
  `score` decimal(4,2) NOT NULL,
  `vote_count` int(10) unsigned NOT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `fake_read` int(10) unsigned DEFAULT NULL COMMENT '假浏览量',
  PRIMARY KEY (`id`),
  KEY `votes_article_id_index` (`article_id`),
  KEY `votes_movie_id_index` (`movie_id`),
  KEY `votes_show_id_index` (`show_id`),
  CONSTRAINT `votes_article_id_foreign` FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`) ON DELETE CASCADE,
  CONSTRAINT `votes_movie_id_foreign` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`) ON DELETE CASCADE,
  CONSTRAINT `votes_show_id_foreign` FOREIGN KEY (`show_id`) REFERENCES `shows` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



# Dump of table weibos
# ------------------------------------------------------------

CREATE TABLE `weibos` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `wb_id` bigint(20) unsigned NOT NULL,
  `body` text COLLATE utf8_unicode_ci NOT NULL,
  `images` text COLLATE utf8_unicode_ci NOT NULL,
  `type` enum('favorate','myself','friend','trash') COLLATE utf8_unicode_ci NOT NULL,
  `is_shown` tinyint(3) unsigned NOT NULL,
  `weibor` char(255) COLLATE utf8_unicode_ci NOT NULL,
  `avatar` char(255) COLLATE utf8_unicode_ci NOT NULL,
  `media_url` char(255) COLLATE utf8_unicode_ci NOT NULL,
  `long_url` char(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `media_type` enum('miaopai','youku','bili','acfun','text','miaopai','xiakaxiu','vlook') COLLATE utf8_unicode_ci DEFAULT 'text',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `weibos_wb_id_unique` (`wb_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
