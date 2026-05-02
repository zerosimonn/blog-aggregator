-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: ListFeeds :many
SELECT feeds.name, feeds.url, users.name AS user_name FROM feeds
JOIN users ON feeds.user_id = users.id;

-- name: GetFeedByURL :one
SELECT feeds.name, feeds.id FROM feeds
WHERE feeds.url = $1;

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*, feeds.name AS feed_name, users.name AS user_name
FROM feed_follows
INNER JOIN feeds ON feed_follows.feed_id = feeds.id
INNER JOIN users ON feed_follows.user_id = users.id
WHERE feed_follows.user_id = $1;