CREATE DATABASE IF NOT EXISTS `alta_online_shop` CHARACTER set utf8mb4 COLLATE utf8mb4_general_ci;
CREATE TABLE `product_types`(
    `id` BigInt(0) AUTO_INCREMENT NOT NULL,
    `name` VarChar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `created_at` DateTime NOT NULL,
    `updated_at` DateTime NOT NULL,
    PRIMARY KEY (`id`)
);
CREATE TABLE `operators`(
    `id` BigInt(0) AUTO_INCREMENT NOT NULL,
    `name` VarChar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `created_at` DateTime NOT NULL,
    `updated_at` DateTime NOT NULL,
    PRIMARY KEY (`id`)
);
CREATE TABLE `payment_methods`(
    `id` BigInt(0) AUTO_INCREMENT NOT NULL,
    `name` VarChar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `created_at` DateTime NOT NULL,
    `updated_at` DateTime NOT NULL,
    PRIMARY KEY (`id`)
);
CREATE TABLE `users`(
    `id` BigInt(0) AUTO_INCREMENT NOT NULL,
    `name` VarChar(125) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `address` Text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `birth_date` Date NOT NULL,
    `status` Enum('BELUM KAWIN', 'KAWIN TERCATAT', 'KAWIN BELUM TERCATAT', 'CERAI HIDUP', 'CERAI MATI') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `gender` Enum('LAKI-LAKI', 'PEREMPUAN') CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `created_at` DateTime NOT NULL,
    `updated_at` DateTime NOT NULL,
    PRIMARY KEY (`id`)
);
CREATE TABLE `products`(
    `id` BigInt(0) AUTO_INCREMENT NOT NULL,
    `name` VarChar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `price` BigInt(0) NOT NULL,
    `product_type_id` BigInt(0) NOT NULL,
    `operator_id` BigInt(0) NOT NULL,
    `created_at` DateTime NOT NULL,
    `updated_at` DateTime NOT NULL,
    PRIMARY KEY (`id`)
);
CREATE TABLE `product_descriptions`(
    `id` BigInt(0) AUTO_INCREMENT NOT NULL,
    `description` Text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `created_at` DateTime NOT NULL,
    `updated_at` DateTime NOT NULL,
    PRIMARY KEY (`id`)
);
CREATE TABLE `transactions`(
    `id` BigInt(0) AUTO_INCREMENT NOT NULL,
    `user_id` BigInt(0) NOT NULL,
    `payment_method_id` BigInt(0) NOT NULL,
    `total_qty` INT(0) NOT NULL,
    `total_price` BigInt(0) NOT NULL,
    `created_at` DateTime NOT NULL,
    `updated_at` DateTime NOT NULL,
    PRIMARY KEY (`id`)
);
CREATE TABLE `transaction_details`(
    `id` BigInt(0) AUTO_INCREMENT NOT NULL,
    `transaction_id` BigInt(0) NOT NULL,
    `product_id` BigInt(0) NOT NULL,
    `price` BigInt(0) NOT NULL,
    `qty` BigInt(0) NOT NULL,
    `created_at` DateTime NOT NULL,
    `updated_at` DateTime NOT NULL,
    PRIMARY KEY (`id`)
);
CREATE TABLE `kurir`(
    `id` BigInt(0) AUTO_INCREMENT NOT NULL,
    `name` VarChar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `created_at` DateTime NOT NULL,
    `updated_at` DateTime NOT NULL,
    PRIMARY KEY (`id`)
);
ALTER TABLE `kurir`
ADD `ongkos_dasar` BIGINT(0) NOT NULL
AFTER `name`;
ALTER TABLE `kurir`
    RENAME TO `shipping`;
DROP TABLE `shipping`;
ALTER TABLE `transactions`
ADD CONSTRAINT `lnk_users_transactions` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE Cascade ON UPDATE Cascade;
ALTER TABLE `transactions`
ADD CONSTRAINT `lnk_payment_methods_transactions` FOREIGN KEY (`payment_method_id`) REFERENCES `payment_methods`(`id`) ON DELETE Cascade ON UPDATE Cascade;
ALTER TABLE `transaction_details`
ADD CONSTRAINT `lnk_transactions_transaction_details` FOREIGN KEY (`transaction_id`) REFERENCES `transactions`(`id`) ON DELETE Cascade ON UPDATE Cascade;
ALTER TABLE `transaction_details`
ADD CONSTRAINT `lnk_products_transaction_details` FOREIGN KEY (`product_id`) REFERENCES `products`(`id`) ON DELETE Cascade ON UPDATE Cascade;
ALTER TABLE `products`
ADD CONSTRAINT `lnk_operators_products` FOREIGN KEY (`operator_id`) REFERENCES `operators`(`id`) ON DELETE Cascade ON UPDATE Cascade;
ALTER TABLE `products`
ADD CONSTRAINT `lnk_product_types_products` FOREIGN KEY (`product_type_id`) REFERENCES `product_types`(`id`) ON DELETE Cascade ON UPDATE Cascade;
ALTER TABLE `product_descriptions`
ADD CONSTRAINT `lnk_products_product_descriptions` FOREIGN KEY (`id`) REFERENCES `products`(`id`) ON DELETE Cascade ON UPDATE Cascade;
CREATE TABLE `payment_method_descriptions`(
    `id` BigInt(0) AUTO_INCREMENT NOT NULL,
    `description` Text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `created_at` DateTime NOT NULL,
    `updated_at` DateTime NOT NULL,
    PRIMARY KEY (`id`)
);
ALTER TABLE `payment_methods`
ADD CONSTRAINT `lnk_payment_methods_payment_method_descriptions` FOREIGN KEY (`id`) REFERENCES `payment_methods_descriptions`(`id`) ON DELETE Cascade ON UPDATE Cascade;
CREATE TABLE `user_addresses`(
    `id` BigInt(0) AUTO_INCREMENT NOT NULL,
    `user_id` BigInt(0) NOT NULL,
    `address` Text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `created_at` DateTime NOT NULL,
    `updated_at` DateTime NOT NULL,
    PRIMARY KEY (`id`)
);
ALTER TABLE `user_addresses`
ADD CONSTRAINT `lnk_users_user_addresses` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE Cascade ON UPDATE Cascade;
CREATE TABLE `user_payment_method_details` (
    `user_id` BIGINT(0) NOT NULL,
    `payment_method_id` BIGINT(0) NOT NULL
);
ALTER TABLE `users` DROP COLUMN `address`;
ALTER TABLE `user_payment_method_details`
ADD CONSTRAINT `lnk_users_user_payment_method_details` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE Cascade ON UPDATE Cascade;
ALTER TABLE `user_payment_method_details`
ADD CONSTRAINT `lnk_users_payment_methods_payment_method_details` FOREIGN KEY (`payment_method_id`) REFERENCES `payment_methods`(`id`) ON DELETE Cascade ON UPDATE Cascade;