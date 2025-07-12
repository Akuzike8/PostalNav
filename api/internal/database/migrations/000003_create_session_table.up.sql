CREATE TABLE IF NOT EXISTS `session` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `session_no` VARCHAR(50) NOT NULL UNIQUE,
    `user_no` VARCHAR(50) NOT NULL,
    `device` VARCHAR(50) NOT NULL,
    `token` VARCHAR(50) NOT NULL,
    `status` ENUM('active', 'inactive','expired') NOT NULL DEFAULT 'active',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


CREATE TRIGGER IF NOT EXISTS `before_insert_session` BEFORE INSERT ON `session`
    FOR EACH ROW
    BEGIN
        IF NEW.session_no IS NULL OR NEW.session_no = '' THEN
            SET NEW.session_no = nanoid();
            INSERT INTO `audit` (`action`, `table_name`, `record_no`)
            VALUES ('create', 'session', NEW.session_no);
        END IF;
    END;


