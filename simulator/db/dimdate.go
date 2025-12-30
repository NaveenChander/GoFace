package db

import (
	"fmt"
	"time"

	"github.com/NaveenChander/GoFace/simulator/models"
)

/*.          SQL Queries         .*/

func InsertDimDateRange(date models.DIMDate) {
	const query = `
		INSERT INTO dimdate (
			datesk,
			fulldate,
			dayofweek,
			month,
			quarter,
			year
		)
		VALUES ($1, $2, $3, $4, $5, $6);
	`

	_, err := DbContext.Pool.Exec(DbContext.Ctx, query, date.Datesk, date.Fulldate, date.Dayofweek, date.Month, date.Quarter, date.Year)
	if err != nil {
		fmt.Printf("Error inserting date %s: %v\n", date.Fulldate, err)
	}

}

func GetMaxFullDate() string {
	const query = `
		SELECT fulldate
		FROM dimdate
		ORDER BY fulldate DESC
		LIMIT 1;
	`
	var maxDate time.Time
	err := DbContext.Pool.QueryRow(DbContext.Ctx, query).Scan(&maxDate)
	if err != nil {
		fmt.Printf("Error fetching max fulldate: %v\n", err)
		return ""
	}

	return maxDate.Format("2006-01-02")
}
