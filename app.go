package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"sqlc-tutorial/books"

	_ "github.com/go-sql-driver/mysql"
)

func run() error {
	ctx := context.Background()
	fmt.Print(ctx)

	db, err := sql.Open("mysql", "test:test@/test?parseTime=true")
	if err != nil {
		return err
	}

	queries := books.New(db)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	createdAuthor, err := queries.CreateAuthor(ctx, books.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		return err
	}

	insertedAuthorID, err := createdAuthor.LastInsertId()
	if err != nil {
		return err
	}
	log.Println(insertedAuthorID)

	// create a book
	createdBook, err := queries.CreateBook(ctx, books.CreateBookParams{
		Title:       "The Go Programming Language",
		Description: sql.NullString{String: "The Go Programming Language is the authoritative resource for any programmer who wants to learn Go.", Valid: true},
		AuthorID:    insertedAuthorID,
	})
	if err != nil {
		return err
	}

	insertedBookID, err := createdBook.LastInsertId()
	if err != nil {
		return err
	}
	log.Println(insertedBookID)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthorID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthorID, fetchedAuthor.ID))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
