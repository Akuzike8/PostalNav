CREATE TABLE IF NOT EXISTS `audit` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `action` ENUM('create', 'update', 'delete') NOT NULL,
    `table_name` VARCHAR(255) NOT NULL,
    `record_no` VARCHAR(50) NOT NULL,
    `old_value` TEXT,
    `new_value` TEXT,
    `audit_no` VARCHAR(255) NOT NULL UNIQUE,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TRIGGER IF NOT EXISTS `before_insert_audit` BEFORE INSERT ON `audit`
    FOR EACH ROW
    BEGIN
        IF NEW.audit_no IS NULL OR NEW.audit_no = '' THEN
            SET NEW.audit_no = nanoid();
        END IF;
    END;

