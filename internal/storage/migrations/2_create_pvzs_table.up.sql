CREATE TYPE city_enum AS ENUM (
    'Москва',
    'Санкт-Петербург',
    'Казань'
);

CREATE TABLE pvzs (
    id SERIAL PRIMARY KEY,
    registration_date TIMESTAMP NOT NULL DEFAULT NOW(),
    city city_enum NOT NULL
);