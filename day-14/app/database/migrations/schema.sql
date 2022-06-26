-- SQLite
CREATE TABLE IF NOT EXISTS `users` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `email` VARCHAR UNIQUE,
    `password` VARCHAR,
    `created_at` DATETIME DEFAULT (DATETIME('NOW')),
    `updated_at` DATETIME DEFAULT (DATETIME('NOW')),
    `deleted_at` DATETIME
);
CREATE TABLE IF NOT EXISTS `books` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `title` VARCHAR,
    `created_at` DATETIME DEFAULT (DATETIME('NOW')),
    `updated_at` DATETIME DEFAULT (DATETIME('NOW')),
    `deleted_at` DATETIME
);