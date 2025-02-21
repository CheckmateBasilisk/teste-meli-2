-- create
-- name: CreateCustomer :one
INSERT INTO customer (
    id, login, name, password
) VALUES (
    uuid_generate_v4(), $1, $2, $3
)
    RETURNING *
;

-- read one
-- name: ReadByCustomerId :one
SELECT * FROM customer
    WHERE
        id = $1
    LIMIT 1
;

-- read many
-- name: ReadCustomer :many
SELECT * FROM customer
    ORDER BY name
;

-- update
-- name: UpdateCustomer :one
UPDATE customer
SET
    login = $2,
	name = $3,
    password = $4

    WHERE
        id = $1
    RETURNING *
;

-- delete
-- name: DeleteCustomer :exec
DELETE FROM customer
    WHERE
        id = $1
;

