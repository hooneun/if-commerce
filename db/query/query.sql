-- name: CreateUser :execresult
INSERT INTO users (
  email, password
) VALUES (
  ?, ?
);

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = ? LIMIT 1;


-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ? LIMIT 1;
