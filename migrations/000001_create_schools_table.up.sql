--Filename: 000001_create_schools_table.up.sql

CREATE TABLE IF NOT EXISTS school (
id bigserial PRIMARY KEY,
created_at timestamp (0) with time zone NOT NULL DEFAULT NOW(),
version integer NOT NULL DEFAULT 1
);