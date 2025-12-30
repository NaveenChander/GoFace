package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NaveenChander/GoFace/simulator/db"
	"github.com/NaveenChander/GoFace/simulator/models"
)

// CreateFactTransactionSessionHandler handles creation of session records.
func CreateFactTransactionSessionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Transaction Session endpoint hit")

	var session models.FactTransactionSession

	// Decode JSON body
	err := json.NewDecoder(r.Body).Decode(&session)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Basic validation (PropertySK, GamingDateSK, Username, ClientName are required)
	if session.PropertySK == 0 || session.GamingDateSK == 0 || session.Username == "" || session.ClientName == "" {
		http.Error(w, "Missing required fields (propertySk, gamingDateSk, username, clientName)", http.StatusBadRequest)
		return
	}

	// Insert into DB
	if err := db.InsertFactTransactionSession(session); err != nil {
		fmt.Printf("Error inserting session: %v\n", err)
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Transaction Session created successfully",
		"data":    session,
	})
}
