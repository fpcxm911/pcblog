/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80016
 Source Host           : localhost:3306
 Source Schema         : pcblog

 Target Server Type    users: MySQL
 Target Server Version : 80016
 File Encoding         : 65001

 Date: 15/12/2021 18:22
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
SET GLOBAL sql_mode='STRICT_TRANS_TABLES';
-- ----------------------------
-- Table structure for blog_comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `comment_id` bigint NOT NULL AUTO_INCREMENT COMMENT '评论ID',
  `user_id` bigint NOT NULL COMMENT '发表用户ID',
  `article_id` bigint NOT NULL COMMENT '评论博文ID',
  `comment_like_count` bigint NOT NULL DEFAULT 0 COMMENT '点赞数',
  `comment_date` int(10) NOT NULL DEFAULT 0 COMMENT '评论日期',
  `comment_content` text NOT NULL COMMENT '评论内容',
  `parent_comment_id` bigint NOT NULL DEFAULT 0 COMMENT '父评论ID',
  PRIMARY KEY (`comment_id`) USING BTREE,
  INDEX `article_id`(`article_id`) USING BTREE,
  INDEX `comment_date`(`comment_date`) USING BTREE,
  INDEX `parent_comment_id`(`parent_comment_id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for user_follower
-- ----------------------------
DROP TABLE IF EXISTS `user_follower`;
CREATE TABLE `user_follower`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '标识ID',
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `user_follower` bigint NOT NULL COMMENT '关注者ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  CONSTRAINT `user_follower_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `user_follower_ibfk_2` FOREIGN KEY (`user_follower`) REFERENCES `users` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for user_subscriber
-- ----------------------------
DROP TABLE IF EXISTS `user_subscriber`;
CREATE TABLE `user_subscriber`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '标识ID',
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `user_subscriber` bigint NOT NULL COMMENT '关注的ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  CONSTRAINT `user_subscriber_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `user_subscriber_ibfk_2` FOREIGN KEY (`user_subscriber`) REFERENCES `users` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `user_id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户自增ID',
  `username` varchar(20) NOT NULL COMMENT '账号',
  `password` varchar(15) NOT NULL COMMENT '用户密码',
  `createtime` int(10) NOT NULL DEFAULT 0 COMMENT '注册时间',
  PRIMARY KEY (`user_id`) USING BTREE,
  INDEX `username`(`username`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for articles
-- ----------------------------
DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles`  (
  `article_id` bigint NOT NULL AUTO_INCREMENT COMMENT '文章ID',
  `user_id` bigint NOT NULL COMMENT '发表用户ID',
  `article_title` text NOT NULL COMMENT '博文标题',
  `article_content` longtext NOT NULL COMMENT '博文内容',
  `article_views` bigint NOT NULL DEFAULT 0 COMMENT '浏览量',
  `article_comment_count` bigint NOT NULL DEFAULT 0 COMMENT '评论总数',
  `article_date` int(10) NOT NULL DEFAULT 0 COMMENT '发表时间',
  `article_like_count` bigint NOT NULL DEFAULT 0 COMMENT '点赞数',
  PRIMARY KEY (`article_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  CONSTRAINT `articles_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for inbox
-- ----------------------------
DROP TABLE IF EXISTS `inbox`;
CREATE TABLE `inbox`  (
  `article_id` bigint NOT NULL AUTO_INCREMENT COMMENT '文章ID',
  `user_id` bigint NOT NULL COMMENT '发表用户ID',
  `article_title` text NOT NULL COMMENT '博文标题',
  `article_content` longtext NOT NULL COMMENT '博文内容',
  `article_views` bigint NOT NULL DEFAULT 0 COMMENT '浏览量',
  `article_comment_count` bigint NOT NULL DEFAULT 0 COMMENT '评论总数',
  `article_date` int(10) NOT NULL DEFAULT 0 COMMENT '发表时间',
  `article_like_count` bigint NOT NULL DEFAULT 0 COMMENT '点赞数',
  PRIMARY KEY (`article_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  CONSTRAINT `inbox_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `inbox_ibfk_2` FOREIGN KEY (`article_id`) REFERENCES `ariticles` (`article_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `test`;
CREATE TABLE `test` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `username` varchar(30) NOT NULL COMMENT '账号',
  `password` varchar(100) NOT NULL COMMENT '密码',
  `createtime` int(10) NOT NULL DEFAULT 0 COMMENT '创建时间',
   PRIMARY KEY (`id`)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
