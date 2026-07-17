-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password)
VALUES (
  gen_random_uuid(), NOW(), NOW(), $1, $2
)
RETURNING *;

-- name: DeleteUsers :exec
DELETE FROM users;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserFromRefreshToken :one
SELECT u.*
FROM users u
JOIN refresh_tokens rt ON u.id = rt.user_id
WHERE rt.token = $1 AND rt.expires_at > NOW() AND rt.revoked_at IS NULL;

-- name: UpdateUser :one
UPDATE users
SET updated_at = NOW(), email = $2, hashed_password = $3
WHERE id = $1
RETURNING *;

-- name: SetUserChirpyRed :one
UPDATE users
SET updated_at = NOW(), is_chirpy_red = TRUE
WHERE id = $1
RETURNING *;