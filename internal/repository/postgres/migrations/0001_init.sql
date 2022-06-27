-- +migrate Up
CREATE TABLE IF NOT EXISTS urls
(
    id serial primary key,
    long_url text unique not null,
    short_url character varying(10) unique
);

-- +migrate Down
DROP TABLE IF EXISTS urls;