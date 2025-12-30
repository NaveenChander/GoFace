package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NaveenChander/GoFace/simulator/db"
	"github.com/NaveenChander/GoFace/simulator/models"
)

// CreateFactTransactionInstrumentDenominationHandler handles creation of denomination records.
func CreateFactTransactionInstrumentDenominationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Transaction Instrument Denomination endpoint hit")

	var denom models.FactTransactionInstrumentDenomination

	// Decode JSON body
	err := json.NewDecoder(r.Body).Decode(&denom)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Basic validation
	if denom.TransactionInstrumentSK == 0 || denom.DenominationCount == 0 || denom.DenominationValue == 0 {
		http.Error(w, "Missing required fields (transactionInstrumentSk, denominationCount, denominationValue)", http.StatusBadRequest)
		return
	}

	// Insert into DB
	if err := db.InsertFactTransactionInstrumentDenomination(denom); err != nil {
		fmt.Printf("Error inserting denomination: %v\n", err)
		http.Error(w, "Failed to create denomination record", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Transaction Instrument Denomination created successfully",
		"data":    denom,
	})
}
