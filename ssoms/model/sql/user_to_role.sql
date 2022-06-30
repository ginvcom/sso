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