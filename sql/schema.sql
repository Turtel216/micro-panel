CREATE DATABASE IF NOT EXISTS micropanel;

USE micropanel;

CREATE TABLE `products` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `image` varchar(255) NOT NULL,
  `category` varchar(255) NOT NULL,
  `description` text,
  `rating` decimal(3,2) NOT NULL,
  `num_reviews` int NOT NULL DEFAULT 0,
  `price` decimal(10,2) NOT NULL,
  `count_in_stock` int NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime
);

CREATE TABLE `users` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `is_admin` bool NOT NULL DEFAULT false,
  UNIQUE(email)
);

CREATE TABLE `orders` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `payment_method` varchar(255) NOT NULL,
  `tax_price` decimal(10,2) NOT NULL,
  `shipping_price` decimal(10,2) NOT NULL,
  `total_price` decimal(10,2) NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT `user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);

CREATE TABLE `order_items` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `order_id` int NOT NULL,
  `product_id` int NOT NULL,
  `quantity` int NOT NULL,
  `price` decimal(10,2) NOT NULL,
  FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`),
  FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
);

CREATE TABLE `sessions` (
  `id` varchar(255) PRIMARY KEY NOT NULL,
  `user_email` varchar(255) NOT NULL,
  `refresh_token` varchar(512) NOT NULL,
  `is_revoked` bool NOT NULL DEFAULT false,
  `created_at` datetime DEFAULT (now()),
  `expires_at` datetime
);
