USE `uam`;

/* 角色 */
DROP TABLE IF EXISTS `uam_role`;
CREATE TABLE `uam_role`
(
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `client_id` int NOT NULL DEFAULT 0 COMMENT '所属客户端ID',
  `name` varchar(64) NOT NULL COMMENT '角色名称',
  `desc` varchar(256) NOT NULL DEFAULT '' COMMENT '角色描述',
  `editable` tinyint NOT NULL DEFAULT 1 COMMENT '是否允许通过后台编辑',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uidx_key` (`client_id`, `name`)
) ENGINE=InnoDB AUTO_INCREMENT=127 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='角色';
