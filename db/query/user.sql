
-- name: CreateUser :one
INSERT INTO users (
  user_name,
  full_name, 
  hashed_password,
  email
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE user_name = $1 LIMIT 1;
