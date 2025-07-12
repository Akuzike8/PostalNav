CREATE TABLE IF NOT EXISTS `apikey` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `apikey_no` VARCHAR(255) NOT NULL UNIQUE,
    `key` VARCHAR(255) NOT NULL UNIQUE,
    `status` ENUM('active', 'inactive','revoked') NOT NULL DEFAULT 'active',
    `merchant` VARCHAR(50) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


CREATE TRIGGER IF NOT EXISTS `before_insert_apikey` BEFORE INSERT ON `apikey`
    FOR EACH ROW
    BEGIN
        IF NEW.apikey_no IS NULL OR NEW.apikey_no = '' THEN
            SET NEW.apikey_no = nanoid();
        END IF;
        IF NEW.key IS NULL OR NEW.key = '' THEN
            SET NEW.key = UUID();
            INSERT INTO `audit` (`action`, `table_name`, `record_no`)
            VALUES ('create', 'apikey', NEW.apikey_no);

        END IF;
    END;


