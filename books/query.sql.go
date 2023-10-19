// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query.sql

package books

import (
	"context"
	"database/sql"
)

const createAuthor = `-- name: CreateAuthor :execresult
INSERT INTO authors (name, bio)
VALUES (?, ?)
`

type CreateAuthorParams struct {
	Name string
	Bio  sql.NullString
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createAuthor, arg.Name, arg.Bio)
}

const createBook = `-- name: CreateBook :execresult
INSERT INTO books (title, description, author_id)
VALUES (?, ?, ?)
`

type CreateBookParams struct {
	Title       string
	Description sql.NullString
	AuthorID    int64
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createBook, arg.Title, arg.Description, arg.AuthorID)
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE
FROM authors
WHERE id = ?
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAuthor, id)
	return err
}

const getAuthor = `-- name: GetAuthor :one
SELECT id, name, bio
FROM authors
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetAuthor(ctx context.Context, id int64) (Author, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, id)
	var i Author
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}

const getBook = `-- name: GetBook :one
SELECT id, title, author_id, description
FROM books
WHERE id = ?
`

func (q *Queries) GetBook(ctx context.Context, id int64) (Book, error) {
	row := q.db.QueryRowContext(ctx, getBook, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.AuthorID,
		&i.Description,
	)
	return i, err
}

const listAuthors = `-- name: ListAuthors :many
SELECT id, name, bio
FROM authors
ORDER BY name
`

func (q *Queries) ListAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, listAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(&i.ID, &i.Name, &i.Bio); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listBooks = `-- name: ListBooks :many
SELECT id, title, author_id, description
FROM books
ORDER BY title
`

func (q *Queries) ListBooks(ctx context.Context) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, listBooks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.AuthorID,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAuthor = `-- name: UpdateAuthor :exec
UPDATE authors
SET name = ?,
    bio  = ?
WHERE id = ?
`

type UpdateAuthorParams struct {
	Name string
	Bio  sql.NullString
	ID   int64
}

func (q *Queries) UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) error {
	_, err := q.db.ExecContext(ctx, updateAuthor, arg.Name, arg.Bio, arg.ID)
	return err
}