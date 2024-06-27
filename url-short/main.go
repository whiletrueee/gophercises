package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"

	"gophercises/url-short/db/pg_db"

	"github.com/aidarkhanov/nanoid"
	_ "github.com/go-sql-driver/mysql"
	"go.bryk.io/pkg/ulid"
)

func main() {
	// Define the flags
	original_url := flag.String("ou", "", "The original URL")
	flag.Parse()
	if *original_url == "" {
		fmt.Println("Please provide the original URL")
		return
	}
	// Open up our database connection.
	dbConn, err := sql.Open("mysql", "root:231304@tcp(0.0.0.0:3306)/url")
	if err != nil {
		panic(err.Error())
	}
	defer dbConn.Close()

	queries := pg_db.New(dbConn)

	ctx := context.Background()
	// Create a table

	ulid, err := ulid.New()
	if err != nil {
		panic(err)
	}
	sk, err := nanoid.Generate("qwert", 5)
	if err != nil {
		panic(err)
	}
	shorten_url, err := queries.CreateShortURL(ctx, pg_db.CreateShortURLParams{
		ID:          ulid.String(),
		OriginalUrl: *original_url,
		ShortKey:    sk,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("shorten_url", sk, shorten_url)

}
