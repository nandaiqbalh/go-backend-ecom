CREATE TABLE IF NOT EXISTS orders (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `userId` INT UNSIGNED NOT NULL,
    `total` DECIMAL(10, 2) NOT NULL,
    `status` ENUM('pending', 'completed', 'cancelled') NOT NULL DEFAULT 'pending',
    `address` VARCHAR(255) NOT NULL,
    `createdAt` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    FOREIGN KEY (`userId`) REFERENCES users(`id`) ON DELETE CASCADE
);