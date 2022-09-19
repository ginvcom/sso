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