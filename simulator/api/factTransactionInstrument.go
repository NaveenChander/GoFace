package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NaveenChander/GoFace/simulator/db"
	"github.com/NaveenChander/GoFace/simulator/models"
)

// CreateFactTransactionInstrumentHandler handles creation of transaction instrument records.
func CreateFactTransactionInstrumentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Transaction Instrument endpoint hit")

	var inst models.FactTransactionInstrument

	// Decode JSON body
	err := json.NewDecoder(r.Body).Decode(&inst)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Basic validation
	if inst.TransactionSK == 0 || inst.InstrumentTypeSK == 0 {
		http.Error(w, "Missing required fields (transactionSk, instrumentTypeSk)", http.StatusBadRequest)
		return
	}

	// Insert into DB
	if err := db.InsertFactTransactionInstrument(inst); err != nil {
		fmt.Printf("Error inserting transaction instrument: %v\n", err)
		http.Error(w, "Failed to create transaction instrument", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Transaction Instrument created successfully",
		"data":    inst,
	})
}
