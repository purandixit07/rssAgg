-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPostsByUser :many
SELECT posts.* FROM posts
JOIN feedFollows ON posts.feed_id = feedFollows.feed_id
WHERE feedFollows.user_id = $1 
ORDER BY posts.published_at DESC 
LIMIT $2;