CREATE TABLE `users` (
  `id` integer PRIMARY KEY,
  `username` varchar(255),
  `email` varchar(255),
  `phone_number` varchar(255),
  `role` varchar(255),
  `photo` varchar(255),
  `status` integer,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `companies` (
  `id` integer PRIMARY KEY,
  `product_id` integer,
  `name` varchar(255),
  `email` varchar(255) UNIQUE,
  `phone_number` varchar(255),
  `website` varchar(255),
  `logo` varchar(255),
  `agent_id` integer,
  `status` integer,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `products` (
  `id` integer PRIMARY KEY,
  `name` varchar(255),
  `status` integer,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `company_member` (
  `id` integer PRIMARY KEY,
  `user_id` integer,
  `company_id` integer,
  `role` varchar(255),
  `status` integer,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `uploads` (
  `id` integer PRIMARY KEY,
  `filename` varchar(255),
  `size` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `settings` (
  `id` integer PRIMARY KEY,
  `setting_key` varchar(255),
  `setting_value` text,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `invoices` (
  `id` integer PRIMARY KEY,
  `company_id` integer,
  `name` varchar(255),
  `status` integer,
  `payment_method` integer,
  `invoice_date` timestamp,
  `invoice_file` varchar(255),
  `receipt` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `permissions` (
  `id` integer PRIMARY KEY,
  `user_id` integer,
  `created_at` timestamp,
  `updated_at` timestamp
);

ALTER TABLE `companies` ADD FOREIGN KEY (`agent_id`) REFERENCES `users` (`id`);

ALTER TABLE `companies` ADD FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

ALTER TABLE `company_member` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `company_member` ADD FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`);

ALTER TABLE `permissions` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
