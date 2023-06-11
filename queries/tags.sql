-- name: GetTags :many
SELECT id, tagname FROM tags;

-- name: CreateTag :one
INSERT INTO tags (
  tagname
) VALUES (
  $1
)
RETURNING id, tagname;

-- name: GetTagByPostID :many
SELECT 
	b.id,
	b.tagname
FROM post_tags a JOIN tags b
ON a.tagID = b.id
WHERE a.postid = $1;

-- name: UpdateTag :one
UPDATE tags
  set tagname = $2
WHERE id = $1
RETURNING id, tagname;

-- name: DeleteTag :exec
DELETE FROM tags
WHERE id = $1;

-- name: GetTag :one
SELECT id, tagname FROM tags
WHERE id = $1 LIMIT 1;