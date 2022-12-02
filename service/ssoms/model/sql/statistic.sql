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