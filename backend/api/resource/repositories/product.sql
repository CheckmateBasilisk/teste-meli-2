-- create
-- name: CreateProduct :one
INSERT INTO product (
    id, barcode, name, price, stock, rating, descr, image
) VALUES (
    uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7
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
	stock = $5,
    rating = $6,
    descr = $7,
    image = $8

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


-- read many
-- name: ReadByProductName :many
SELECT * FROM product
    WHERE LOWER(name) LIKE CONCAT('%%',$1::TEXT,'%%')
    ORDER BY name DESC
;
