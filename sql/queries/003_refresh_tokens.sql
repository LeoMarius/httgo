-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (token, created_at, updated_at, user_id, expires_at, revoked_at)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    NOW() + INTERVAL '60 days', -- revoked_at)
    NULL
)
RETURNING *;



-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens
SET 
    updated_at = NOW(),
    revoked_at = NOW()
WHERE token = $1;



-- name: DeleteAllTokens :exec
DELETE FROM refresh_tokens;


-- name: FindToken :one
SELECT *
FROM refresh_tokens
WHERE token = $1;