-- name: GetPosts :many
SELECT id, userid, title, description FROM posts
ORDER BY created_at DESC;

-- name: CreatePost :one
INSERT INTO posts (
  userid, title, description
) VALUES (
  $1,$2,$3
)
RETURNING id, userid, title, description;

-- name: UpdatePost :one
UPDATE posts
  set title = $2,
  description = $3
WHERE id = $1
RETURNING title, description;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;

-- name: GetPost :one
SELECT id, userid, title, description FROM posts
WHERE id = $1 LIMIT 1;