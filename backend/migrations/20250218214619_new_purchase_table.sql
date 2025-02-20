-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TYPE PURCHASE_STATUS AS ENUM (
    'completed',
    'pending',
    'cancelled'
);

CREATE TABLE IF NOT EXISTS purchase (
    id UUID PRIMARY KEY,

    status PURCHASE_STATUS NOT NULL DEFAULT  'pending',
    customer_id UUID NOT NULL,
    latest_update DATE DEFAULT CURRENT_DATE
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE IF EXISTS purchase;
