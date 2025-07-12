CREATE TABLE IF NOT EXISTS `user` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_no` VARCHAR(50) NOT NULL UNIQUE,
    `email` VARCHAR(50) NOT NULL UNIQUE,
    `phone` VARCHAR(20) NOT NULL UNIQUE,
    `firstname`  VARCHAR(255) NOT NULL,
    `lastname`  VARCHAR(255) NOT NULL,
    `type` ENUM('service', 'normal') NOT NULL DEFAULT 'normal',
    `status` ENUM('active', 'inactive') NOT NULL DEFAULT 'active',
    `last_active` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


CREATE TRIGGER IF NOT EXISTS `before_insert_user` BEFORE INSERT ON `user`
    FOR EACH ROW
    BEGIN
        IF NEW.user_no IS NULL OR NEW.user_no = '' THEN
            SET NEW.user_no = nanoid();
            INSERT INTO `audit` (`action`, `table_name`, `record_no`)
            VALUES ('create', 'user', NEW.user_no);
        END IF;
    END;

