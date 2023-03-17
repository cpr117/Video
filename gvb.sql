/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80024
 Source Host           : localhost:3306
 Source Schema         : mybilibili

 Target Server Type    : MySQL
 Target Server Version : 80024
 File Encoding         : 65001

 Date: 15/03/2023 14:41:40
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for anime
-- ----------------------------
DROP TABLE IF EXISTS `anime`;
CREATE TABLE `anime`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '\'动漫标题\'',
  `img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '\'动漫封面\'',
  `category_ids` bigint NOT NULL COMMENT '分类 ID',
  `desc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '描述',
  `avatar` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '封面图片地址',
  `year` bigint NULL DEFAULT NULL COMMENT '年份',
  `season` bigint NULL DEFAULT NULL COMMENT '季度',
  `episode` bigint NULL DEFAULT NULL COMMENT '集数',
  `is_delete` tinyint NULL DEFAULT 0 COMMENT '是否删除',
  `score` decimal(3, 1) NULL DEFAULT NULL COMMENT '评分',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `need_member` tinyint NULL DEFAULT 0 COMMENT '是否需要会员',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '番剧链接',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of anime
-- ----------------------------

-- ----------------------------
-- Table structure for anime_category
-- ----------------------------
DROP TABLE IF EXISTS `anime_category`;
CREATE TABLE `anime_category`  (
  `anime_id` bigint NULL DEFAULT NULL,
  `category_id` bigint NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of anime_category
-- ----------------------------

-- ----------------------------
-- Table structure for bili_config
-- ----------------------------
DROP TABLE IF EXISTS `bili_config`;
CREATE TABLE `bili_config`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `config` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '配置',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of bili_config
-- ----------------------------
INSERT INTO `bili_config` VALUES (1, '2023-02-21 11:36:33.048', '2023-02-21 11:36:33.048', '{\"website_avatar\":\"\",\"website_name\":\"VideoWeb\",\"website_intro\":\"这是一个视频网站 试图模仿B站\",\"website_author\":\"cpr\",\"website_notice\":\"\",\"website_create_time\":\"2023-02-21\",\"website_record\":\"京ICP备00000000号\",\"social_login_list\":[\"QQ\",\"Github\",\"Gitee\"],\"social_url_list\":[\"https://www.qq.com\",\"https://www.github.com\",\"https://www.gitee.com\"],\"qq\":\"\",\"github\":\"\",\"gitee\":\"\",\"tourist_avatar\":\"https://www.baidu.com/img/bd_logo1.png\",\"user_avatar\":\"https://www.baidu.com/img/bd_logo1.png\",\"article_cover\":\"https://www.baidu.com/img/bd_logo1.png\",\"is_comment_review\":0,\"is_message_review\":0,\"is_email_notice\":0,\"is_reward\":0,\"wechat_qrcode\":\"https://www.baidu.com/img/bd_logo1.png\",\"alipay_ode\":\"https://www.baidu.com/img/bd_logo1.png\"}');
INSERT INTO `bili_config` VALUES (2, '2023-02-21 15:35:39.057', '2023-02-21 15:35:39.057', '{\"website_avatar\":\"\",\"website_name\":\"VideoWeb\",\"website_intro\":\"这是一个视频网站 试图模仿B站\",\"website_author\":\"cpr\",\"website_notice\":\"\",\"website_create_time\":\"2023-02-21\",\"website_record\":\"京ICP备00000000号\",\"social_login_list\":[\"QQ\",\"Github\",\"Gitee\"],\"social_url_list\":[\"https://www.qq.com\",\"https://www.github.com\",\"https://www.gitee.com\"],\"qq\":\"\",\"github\":\"\",\"gitee\":\"\",\"tourist_avatar\":\"https://www.baidu.com/img/bd_logo1.png\",\"user_avatar\":\"https://www.baidu.com/img/bd_logo1.png\",\"article_cover\":\"https://www.baidu.com/img/bd_logo1.png\",\"is_comment_review\":0,\"is_message_review\":0,\"is_email_notice\":0,\"is_reward\":0,\"wechat_qrcode\":\"https://www.baidu.com/img/bd_logo1.png\",\"alipay_ode\":\"https://www.baidu.com/img/bd_logo1.png\"}');
INSERT INTO `bili_config` VALUES (3, '2023-02-21 15:39:19.385', '2023-02-21 15:39:19.385', '{\"website_avatar\":\"\",\"website_name\":\"VideoWeb\",\"website_intro\":\"这是一个视频网站 试图模仿B站\",\"website_author\":\"cpr\",\"website_notice\":\"\",\"website_create_time\":\"2023-02-21\",\"website_record\":\"京ICP备00000000号\",\"social_login_list\":[\"QQ\",\"Github\",\"Gitee\"],\"social_url_list\":[\"https://www.qq.com\",\"https://www.github.com\",\"https://www.gitee.com\"],\"qq\":\"\",\"github\":\"\",\"gitee\":\"\",\"tourist_avatar\":\"https://www.baidu.com/img/bd_logo1.png\",\"user_avatar\":\"https://www.baidu.com/img/bd_logo1.png\",\"article_cover\":\"https://www.baidu.com/img/bd_logo1.png\",\"is_comment_review\":0,\"is_message_review\":0,\"is_email_notice\":0,\"is_reward\":0,\"wechat_qrcode\":\"https://www.baidu.com/img/bd_logo1.png\",\"alipay_ode\":\"https://www.baidu.com/img/bd_logo1.png\"}');
INSERT INTO `bili_config` VALUES (4, '2023-02-21 15:41:42.875', '2023-02-21 15:41:42.875', '{\"website_avatar\":\"\",\"website_name\":\"VideoWeb\",\"website_intro\":\"这是一个视频网站 试图模仿B站\",\"website_author\":\"cpr\",\"website_notice\":\"\",\"website_create_time\":\"2023-02-21\",\"website_record\":\"京ICP备00000000号\",\"social_login_list\":[\"QQ\",\"Github\",\"Gitee\"],\"social_url_list\":[\"https://www.qq.com\",\"https://www.github.com\",\"https://www.gitee.com\"],\"qq\":\"\",\"github\":\"\",\"gitee\":\"\",\"tourist_avatar\":\"https://www.baidu.com/img/bd_logo1.png\",\"user_avatar\":\"https://www.baidu.com/img/bd_logo1.png\",\"article_cover\":\"https://www.baidu.com/img/bd_logo1.png\",\"is_comment_review\":0,\"is_message_review\":0,\"is_email_notice\":0,\"is_reward\":0,\"wechat_qrcode\":\"https://www.baidu.com/img/bd_logo1.png\",\"alipay_ode\":\"https://www.baidu.com/img/bd_logo1.png\"}');

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (12, 'p', 'admin', '/v1/api/admin/create', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (11, 'p', 'admin', '/v1/api/admin/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (13, 'p', 'admin', '/v1/api/admin/update', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (16, 'p', 'admin', '/v1/api/comment/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (14, 'p', 'admin', '/v1/api/partition/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (15, 'p', 'admin', '/v1/api/partition/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (9, 'p', 'admin', '/v1/api/user/delete', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (8, 'p', 'admin', '/v1/api/user/disable', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (7, 'p', 'admin', '/v1/api/user/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (10, 'p', 'admin', '/v1/api/user/role', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (6, 'p', 'admin', '/v1/api/video/delete', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (4, 'p', 'admin', '/v1/api/video/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (5, 'p', 'admin', '/v1/api/video/update', 'POST', '', '', '');

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '分类名称',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of category
-- ----------------------------

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint NULL DEFAULT NULL COMMENT '评论用户id',
  `reply_comment_id` bigint NULL DEFAULT NULL COMMENT '回复评论id',
  `reply_root_id` bigint NULL DEFAULT NULL COMMENT '回复根评论id',
  `video_id` bigint NULL DEFAULT NULL COMMENT '评论主题id',
  `video_type` bigint NULL DEFAULT NULL COMMENT '评论主题类型',
  `content` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '评论内容',
  `is_delete` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除(0.否 1.是)',
  `is_review` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否审核(0.否 1.是)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comment
-- ----------------------------

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `msg_type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '消息类型(0-系统消息,1-用户私信)',
  `send_id` bigint NULL DEFAULT NULL COMMENT '发送者id',
  `receive_id` bigint NULL DEFAULT NULL COMMENT '接收者id',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '留言内容',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of message
-- ----------------------------

-- ----------------------------
-- Table structure for operation_log
-- ----------------------------
DROP TABLE IF EXISTS `operation_log`;
CREATE TABLE `operation_log`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `opt_module` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '操作模块',
  `opt_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '操作类型',
  `opt_method` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '操作方法',
  `opt_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '操作URL',
  `opt_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '操作描述',
  `request_param` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '请求参数',
  `request_method` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '请求方法',
  `response_data` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '响应数据',
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户ID',
  `nickname` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '用户昵称',
  `ip_address` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '操作IP',
  `ip_source` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '操作地址',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of operation_log
-- ----------------------------
INSERT INTO `operation_log` VALUES (1, '2023-03-15 14:16:19.189', '2023-03-15 14:16:19.189', '', '新增或修改', 'VideoWeb/api/back.(*Partition).AddPartition-fm', '/v1/api/partition/add', '新增或修改', '{\r\n    \"name\": \"生活区\"\r\n}', 'POST', '{\"code\":0,\"message\":\"OK\",\"data\":0}', 2, 'user:admin1', '172.29.176.1', '内网IP');
INSERT INTO `operation_log` VALUES (2, '2023-03-15 14:21:16.715', '2023-03-15 14:21:16.715', '', '新增或修改', 'VideoWeb/api/back.(*Partition).AddPartition-fm', '/v1/api/partition/add', '新增或修改', '{\r\n    \"name\": \"鬼畜区\"\r\n}', 'POST', '{\"code\":0,\"message\":\"OK\",\"data\":0}', 2, 'user:admin1', '172.29.176.1', '内网IP');

-- ----------------------------
-- Table structure for partition
-- ----------------------------
DROP TABLE IF EXISTS `partition`;
CREATE TABLE `partition`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of partition
-- ----------------------------
INSERT INTO `partition` VALUES (1, '2023-03-15 14:16:19.180', '2023-03-15 14:16:19.180', '生活区');
INSERT INTO `partition` VALUES (2, '2023-03-15 14:21:16.710', '2023-03-15 14:21:16.710', '鬼畜区');

-- ----------------------------
-- Table structure for resource
-- ----------------------------
DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '资源路径(接口URL)',
  `request_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '请求方式',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '资源名(接口名)',
  `parent_id` bigint NULL DEFAULT NULL COMMENT '父权限id',
  `is_anonymous` tinyint(1) NULL DEFAULT NULL COMMENT '是否匿名访问(0-否 1-是)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of resource
-- ----------------------------

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '角色名',
  `label` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '角色标签',
  `is_disable` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否禁用(0-否 1-是)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (1, NULL, NULL, 'admin', 'admin', 0);
INSERT INTO `role` VALUES (2, NULL, NULL, 'test', 'test', 0);

-- ----------------------------
-- Table structure for submission
-- ----------------------------
DROP TABLE IF EXISTS `submission`;
CREATE TABLE `submission`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `sub_type` tinyint(1) NOT NULL COMMENT '提交类型',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of submission
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `coins` bigint NULL DEFAULT NULL COMMENT '硬币数',
  `is_member` tinyint(1) NULL DEFAULT NULL COMMENT '是否是会员',
  `followers` json NULL COMMENT '粉丝id列表',
  `favor_videos` json NULL COMMENT '收藏视频id列表',
  `fans_ids` json NULL COMMENT '关注的用户id列表',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '头像',
  `nick_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '昵称',
  `personal_sign` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '个性签名',
  `is_disable` tinyint(1) NULL DEFAULT NULL COMMENT '是否禁用(0-否 1-是)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '2023-02-21 15:54:49.646', '2023-02-21 15:54:49.646', 0, 0, NULL, NULL, NULL, 'https://www.baidu.com/img/bd_logo1.png', 'user:cpr1117', '', 0);
INSERT INTO `user` VALUES (2, '2023-03-09 21:50:02.589', '2023-03-09 21:50:02.589', 0, 0, NULL, NULL, NULL, 'https://www.baidu.com/img/bd_logo1.png', 'user:admin1', '', 0);

-- ----------------------------
-- Table structure for user_auth
-- ----------------------------
DROP TABLE IF EXISTS `user_auth`;
CREATE TABLE `user_auth`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `user_info_id` bigint NULL DEFAULT NULL COMMENT '用户信息ID',
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户名',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '密码',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '邮箱',
  `login_type` tinyint(1) NULL DEFAULT NULL COMMENT '登录类型',
  `ip_address` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '登录IP地址',
  `ip_source` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'IP来源',
  `last_login_time` datetime(3) NULL DEFAULT NULL COMMENT '上次登录时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_auth
-- ----------------------------
INSERT INTO `user_auth` VALUES (1, '2023-02-21 15:54:49.706', '2023-03-09 20:05:14.733', 1, 'cpr1117', '$2a$10$t0fZaRY6tJoQa3NLbgJSLOg6ANsk8FGNPcWQ1RNzCC5t4yj9siypy', '498047949@qq.com', 1, '192.168.127.1', '内网IP', '2023-03-09 20:05:14.732');
INSERT INTO `user_auth` VALUES (2, '2023-03-09 21:50:02.649', '2023-03-15 14:11:25.820', 2, 'admin1', '$2a$10$aQyDeZcwXb0b1X41WCVOj.j16GF34NCPPp5kNYaK8NoOyVaYNBj9G', '12312312312@qq.com', 1, '172.29.176.1', '内网IP', '2023-03-15 14:11:25.819');

-- ----------------------------
-- Table structure for user_fans
-- ----------------------------
DROP TABLE IF EXISTS `user_fans`;
CREATE TABLE `user_fans`  (
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户id',
  `fan_id` bigint NULL DEFAULT NULL COMMENT '粉丝id'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_fans
-- ----------------------------

-- ----------------------------
-- Table structure for user_favor
-- ----------------------------
DROP TABLE IF EXISTS `user_favor`;
CREATE TABLE `user_favor`  (
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户id',
  `video_id` bigint NULL DEFAULT NULL COMMENT '视频id'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_favor
-- ----------------------------

-- ----------------------------
-- Table structure for user_follow
-- ----------------------------
DROP TABLE IF EXISTS `user_follow`;
CREATE TABLE `user_follow`  (
  `user_id` bigint NULL DEFAULT NULL COMMENT '用户id',
  `follow_id` bigint NULL DEFAULT NULL COMMENT '关注的用户id'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_follow
-- ----------------------------

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role`  (
  `user_id` bigint NULL DEFAULT NULL,
  `role_id` bigint NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_role
-- ----------------------------
INSERT INTO `user_role` VALUES (1, 3);
INSERT INTO `user_role` VALUES (2, 1);

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '视频名',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '视频地址',
  `cover` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '视频封面',
  `desc` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '视频描述',
  `partition_id` bigint NOT NULL COMMENT '视频分类id',
  `status` tinyint(1) NOT NULL COMMENT '视频状态',
  `author_id` bigint NULL DEFAULT NULL COMMENT '视频作者id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of video
-- ----------------------------
INSERT INTO `video` VALUES (1, '2023-03-09 20:31:48.850', '2023-03-09 20:31:48.850', '123321', 'rqgk41m3z.hd-bkt.clouddn.com/1678365108ac3fab9d27adda60be76248b8d98fa1b.mkv', 'req.VideoCover', 'descrption', 2, 0, 2);
INSERT INTO `video` VALUES (2, '2023-03-09 20:33:19.079', '2023-03-09 20:33:19.079', 'cpr', 'rqgk41m3z.hd-bkt.clouddn.com/16783651974468176623c4cba6a32d79e4432ec2b8.mp4', 'req.VideoCover', 'geebar', 1, 0, 1);

SET FOREIGN_KEY_CHECKS = 1;
