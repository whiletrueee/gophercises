package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"

	"gophercises/urlshort/db/pg_db"

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
	fmt.Println("shorten_url", shorten_url)

	// Execute the query

	// mux := defaultMux()

	// pathsToUrls := map[string]string{
	// 	"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	// 	"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	// }

	// mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// yaml := `
	// - path: /urlshort
	//   url: https://github.com/gophercises/urlshort
	// - path: /urlshort-final
	//   url: https://github.com/gophercises/urlshort/tree/solution
	// `

	// yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Starting server on 8080")
	// http.ListenAndServe(":8080", yamlHandler)

}

// func defaultMux() *http.ServeMux {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", hello)
// 	return mux
// }

// func hello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, world!")
// }

// func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		path := r.URL.Path
// 		if dest, ok := pathsToUrls[path]; ok {
// 			http.Redirect(w, r, dest, http.StatusFound)
// 			return
// 		}
// 		fallback.ServeHTTP(w, r)
// 	}
// }
