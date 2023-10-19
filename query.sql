-- name: GetAuthor :one
SELECT *
FROM authors
WHERE id = ?
LIMIT 1;

-- name: ListAuthors :many
SELECT *
FROM authors
ORDER BY name;

-- name: CreateAuthor :execresult
INSERT INTO authors (name, bio)
VALUES (?, ?);

-- name: DeleteAuthor :exec
DELETE
FROM authors
WHERE id = ?;

-- name: UpdateAuthor :exec
UPDATE authors
SET name = ?,
    bio  = ?
WHERE id = ?;

-- name: GetBook :one
SELECT *
FROM books
WHERE id = ?;

-- name: ListBooks :many
SELECT *
FROM books
ORDER BY title;

-- name: CreateBook :execresult
INSERT INTO books (title, description, author_id)
VALUES (?, ?, ?);