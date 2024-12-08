CREATE TABLE orders (
    order_id            SERIAL PRIMARY KEY,
    order_uid           TEXT UNIQUE,
    order_data          JSON
);