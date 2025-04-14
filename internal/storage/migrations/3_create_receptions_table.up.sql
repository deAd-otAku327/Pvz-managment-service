CREATE TYPE reception_status_enum AS ENUM (
    'in_progress',
    'close'
);

CREATE TABLE receptions (
    id SERIAL PRIMARY KEY,
    date_time TIMESTAMP NOT NULL DEFAULT NOW(),
    pvz_id INTEGER NOT NULL REFERENCES pvzs(id) ON DELETE RESTRICT,
    status reception_status_enum NOT NULL DEFAULT 'in_progress'
);