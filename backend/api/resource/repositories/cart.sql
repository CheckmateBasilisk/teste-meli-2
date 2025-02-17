-- create
-- name: CreateCart :one
INSERT INTO cart (
    id, product_id, customer_id, amount
) VALUES (
    uuid_generate_v4(), $1, $2, $3
)
    RETURNING *
;

-- read one
-- name: ReadByCartId :one
SELECT * FROM cart
    WHERE
        id = $1
    LIMIT 1
;

-- read many
-- name: ReadCart :many
SELECT * FROM cart
    ORDER BY product_id
;

-- update
-- name: UpdateCart :one
UPDATE cart
SET
    amount = $2

    WHERE
        id = $1
    RETURNING *
;

-- delete
-- name: DeleteCart :exec
DELETE FROM cart
    WHERE
        id = $1
;



