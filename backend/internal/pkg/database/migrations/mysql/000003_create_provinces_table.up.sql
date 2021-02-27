CREATE TABLE `provinces` (
    `id` bigint unsigned NOT NULL UNIQUE AUTO_INCREMENT,

    `name` varchar (50) NOT NULL,
    `department_id` bigint unsigned NOT NULL,

    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `deleted_at` datetime(3) DEFAULT NULL,
    
    PRIMARY KEY (`id`),
    KEY `idx_provinces_deleted_at` (`deleted_at`)
); -- ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
