 CREATE TABLE aqi_details
(
    city_name   TEXT,
    city_aqi     INT,
    created_at   TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);