CREATE TABLE `object_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `object_uuid` char(12) NOT NULL,
  `user_uuid` char(12) NOT NULL COMMENT '用户uuid',
  `user_name` varchar(255) NOT NULL COMMENT '用户姓名',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '类型: 1系统, 2菜单, 3操作(接口)',
  `log_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '日志类型: 1添加, 2更新, 3删除',
  `log_summary` varchar(1000) NOT NULL DEFAULT '' COMMENT '日志描述',
  `log_data` varchar(1000) NOT NULL DEFAULT '' COMMENT '日志内容',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_user_uuid` (`user_uuid`),
  KEY `idx_type` (`type`),
  KEY `idx_log_type` (`log_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;