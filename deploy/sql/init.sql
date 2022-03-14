CREATE TABLE `leaf` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `domain_id` int(10) DEFAULT 1 COMMENT '业务id',
  `max_id` bigint(20) unsigned DEFAULT 0 COMMENT '最大id',
  `step` int(10) unsigned DEFAULT 0 COMMENT '步长',
  `status` tinyint(4) COLLATE utf8mb4_bin DEFAULT 1 COMMENT '状态',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_domain_id` (`domain_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='自增序列号';