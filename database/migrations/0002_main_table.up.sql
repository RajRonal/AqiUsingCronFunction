DROP TABLE  aqi_details;
CREATE TABLE  aqi_details
(
    city_name   TEXT,
    ozone_value DOUBLE PRECISION,
    ozone_last_updated  timestamp,
    ozone_units   TEXT,
    co2_value    DOUBLE PRECISION,
    co2_last_updated TIMESTAMP,
    co2_units    TEXT,
    ch4_value     DOUBLE PRECISION,
    ch4_last_updated  TIMESTAMP,
    ch4_unit           TEXT,
    water_vapour_value  INT,
    water_vapour_last_updated  TIMESTAMP,
    water_vapour_units    TEXT
)