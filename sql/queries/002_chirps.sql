-- name: CreateChirp :one
INSERT INTO chirps (id, created_at, updated_at, body, user_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;


-- name: GetAllChirpsASC :many
SELECT *
FROM chirps
ORDER BY created_at ASC;

-- name: GetAllChirpsDESC :many
SELECT *
FROM chirps
ORDER BY created_at DESC;


-- name: GetChirpsByUserIDASC :many
SELECT *
FROM chirps
WHERE user_id = $1
ORDER BY created_at ASC;


-- name: GetChirpsByUserIDDESC :many
SELECT *
FROM chirps
WHERE user_id = $1
ORDER BY created_at DESC;



-- name: GetChirpsByID :one
SELECT *
FROM chirps
WHERE id = $1;

-- name: DeleteChirpsByID :exec
DELETE
FROM chirps
WHERE id = $1;


-- name: DeleteAllChirps :exec
DELETE FROM chirps;

