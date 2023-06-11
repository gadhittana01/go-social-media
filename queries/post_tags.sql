-- name: CreatePostTag :one
INSERT INTO post_tags (
  postid, tagid
) VALUES (
  $1,$2
)
RETURNING id, postid, tagid;

-- name: DeletePostTag :exec
DELETE FROM post_tags
WHERE postid = $1;