-- -------ssoms的表结构----- --
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uuid` char(12) NOT NULL,
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `name` varchar(255) NOT NULL COMMENT '姓名',
  `mobile` varchar(20) NOT NULL COMMENT '手机号',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `salt` varchar(255) NOT NULL COMMENT '密码盐',
  `gender` tinyint unsigned NOT NULL DEFAULT 3 COMMENT '性别: 1男, 2女, 3未知',
  `birth` timestamp NOT NULL COMMENT '生日',
  `introduction` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '状态: 1正常',
  `is_delete` tinyint NOT NULL DEFAULT 0 COMMENT '是否删除: 0正常, 1删除',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `udx_uuid` (`uuid`),
  KEY `idx_name` (`name`),
  KEY `idx_mobile` (`mobile`),
  KEY `idx_status` (`status`),
  KEY `idx_is_delete` (`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `role_uuid` char(12) NOT NULL,
  `role_name` varchar(255) NOT NULL DEFAULT '',
  `summary` varchar(255) NOT NULL DEFAULT '',
  `is_delete` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除: 0正常, 1删除',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `udx_role_uuid` (`role_uuid`),
  KEY `idx_role_name` (`role_name`),
  KEY `idx_is_delete` (`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `user_to_role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_uuid` char(12) NOT NULL,
  `role_uuid` char(12) NOT NULL,
  `is_delete` tinyint NOT NULL DEFAULT 0 COMMENT '是否删除: 0正常, 1删除',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `udx_user_to_role` (`user_uuid`,`role_uuid`),
  KEY `idx_user_uuid` (`user_uuid`),
  KEY `idx_role_uuid` (`role_uuid`),
  KEY `idx_is_delete` (`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `object` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uuid` char(12) NOT NULL,
  `object_name` varchar(255) NOT NULL,
  `identifier` varchar(255) NOT NULL DEFAULT '',
  `key` varchar(255) NOT NULL COMMENT '操作对象的systemCode, 菜单的path, 操作的uri',
  `sort` int NOT NULL DEFAULT '0',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '类型: 1系统, 2菜单, 3操作(接口)',
  `sub_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '子分类',
  `extra` varchar(1024) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '扩展字段，建议封装成 JSON 字符串',
  `icon` varchar(255) NOT NULL DEFAULT '' COMMENT '图标',
  `status` tinyint(1) NOT NULL COMMENT '状态',
  `puuid` char(12) NOT NULL COMMENT '父级uuid',
  `top_key` varchar(255) NOT NULL DEFAULT '' COMMENT '操作对象的所属systemCode',
  `is_delete` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除: 0正常, 1删除',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `udx_uuid` (`uuid`),
  KEY `idx_object_name` (`object_name`),
  KEY `idx_key` (`key`),
  KEY `idx_status` (`status`),
  KEY `idx_type` (`type`),
  KEY `idx_sub_type` (`sub_type`),
  KEY `idx_puuid` (`puuid`),
  KEY `idx_is_delete` (`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `statistic` (
  `id` int NOT NULL AUTO_INCREMENT,
  `month` char(12) NOT NULL COMMENT '月份, YYYY-MM格式',
  `user_amount` int NOT NULL DEFAULT 0 COMMENT '用户数量',
  `role_amount` int NOT NULL DEFAULT 0 COMMENT '角色数量',
  `system_amount` int NOT NULL DEFAULT 0 COMMENT '系统数量',
  `menu_amount` int NOT NULL DEFAULT 0 COMMENT '菜单数量',
  `action_amount` int NOT NULL DEFAULT 0 COMMENT '操作数量',
  `permission_amount` int NOT NULL DEFAULT 0 COMMENT '授权数量',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `udx_month` (`month`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `object_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `object_uuid` char(12) NOT NULL,
  `user_uuid` char(12) NOT NULL COMMENT '用户uuid',
  `user_name` varchar(255) NOT NULL COMMENT '用户姓名',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '类型: 1系统, 2菜单, 3操作(接口)',
  `log_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '日志类型: 1添加, 2更新, 3删除, 4导入添加, 5导入更新',
  `log_summary` varchar(1000) NOT NULL DEFAULT '' COMMENT '日志描述',
  `log_data` varchar(1000) NOT NULL DEFAULT '' COMMENT '日志内容',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_user_uuid` (`user_uuid`),
  KEY `idx_type` (`type`),
  KEY `idx_log_type` (`log_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `permission` (
  `id` int NOT NULL AUTO_INCREMENT,
  `role_uuid` char(12) NOT NULL,
  `object_uuid` char(12) NOT NULL,
  `type` tinyint(1) NOT NULL,
  `top_key` varchar(255) NOT NULL,
  `is_delete` tinyint(1) NOT NULL DEFAULT 0,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `udx_permission` (`role_uuid`,`object_uuid`),
  KEY `idx_role_uuid` (`role_uuid`),
  KEY `idx_type` (`type`),
  KEY `idx_top_key` (`top_key`),
  KEY `idx_is_delete` (`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ------------------------ --
-- -------ssoms的数据------- --
INSERT INTO `user` (`uuid`, `avatar`, `name`, `mobile`, `password`, `salt`, `gender`, `birth`, `introduction`, `status`, `is_delete`, `create_time`, `update_time`) VALUES ('0kfxnj6x0spb', '6165845514057090.png', '超管', '12345000001', 'c797f80662508b1e53b6e45800638945', 'qt0b79', 1, '2000-01-13 00:00:00', '你是我的好呀好朋友1', 1, 0, '2022-06-29 10:13:58', '2022-11-17 18:12:49');
INSERT INTO `role`(`role_uuid`, `role_name`, `summary`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', '管理员角色', '这是概述1', 0, '2022-06-30 07:14:05', '2022-07-22 10:20:50');
INSERT INTO `user_to_role` (`user_uuid`, `role_uuid`, `is_delete`, `create_time`, `update_time`) VALUES ('0kfxnj6x0spb', '82fbnyvb5vpb', 0, '2022-07-01 09:41:56', '2022-08-20 11:58:48');

INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('8yzo4kzfucsb', '菜单&操作', '', '/object', 0, 2, 1, '', 'icon-unorderedlist', 1, '', 0, 'ssoms', '2022-07-25 10:32:43', '2022-11-18 18:11:52');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('krq246gvhjub', '个人中心', '', '/profile', 0, 2, 3, '', '', 1, '', 0, 'ssoms', '2022-08-16 11:33:17', '2022-08-16 19:33:32');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('kreppg8md1sb', '单点管理系统', 'http://localhost:3000', 'ssoms', 100, 1, 0, '', '6165846916250998.png', 1, '', 0, 'ssoms', '2022-07-22 05:24:54', '2022-11-18 10:27:27');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('wsb9wzmy9u3c', '后台系统', '', '/system', 0, 2, 1, '', 'icon-gold', 1, '', 0, 'ssoms', '2022-11-18 07:18:12', '2022-11-18 18:11:45');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('g89tkuzdfdsb', '用户管理', '', '/user', 0, 2, 1, '', 'icon-user', 1, '', 0, 'ssoms', '2022-07-25 14:27:30', '2022-08-20 22:39:16');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('sd1w57qgfdsb', '角色管理', '', '/role', 0, 2, 1, '', 'icon-team', 1, '', 0, 'ssoms', '2022-07-25 14:28:21', '2022-08-20 22:39:06');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('kfl2fmf3jdsb', '对象列表', '', '/object', 0, 3, 1, '', '', 1, '8yzo4kzfucsb', 0, 'ssoms', '2022-07-25 15:09:03', '2022-08-06 07:53:33');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('wo8oeg4y1fsb', '操作详情', '', '/object/:uuid', 0, 3, 1, '', '', 1, '8yzo4kzfucsb', 0, 'ssoms', '2022-07-26 01:23:58', '2022-08-06 07:53:33');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('4fqt9v682fsb', '添加对象', '', '/object', 0, 3, 2, '', '', 1, '8yzo4kzfucsb', 0, 'ssoms', '2022-07-26 01:27:06', '2022-08-06 07:53:33');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('sln6ra37gfsb', '编辑对象', '', '/object/:uuid', 0, 3, 3, '', '', 1, '8yzo4kzfucsb', 0, 'ssoms', '2022-07-26 04:03:42', '2022-08-06 07:53:33');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('wsvssh0gpfsb', '删除对象', '', '/object/:uuid', 0, 3, 5, '', '', 1, '8yzo4kzfucsb', 0, 'ssoms', '2022-07-26 05:47:22', '2022-08-06 07:53:33');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('k7vcb745vcsb', '对象授权', '', '/permission', 0, 2, 3, '', '', 1, '8yzo4kzfucsb', 0, 'ssoms', '2022-07-25 10:40:32', '2022-08-06 07:53:33');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('c9bcohdat75c', '导入对象', '', '/object/import', 0, 3, 2, '', '', 1, '8yzo4kzfucsb', 'ssoms', 0, '2022-12-02 04:29:26', '2022-12-02 04:29:26');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('knffjw6tup3c', '角色授权', '', '/permission/role/:roleUUID/grant', 0, 3, 3, '', '', 1, 'c5gxcduqwpsb', 0, 'ssoms', '2022-11-17 01:34:06', '2022-11-17 09:34:53');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('kbq34p6ergsb', '添加用户', '', '/user', 0, 3, 2, '', '', 1, 'g89tkuzdfdsb', 0, 'ssoms', '2022-07-26 12:52:47', '2022-08-06 07:53:33');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('47cp2hq92r3c', '分配角色', '', '/assignRole', 0, 3, 2, '', '', 1, 'g89tkuzdfdsb', 0, 'ssoms', '2022-11-17 09:41:17', '2022-11-17 09:41:17');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('g07yw2xhrgsb', '删除用户', '', '/user/:uuid', 0, 3, 5, '', '', 1, 'g89tkuzdfdsb', 0, 'ssoms', '2022-07-26 12:53:56', '2022-08-06 07:53:33');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('8m2rqsogrgsb', '编辑用户', '', '/user/:uuid', 0, 3, 3, '', '', 1, 'g89tkuzdfdsb', 0, 'ssoms', '2022-07-26 12:53:33', '2022-08-06 07:53:33');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('oi2htj4hkgsb', '用户详情', '', '/user/:uuid', 0, 3, 1, '', '', 1, 'g89tkuzdfdsb', 0, 'ssoms', '2022-07-26 11:35:13', '2022-08-06 07:53:33');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('4700r9lkjgsb', '用户列表', '', '/user', 0, 3, 1, '', '', 1, 'g89tkuzdfdsb', 0, 'ssoms', '2022-07-26 11:25:05', '2022-08-06 07:53:33');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('0oqc6a911r3c', '用户移出角色', '', '/role/:roleUUID/deassignUser', 0, 3, 4, '', '', 1, 'g8t8ym1xfdsb', 0, 'ssoms', '2022-11-17 09:27:26', '2022-11-17 09:27:26');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('0ggb8kdx0r3c', '角色分配的用户列表', '', '/assignedUsers', 0, 3, 1, '', '', 1, 'g8t8ym1xfdsb', 0, 'ssoms', '2022-11-17 09:26:13', '2022-11-17 09:26:13');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('cl087dmxrgsb', '角色列表', '', '/role', 0, 3, 1, '', '', 1, 'sd1w57qgfdsb', 0, 'ssoms', '2022-07-26 12:58:50', '2022-08-06 07:53:33');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('c5gxcduqwpsb', '角色授权', '', '/role/permission', 0, 2, 3, '', '', 1, 'sd1w57qgfdsb', 0, 'ssoms', '2022-07-29 02:24:50', '2022-08-06 07:53:33');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('omxb5f1bsgsb', '删除角色', '', '/role/:roleUUID', 0, 3, 5, '', '', 1, 'sd1w57qgfdsb', 0, 'ssoms', '2022-07-26 13:03:00', '2022-08-06 07:53:33');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('gs8m223asgsb', '编辑角色', '', '/role/:roleUUID', 0, 3, 3, '', '', 1, 'sd1w57qgfdsb', 0, 'ssoms', '2022-07-26 13:02:43', '2022-08-06 07:53:33');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('cl86nuz7sgsb', '添加角色', '', '/role', 0, 3, 2, '', '', 1, 'sd1w57qgfdsb', 0, 'ssoms', '2022-07-26 13:02:03', '2022-08-06 07:53:33');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('4vyuwln0sgsb', '角色详情', '', '/role/:roleUUID', 0, 3, 1, '', '', 1, 'sd1w57qgfdsb', 0, 'ssoms', '2022-07-26 12:59:46', '2022-08-06 07:53:33');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('g8t8ym1xfdsb', '角色分配的用户列表', '', '/role/assignedUsers', 0, 2, 3, '', '', 1, 'sd1w57qgfdsb', 0, 'ssoms', '2022-07-25 14:33:26', '2022-08-19 09:08:32');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('c5g1wyqeau3c', '系统列表', '', '/system', 0, 3, 1, '', '', 1, 'wsb9wzmy9u3c', 0, 'ssoms', '2022-11-18 07:23:13', '2022-11-18 07:23:13');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('o6tlp8vfau3c', '系统详情', '', '/system/:uuid', 0, 3, 1, '', '', 1, 'wsb9wzmy9u3c', 0, 'ssoms', '2022-11-18 07:23:34', '2022-11-18 07:23:34');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('o2ejmzshau3c', '添加系统', '', '/system', 0, 3, 2, '', '', 1, 'wsb9wzmy9u3c', 0, 'ssoms', '2022-11-18 07:24:10', '2022-11-18 07:24:10');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('4vi6v7bjau3c', '编辑系统', '', '/system/:uuid', 0, 3, 3, '', '', 1, 'wsb9wzmy9u3c', 0, 'ssoms', '2022-11-18 07:24:38', '2022-11-18 07:24:38');
INSERT INTO `object` (`uuid`, `object_name`, `identifier`, `key`, `sort`, `type`, `sub_type`, `extra`, `icon`, `status`, `puuid`, `is_delete`, `top_key`, `create_time`, `update_time`) VALUES ('0s5ra4ekau3c', '删除系统', '', '/system/:uuid', 0, 3, 5, '', '', 1, 'wsb9wzmy9u3c', 0, 'ssoms', '2022-11-18 07:24:58', '2022-11-18 07:24:58');

INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', '8yzo4kzfucsb', 2, 'ssoms', 0, '2022-08-15 12:14:29', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', 'k7vcb745vcsb', 2, 'ssoms', 0, '2022-08-15 12:14:29', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', 'kfl2fmf3jdsb', 3, 'ssoms', 0, '2022-08-15 12:14:29', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', 'wo8oeg4y1fsb', 3, 'ssoms', 0, '2022-08-15 12:14:29', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', '4fqt9v682fsb', 3, 'ssoms', 0, '2022-08-15 12:14:29', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', 'sln6ra37gfsb', 3, 'ssoms', 0, '2022-08-15 12:14:29', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', 'wsvssh0gpfsb', 3, 'ssoms', 0, '2022-08-15 12:14:29', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', 'c9bcohdat75c', 3, 'ssoms', 0, '2022-12-02 04:29:37', '2022-12-02 12:29:38');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', 'sd1w57qgfdsb', 2, 'ssoms', 0, '2022-08-15 12:14:47', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', 'g8t8ym1xfdsb', 2, 'ssoms', 0, '2022-08-15 12:14:47', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', 'c5gxcduqwpsb', 2, 'ssoms', 0, '2022-08-15 12:14:47', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', 'omxb5f1bsgsb', 3, 'ssoms', 0, '2022-08-15 12:14:53', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', '4vyuwln0sgsb', 3, 'ssoms', 0, '2022-08-15 12:14:57', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', 'cl087dmxrgsb', 3, 'ssoms', 0, '2022-08-15 12:15:00', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', 'cl86nuz7sgsb', 3, 'ssoms', 0, '2022-08-15 12:15:00', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', 'gs8m223asgsb', 3, 'ssoms', 0, '2022-08-15 12:15:00', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', 'g89tkuzdfdsb', 2, 'ssoms', 0, '2022-08-15 12:15:16', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', '4700r9lkjgsb', 3, 'ssoms', 0, '2022-08-15 12:15:54', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', 'oi2htj4hkgsb', 3, 'ssoms', 0, '2022-08-15 12:15:54', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', 'kbq34p6ergsb', 3, 'ssoms', 0, '2022-08-15 12:15:54', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', '8m2rqsogrgsb', 3, 'ssoms', 0, '2022-08-15 12:15:54', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', 'g07yw2xhrgsb', 3, 'ssoms', 0, '2022-08-15 12:15:54', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', 'krq246gvhjub', 2, 'ssoms', 0, '2022-08-16 11:42:41', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', 'knffjw6tup3c', 3, 'ssoms', 0, '2022-11-17 01:37:25', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', '0ggb8kdx0r3c', 3, 'ssoms', 0, '2022-11-17 09:26:47', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', '0oqc6a911r3c', 3, 'ssoms', 0, '2022-11-17 09:27:35', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', '47cp2hq92r3c', 3, 'ssoms', 0, '2022-11-17 09:45:25', '2022-11-18 07:25:19');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', 'wsb9wzmy9u3c', 2, 'ssoms', 0, '2022-11-18 07:25:19', '2022-11-18 15:25:20');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', 'c5g1wyqeau3c', 3, 'ssoms', 0, '2022-11-18 07:25:19', '2022-11-18 15:25:20');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', 'o6tlp8vfau3c', 3, 'ssoms', 0, '2022-11-18 07:25:19', '2022-11-18 15:25:20');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', 'o2ejmzshau3c', 3, 'ssoms', 0, '2022-11-18 07:25:19', '2022-11-18 15:25:20');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', '4vi6v7bjau3c', 3, 'ssoms', 0, '2022-11-18 07:25:19', '2022-11-18 15:25:20');
INSERT INTO `permission` (`role_uuid`, `object_uuid`, `type`, `top_key`, `is_delete`, `create_time`, `update_time`) VALUES ('82fbnyvb5vpb', '0s5ra4ekau3c', 3, 'ssoms', 0, '2022-11-18 07:25:19', '2022-11-18 15:25:20');