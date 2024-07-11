package main

import (
	"log"
	"net/http"
	"receipt-processor-challenge-fetch/handlers"
	"receipt-processor-challenge-fetch/storage"

	"github.com/gorilla/mux"
)

func main() {
	store := storage.NewMemoryStore()

	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", handlers.ProcessReceipt(store)).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", handlers.GetPoints(store)).Methods("GET")

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
