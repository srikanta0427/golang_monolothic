package handler

import (
	"database/sql"
	"log"
	"net/http"
)

func Listing(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		const query = `SELECT * FROM listings`
		rows, err := db.Query(query)
		if err != nil {
			log.Printf("query %v", err)
			http.Error(w, "Query Error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		w.Write([]byte("Listing data"))
	}
}
