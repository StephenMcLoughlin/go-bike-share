CREATE TABLE dock (
    dock_id SERIAL PRIMARY KEY NOT NULL,
    location VARCHAR(255)
);

CREATE TABLE user (
    id SERIAL PRIMARY KEY NOT NULL,
    first_name VARCHAR(20) NOT NULL,
    last_name VARCHAR(20) NOT NULL,
    email VARCHAR(40) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

-- CREATE TYPE bike_status AS ENUM ('available', 'in_use', 'under_maintenance');

CREATE TABLE bike (
    id SERIAL PRIMARY KEY NOT NULL,
    dock_id INT REFERENCES dock(id),
    user_id INT REFERENCES user(id),
    -- status bike_status 
);

CREATE TABLE bike_history (
    id SERIAL PRIMARY KEY NOT NULL,
    bike_id INT REFERENCES bike (id),
    user_id INT REFERENCES user (id),
    start_time TIMESTAMP,
    end_time TIMESTAMP
);

CREATE TABLE trip (
    id SERIAL PRIMARY KEY NOT NULL,
    user_id INT REFERENCES user(id),
    bike_id INT REFERENCES bike(id),
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    start_dock_id INT REFERENCES dock(id),
    end_dock_id INT REFERENCES dock(id)
)

-- COPY dock(dockid, bikeid) FROM '/docker-entrypoint-initdb.d/dock.csv' DELIMITER ',' CSV HEADER;

