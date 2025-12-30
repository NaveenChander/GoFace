package db

import (
	"errors"
	"fmt"

	"github.com/NaveenChander/GoFace/simulator/models"
)

/*.          SQL Queries         .*/

func InsertPatron(patron models.Patron) error {

	const query = `
		INSERT INTO Patron (
			SubjectNumber,
			FirstName,
			LastName
		)
		VALUES ($1, $2, $3);
	`

	_, err := DbContext.Pool.Exec(DbContext.Ctx, query, patron.SubjectNumber, patron.FirstName, patron.LastName)
	return err
}

// GetMaxID returns the max integer value from a given table and column.
func GetMaxID(tableName, columnName string) (int, error) {

	// Build SQL safely using identifiers, not parameters.
	query := fmt.Sprintf(`SELECT COALESCE(MAX(%s), 0) FROM %s`, columnName, tableName)

	var maxID int
	err := DbContext.Pool.QueryRow(DbContext.Ctx, query).Scan(&maxID)
	if err != nil {
		return 0, errors.New("failed to query max ID: " + err.Error())
	}

	return maxID, nil
}
