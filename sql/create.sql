CREATE TABLE IF NOT EXISTS weather
(
    id SERIAL,
    name TEXT NOT NULL,
    region TEXT NOT NULL,
    country TEXT NOT NULL,
    temp_c NUMERIC NOT NULL,
    temp_f NUMERIC NOT NULL,
    feelslike_c NUMERIC NOT NULL,
    feelslike_f NUMERIC NOT NULL,
    wind_mph NUMERIC NOT NULL,
    wind_kph NUMERIC NOT NULL,
    humidity SMALLINT NOT NULL,
    created_on TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,

    PRIMARY KEY (Id)
)