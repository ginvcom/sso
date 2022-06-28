CREATE TABLE `role_nheritance`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_uuid` char(12) NOT NULL,
  `inheritance_role_uuid` char(12) NOT NULL,
  `create_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`),
  INDEX `idx_role_uuid`(`role_uuid`),
  INDEX `idx_inheritance_role_uuid`(`inheritance_role_uuid`),
  INDEX `uidx_role_nheritance`(`role_uuid`, `inheritance_role_uuid`)
);