CREATE TABLE IF NOT EXISTS postalcodes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    country_code VARCHAR(10) NOT NULL,
    postcode VARCHAR(10) NOT NULL,
    place_name VARCHAR(100) NOT NULL,
    state_region VARCHAR(100) NOT NULL,
    state_region_code VARCHAR(10) NOT NULL,
    county_province VARCHAR(100) NOT NULL,
    county_province_code VARCHAR(10) NOT NULL,
    community VARCHAR(100) NOT NULL,
    community_code VARCHAR(10) NOT NULL,
    latitude VARCHAR(10) NOT NULL,
    longtitude VARCHAR(10) NOT NULL,
    accuracy VARCHAR(100) NOT NULL,
    h3Id VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
