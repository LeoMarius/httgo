-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;

-- name: GetUserEmail :one
SELECT *
FROM users
WHERE email = $1;


-- name: GetUserID :one
SELECT *
FROM users
WHERE id = $1;

-- name: UpdateUsers :exec
UPDATE users
SET 
    email = $1,
    hashed_password = $2,
    updated_at = NOW()
WHERE id = $3;

-- name: UpgradeRed :exec
UPDATE users
SET 
    updated_at = NOW(),
    is_chirpy_red = TRUE
WHERE id = $1;


-- name: DeleteUsers :exec
DELETE FROM users;

