CREATE TYPE product_type_enum AS ENUM (
    'одежда',
    'обувь',
    'электроника'
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    date_time TIMESTAMP NOT NULL DEFAULT NOW(),
    reception_id INTEGER NOT NULL REFERENCES receptions(id) ON DELETE RESTRICT,
    type product_type_enum NOT NULL
);