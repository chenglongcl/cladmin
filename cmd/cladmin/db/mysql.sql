/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50721 (5.7.21-log)
 Source Host           : localhost:3306
 Source Schema         : cl_admin

 Target Server Type    : MySQL
 Target Server Version : 50721 (5.7.21-log)
 File Encoding         : 65001

 Date: 27/10/2023 10:11:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_article
-- ----------------------------
DROP TABLE IF EXISTS `sys_article`;
CREATE TABLE `sys_article` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL,
  `cate_id` int(11) unsigned NOT NULL,
  `title` varchar(255) NOT NULL,
  `thumb` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `release_time` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_tb_article_deletedAt` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_article
-- ----------------------------
BEGIN;
INSERT INTO `sys_article` (`id`, `user_id`, `cate_id`, `title`, `thumb`, `content`, `release_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 1, 6, '标题二修改', 'http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20190506/0db7fc99-d70a-49d8-8677-60968acb8471.jpg', '<p>内容修改</p>', '2019-05-03 11:29:02', '2019-05-05 09:11:20', '2019-07-18 11:31:44', NULL);
INSERT INTO `sys_article` (`id`, `user_id`, `cate_id`, `title`, `thumb`, `content`, `release_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 1, 6, '标题三', 'http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20190506/16a47da4-cb6b-47fe-8857-34dfa1f72128.jpg', '<p>内容</p>', '2019-05-05 08:28:54', '2019-05-05 09:13:07', '2023-10-26 17:33:38', NULL);
INSERT INTO `sys_article` (`id`, `user_id`, `cate_id`, `title`, `thumb`, `content`, `release_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 1, 10, '23123', 'http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20190506/597c1c9f-90a5-4596-8db9-1e070dd92259.jpg', '<p style=\"text-align: center;\"><img src=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20190506/f9763bd6-d902-478d-8ff8-5963075611f8.jpg\"/></p><p style=\"text-align: center;\"><img src=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20190506/795261af-3ca4-402a-8032-102a2470dfa1.gif\"/></p><p style=\"text-align: center;\">asdas</p><p style=\"text-align: center;\"><embed type=\"application/x-shockwave-flash\" class=\"edui-faked-video\" pluginspage=\"http://www.macromedia.com/go/getflashplayer\" src=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20190506/fa28bf52-04f5-4312-8749-4f1cc4aaeeab.mp4\" width=\"420\" height=\"280\" wmode=\"transparent\" play=\"true\" loop=\"false\" menu=\"false\" allowscriptaccess=\"never\" allowfullscreen=\"true\"/></p>', '2019-05-05 09:28:37', '2019-05-06 10:22:24', '2023-10-26 17:33:45', NULL);
INSERT INTO `sys_article` (`id`, `user_id`, `cate_id`, `title`, `thumb`, `content`, `release_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, 1, 11, '标题四', 'http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20190506/8368ad28-04a0-4be6-8496-59e350c2ee38.jpg', '<p style=\"text-align: left;\">实打实大师的大三</p><p style=\"text-align: left;\">啊大神</p><p style=\"text-align: center;\"><img src=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20190506/2870b84f-ef29-4478-87d9-20ead2c19879.jpg\"/></p>', '2019-05-06 07:28:13', '2019-05-06 10:34:32', '2023-10-26 17:33:52', NULL);
INSERT INTO `sys_article` (`id`, `user_id`, `cate_id`, `title`, `thumb`, `content`, `release_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (7, 1, 6, '标题五修改', 'http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20190506/4e0b5ca1-9143-4779-8303-318aa5589805.jpg', '<p style=\"text-align: center;\">21312312隐隐约约b</p><p style=\"text-align: center;\"><embed type=\"application/x-shockwave-flash\" class=\"edui-faked-video\" pluginspage=\"http://www.macromedia.com/go/getflashplayer\" src=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20190506/af91f0d2-f119-44e7-80a0-860e08fd6885.mp4\" width=\"420\" height=\"280\" wmode=\"transparent\" play=\"true\" loop=\"false\" menu=\"false\" allowscriptaccess=\"never\" allowfullscreen=\"true\"/></p>', '2019-05-06 11:28:00', '2019-05-06 10:35:51', '2019-07-01 10:11:49', NULL);
INSERT INTO `sys_article` (`id`, `user_id`, `cate_id`, `title`, `thumb`, `content`, `release_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (8, 1, 12, '文章阿 ', 'http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20190506/5dddd078-b073-4ae7-819e-252b4b18e5ad.jpg', '<p>实打实的</p>', '2019-05-06 16:18:31', '2019-05-06 16:16:28', '2023-10-26 17:34:09', NULL);
INSERT INTO `sys_article` (`id`, `user_id`, `cate_id`, `title`, `thumb`, `content`, `release_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (9, 1, 10, '文章啊啊啊阿', 'http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20190506/3d0a9d14-c670-4e7f-8639-8e59141b0fc9.jpg', '<p style=\"text-align: center;\"><img src=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20190506/caa25884-8b79-4dde-8d80-4d891f92f651.jpg\"/></p><p style=\"text-align: center;\"><img src=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20190506/f67d9200-afc5-4ed1-8707-f00250ad7981.jpg\"/></p><p style=\"text-align: center;\"><img src=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20190506/cc75398b-0b0e-4d24-88c0-b5b4b2782c33.jpg\"/></p><p style=\"text-align: center;\"><img src=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20190506/887cfde4-0cf9-4eef-89e7-9dce63a9ac82.jpg\"/></p>', '2019-05-06 16:04:18', '2019-05-06 16:18:15', '2023-10-26 17:34:22', NULL);
INSERT INTO `sys_article` (`id`, `user_id`, `cate_id`, `title`, `thumb`, `content`, `release_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (10, 1, 11, '啊啊啊啊啊', 'http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20190506/7855e7d0-9c80-4cc6-8526-441d9562c8ab.jpg', '<p style=\"text-align: center;\"><img src=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20190506/3a7e6c3f-c151-4e48-8a8a-f453cc2665d8.jpg\"/></p><p style=\"text-align: center;\"><img src=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20190506/70d39b8c-0fda-4796-87d2-9225e09db9a6.jpg\"/></p>', '2019-05-06 16:04:33', '2019-05-06 16:33:36', '2023-10-26 17:34:28', NULL);
INSERT INTO `sys_article` (`id`, `user_id`, `cate_id`, `title`, `thumb`, `content`, `release_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (11, 1, 6, '测试文章', 'http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20231006/509c0449-3941-4297-8606-bb6e8e728efb.jpg', '<p style=\"text-align: center;\">测试文章</p><p style=\"text-align: center;\"><img src=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20231006/8deb4359-8f86-4086-8721-39dc503bb7ea.jpg\"/></p><p style=\"text-align: center;\"><img src=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20231006/9ffdbecf-c2a9-4433-8d49-64cdda691b92.jpg\"/></p><p><br/></p>', '2023-10-06 15:33:25', '2023-10-06 15:33:26', '2023-10-06 15:33:26', NULL);
INSERT INTO `sys_article` (`id`, `user_id`, `cate_id`, `title`, `thumb`, `content`, `release_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (12, 1, 6, '测试', 'http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20231011/7429d3bc-d39f-43ec-8e4d-cbf47598c097.jpg', '<p style=\"text-indent: 2em;\">IT之家 10 月 11 日消息，根据国外科技媒体 9to5Mac 和诸多网友提交的反馈，iPhone 即便处于充电状态，在夜间会出现自动关机，用户起床看到的是 iPhone 密码屏幕页面。</p>\n<p style=\"text-indent: 0em; text-align: center;\"><img src=\"https://inews.gtimg.com/om_bt/OOA4SKCdsCnyOXWU6ygaJu1WgvQNY75rDofblU8nfDOv4AA/641\" alt=\"图片\" /></p>\n<p style=\"text-indent: 2em;\">这个问题并非个例，该媒体编辑使用的 iPhone 15 Pro Max 在没有启用电池优化的情况下，并且使用 MagSafe 过夜充电的情况下，夜间中途手机会自动关机一段时间。</p>\n<p style=\"text-indent: 2em;\">Reddit 上也有部分用户反馈存在这个问题，网友 @rathan_lesage 表示他的 iPhone 在凌晨 3 点到 7 点之间自动关机，在闹钟响起的时候，必须要重新输入密码。</p>\n<p style=\"text-indent: 0em; text-align: center;\"><img src=\"https://inews.gtimg.com/om_bt/ObhTuxxgH-rg9n3qFCIrNqdWS-vMM0oS3lYhKp-9EbXpIAA/641\" alt=\"图片\" /></p>\n<p style=\"text-indent: 2em;\">@rathan_lesage 设置了充电优化，将充电电量限制在 80%。不过在夜间自动重启之后，充电直接恢复到 100%。</p>\n<p style=\"text-indent: 2em;\">IT之家查询微博，目前并未发现有用户反馈存在夜间自动关机的情况。</p>', '2023-10-11 11:07:31', '2023-10-11 11:07:32', '2023-10-11 11:07:32', NULL);
INSERT INTO `sys_article` (`id`, `user_id`, `cate_id`, `title`, `thumb`, `content`, `release_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (13, 1, 6, '测试修改', 'http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20231011/8d733af1-4e24-4d4e-880f-8f6b474b514a.jpg', '<p style=\"text-indent: 0em; text-align: center;\"><img src=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20231011/b682ed94-6e76-4864-80ab-7d90d3c78081.jpg\" /><img src=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20231011/8c59c6ea-9822-47b0-8cb1-804add6cf513.jpg\" /></p>\n<p style=\"text-indent: 2em;\">以 5 月 1 日为界，回归一季度业绩期后至今，我们覆盖的国内泛电商公司的股价表现，除了一只独秀的拼多多和勉强跑出正收益的阿里集团外，其他各家股价不仅都是下跌的，且也都跑输了中概金龙指数同期的走势。即便不愿也无法否认，在过去的近 5 个月内，国内泛电商公司在绝对收益和相对收益上两个维度上的表现都是偏弱的。</p>\n<p style=\"text-indent: 2em;\">以 5 月 1 日为界，回归一季度业绩期后至今，我们覆盖的国内泛电商公司的股价表现，除了一只独秀的拼多多和勉强跑出正收益的阿里集团外，其他各家股价不仅都是下跌的，且也都跑输了中概金龙指数同期的走势。即便不愿也无法否认，在过去的近 5 个月内，国内泛电商公司在绝对收益和相对收益上两个维度上的表现都是偏弱的。</p>\n<p style=\"text-indent: 0em; text-align: center;\"><video src=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20231011/f65275ae-a5d9-45a0-8eaf-6e884e0c1508.mp4\" poster=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20231011/909e9191-f6d7-4b66-80e2-a1c22b33c7e9.jpeg\" controls=\"controls\" data-mce-src=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20231011/f65275ae-a5d9-45a0-8eaf-6e884e0c1508.mp4\"></video></p>\n<p style=\"text-indent: 0em; text-align: center;\"><video src=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20231011/2e64dc60-6186-4122-8ecf-de7db406ca36.mp4\" poster=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20231011/16b49c4c-da4e-4641-80ec-192e836ae9e5.jpg\" controls=\"controls\" data-mce-src=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20231011/2e64dc60-6186-4122-8ecf-de7db406ca36.mp4\"></video></p>\n<p style=\"text-indent: 2em;\">以 5 月 1 日为界，回归一季度业绩期后至今，我们覆盖的国内泛电商公司的股价表现，除了一只独秀的拼多多和勉强跑出正收益的阿里集团外，其他各家股价不仅都是下跌的，且也都跑输了中概金龙指数同期的走势。即便不愿也无法否认，在过去的近 5 个月内，国内泛电商公司在绝对收益和相对收益上两个维度上的表现都是偏弱的。</p>\n<p style=\"text-indent: 0em; text-align: center;\"><audio src=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20231011/4e9af023-6afe-4b1d-8235-17ea2bd6ba5c.m4a\" controls=\"controls\"></audio></p>\n<p style=\"text-indent: 2em;\"><a href=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20231011/3b9ef2ea-8b37-4277-868c-debcf07575c3.xlsx\" target=\"_blank\" rel=\"noopener\">题目格式.xlsx</a></p>\n<p style=\"text-indent: 2em;\"><a href=\"http://aisyweixinpic.oss-cn-shanghai.aliyuncs.com/20231011/19ca3eee-1685-4eba-8694-6981a5a374ea.pdf\" target=\"_blank\" rel=\"noopener\">2021.6湖北省广播电视工程评审条件.pdf</a></p>\n<p style=\"text-indent: 2em;\"></p>', '2023-10-11 11:35:00', '2023-10-11 11:35:00', '2023-10-11 15:07:23', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_bulletin
-- ----------------------------
DROP TABLE IF EXISTS `sys_bulletin`;
CREATE TABLE `sys_bulletin` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `tag` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `ids_tb_public_notice_deletedAt` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_bulletin
-- ----------------------------
BEGIN;
INSERT INTO `sys_bulletin` (`id`, `title`, `tag`, `content`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, '瓜迪奥拉：德泽尔比是我在全世界范围内最钦佩的主教练之一', '1', '内容内容', '2023-10-25 23:37:04', '2023-10-25 23:37:04', NULL);
INSERT INTO `sys_bulletin` (`id`, `title`, `tag`, `content`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, '【关键先生】罗德里本场参与曼城全部三粒进球，回归后曼城两连胜', '2', '内容内容', '2023-10-26 08:26:22', '2023-10-26 08:26:22', NULL);
INSERT INTO `sys_bulletin` (`id`, `title`, `tag`, `content`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'TA：德布劳内2025年将是34岁，曼城可能不愿提供超过1年新合同', '3', '内容内容', '2023-10-26 13:46:39', '2023-10-26 13:46:39', NULL);
INSERT INTO `sys_bulletin` (`id`, `title`, `tag`, `content`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, '老雷：到目前为止，理查利森的表现仍令我无法信服', '1', '内容内容', '2023-10-26 15:21:06', '2023-10-26 15:21:06', NULL);
INSERT INTO `sys_bulletin` (`id`, `title`, `tag`, `content`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, '热刺旧将：为什么不能夺得英超冠军？就因为我们是热刺？', '3', '内容内容', '2023-10-26 15:57:31', '2023-10-26 15:57:31', NULL);
INSERT INTO `sys_bulletin` (`id`, `title`, `tag`, `content`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, '惨昔日多特三杰近况：哈兰德贝皇高光！桑乔跟青训队吃饭...', '2', '内容内容', '2023-10-26 16:11:24', '2023-10-26 16:11:24', NULL);
INSERT INTO `sys_bulletin` (`id`, `title`, `tag`, `content`, `created_at`, `updated_at`, `deleted_at`) VALUES (7, '瓜帅谈论表现出色球队突然被问：那曼联呢❓直接被整笑：他们迟早会赢', '1', '内容内容', '2019-06-10 09:18:49', '2019-06-10 09:18:49', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_captcha
-- ----------------------------
DROP TABLE IF EXISTS `sys_captcha`;
CREATE TABLE `sys_captcha` (
  `uuid` char(36) CHARACTER SET utf8 NOT NULL COMMENT 'uuid',
  `code` varchar(6) CHARACTER SET utf8 NOT NULL COMMENT '验证码',
  `expire_time` datetime DEFAULT NULL COMMENT '过期时间',
  PRIMARY KEY (`uuid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统验证码';

-- ----------------------------
-- Records of sys_captcha
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_category
-- ----------------------------
DROP TABLE IF EXISTS `sys_category`;
CREATE TABLE `sys_category` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) unsigned NOT NULL COMMENT '父菜单ID，一级菜单为0',
  `name` varchar(50) CHARACTER SET utf8 NOT NULL COMMENT '菜单名称',
  `icon` varchar(50) CHARACTER SET utf8 NOT NULL COMMENT '菜单图标',
  `order_num` int(11) NOT NULL COMMENT '排序',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_tb_category_deletedAt` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of sys_category
-- ----------------------------
BEGIN;
INSERT INTO `sys_category` (`id`, `parent_id`, `name`, `icon`, `order_num`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 0, '时事', '', 0, '2019-04-19 17:47:48', '2023-10-26 17:32:51', NULL);
INSERT INTO `sys_category` (`id`, `parent_id`, `name`, `icon`, `order_num`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 0, '视频分类', '', 0, '2019-04-19 17:48:40', '2019-04-19 17:48:40', '2023-10-26 17:31:58');
INSERT INTO `sys_category` (`id`, `parent_id`, `name`, `icon`, `order_num`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 0, '图片分类', '', 0, '2019-04-19 17:48:44', '2019-04-19 17:48:44', '2023-10-26 17:31:53');
INSERT INTO `sys_category` (`id`, `parent_id`, `name`, `icon`, `order_num`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, 3, '国内', '', 0, '2019-04-28 15:21:02', '2023-10-26 17:30:23', NULL);
INSERT INTO `sys_category` (`id`, `parent_id`, `name`, `icon`, `order_num`, `created_at`, `updated_at`, `deleted_at`) VALUES (8, 4, '本地', '', 0, '2019-04-28 15:21:54', '2023-10-26 17:30:45', '2023-10-26 17:31:38');
INSERT INTO `sys_category` (`id`, `parent_id`, `name`, `icon`, `order_num`, `created_at`, `updated_at`, `deleted_at`) VALUES (9, 0, '其他分类', '', 0, '2019-04-28 15:22:43', '2019-07-18 11:26:12', '2023-10-26 17:31:50');
INSERT INTO `sys_category` (`id`, `parent_id`, `name`, `icon`, `order_num`, `created_at`, `updated_at`, `deleted_at`) VALUES (10, 3, '国际', '', 0, '2023-10-26 17:30:34', '2023-10-26 17:30:34', NULL);
INSERT INTO `sys_category` (`id`, `parent_id`, `name`, `icon`, `order_num`, `created_at`, `updated_at`, `deleted_at`) VALUES (11, 0, '体育', '', 0, '2023-10-26 17:33:01', '2023-10-26 17:33:01', NULL);
INSERT INTO `sys_category` (`id`, `parent_id`, `name`, `icon`, `order_num`, `created_at`, `updated_at`, `deleted_at`) VALUES (12, 0, '游戏', '', 0, '2023-10-26 17:33:06', '2023-10-26 17:33:06', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `param_key` varchar(255) CHARACTER SET utf8 NOT NULL COMMENT 'key',
  `param_value` longtext CHARACTER SET utf8 NOT NULL COMMENT 'value',
  `type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '1字符串值类型 2字符串JSON类型',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态   0：隐藏   1：显示',
  `remark` varchar(255) CHARACTER SET utf8 NOT NULL COMMENT '备注',
  `locked` tinyint(1) NOT NULL COMMENT '锁定  0：否  1：是',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `param_key` (`param_key`) USING BTREE,
  KEY `idx_tb_config_deletedAt` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统配置信息表';

-- ----------------------------
-- Records of sys_config
-- ----------------------------
BEGIN;
INSERT INTO `sys_config` (`id`, `param_key`, `param_value`, `type`, `status`, `remark`, `locked`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'CLOUD_STORAGE_ALI_CONFIG_KEY', '{\"aliyunAccessKeyId\":\"\",\"aliyunAccessKeySecret\":\"\",\"aliyunBucketName\":\"\",\"aliyunEndPoint\":\"\",\"ossType\":\"1\"}', 2, 1, '阿里云OSS配置', 1, '2023-07-18 15:59:34', '2023-10-24 11:42:55', NULL);
INSERT INTO `sys_config` (`id`, `param_key`, `param_value`, `type`, `status`, `remark`, `locked`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'ALI_STS_CONFIG_KEY', '{\"aliyunAccessKeyId\":\"\",\"aliyunAccessKeySecret\":\"\",\"aliyunEndPoint\":\"\",\"aliyunOSSBucket\":\"\",\"aliyunOSSRegion\":\"\",\"aliyunRoleArn\":\"\",\"aliyunRoleSessionName\":\"\"}', 2, 1, '阿里云STS配置', 1, '2023-09-06 10:35:31', '2023-10-26 16:12:39', NULL);
INSERT INTO `sys_config` (`id`, `param_key`, `param_value`, `type`, `status`, `remark`, `locked`, `created_at`, `updated_at`, `deleted_at`) VALUES (16, 'TEST', 'test', 1, 1, '测试参数', 0, '2023-10-26 16:57:47', '2023-10-26 17:00:18', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `parent_id` int(11) unsigned NOT NULL COMMENT '上级ID',
  `name` varchar(255) NOT NULL COMMENT '部门名称',
  `sort` tinyint(4) NOT NULL COMMENT '排序',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_deletedAt` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
BEGIN;
INSERT INTO `sys_dept` (`id`, `parent_id`, `name`, `sort`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 0, 'XXX集团', 1, '2023-10-24 17:42:24', '2023-10-26 10:04:43', NULL);
INSERT INTO `sys_dept` (`id`, `parent_id`, `name`, `sort`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 1, '研发部', 0, '2023-10-24 17:44:32', '2023-10-26 10:05:29', NULL);
INSERT INTO `sys_dept` (`id`, `parent_id`, `name`, `sort`, `created_at`, `updated_at`, `deleted_at`) VALUES (10, 1, '市场部', 0, '2023-10-26 10:05:21', '2023-10-26 10:05:21', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `dict_type_id` int(11) unsigned NOT NULL COMMENT '字典类型ID',
  `dict_label` varchar(255) NOT NULL COMMENT '字典标签',
  `dict_value` varchar(255) NOT NULL COMMENT '字典值',
  `remark` varchar(255) NOT NULL COMMENT '备注',
  `sort` int(11) unsigned NOT NULL COMMENT '排序',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uk_dict_type_value` (`dict_type_id`,`dict_value`),
  KEY `idx_sort` (`sort`),
  KEY `idx_deletedAt` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COMMENT='字典数据';

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
BEGIN;
INSERT INTO `sys_dict_data` (`id`, `dict_type_id`, `dict_label`, `dict_value`, `remark`, `sort`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 3, '男', '0', '', 1, '2023-10-16 09:35:57', '2023-10-16 09:46:17', NULL);
INSERT INTO `sys_dict_data` (`id`, `dict_type_id`, `dict_label`, `dict_value`, `remark`, `sort`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 3, '女', '1', '', 2, '2023-10-16 09:46:38', '2023-10-16 09:46:59', NULL);
INSERT INTO `sys_dict_data` (`id`, `dict_type_id`, `dict_label`, `dict_value`, `remark`, `sort`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 3, '保密', '2', '', 3, '2023-10-16 09:48:27', '2023-10-16 15:46:30', NULL);
INSERT INTO `sys_dict_data` (`id`, `dict_type_id`, `dict_label`, `dict_value`, `remark`, `sort`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 4, '公告', '1', '', 1, '2023-10-16 09:55:16', '2023-10-16 09:55:16', NULL);
INSERT INTO `sys_dict_data` (`id`, `dict_type_id`, `dict_label`, `dict_value`, `remark`, `sort`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, 4, '会议', '2', '', 2, '2023-10-16 09:55:24', '2023-10-16 09:55:24', NULL);
INSERT INTO `sys_dict_data` (`id`, `dict_type_id`, `dict_label`, `dict_value`, `remark`, `sort`, `created_at`, `updated_at`, `deleted_at`) VALUES (7, 4, '其他', '3', '', 3, '2023-10-16 09:55:33', '2023-10-16 09:55:33', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `dict_type` varchar(100) NOT NULL COMMENT '字典类型',
  `dict_name` varchar(255) NOT NULL COMMENT '字典名称',
  `remark` varchar(255) NOT NULL COMMENT '备注',
  `sort` int(11) unsigned NOT NULL COMMENT '排序',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_dict_type` (`dict_type`),
  KEY `idx_deletedAt` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COMMENT='字典类型';

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
BEGIN;
INSERT INTO `sys_dict_type` (`id`, `dict_type`, `dict_name`, `remark`, `sort`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'gender', '性别', '性别备注', 0, '2023-10-13 15:30:26', '2023-10-13 15:32:16', NULL);
INSERT INTO `sys_dict_type` (`id`, `dict_type`, `dict_name`, `remark`, `sort`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'notice_type', '站内通知-类型', '站内通知-类型-备注', 0, '2023-10-13 15:30:44', '2023-10-13 15:30:44', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_log`;
CREATE TABLE `sys_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8 DEFAULT NULL COMMENT '用户名',
  `operation` varchar(50) CHARACTER SET utf8 DEFAULT NULL COMMENT '用户操作',
  `method` varchar(200) CHARACTER SET utf8 DEFAULT NULL COMMENT '请求方法',
  `params` varchar(5000) CHARACTER SET utf8 DEFAULT NULL COMMENT '请求参数',
  `time` bigint(20) NOT NULL COMMENT '执行时长(毫秒)',
  `ip` varchar(64) CHARACTER SET utf8 DEFAULT NULL COMMENT 'IP地址',
  `create_date` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统日志';

-- ----------------------------
-- Records of sys_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) unsigned NOT NULL COMMENT '父菜单ID，一级菜单为0',
  `name` varchar(50) CHARACTER SET utf8 NOT NULL COMMENT '菜单名称',
  `url` varchar(200) CHARACTER SET utf8 NOT NULL COMMENT '菜单URL',
  `perms` varchar(500) CHARACTER SET utf8 NOT NULL COMMENT '授权(多个用逗号分隔，如：user:list,user:create)',
  `type` int(11) NOT NULL COMMENT '类型   0：目录   1：菜单   2：按钮',
  `icon` varchar(50) CHARACTER SET utf8 NOT NULL COMMENT '菜单图标',
  `order_num` int(11) NOT NULL COMMENT '排序',
  `is_tab` tinyint(1) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_tb_menu_deletedAt` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=72 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='菜单管理';

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 0, '系统管理', '', '', 0, 'system', 0, 0, 1, '0000-00-00 00:00:00', '2019-06-11 16:45:44', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 1, '管理员列表', 'sys/user', '', 1, 'admin', 1, 1, 1, '0000-00-00 00:00:00', '2023-10-06 15:08:20', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 1, '角色管理', 'sys/role', '', 1, 'role', 2, 1, 1, '0000-00-00 00:00:00', '2023-10-06 15:09:18', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 1, '菜单管理', 'sys/menu', '', 1, 'menu', 3, 1, 1, '0000-00-00 00:00:00', '2023-10-06 15:09:26', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 1, 'SQL监控', 'http://localhost:8080/renren-fast/druid/sql.html', '', 1, 'sql', 4, 1, 1, '0000-00-00 00:00:00', '2023-10-06 15:14:06', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (15, 2, '查看', '', 'sys:user:list,sys:user:info', 2, '', 0, 0, 1, '0000-00-00 00:00:00', '2019-04-08 17:15:42', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (16, 2, '新增', '', 'sys:user:save,sys:role:select', 2, '', 0, 0, 1, '0000-00-00 00:00:00', '2019-04-08 17:15:42', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (17, 2, '修改', '', 'sys:user:update,sys:role:select', 2, '', 0, 0, 1, '0000-00-00 00:00:00', '2019-04-08 17:15:42', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (18, 2, '删除', '', 'sys:user:delete', 2, '', 0, 0, 1, '0000-00-00 00:00:00', '2019-04-08 17:15:42', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (19, 3, '查看', '/v1/roles/get', 'sys:role:info', 2, '', 0, 0, 1, '0000-00-00 00:00:00', '2019-06-11 16:45:44', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (20, 3, '新增', '/v1/roles/create', 'sys:role:save', 2, '', 0, 0, 1, '0000-00-00 00:00:00', '2019-06-11 16:45:44', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (21, 3, '修改', '/v1/roles/update', 'sys:role:update', 2, '', 0, 0, 1, '0000-00-00 00:00:00', '2019-06-11 16:45:44', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (22, 3, '删除', '/v1/roles/delete', 'sys:role:delete', 2, '', 0, 0, 1, '0000-00-00 00:00:00', '2019-06-11 16:45:44', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (23, 4, '查看', '/v1/menus/get', 'sys:menu:info', 2, '', 0, 0, 1, '0000-00-00 00:00:00', '2019-04-11 10:58:38', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (24, 4, '新增', '/v1/menus/create', 'sys:menu:save', 2, '', 0, 0, 1, '0000-00-00 00:00:00', '2019-04-11 10:58:38', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (25, 4, '修改', '/v1/menus/update', 'sys:menu:update', 2, '', 0, 0, 1, '0000-00-00 00:00:00', '2019-04-11 10:58:38', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (26, 4, '删除', '/v1/menus/delete', 'sys:menu:delete', 2, '', 0, 0, 1, '0000-00-00 00:00:00', '2019-04-11 10:58:38', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (27, 1, '参数管理', 'sys/config', '', 1, 'config', 6, 1, 1, '0000-00-00 00:00:00', '2023-10-24 09:59:56', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (29, 1, '其他', 'sys/other', '', 1, 'log', 7, 0, 0, '0000-00-00 00:00:00', '2023-10-12 14:43:38', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (30, 1, '文件上传', 'oss/oss', 'sys:oss:all', 1, 'oss', 8, 1, 1, '0000-00-00 00:00:00', '2023-10-24 15:37:50', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (32, 3, '列表', '/v1/roles/list', 'sys:role:list', 2, '', 0, 0, 1, '2019-04-10 15:19:54', '2019-06-11 16:45:44', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (33, 4, '列表', '/v1/menus/list', 'sys:menu:list', 2, '', 0, 0, 1, '2019-04-10 16:00:32', '2019-04-11 10:58:38', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (34, 0, '内容管理', '', '', 0, 'zonghe', 0, 0, 1, '0000-00-00 00:00:00', '2019-07-01 10:11:10', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (35, 34, '分类管理', 'content/category', '', 1, 'menu', 0, 1, 1, '2019-04-28 15:16:24', '2023-10-06 15:09:45', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (36, 35, '查看', '/v1/categoies/get', 'sys:category:info', 2, '', 0, 0, 1, '2019-04-28 15:17:36', '2019-07-01 10:11:10', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (37, 35, '新增', '/v1/categories/create', 'sys:category:save', 2, '', 0, 0, 1, '2019-04-28 15:18:38', '2019-07-01 10:11:10', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (38, 35, '修改', '/v1/categories/update', 'sys:category:update', 2, '', 0, 0, 1, '2019-04-28 15:19:06', '2019-07-01 10:11:10', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (39, 35, '删除', '/v1/categories/delete', 'sys:category:delete', 2, '', 0, 0, 1, '2019-04-28 15:19:40', '2019-07-01 10:11:10', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (40, 35, '列表', '/v1/categories/list', 'sys:category:list', 2, '', 0, 0, 1, '2019-04-28 15:20:42', '2019-07-01 10:11:10', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (41, 34, '文章管理', 'content/article', '', 1, 'log', 0, 1, 1, '2019-04-29 09:43:03', '2023-10-06 15:09:54', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (42, 41, '查看', '/v1/articles/get', 'sys:article:info', 2, '', 0, 0, 1, '2019-04-29 09:57:23', '2019-07-01 10:11:10', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (43, 41, '新增', '/v1/articles/create', 'sys:article:save', 2, '', 0, 0, 1, '2019-04-29 09:58:04', '2019-07-01 10:11:10', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (44, 41, '修改', '/v1/articles/update', 'sys:article:update', 2, '', 0, 0, 1, '2019-04-29 09:58:50', '2019-07-01 10:11:10', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (45, 41, '删除', '/v1/articles/delete', 'sys:article:delete', 2, '', 0, 0, 1, '2019-04-29 09:59:53', '2019-07-01 10:11:10', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (46, 41, '列表', '/v1/articles/list', 'sys:article:list', 2, '', 0, 0, 1, '2019-04-29 10:01:28', '2019-07-01 10:11:10', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (47, 2, '注销登录', '/v1/users/logout', 'sys:user:logout', 2, '', 0, 0, 1, '2019-09-16 18:00:59', '2023-10-26 17:52:14', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (48, 34, '发布文章', 'content/article-add-or-update', '', 1, 'bianji', 0, 0, 1, '2023-10-06 15:19:12', '2023-10-12 11:08:08', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (49, 29, '公告列表', '/v1/bulletin/list', 'sys:bulletin:list', 2, '', 0, 0, 1, '2023-10-12 14:46:14', '2023-10-13 16:28:13', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (50, 1, '字典管理', 'sys/dict-type', '', 1, 'read', 5, 1, 1, '2023-10-13 10:24:24', '2023-10-24 15:37:33', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (51, 50, '列表', '/v1/dictType/list', 'sys:dictType:list', 2, '', 0, 0, 1, '2023-10-13 10:25:10', '2023-10-13 10:25:10', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (52, 50, '新增', '/v1/dictType/create', 'sys:dictType:save', 2, '', 0, 0, 1, '2023-10-13 10:28:07', '2023-10-13 10:30:13', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (53, 50, '查看', '/v1/dictType/get', 'sys:dictType:info', 2, '', 0, 0, 1, '2023-10-13 14:30:18', '2023-10-13 14:30:18', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (54, 50, '修改', '/v1/dictType/update', 'sys:dictType:update', 2, '', 0, 0, 1, '2023-10-13 14:31:27', '2023-10-13 14:31:27', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (55, 50, '删除', '/v1/dictType/delete', 'sys:dictType:delete', 2, '', 0, 0, 1, '2023-10-13 14:32:01', '2023-10-13 14:32:01', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (56, 50, '新增字典值', '/v1/dictData/create', 'sys:dictData:save', 2, '', 0, 0, 1, '2023-10-13 17:27:19', '2023-10-13 17:27:19', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (57, 50, '删除字典值', '/v1/dictData/delete', 'sys:dictData:delete', 2, '', 0, 0, 1, '2023-10-13 17:28:20', '2023-10-13 17:28:20', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (58, 50, '列表字典值', '/v1/dictData/list', 'sys:dictData:list', 2, '', 0, 0, 1, '2023-10-16 09:58:46', '2023-10-16 10:00:01', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (59, 50, '查看字典值', '/v1/dictData/get', 'sys:dictData:list', 2, '', 0, 0, 1, '2023-10-16 10:00:36', '2023-10-16 10:00:36', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (60, 50, '修改字典值', '/v1/dictData/update', 'sys:dictData:update', 2, '', 0, 0, 1, '2023-10-16 10:01:16', '2023-10-16 10:01:16', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (61, 27, '新增', '/v1/config/create', 'sys:config:save', 2, '', 0, 0, 1, '2023-10-24 10:00:36', '2023-10-24 10:00:36', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (62, 27, '修改', '/v1/config/update', 'sys:config:update', 2, '', 0, 0, 1, '2023-10-24 10:01:14', '2023-10-24 10:01:14', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (63, 27, '查看', '/v1/config/get', 'sys:config:info', 2, '', 0, 0, 1, '2023-10-24 10:01:44', '2023-10-24 10:01:44', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (64, 27, '列表', '/v1/config/list', 'sys:config:list', 2, '', 0, 0, 1, '2023-10-24 10:02:10', '2023-10-24 10:02:10', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (65, 27, '删除', '/v1/config/delete', 'sys:config:delete', 2, '', 0, 0, 1, '2023-10-24 10:02:44', '2023-10-24 10:02:44', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (66, 1, '部门管理', '/sys/dept', '', 1, 'cluster', 2, 1, 1, '2023-10-24 17:31:10', '2023-10-26 09:51:53', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (67, 66, '新增', '/v1/dept/create', 'sys:dept:save', 2, '', 0, 0, 1, '2023-10-24 17:33:46', '2023-10-24 17:33:46', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (68, 66, '修改', '/v1/dept/update', 'sys:dept:update', 2, '', 0, 0, 1, '2023-10-24 17:34:09', '2023-10-24 17:34:09', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (69, 66, '查看', '/v1/dept/get', 'sys:dept:info', 2, '', 0, 0, 1, '2023-10-24 17:34:32', '2023-10-24 17:34:32', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (70, 66, '列表', '/v1/dept/list', 'sys:dept:list', 2, '', 0, 0, 1, '2023-10-24 17:34:50', '2023-10-24 17:34:50', NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `url`, `perms`, `type`, `icon`, `order_num`, `is_tab`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (71, 66, '删除', '/v1/dept/delete', 'sys:dept:delete', 2, '', 0, 0, 1, '2023-10-24 17:35:08', '2023-10-24 17:35:08', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_oss
-- ----------------------------
DROP TABLE IF EXISTS `sys_oss`;
CREATE TABLE `sys_oss` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `url` varchar(200) CHARACTER SET utf8 DEFAULT NULL COMMENT 'URL地址',
  `create_date` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='文件上传';

-- ----------------------------
-- Records of sys_oss
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `role_name` varchar(100) CHARACTER SET utf8 NOT NULL COMMENT '角色名称',
  `remark` varchar(100) CHARACTER SET utf8 NOT NULL COMMENT '备注',
  `menu_id_list` text CHARACTER SET utf8 NOT NULL COMMENT '配合前端tree半选 666666为临时KEY分隔符',
  `create_user_id` int(11) unsigned NOT NULL COMMENT '创建者ID',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_tb_role_deletedAt` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='角色';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_role` (`id`, `role_name`, `remark`, `menu_id_list`, `create_user_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, '角色一', '角色一说明', '[1,50,51,52,53,54,55,56,57,58,59,60,29,49,40,41,42,43,44,45,46,48,34,35]', 0, '2019-03-26 09:25:17', '2023-10-16 11:51:07', NULL);
INSERT INTO `sys_role` (`id`, `role_name`, `remark`, `menu_id_list`, `create_user_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, '角色二', '角色一说明', '[29,49,30,34,35,36,37,38,39,40,41,42,43,44,45,46,48,1]', 0, '2019-03-26 09:27:01', '2023-10-26 17:28:51', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` int(11) unsigned NOT NULL COMMENT '角色ID',
  `menu_id` int(11) unsigned NOT NULL COMMENT '菜单ID',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uni_role_id_menu_id` (`role_id`,`menu_id`) USING BTREE,
  KEY `idx_tb_role_menu_deletedAt` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=299 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='角色与菜单对应关系';

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (257, 1, 1, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (258, 1, 29, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (259, 1, 34, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (260, 1, 35, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (261, 1, 40, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (262, 1, 41, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (263, 1, 42, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (264, 1, 43, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (265, 1, 44, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (266, 1, 45, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (267, 1, 46, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (268, 1, 48, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (269, 1, 49, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (270, 1, 50, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (271, 1, 51, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (272, 1, 52, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (273, 1, 53, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (274, 1, 54, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (275, 1, 55, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (276, 1, 56, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (277, 1, 57, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (278, 1, 58, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (279, 1, 59, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (280, 1, 60, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (281, 2, 1, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (282, 2, 29, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (283, 2, 30, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (284, 2, 34, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (285, 2, 35, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (286, 2, 36, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (287, 2, 37, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (288, 2, 38, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (289, 2, 39, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (290, 2, 40, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (291, 2, 41, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (292, 2, 42, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (293, 2, 43, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (294, 2, 44, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (295, 2, 45, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (296, 2, 46, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (297, 2, 48, NULL, NULL, NULL);
INSERT INTO `sys_role_menu` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (298, 2, 49, NULL, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8 NOT NULL COMMENT '用户名',
  `password` varchar(100) CHARACTER SET utf8 NOT NULL COMMENT '密码',
  `email` varchar(100) CHARACTER SET utf8 NOT NULL COMMENT '邮箱',
  `mobile` varchar(100) CHARACTER SET utf8 NOT NULL COMMENT '手机号',
  `gender` tinyint(4) NOT NULL COMMENT '性别',
  `dept_id` int(11) unsigned NOT NULL COMMENT '部门ID',
  `status` tinyint(4) NOT NULL COMMENT '状态  0：禁用   1：正常',
  `super_admin` tinyint(1) NOT NULL COMMENT '超级管理员   0：否   1：是',
  `create_user_id` int(11) unsigned NOT NULL COMMENT '创建者ID',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `username` (`username`) USING BTREE,
  KEY `idx_tb_users_deletedAt` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统用户';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
BEGIN;
INSERT INTO `sys_user` (`id`, `username`, `password`, `email`, `mobile`, `gender`, `dept_id`, `status`, `super_admin`, `create_user_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'admin', '$2a$10$Maf8Y4P7yNk7p3bJb2jemeqhd23jhxksVxIHPj/lYdWV7yABjtc.6', 'admin@cladmin.com', '15971884095', 0, 1, 1, 1, 1, '2019-03-25 15:03:49', '2023-10-26 14:55:37', NULL);
INSERT INTO `sys_user` (`id`, `username`, `password`, `email`, `mobile`, `gender`, `dept_id`, `status`, `super_admin`, `create_user_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'admin4', '$2a$10$6kN.aUmzbw1NMoLus2mWVOuG.KUFV/SyS1qfKYrdwGQDVmoH2fYy.', 'admin2@cladmin.com', '15971884888', 0, 2, 1, 0, 1, '2019-03-26 09:27:12', '2023-10-26 17:29:27', NULL);
INSERT INTO `sys_user` (`id`, `username`, `password`, `email`, `mobile`, `gender`, `dept_id`, `status`, `super_admin`, `create_user_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 'admin5', '$2a$10$H.99jz8X3.jbrUa2PYMqcuU/CjljeyGAhNmjzFbeKI.AHat4t7cJq', 'admin5@cladmin.com', '15971884888', 1, 10, 1, 0, 1, '2019-03-26 14:50:44', '2023-10-26 10:07:53', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL COMMENT '用户ID',
  `role_id` int(11) unsigned NOT NULL COMMENT '角色ID',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uni_user_id_role_id` (`user_id`,`role_id`) USING BTREE,
  KEY `idx_tb_user_role_deletedAt` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户与角色对应关系';

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_role` (`id`, `user_id`, `role_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (36, 5, 1, NULL, NULL, NULL);
INSERT INTO `sys_user_role` (`id`, `user_id`, `role_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (40, 4, 1, NULL, NULL, NULL);
INSERT INTO `sys_user_role` (`id`, `user_id`, `role_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (41, 4, 2, NULL, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_token
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_token`;
CREATE TABLE `sys_user_token` (
  `user_id` int(11) unsigned NOT NULL,
  `token` varchar(255) NOT NULL COMMENT 'token',
  `expire_time` datetime NOT NULL COMMENT '过期时间',
  `refresh_time` datetime NOT NULL COMMENT '更新时间',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`user_id`) USING BTREE,
  UNIQUE KEY `idx_tb_user_token_token` (`token`) USING BTREE,
  KEY `idx_tb_vote_deletedAt` (`deleted_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统用户Token';

-- ----------------------------
-- Records of sys_user_token
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_token` (`user_id`, `token`, `expire_time`, `refresh_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTgzNzk2MDIsImlhdCI6MTY5ODM3MjQwMiwiaWQiOjEsIm5iZiI6MTY5ODM3MjQwMiwicmV4cCI6MTY5OTY2ODQwMiwic3ViIjoxLCJzdXBlckFkbWluIjp0cnVlLCJ1c2VybmFtZSI6ImFkbWluIn0.5pK_jR3fU9EXuoXXOX40tGdzs55u2dyTLtxspwZqY8w', '2023-10-27 12:06:42', '2023-11-11 10:06:42', '2019-09-16 15:59:05', '2023-07-18 15:37:25', NULL);
INSERT INTO `sys_user_token` (`user_id`, `token`, `expire_time`, `refresh_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Njg2MzQ5MzMsImlhdCI6MTU2ODYyNzczMywiaWQiOjIsIm5iZiI6MTU2ODYyNzczMywicmV4cCI6MTU2OTkyMzczMywic3ViIjoyLCJ1c2VybmFtZSI6ImFkbWluMiJ9.18zGujW-39_lIRhJ1J13sj8FLoznSLPQU1aA9vrwIyc', '2019-09-16 19:55:33', '2019-10-01 17:55:33', '2019-09-16 17:36:04', '2019-09-16 17:55:33', NULL);
INSERT INTO `sys_user_token` (`user_id`, `token`, `expire_time`, `refresh_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Njg2MzUxMDEsImlhdCI6MTU2ODYyNzkwMSwiaWQiOjMsIm5iZiI6MTU2ODYyNzkwMSwicmV4cCI6MTU2OTkyMzkwMSwic3ViIjozLCJ1c2VybmFtZSI6ImFkbWluMyJ9.ZWBMbz0llHVW5BPcRj-cF2Uqd5RUZ9P_tnv1heY78HQ', '2019-09-16 19:58:21', '2019-10-01 17:58:21', '2019-09-16 17:56:09', '2019-09-16 17:58:22', NULL);
INSERT INTO `sys_user_token` (`user_id`, `token`, `expire_time`, `refresh_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTgzMjE2MjksImlhdCI6MTY5ODMxNDQyOSwiaWQiOjQsIm5iZiI6MTY5ODMxNDQyOSwicmV4cCI6MTY5OTYxMDQyOSwic3ViIjo0LCJzdXBlckFkbWluIjpmYWxzZSwidXNlcm5hbWUiOiJhZG1pbjQifQ.9Sy9_rm4t3pvwRVg_zsvSvoaL7Lra-YaSSTzYrIl_FI', '2023-10-26 20:00:29', '2023-11-10 18:00:29', '2023-07-18 16:49:38', '2023-07-18 16:49:38', NULL);
INSERT INTO `sys_user_token` (`user_id`, `token`, `expire_time`, `refresh_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTgyOTQxOTgsImlhdCI6MTY5ODI4Njk5OCwiaWQiOjUsIm5iZiI6MTY5ODI4Njk5OCwicmV4cCI6MTY5OTU4Mjk5OCwic3ViIjo1LCJ1c2VybmFtZSI6ImFkbWluNSJ9.HW0JyHamm5pOlVWvzBSTkVBwxjkl7AmIvupkifs-V2M', '2023-10-26 12:23:18', '2023-11-10 10:23:18', '2023-07-05 17:09:30', '2023-07-10 09:38:57', NULL);
INSERT INTO `sys_user_token` (`user_id`, `token`, `expire_time`, `refresh_time`, `created_at`, `updated_at`, `deleted_at`) VALUES (7, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTgzNzY0NzcsImlhdCI6MTY5ODM2OTI3NywiaWQiOjcsIm5iZiI6MTY5ODM2OTI3NywicmV4cCI6MTY5OTY2NTI3Nywic3ViIjo3LCJzdXBlckFkbWluIjp0cnVlLCJ1c2VybmFtZSI6ImFkbWluNiJ9.uGL6XcNzb6srDCsuTX870rTNtcS6SUkc_sk5xGgVUcU', '2023-10-27 11:14:37', '2023-11-11 09:14:37', '2023-10-27 09:14:37', '2023-10-27 09:14:37', NULL);
COMMIT;

-- ----------------------------
-- Table structure for tb_user
-- ----------------------------
DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user` (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8 NOT NULL COMMENT '用户名',
  `mobile` varchar(20) CHARACTER SET utf8 NOT NULL COMMENT '手机号',
  `password` varchar(64) CHARACTER SET utf8 DEFAULT NULL COMMENT '密码',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`user_id`) USING BTREE,
  UNIQUE KEY `username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户';

-- ----------------------------
-- Records of tb_user
-- ----------------------------
BEGIN;
INSERT INTO `tb_user` (`user_id`, `username`, `mobile`, `password`, `create_time`) VALUES (1, 'mark', '13612345678', '8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918', '2017-03-23 22:37:41');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
