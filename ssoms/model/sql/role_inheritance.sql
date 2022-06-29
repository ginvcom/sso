CREATE TABLE `role_inheritance`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_uuid` char(12) NOT NULL,
  `inheritance_role_uuid` char(12) NOT NULL,
  `is_delete` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除: 0正常, 1删除',
  `create_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`),
  INDEX `idx_role_uuid`(`role_uuid`),
  INDEX `idx_inheritance_role_uuid`(`inheritance_role_uuid`),
  INDEX `uidx_role_nheritance`(`role_uuid`, `inheritance_role_uuid`),
  KEY `idx_is_delete` (`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;