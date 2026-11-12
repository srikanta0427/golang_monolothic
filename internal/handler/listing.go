package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type listings struct {
	Id string	`json:"id"`
	Title string	`json:"title"`
	Description string	`json:"description"`
	Price int	`json:"price"`
	City string	`json:"city"`
}

func List(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		const query = `SELECT * FROM listings`
		rows, err := db.Query(query)
		if err != nil {
			log.Printf("query %v", err)
			http.Error(w, "Query Error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		listing := []listings{}
		for rows.Next(){
			var l listings
			if err := rows.Scan(&l.Id, &l.Title, &l.Description, &l.Price, &l.City); err != nil {
				log.Printf("scan %v", err)
				http.Error(w, "Scan Error", http.StatusInternalServerError)
				return
			}
			listing = append(listing, l)
		}
		if err := rows.Err(); err != nil {
			log.Printf("rows %v", err)
			http.Error(w, "Rows Error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content_Type","application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(listing)
	}
}
