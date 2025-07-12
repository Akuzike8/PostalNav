

DROP FUNCTION IF EXISTS nanoid_custom;
CREATE FUNCTION nanoid_custom(
    size INT,
    alphabet TEXT
)
    RETURNS TEXT
    LANGUAGE SQL
    NOT DETERMINISTIC
    SQL SECURITY INVOKER
    READS SQL DATA
BEGIN
    DECLARE idBuilder TEXT DEFAULT '';
    DECLARE counter INT DEFAULT 0;
    DECLARE alphabetIndex INT;
    DECLARE alphabetArray VARCHAR(64);
    DECLARE alphabetLength INT;
    DECLARE mask INT;
    DECLARE step INT;

    SET size = IFNULL(size, 21);
    SET alphabet = IFNULL(alphabet, '_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ');

    SET alphabetArray = alphabet;
    SET alphabetLength = CHAR_LENGTH(alphabet);
    SET mask = POW(2, FLOOR(LOG(alphabetLength - 1) / LOG(2))) - 1;
    SET step = CEILING(1.6 * mask * size / alphabetLength);

    create_loop:
    WHILE CHAR_LENGTH(idBuilder) < size
        DO
            SET counter = 0;
            add_character:
            WHILE counter < step
                DO
                    SET alphabetIndex = MOD(CONV(HEX(RANDOM_BYTES(1)), 16, 10), alphabetLength) + 1;
                    IF alphabetIndex <= alphabetLength THEN
                        IF CHAR_LENGTH(idBuilder) < size THEN
                            SET idBuilder = CONCAT(idBuilder, SUBSTRING(alphabetArray, alphabetIndex, 1));
                        ELSE
                            LEAVE create_loop;
                        END IF;
                    END IF;
                    SET counter = counter + 1;
                END WHILE;
        END WHILE;

    RETURN idBuilder;
END;

DROP FUNCTION IF EXISTS nanoid_simple;
CREATE FUNCTION nanoid_simple(
    size INT
)
    RETURNS TEXT
    LANGUAGE SQL
    NOT DETERMINISTIC
    SQL SECURITY INVOKER
    READS SQL DATA
BEGIN
    SET size = IFNULL(size, 21);
    RETURN nanoid_custom(size, '_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ');
END;

DROP FUNCTION IF EXISTS nanoid;
CREATE FUNCTION nanoid()
    RETURNS VARCHAR(21)
    LANGUAGE SQL
    NOT DETERMINISTIC
    SQL SECURITY INVOKER
    READS SQL DATA
BEGIN
    RETURN nanoid_simple(21);
END;
