package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/NaveenChander/GoFace/simulator/db"
	"github.com/NaveenChander/GoFace/simulator/models"
)

// CreateFactTransactionHandler handles creation of a single fact transaction.
func CreateFactTransactionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Fact Transaction endpoint hit")

	var txn models.FactTransaction

	// Decode JSON body
	err := json.NewDecoder(r.Body).Decode(&txn)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields (basic check)
	if txn.SourceTransactionNumber == "" || txn.GamingDateSK == 0 || txn.TransactionDateSK == 0 {
		http.Error(w, "Missing required fields (sourceTransactionNumber, gamingDateSk, transactionDateSk)", http.StatusBadRequest)
		return
	}

	// Default time if zero
	if txn.TransactionUTCDateTime.IsZero() {
		txn.TransactionUTCDateTime = time.Now().UTC()
	}

	// Insert into DB
	if err := db.InsertFactTransaction(txn); err != nil {
		fmt.Printf("Error inserting transaction: %v\n", err)
		http.Error(w, "Failed to create transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Transaction created successfully",
		"data":    txn,
	})
}
