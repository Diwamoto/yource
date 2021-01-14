-- Adminer 4.7.8 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ユーザid',
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'メールアドレス',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '名前',
  `phone` int(12) DEFAULT NULL COMMENT '電話番号',
  `status` tinyint(1) DEFAULT NULL COMMENT '0:無効 1:有効',
  `created` datetime DEFAULT NULL COMMENT '作成日時',
  `modified` datetime DEFAULT NULL COMMENT '更新日時',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


DROP TABLE IF EXISTS `user_profile`;
CREATE TABLE `user_profile` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ユーザプロフィールID',
  `user_id` int(11) NOT NULL COMMENT 'ユーザID',
  `profile` varchar(2000) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '自己紹介(html)',
  `birthday` varchar(2000) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '生年月日',
  `from` varchar(2000) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '出身地',
  `job` varchar(2000) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '職業',
  `twitter` varchar(2000) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'twitter url',
  `facebook` varchar(2000) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'facebook url',
  `instagram` varchar(2000) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'instagram url',
  `other` varchar(2000) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'other sns url',
  `created` datetime DEFAULT NULL COMMENT '作成日時',
  `modified` datetime DEFAULT NULL COMMENT '更新日時',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `user_profile_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- 2021-01-07 13:32:37
