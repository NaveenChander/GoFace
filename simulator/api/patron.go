package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/NaveenChander/GoFace/simulator/db"
	"github.com/NaveenChander/GoFace/simulator/models"
	"github.com/go-faker/faker/v4"
)

// CreateBulkPatronHandler handles bulk creation of patrons.
func CreateBulkPatronHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Bulk Patron Creation endpoint hit")
	cnt := r.URL.Query().Get("count")
	numPatrons, err := strconv.Atoi(cnt)
	if err != nil || numPatrons <= 0 {
		http.Error(w, "Invalid 'count' query parameter", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Here you would typically insert the patrons into the database.
	// For demonstration, we'll just return the received patrons.

	for i := range numPatrons {
		err := CreateAndInsertPatron()
		if err != nil {
			fmt.Printf("Error creating and inserting patron: %v\n", err)
		} else {
			fmt.Printf("Successfully inserted patron %d\n", i+1)
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Bulk patron creation completed"})
}

func NewPatron(subjectNumber, firstName, lastName string) models.Patron {
	return models.Patron{
		SubjectNumber: subjectNumber,
		FirstName:     firstName,
		LastName:      lastName,
	}
}

func CreateAndInsertPatron() error {
	subjectID, err := db.GetMaxID("Patron", "PatronID")
	if err != nil {
		return fmt.Errorf("failed to get max PatronID: %w", err)
	}

	firstName := faker.FirstName()
	lastName := faker.LastName()
	subjectID++
	subjectNumber := BuildSubjectNumber("AML", subjectID)
	patron := NewPatron(subjectNumber, firstName, lastName)

	if err := db.InsertPatron(patron); err != nil {
		return fmt.Errorf("failed to insert patron: %w", err)
	}

	return nil
}

/*.        Helper Functions       .*/

func BuildSubjectNumber(prefix string, num int) string {
	numStr := strconv.Itoa(num)
	totalLength := 10

	// Calculate how many zeros are needed in the middle
	zeroCount := totalLength - len(prefix) - len(numStr)
	if zeroCount < 0 {
		// If parameters exceed 10 characters, return as-is or handle as needed
		return prefix + numStr
	}

	zeros := strings.Repeat("0", zeroCount)
	return prefix + zeros + numStr
}
