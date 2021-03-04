CREATE TABLE `departments` (
    `id` bigint unsigned NOT NULL UNIQUE AUTO_INCREMENT,

    `name` varchar (20) NOT NULL UNIQUE,

    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `deleted_at` datetime(3) DEFAULT NULL,
    
    PRIMARY KEY (`id`),
    KEY `idx_departments_deleted_at` (`deleted_at`)
); -- ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
