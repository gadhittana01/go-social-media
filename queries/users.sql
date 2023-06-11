-- name: GetUsers :many
SELECT id, fullname FROM users
ORDER BY created_at DESC;

-- name: CreateUser :one
INSERT INTO users (
  fullname
) VALUES (
  $1
)
RETURNING id, fullname;

-- name: UpdateUser :one
UPDATE users
  set fullname = $2
WHERE id = $1
RETURNING id, fullname;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: GetUser :one
SELECT id, fullname FROM users
WHERE id = $1 LIMIT 1;