CREATE TABLE IF NOT EXISTS `country` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `country_no` VARCHAR(50) NOT NULL UNIQUE,
    `name` VARCHAR(255) NOT NULL,
    `code` VARCHAR(10) NOT NULL UNIQUE,
    `currency` VARCHAR(10) NOT NULL,
    `status` ENUM('active', 'inactive') NOT NULL DEFAULT 'active',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


CREATE TRIGGER IF NOT EXISTS `before_insert_country` BEFORE INSERT ON `country`
    FOR EACH ROW
    BEGIN
        IF NEW.country_no IS NULL OR NEW.country_no = '' THEN
            SET NEW.country_no = nanoid();
            INSERT INTO `audit` (`action`, `table_name`, `record_no`)
            VALUES ('create', 'country', NEW.country_no);
        END IF;

    END;

