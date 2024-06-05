-- name: GetURL :one
SELECT * FROM url_shortener
WHERE original_url = ?;

-- name: CreateShortURL :execresult
INSERT INTO url_shortener (
  id, original_url, shortkey
) VALUES (
 ?, ?, ?
);