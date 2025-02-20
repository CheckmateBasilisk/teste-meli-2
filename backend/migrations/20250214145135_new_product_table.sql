-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS product (
    id UUID PRIMARY KEY,

    barcode TEXT NOT NULL,
    name TEXT NOT NULL,
    price INTEGER NOT NULL,
    stock INTEGER NOT NULL,
    rating INTEGER NOT NULL,
    descr TEXT NOT NULL,

    -- TODO: this will just point to random images @https://picsum.photos/id/{image}/300/200
    image TEXT NOT NULL
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE IF EXISTS product;
