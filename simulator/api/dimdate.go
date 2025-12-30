package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/NaveenChander/GoFace/simulator/db"
	"github.com/NaveenChander/GoFace/simulator/models"
)

func InsertDimDateRange(w http.ResponseWriter, r *http.Request) {

	// Implementation for inserting date range into dimdate table{
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("InsertDimDateRange endpoint hit")

	//startDateStr := time.Date(2023, 01, 01, 00, 00, 00, 00, time.UTC).Format("2006-01-02")
	startDateStr := db.GetMaxFullDate()
	fmt.Println("Starting from date: ", startDateStr)
	if startDateStr == "" {
		startDateStr = "2023-01-01"
	} else {
		// Increment the start date by one day to avoid duplicate entry
		parsedDate, _ := time.Parse("2006-01-02", startDateStr)
		nextDate := parsedDate.AddDate(0, 0, 1)
		startDateStr = nextDate.Format("2006-01-02")
	}
	endDateStr := time.Now().Format("2006-01-02")
	//endDateStr := time.Date(2023, 03, 01, 00, 00, 00, 00, time.UTC).Format("2006-01-02")

	for date := startDateStr; date <= endDateStr; {
		// Insert date into dimdate table
		fmt.Printf("Inserting date: %s\n", date)
		parsedDate, _ := time.Parse("2006-01-02", date)

		dimDate := models.DIMDate{
			Datesk:    int(parsedDate.Unix()), // Example: using Unix timestamp as surrogate key
			Fulldate:  date,
			Dayofweek: int(parsedDate.Weekday() + 1), // Go's Weekday starts from 0 (Sunday)
			Month:     int(parsedDate.Month()),
			Quarter:   (int(parsedDate.Month())-1)/3 + 1,
			Year:      parsedDate.Year(),
		}

		db.InsertDimDateRange(dimDate)

		// Move to the next date
		nextDate, _ := time.Parse("2006-01-02", date)
		nextDate = nextDate.AddDate(0, 0, 1)
		date = nextDate.Format("2006-01-02")
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Date range insertion completed")
}
