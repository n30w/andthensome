package main

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	RedditPosts, RedditComments, DBPosts, DBComments *Table[Row[id, text]]
	Key                                              *Key
	*PlanetscaleDB
}

// UpdateHandler handles updating SQL database requests
func (s *Server) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	GrabSaved(s.RedditPosts, s.RedditComments, s.Key)

	err = s.RetrieveSQL(s.DBPosts, s.DBComments)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s\n", err)))
		log.Fatal(err)
	}

	err = s.UpdateSQL(s.DBPosts, s.RedditPosts, add)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s\n", err)))
		log.Fatal(err)
	}

	err = s.UpdateSQL(s.DBComments, s.RedditComments, add)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%s\n", err)))
		log.Fatal(err)
	}

	w.Write([]byte(Result.Sprintf("Successfully updated Planetscale Database\n")))
	ClearTable(s.RedditPosts, s.RedditComments, s.DBPosts, s.DBComments)
}

// PopulateHandler handles populating tables requests
func (s *Server) PopulateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(Result.Sprintf("Successfully populated Planetscale Database\n")))
	GrabSaved(s.RedditPosts, s.RedditComments, s.Key)

	if err := s.InsertToSQL(s.RedditPosts.Name, s.RedditPosts.Rows[:]); err != nil {
		fmt.Println(err)
	}

	if err := s.InsertToSQL(s.RedditComments.Name, s.RedditComments.Rows[:]); err != nil {
		fmt.Println(err)
	}
}

// ClearTableHandler handles clearing tables requests
func (s *Server) ClearTableHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(Result.Sprintf("Cleared all rows from all tables\n")))
	if err := s.UpdateSQL(s.DBPosts, s.RedditPosts, delete); err != nil {
		fmt.Println(err)
	}

	if err := s.UpdateSQL(s.DBComments, s.RedditComments, delete); err != nil {
		fmt.Println(err)
	}
}
