package handlers

import (
	"encoding/json"
	"net/http"
	"receipt-processor-challenge-fetch/models"
	"receipt-processor-challenge-fetch/storage"

	"github.com/gorilla/mux"
)

func GetPoints(store *storage.MemoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		points, exists := store.GetPoints(id)
		if !exists {
			http.Error(w, "No receipt found for that id", http.StatusNotFound)
			return
		}

		response := models.PointsResponse{Points: points}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
