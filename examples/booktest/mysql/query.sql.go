// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: query.sql

package booktest

import (
	"context"
	"database/sql"
	"time"
)

const booksByTags = `-- name: BooksByTags :many
SELECT
  book_id,
  title,
  name,
  isbn,
  tags
FROM books
LEFT JOIN authors ON books.author_id = authors.author_id
WHERE tags = ?
`

type BooksByTagsRow struct {
	BookID int32
	Title  string
	Name   sql.NullString
	Isbn   string
	Tags   string
}

func (q *Queries) BooksByTags(ctx context.Context, tags string) ([]BooksByTagsRow, error) {
	rows, err := q.db.QueryContext(ctx, booksByTags, tags)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BooksByTagsRow
	for rows.Next() {
		var i BooksByTagsRow
		if err := rows.Scan(
			&i.BookID,
			&i.Title,
			&i.Name,
			&i.Isbn,
			&i.Tags,
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

const booksByTitleYear = `-- name: BooksByTitleYear :many
SELECT book_id, author_id, isbn, book_type, title, yr, available, tags FROM books
WHERE title = ? AND yr = ?
`

type BooksByTitleYearParams struct {
	Title string
	Yr    int32
}

func (q *Queries) BooksByTitleYear(ctx context.Context, arg BooksByTitleYearParams) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, booksByTitleYear, arg.Title, arg.Yr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.BookID,
			&i.AuthorID,
			&i.Isbn,
			&i.BookType,
			&i.Title,
			&i.Yr,
			&i.Available,
			&i.Tags,
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

const createAuthor = `-- name: CreateAuthor :execresult
INSERT INTO authors (name) VALUES (?)
`

func (q *Queries) CreateAuthor(ctx context.Context, name string) (sql.Result, error) {
	return q.db.ExecContext(ctx, createAuthor, name)
}

const createBook = `-- name: CreateBook :execresult
INSERT INTO books (
    author_id,
    isbn,
    book_type,
    title,
    yr,
    available,
    tags
) VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
)
`

type CreateBookParams struct {
	AuthorID  int32
	Isbn      string
	BookType  BooksBookType
	Title     string
	Yr        int32
	Available time.Time
	Tags      string
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createBook,
		arg.AuthorID,
		arg.Isbn,
		arg.BookType,
		arg.Title,
		arg.Yr,
		arg.Available,
		arg.Tags,
	)
}

const deleteAuthorBeforeYear = `-- name: DeleteAuthorBeforeYear :exec
DELETE FROM books
WHERE yr < ? AND author_id = ?
`

type DeleteAuthorBeforeYearParams struct {
	Yr       int32
	AuthorID int32
}

func (q *Queries) DeleteAuthorBeforeYear(ctx context.Context, arg DeleteAuthorBeforeYearParams) error {
	_, err := q.db.ExecContext(ctx, deleteAuthorBeforeYear, arg.Yr, arg.AuthorID)
	return err
}

const deleteBook = `-- name: DeleteBook :exec
DELETE FROM books
WHERE book_id = ?
`

func (q *Queries) DeleteBook(ctx context.Context, bookID int32) error {
	_, err := q.db.ExecContext(ctx, deleteBook, bookID)
	return err
}

const getAuthor = `-- name: GetAuthor :one
SELECT author_id, name FROM authors
WHERE author_id = ?
`

func (q *Queries) GetAuthor(ctx context.Context, authorID int32) (Author, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, authorID)
	var i Author
	err := row.Scan(&i.AuthorID, &i.Name)
	return i, err
}

const getBook = `-- name: GetBook :one
SELECT book_id, author_id, isbn, book_type, title, yr, available, tags FROM books
WHERE book_id = ?
`

func (q *Queries) GetBook(ctx context.Context, bookID int32) (Book, error) {
	row := q.db.QueryRowContext(ctx, getBook, bookID)
	var i Book
	err := row.Scan(
		&i.BookID,
		&i.AuthorID,
		&i.Isbn,
		&i.BookType,
		&i.Title,
		&i.Yr,
		&i.Available,
		&i.Tags,
	)
	return i, err
}

const updateBook = `-- name: UpdateBook :exec
UPDATE books
SET title = ?, tags = ?
WHERE book_id = ?
`

type UpdateBookParams struct {
	Title  string
	Tags   string
	BookID int32
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) error {
	_, err := q.db.ExecContext(ctx, updateBook, arg.Title, arg.Tags, arg.BookID)
	return err
}

const updateBookISBN = `-- name: UpdateBookISBN :exec
UPDATE books
SET title = ?, tags = ?, isbn = ?
WHERE book_id = ?
`

type UpdateBookISBNParams struct {
	Title  string
	Tags   string
	Isbn   string
	BookID int32
}

func (q *Queries) UpdateBookISBN(ctx context.Context, arg UpdateBookISBNParams) error {
	_, err := q.db.ExecContext(ctx, updateBookISBN,
		arg.Title,
		arg.Tags,
		arg.Isbn,
		arg.BookID,
	)
	return err
}
