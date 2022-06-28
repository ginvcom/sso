CREATE TABLE `user_to_role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_uuid` int NOT NULL,
  `role_uuid` int NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `udx_user_to_role` (`user_uuid`,`role_uuid`),
  KEY `idx_user_uuid` (`user_uuid`),
  KEY `idx_role_uuid` (`role_uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;