package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NaveenChander/GoFace/simulator/db"
	"github.com/NaveenChander/GoFace/simulator/models"
)

// CreateFactTransactionInstrumentDetailHandler handles creation of instrument detail records.
func CreateFactTransactionInstrumentDetailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Transaction Instrument Detail endpoint hit")

	var detail models.FactTransactionInstrumentDetail

	// Decode JSON body
	err := json.NewDecoder(r.Body).Decode(&detail)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Basic validation
	if detail.TransactionInstrumentSK == 0 || detail.Value == "" {
		http.Error(w, "Missing required fields (transactionInstrumentSk, value)", http.StatusBadRequest)
		return
	}

	// Insert into DB
	if err := db.InsertFactTransactionInstrumentDetail(detail); err != nil {
		fmt.Printf("Error inserting instrument detail: %v\n", err)
		http.Error(w, "Failed to create instrument detail", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Transaction Instrument Detail created successfully",
		"data":    detail,
	})
}
