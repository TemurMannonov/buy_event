CREATE TABLE IF NOT EXISTS customer (
    id              UUID PRIMARY KEY,
    phone_number    VARCHAR UNIQUE NOT NULL,
    email           VARCHAR UNIQUE NOT NULL
);


CREATE TABLE IF NOT EXISTS "order" (
    id              UUID PRIMARY KEY,
    customer_id     UUID NOT NULL REFERENCES customer (id) ON DELETE RESTRICT,
    products        TEXT NOT NULL,
    total_price     NUMERIC(10, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS log (
    id          UUID PRIMARY KEY,
    message     TEXT NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)