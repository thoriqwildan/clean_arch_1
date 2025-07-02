CREATE TABLE `payment_methods` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) UNIQUE,
    `desc` varchar(255) NULL,
    `order_num` int DEFAULT 1,
    `user_action` varchar(255),
    `created_at` datetime(3) NOT NULL,
    `updated_at` datetime(3) NOT NULL,
    `code` varchar(255) NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `payment_channels` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `payment_method_id` bigint unsigned,
    `code` varchar(255) UNIQUE,
    `name` varchar(255) UNIQUE,
    `icon_url` varchar(255) NULL,
    `order_num` bigint NULL DEFAULT 1,
    `lib_name` varchar(255) NULL,
    `user_action` varchar(255),
    `created_at` datetime(3) NOT NULL,
    `updated_at` datetime(3) NOT NULL,
    `mdr` varchar(255) DEFAULT '0',
    `fixed_fee` double DEFAULT 0,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_payment_methods_channels` FOREIGN KEY (`payment_method_id`) REFERENCES `payment_methods` (`id`) ON UPDATE CASCADE ON DELETE CASCADE
);