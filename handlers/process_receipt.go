package handlers

import (
	"encoding/json"
	"net/http"
	"receipt-processor-challenge-fetch/models"
	"receipt-processor-challenge-fetch/services"
	"receipt-processor-challenge-fetch/storage"

	"github.com/google/uuid"
)

func ProcessReceipt(store *storage.MemoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var receipt models.Receipt
		if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
			http.Error(w, "Invalid receipt data", http.StatusBadRequest)
			return
		}

		id := uuid.New().String()
		store.SaveReceipt(id, receipt)

		points := services.CalculatePoints(receipt)
		store.SavePoints(id, points)

		response := models.ReceiptResponse{ID: id}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
