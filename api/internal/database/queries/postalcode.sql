
-- name: GetPostalCode :one
SELECT * FROM postcodes
WHERE id = ? LIMIT 1;

-- name: GetPostalCodes :many
SELECT * FROM postcodes
ORDER BY place_name;

-- name: GetPostalCodeByH3Id :one
SELECT * FROM postcodes
WHERE h3Id = ? LIMIT 1;

-- name: GetPostalCodesByCountry :many
SELECT * FROM postcodes
WHERE country_code = ?  ORDER BY place_name;

-- name: GetPostalCodesPaginated :many
SELECT * FROM postcodes
ORDER BY place_name
LIMIT ? OFFSET ?;

-- name: GetPostalCodesByCountryPaginated :many
SELECT * FROM postcodes
WHERE country_code = ?
ORDER BY place_name
LIMIT ? OFFSET ?;

-- name: CreatePostalCode :exec
INSERT INTO postcodes (
    country_code, postcode, place_name, state_region, state_region_code,
    county_province, county_province_code, community, community_code,
    latitude, longtitude, accuracy, h3Id
) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdatePostalCode :exec
UPDATE postcodes
SET country_code = ?, postcode = ?, place_name = ?, state_region = ?,
    state_region_code = ?, county_province = ?, county_province_code = ?,
    community = ?, community_code = ?, latitude = ?, longtitude = ?,
    accuracy = ?, h3Id = ?
WHERE id = ?;

-- name: DeletePostalCode :exec
DELETE FROM postcodes
WHERE id = ?;

-- name: SearchPostalCode :many
SELECT * FROM postcodes
WHERE place_name LIKE ? OR state_region LIKE ? OR county_province LIKE ?
OR community LIKE ? OR h3Id LIKE ?;

-- name: SearchPostalCodePaginated :many
SELECT * FROM postcodes
WHERE place_name LIKE ? OR state_region LIKE ? OR county_province LIKE ?
OR community LIKE ? OR h3Id LIKE ?
ORDER BY place_name
LIMIT ? OFFSET ?;

