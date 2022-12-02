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