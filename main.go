package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

const PleasePopulateIDs = false

var (
	db *sql.DB
)

func init() {
	var err error
	env := "PROD"
	db, err = sql.Open("mysql", os.Getenv(env))
	if err != nil {
		panic(Warn.Sprint(err))
	}

	if err := db.Ping(); err != nil {
		panic(Warn.Sprint(err))
	}
}

func main() {
	server := &Server{
		RedditPosts:    &Table[Row[id, text]]{Name: "posts"},
		RedditComments: &Table[Row[id, text]]{Name: "comments"},
		DBPosts:        &Table[Row[id, text]]{Name: "posts"},
		DBComments:     &Table[Row[id, text]]{Name: "comments"},
		Key:            &Key{},
		Psdb:           &PlanetscaleDB{db},
	}

	log.Print(Start.Sprint("Server listening on :4000"))

	mux := http.NewServeMux()
	mux.HandleFunc("/", server.UpdateHandler)

	if err := http.ListenAndServe(":4000", mux); err != nil {
		log.Fatal(Warn.Sprint(err))
	}
}
