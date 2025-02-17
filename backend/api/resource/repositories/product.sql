-- create
-- name: CreateProduct :one
INSERT INTO product (
    id, barcode, name, price, stock
) VALUES (
    uuid_generate_v4(), $1, $2, $3, $4
)
    RETURNING *
;

-- read one
-- name: ReadByProductId :one
SELECT * FROM product
    WHERE
        id = $1
    LIMIT 1
;

-- read many
-- name: ReadProduct :many
SELECT * FROM product
    ORDER BY name
;

-- update
-- name: UpdateProduct :one
UPDATE product
SET
    barcode = $2,
	name = $3,
	price = $4,
	stock = $5

    WHERE
        id = $1
    RETURNING *
;

-- delete
-- name: DeleteProduct :exec
DELETE FROM product
    WHERE
        id = $1
;



