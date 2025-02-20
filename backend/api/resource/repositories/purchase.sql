
-- create
-- name: CreatePurchase :one
INSERT INTO purchase (
    id, status, customer_id, latest_update
) VALUES (
    uuid_generate_v4(), $1, $2, CURRENT_DATE
)
    RETURNING *
;

-- read one
-- name: ReadByPurchaseId :one
SELECT * FROM purchase
    WHERE
        id = $1
    LIMIT 1
;

-- read many
-- name: ReadPurchase :many
SELECT * FROM purchase
    ORDER BY latest_update
;

-- update
-- name: UpdatePurchase :one
UPDATE purchase
SET
    status = $2,
    customer_id = $3,
    latest_update = CURRENT_DATE

    WHERE
        id = $1
    RETURNING *
;

-- delete
-- name: DeletePurchase :exec
DELETE FROM purchase
    WHERE
        id = $1
;



